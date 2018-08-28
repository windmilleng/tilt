//+build !skipcontainertests

package build

import (
	"fmt"
	"strings"
	"testing"

	"github.com/windmilleng/tilt/internal/model"
)

func BenchmarkBuildTenSteps(b *testing.B) {
	run := func() {
		f := newDockerBuildFixture(b)
		defer f.teardown()

		steps := []model.Cmd{}
		for i := 0; i < 10; i++ {
			steps = append(steps, model.ToShellCmd(fmt.Sprintf("echo -n %d > hi", i)))
		}

		ref, err := f.b.BuildDockerFromScratch(f.ctx, f.getNameFromTest(), simpleDockerfile, []model.Mount{}, steps, model.Cmd{})
		if err != nil {
			b.Fatal(err)
		}

		expected := []expectedFile{
			expectedFile{path: "hi", contents: "9"},
		}
		f.assertFilesInImage(ref, expected)
	}
	for i := 0; i < b.N; i++ {
		run()
	}
}

func BenchmarkBuildTenStepsInOne(b *testing.B) {
	run := func() {
		f := newDockerBuildFixture(b)
		defer f.teardown()

		allCmds := make([]string, 10)
		for i := 0; i < 10; i++ {
			allCmds[i] = fmt.Sprintf("echo -n %d > hi", i)
		}

		oneCmd := strings.Join(allCmds, " && ")

		steps := []model.Cmd{model.ToShellCmd(oneCmd)}
		ref, err := f.b.BuildDockerFromScratch(f.ctx, f.getNameFromTest(), simpleDockerfile, []model.Mount{}, steps, model.Cmd{})
		if err != nil {
			b.Fatal(err)
		}

		expected := []expectedFile{
			expectedFile{path: "hi", contents: "9"},
		}
		f.assertFilesInImage(ref, expected)
	}
	for i := 0; i < b.N; i++ {
		run()
	}
}

func BenchmarkIterativeBuildTenTimes(b *testing.B) {
	f := newDockerBuildFixture(b)
	defer f.teardown()
	steps := []model.Cmd{model.ToShellCmd("echo 1 >> hi")}
	ref, err := f.b.BuildDockerFromScratch(f.ctx, f.getNameFromTest(), simpleDockerfile, []model.Mount{}, steps, model.Cmd{})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			ref, err = f.b.BuildDockerFromExisting(f.ctx, ref, nil, steps)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
