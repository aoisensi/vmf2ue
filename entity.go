package main

import (
	"strconv"
	"strings"

	"github.com/aoisensi/vmf2ue/vmf"
	cprint "github.com/fatih/color"
)

func writeEntity(entity vmf.Entity) {
	class := entity["classname"].(string)
	id, _ := strconv.Atoi(entity["id"].(string))

	if class == "info_player_teamspawn" && hasOnly("spawn") {
		writeBegin("Actor", "Class=/Game/Unreal4tress/Core/GameModes/Common/BP_PlayerStart.BP_PlayerStart_C Name=BP_PlayerStart_%v", id)
		if entity.Has("TeamNum") {
			team := map[string]string{
				"": "None", "1": "None", "2": "Red", "3": "Blue",
			}[entity.String("TeamNum")]
			if entity.Has("StartDisabled") {
				team = "None"
			}
			writef("Team=%v", team)
		}
		writeBegin("Object", "Name=\"Capsule\"")
		writeOrigin(entity, &Vec3{0, 0, 41.5})
		writeAngles(entity, &Vec3{0, -90, 0})
		writeEnd()
		writef("ActorLable=\"BP_PlayerStart_%v\"", id)
		writef("FolderPath=\"PlayerStarts\"")
		writeEnd()
		return
	}
	if class == "prop_static" && hasOnly("mesh") {
		model := entity.String("model")
		asset, assetOk := bind.Props[model]
		if !assetOk {
			if _, warned := unknownMeshes[model]; !warned {
				unknownMeshes[model] = struct{}{}
				cprint.Yellow("Unknown mesh detect: \"%v\" ID: %v", model, id)
			}

			return
		}
		name := strings.Split(asset.Asset, ".")[1]
		writeBegin("Actor", "Class=/Script/Engine.StaticMeshActor Name=%v_%v", name, id)
		writeBegin("Object", "Name=\"StaticMeshComponent0\"")
		writef("StaticMesh='\"%v\"'", asset.Asset)
		writeOrigin(entity, nil)
		writeAngles(entity, nil)
		writeEnd()
		writef("ActorLabel=\"%v_%v\"", name, id)
		writef("FolderPath=\"Props\"")
		writef("StaticMeshes=\"Solids\"")
		writeEnd()
		return
	}
	if class == "light_environment//ignore" { //ignore
		writeBegin("Actor", "Class=/Script/Engine.SkyLight Name=SkyLight_%v", id)
		writeBegin("Object", "Name=\"SkyLightComponent0\"")
		writef("CastShadow=False")
		writeOrigin(entity, nil)
		writef("Mobility=Static")
		writeEnd()
		writef("ActorLabel=\"SkyLight_%v\"", id)
		writeEnd()
		return
	}
	if class == "light_spot" && hasOnly("light") {
		writeBegin("Actor", "Class=/Script/Engine.SpotLight Name=SpotLight_%v", id)
		writeBegin("Object", "Name=\"LightComponent0\"")
		innerConeAngle := entity.Float("_inner_cone")
		writef("InnerConeAngle=%.6f", innerConeAngle)
		outerConeAngle := entity.Float("_cone")
		writef("outerConeAngle=%.6f", outerConeAngle)
		attenuationRadius := 0.0
		if entity.Has("_zero_percent_distance") {
			attenuationRadius = entity.Float("_zero_percent_distance")
		}
		writef("AttenuationRadius=%.6f", attenuationRadius*SCALE)
		color := entity.FloatSlice("_light")
		writef("Intensity=%.6f", color[3]*BRIGHTNESS)
		writef("LightColor=(R=%.0f,G=%.0f,B=%.0f,A=255)", color[0], color[1], color[2])
		writeOrigin(entity, nil)
		rot := ParseVec3(entity["angles"].(string))
		pitch := entity.Float("pitch")
		writef("RelativeRotation=(Pitch=%.6f,Yaw=%.6f,Roll=%.6f)", pitch, 90.0-rot[1], rot[2])
		writef("Mobility=Static")
		writeEnd()
		writef("ActorLable=\"SpotLight_%v\"", id)
		writef("FolderPath=\"Lights\"")
		writeEnd()
		return
	}
	if _, skip := skipClasses[class]; !skip {
		if _, warned := unknownClasses[class]; !warned && only == "" {
			unknownClasses[class] = struct{}{}
			cprint.Yellow("Unknown entity class detect: \"%v\" ID: %v", class, id)
		}
	}
}

func writeOrigin(entity vmf.Entity, origin *Vec3) {
	pos := ParseVec3(entity["origin"].(string))
	if origin != nil {
		pos = pos.Add(*origin)
	}
	writef("RelativeLocation=(X=%.6f,Y=%.6f,Z=%.6f)", pos[1]*SCALE, pos[0]*SCALE, pos[2]*SCALE)

}

func writeAngles(entity vmf.Entity, angles *Vec3) {
	rot := ParseVec3(entity["angles"].(string))
	if angles != nil {
		rot = rot.Add(*angles)
	}
	writef("RelativeRotation=(Pitch=%.6f,Yaw=%.6f,Roll=%.6f)", rot[2], -rot[1], rot[0])
}
