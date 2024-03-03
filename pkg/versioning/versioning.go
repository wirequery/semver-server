package versioning

import "fmt"

type SemanticVersion struct {
	Major int
	Minor int
	Patch int
}

func New() SemanticVersion {
	return SemanticVersion{
		Major: 0,
		Minor: 1,
		Patch: 0,
	}
}

func IncrementMajor(s SemanticVersion) SemanticVersion {
	s.Major = s.Major + 1
	s.Minor = 0
	s.Patch = 0
	return s
}

func IncrementMinor(s SemanticVersion) SemanticVersion {
	s.Minor = s.Minor + 1
	s.Patch = 0
	return s
}

func IncrementPatch(s SemanticVersion) SemanticVersion {
	s.Patch = s.Patch + 1
	return s
}

func CreateString(s SemanticVersion) string {
	return fmt.Sprintf("%v.%v.%v", s.Major, s.Minor, s.Patch)
}
