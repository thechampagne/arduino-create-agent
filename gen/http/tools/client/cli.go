// Code generated by goa v3.12.4, DO NOT EDIT.
//
// tools HTTP client CLI support package
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package client

import (
	"encoding/json"
	"fmt"

	tools "github.com/arduino/arduino-create-agent/gen/tools"
)

// BuildInstallPayload builds the payload for the tools install endpoint from
// CLI flags.
func BuildInstallPayload(toolsInstallBody string) (*tools.ToolPayload, error) {
	var err error
	var body InstallRequestBody
	{
		err = json.Unmarshal([]byte(toolsInstallBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"checksum\": \"Beatae dolor adipisci nulla et quam voluptas.\",\n      \"name\": \"avrdude\",\n      \"packager\": \"arduino\",\n      \"url\": \"Deserunt voluptatem impedit iusto libero.\",\n      \"version\": \"6.3.0-arduino9\"\n   }'")
		}
	}
	v := &tools.ToolPayload{
		Name:     body.Name,
		Version:  body.Version,
		Packager: body.Packager,
		URL:      body.URL,
		Checksum: body.Checksum,
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the tools remove endpoint from CLI
// flags.
func BuildRemovePayload(toolsRemoveBody string, toolsRemovePackager string, toolsRemoveName string, toolsRemoveVersion string) (*tools.ToolPayload, error) {
	var err error
	var body RemoveRequestBody
	{
		err = json.Unmarshal([]byte(toolsRemoveBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"checksum\": \"Ipsa minima quia.\",\n      \"url\": \"Expedita rem ipsum quasi harum nostrum.\"\n   }'")
		}
	}
	var packager string
	{
		packager = toolsRemovePackager
	}
	var name string
	{
		name = toolsRemoveName
	}
	var version string
	{
		version = toolsRemoveVersion
	}
	v := &tools.ToolPayload{
		URL:      body.URL,
		Checksum: body.Checksum,
	}
	v.Packager = packager
	v.Name = name
	v.Version = version

	return v, nil
}
