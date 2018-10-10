/*

Package controllers_test tests the autogenerated
scaffold outputs. Gock is used to mock the
go-github api calls.

*/
package controllers_test

import (
	"io/ioutil"

	"github.com/goadesign/goa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tinakurian/build-tool-detector/app/test"
	controllers "github.com/tinakurian/build-tool-detector/controllers"
	"gopkg.in/h2non/gock.v1"
)

var _ = Describe("BuildToolDetector", func() {
	Context("Internal Server Error", func() {
		It("Non-existent owner name -- 404 Repository Not Found", func() {
			defer gock.Off()

			bodyString, err := ioutil.ReadFile("../controllers/test/mock/fabric8_launcher_backend/not_found_repo_branch.json")
			Expect(err).Should(BeNil())

			gock.New("https://api.github.com").
				Get("/repos/fabric8-launcherz/launcher-backend/branches/master").
				Reply(404).
				BodyString(string(bodyString))

			service := goa.New("build-tool-detector")
			branch := "master"
			test.ShowBuildToolDetectorNotFound(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "https://github.com/fabric8-launcherz/launcher-backend", &branch)
		})

		It("Non-existent owner name -- 404 Owner Not Found", func() {
			defer gock.Off()

			bodyString, err := ioutil.ReadFile("../controllers/test/mock/fabric8_launcher_backend/not_found_repo_branch.json")
			Expect(err).Should(BeNil())

			gock.New("https://api.github.com").
				Get("/repos/fabric8-launcher/launcher-backendz/branches/master").
				Reply(404).
				BodyString(string(bodyString))

			service := goa.New("build-tool-detector")
			branch := "master"
			test.ShowBuildToolDetectorNotFound(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "https://github.com/fabric8-launcher/launcher-backendz", &branch)
		})

		It("Non-existent branch name -- 404 Branch Not Found", func() {
			defer gock.Off()

			bodyString, err := ioutil.ReadFile("../controllers/test/mock/fabric8_launcher_backend/not_found_branch.json")
			Expect(err).Should(BeNil())

			gock.New("https://api.github.com").
				Get("/repos/fabric8-launcher/launcher-backend/branches/masterz").
				Reply(404).
				BodyString(string(bodyString))

			service := goa.New("build-tool-detector")
			test.ShowBuildToolDetectorNotFound(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "https://github.com/fabric8-launcher/launcher-backend/tree/masterz", nil)
		})

		It("Invalid URL -- 400 Bad Request", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			test.ShowBuildToolDetectorBadRequest(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "fabric8-launcher/launcher-backend", &branch)
		})

		It("Unsupported Git Service -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			branch := "master"
			test.ShowBuildToolDetectorInternalServerError(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "http://gitlab.com/fabric8-launcher/launcher-backend", &branch)
		})

		It("Invalid URL and Branch -- 500 Internal Server Error", func() {
			service := goa.New("build-tool-detector")
			test.ShowBuildToolDetectorBadRequest(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "", nil)
		})
	})

	Context("Okay", func() {
		It("Recognize Unknown - Branch field populated", func() {
			defer gock.Off()

			bodyString, err := ioutil.ReadFile("../controllers/test/mock/fabric8_wit/ok_branch.json")
			Expect(err).Should(BeNil())

			gock.New("https://api.github.com").
				Get("/repos/fabric8-services/fabric8-wit/branches/master").
				Reply(200).
				BodyString(string(bodyString))

			bodyString, err = ioutil.ReadFile("../controllers/test/mock/fabric8_wit/not_found_contents.json")
			Expect(err).Should(BeNil())
			gock.New("https://api.github.com").
				Get("/repos/fabric8-services/fabric8-wit/contents/pom.xml").
				Reply(404).
				BodyString(string(bodyString))
			service := goa.New("build-tool-detector")
			branch := "master"
			_, buildTool := test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "https://github.com/fabric8-services/fabric8-wit", &branch)
			Expect(buildTool.BuildToolType).Should(Equal("unknown"), "buildTool should not be empty")
		})

		It("Recognize Unknown - Branch included in URL", func() {
			defer gock.Off()

			bodyString, err := ioutil.ReadFile("../controllers/test/mock/fabric8_wit/ok_branch.json")
			Expect(err).Should(BeNil())

			gock.New("https://api.github.com").
				Get("/repos/fabric8-services/fabric8-wit/branches/master").
				Reply(200).
				BodyString(string(bodyString))

			bodyString, err = ioutil.ReadFile("../controllers/test/mock/fabric8_wit/not_found_contents.json")
			Expect(err).Should(BeNil())
			gock.New("https://api.github.com").
				Get("/repos/fabric8-services/fabric8-wit/contents/pom.xml").
				Reply(404).
				BodyString(string(bodyString))
			service := goa.New("build-tool-detector")
			_, buildTool := test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "https://github.com/fabric8-services/fabric8-wit/tree/master", nil)
			Expect(buildTool.BuildToolType).Should(Equal("unknown"), "buildTool should not be empty")
		})

		It("Recognize Maven - Branch field populated", func() {
			defer gock.Off()

			bodyString, err := ioutil.ReadFile("../controllers/test/mock/fabric8_launcher_backend/ok_branch.json")
			Expect(err).Should(BeNil())

			gock.New("https://api.github.com").
				Get("/repos/fabric8-launcher/launcher-backend/branches/master").
				Reply(200).
				BodyString(string(bodyString))

			bodyString, err = ioutil.ReadFile("../controllers/test/mock/fabric8_launcher_backend/ok_contents.json")
			Expect(err).Should(BeNil())
			gock.New("https://api.github.com").
				Get("/repos/fabric8-launcher/launcher-backend/contents/pom.xml").
				Reply(200).
				BodyString(string(bodyString))
			service := goa.New("build-tool-detector")
			branch := "master"
			_, buildTool := test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "https://github.com/fabric8-launcher/launcher-backend", &branch)
			Expect(buildTool.BuildToolType).Should(Equal("maven"), "buildTool should not be empty")
		})

		It("Recognize Maven - Branch included in URL", func() {
			defer gock.Off()

			bodyString, err := ioutil.ReadFile("../controllers/test/mock/fabric8_launcher_backend/ok_branch.json")
			Expect(err).Should(BeNil())

			gock.New("https://api.github.com").
				Get("/repos/fabric8-launcher/launcher-backend/branches/master").
				Reply(200).
				BodyString(string(bodyString))

			bodyString, err = ioutil.ReadFile("../controllers/test/mock/fabric8_launcher_backend/ok_contents.json")
			Expect(err).Should(BeNil())
			gock.New("https://api.github.com").
				Get("/repos/fabric8-launcher/launcher-backend/contents/pom.xml").
				Reply(200).
				BodyString(string(bodyString))
			service := goa.New("build-tool-detector")
			_, buildTool := test.ShowBuildToolDetectorOK(GinkgoT(), nil, nil, controllers.NewBuildToolDetectorController(service, "test", "test"), "https://github.com/fabric8-launcher/launcher-backend/tree/master", nil)
			Expect(buildTool.BuildToolType).Should(Equal("maven"), "buildTool should not be empty")
		})
	})
})
