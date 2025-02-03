package utils

import (
	"os/exec"
	"regexp"
	"strings"
)

func TokenizeInput(input string) (cmd string, args []string, argsExist bool) {
	// Break out the input into a slice... this is only used to get the command in the next step
	inputParts := strings.Fields(input)
	var parsedArgs []string

	// Since the first 'part'  of the input needs to be the command...
	// ... we can strip it out of the input and are left with the args.
	rawArgs, _ := strings.CutPrefix(input, inputParts[0])
	// Handle trailing space
	rawArgs = strings.TrimSpace(rawArgs)
	argsFound := rawArgs != ""

	if argsFound {

		regex := regexp.MustCompile(`'([^']*)'|\S+`)
		matches := regex.FindAllStringSubmatch(rawArgs, -1)

		var buffer strings.Builder

		for i, match := range matches {
			// First check to see if this is a quoted string...
			if match[1] != "" {
				// ... if so we add the matched string to a buffer
				buffer.WriteString(match[1])
				// We now need to check the next match...
				// ...to determine if it is adjacent to the current match.
				// To accomplish this we check 2 things...
				// - are we at the end of the `matches` slice
				// - is the next match a quoted string
				if i < len(matches)-1 && matches[i+1][1] != "" {
					// Get the start and end indicies of each match
					matchesIdx := regex.FindAllStringIndex(rawArgs, -1)
					// Get the end index of the current match...
					matchCurrEnd := matchesIdx[i][1]
					// ... and the start index of the next match
					matchNextStart := matchesIdx[i+1][0]

					// Check to see if the strings are adjacent...
					if rawArgs[matchCurrEnd:matchNextStart] == "" {
						// ... if so we just continue until we find either...
						// - a non-adjacent string
						// - an unquoted string
						continue
					} else {
						// if the current and next strings ARE NOT adjacent...
						// ... add the current buffer to parsedArgs and flush
						parsedArgs = append(parsedArgs, buffer.String())
						buffer.Reset()
					}
				} else {
					// If we are either at the end of the slice...
					// ... or the next match isn't quoted...
					// ... add the current buffer to parsedArgs and flush
					parsedArgs = append(parsedArgs, buffer.String())
					buffer.Reset()
				}
			} else {
				// If this is a non quoted string...
				// ... add the current buffer to parsedArgs and flush
				if buffer.Len() > 0 {
					parsedArgs = append(parsedArgs, buffer.String())
					buffer.Reset()
				}
				// ... then we can add this unquoted string to parsedArgs
				parsedArgs = append(parsedArgs, match[0])
			}
		}
	}
	return inputParts[0], parsedArgs, argsFound
}

func CheckCmdType(cmd string) (cmdType string, cmdArgs string) {
	if CheckBuiltin(cmd) {
		return "builtin", ""
	}
	if cmdPath := CheckExecutable(cmd); cmdPath != "" {
		return "executable", cmdPath
	}
	return "", ""
}

func CheckBuiltin(cmd string) bool {
	builtin := map[string]struct{}{
		"cd":   {},
		"echo": {},
		"exit": {},
		"pwd":  {},
		"type": {},
	}
	_, exists := builtin[cmd]
	return exists
}

// TODO: add alias

func CheckExecutable(args string) string {
	cmdPath, err := exec.LookPath(args)
	if err != nil {
		return ""
	}
	return cmdPath
}

func SliceToString(slice []string) string {
	sliceString := strings.TrimSpace(strings.Join(slice, " "))
	return sliceString
}
