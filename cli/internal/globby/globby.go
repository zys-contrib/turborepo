package globby

import (
	"context"

	"github.com/saracen/matcher"
)

// IgnoreFunc checks if a path ought to be ignored
type IgnoreFunc func(path string) bool

func Globby(baseDir string, patterns []string) ([]string, error) {
	globs := []matcher.Matcher{}
	for _, pattern := range patterns {
		globs = append(globs, matcher.New(pattern))
	}
	match := matcher.Multi(globs...)

	matches, err := matcher.Glob(context.Background(), ".", match)
	if err != nil {
		return nil, err
	}
	file := make([]string, len(matches))
	for i := range matches {
		file = append(file, i)
	}
	return file, nil
}

// var filesToBeCached = make(util.Set)
// for _, output := range patterns {
// 	results, err := doublestar.Glob(filepath.Join(baseDir, strings.TrimPrefix(output, "!")))
// 	if err != nil {
// 		return nil, fmt.Errorf("globbing %s: %w", output, err)
// 	}
// 	fmt.Println(results)
// 	for _, result := range results {
// 		if strings.HasPrefix(output, "!") {
// 			filesToBeCached.Delete(result)
// 		} else {
// 			filesToBeCached.Add(result)
// 		}
// 	}
// }

// return filesToBeCached.UnsafeListOfStrings(), nil
// }
