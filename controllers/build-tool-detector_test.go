/*

Package controllers_test is autogenerated
and containing scaffold outputs.

*/
package controllers_test

import (
	"github.com/tinakurian/build-tool-detector/app/test"
	controllers "github.com/tinakurian/build-tool-detector/controllers"
	git "github.com/tinakurian/build-tool-detector/controllers/git"
	"github.com/tinakurian/build-tool-detector/controllers/git/buildtype"
	"github.com/tinakurian/build-tool-detector/controllers/git/github"
	"github.com/goadesign/goa"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("BuildToolDetector", func() {
	Context("Internal Server Error", func() {
		It("Non-existent owner name -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcherz/launcher-backend", nil)
		})

		It("Non-existent repository name -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backendz", nil)
		})

		It("Non-existent branch name -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			branch := "masterz"
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backend", &branch)
		})

		It("Build tool type expected to be Unknown -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-services/fabric8-wit", &branch)
		})
	})

	Context("Okay", func() {
		It("Okay response -- 200 Okay", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backend", &branch)
		})

		It("Non-nil response -- 200 Okay", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			_, buildTool := test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backend", &branch)
			gomega.Expect(buildTool).ShouldNot(gomega.BeNil(), "buildTool should not be empty")
		})

		It("Build tool type to be Maven -- 200 Okay", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			_, buildTool := test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service), "https://github.com/fabric8-launcher/launcher-backend", &branch)
			gomega.Expect(buildTool.BuildToolType).Should(gomega.BeEquivalentTo("maven"), "build type should be equivalent to 'maven'")
		})
	})

	Context("build-tool-detector/controllers/git", func() {
		It("url_parser - GetServiceType() service github", func() {
			_, gitService := git.GetServiceType("https://github.com/owner/repo/tree/branch")

			gomega.Expect(gitService.Service).Should(gomega.BeEquivalentTo("github"), "git service should be equivalent to 'github'")
			gomega.Expect(gitService.Segments[1]).Should(gomega.BeEquivalentTo("owner"), "1st segment from url should be 'owner'")
			gomega.Expect(gitService.Segments[2]).Should(gomega.BeEquivalentTo("repo"), "2nd segment from url should be 'repo'")
			gomega.Expect(gitService.Segments[3]).Should(gomega.BeEquivalentTo("tree"), "third segment from url should be 'tree'")
			gomega.Expect(gitService.Segments[4]).Should(gomega.BeEquivalentTo("branch"), "fourth segment from url should be 'branch'")
		})

		It("url_parser - GetServiceType() service unknown", func() {
			_, gitService := git.GetServiceType("https://test.com/test/test/tree/master")
			gomega.Expect(gitService.Service).Should(gomega.BeEquivalentTo("unknown"), "git service should be equivalent to 'unknown'")
		})

		It("url_parser - GetServiceType() bad request with no owner or repository", func() {
			err, _ := git.GetServiceType("https://test.com")
			gomega.Expect(err.StatusCode).Should(gomega.BeEquivalentTo(http.StatusBadRequest), "git service should be equivalent to 'unknown'")
		})

		It("url_parser - GetServiceType() bad request with no schema or host", func() {
			err, _ := git.GetServiceType("test/test/test")
			gomega.Expect(err.StatusCode).Should(gomega.BeEquivalentTo(http.StatusBadRequest), "git service should be equivalent to 'unknown'")
		})

		It("url_parser - GetTyGetServiceTypepe() bad request whitespace url", func() {
			err, _ := git.GetServiceType(" ")
			gomega.Expect(err.StatusCode).Should(gomega.BeEquivalentTo(http.StatusBadRequest), "git service should be equivalent to 'unknown'")
		})
	})

	Context("build-tool-detector/controllers/git/buildtype", func() {
		It("build_tool_type - Maven()", func() {
			buildToolType := buildtype.Maven()
			gomega.Expect(buildToolType.BuildToolType).Should(gomega.BeEquivalentTo("maven"), "build type should be equivalent to 'maven'")
		})

		It("build_tool_type - Unknown()", func() {
			buildToolType := buildtype.Unknown()
			gomega.Expect(buildToolType.BuildToolType).Should(gomega.BeEquivalentTo("unknown"), "build type should be equivalent to 'maven'")
		})
	})

	Context("build-tool-detector/controllers/git/github", func() {
		It("attributes - GetAttributes() use branch from url", func() {
			_, gitService := git.GetServiceType("https://github.com/owner/repo/tree/branch")
			_, attributes := github.GetAttributes(gitService.Segments, nil)
			gomega.Expect(attributes.Owner).Should(gomega.BeEquivalentTo("owner"), "Owner field should be equivalent to 'owner'")
			gomega.Expect(attributes.Repository).Should(gomega.BeEquivalentTo("repo"), "Repository field should be equivalent to 'repo'")
			gomega.Expect(attributes.Branch).Should(gomega.BeEquivalentTo("branch"), "Branch field should be equivalent to 'branch'")

		})

		It("attributes - GetAttributes() use default branch master", func() {
			_, gitService := git.GetServiceType("https://github.com/owner/repo")
			_, attributes := github.GetAttributes(gitService.Segments, nil)
			gomega.Expect(attributes.Owner).Should(gomega.BeEquivalentTo("owner"), "Owner field should be equivalent to 'owner'")
			gomega.Expect(attributes.Repository).Should(gomega.BeEquivalentTo("repo"), "Repository field should be equivalent to 'repo'")
			gomega.Expect(attributes.Branch).Should(gomega.BeEquivalentTo("master"), "Branch field should be equivalent to 'master'")
		})

		It("attributes - GetAttributes() not enough segments extracted from path", func() {
			segments := []string{"test1", "test2"}
			err, attributes := github.GetAttributes(segments, nil)
			gomega.Expect(err.StatusCode).Should(gomega.BeEquivalentTo(http.StatusBadRequest), "invalid url segments should be equivalent to a Bad Request")
			gomega.Expect(attributes).Should(gomega.Equal(github.Attributes{}), "invalid url segments cause empty attributes")
		})
	})
})
