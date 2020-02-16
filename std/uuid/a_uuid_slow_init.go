// This file is generated by generate-std.joke script. Do not edit manually!

package uuid

import (
	"fmt"
	. "github.com/candid82/joker/core"
	"os"
)

func InternsOrThunks() {
	if VerbosityLevel > 0 {
		fmt.Fprintln(os.Stderr, "Lazily running slow version of uuid.InternsOrThunks().")
	}
	uuidNamespace.ResetMeta(MakeMeta(nil, `Generates UUIDs.`, "1.0"))

	uuidNamespace.InternVar("new", new_,
		MakeMeta(
			NewListFrom(NewVectorFrom()),
			`Creates a new random UUID.`, "1.0").Plus(MakeKeyword("tag"), String{S: "String"}))

}
