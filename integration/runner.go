package integration

var fixture *Fixture
var Runner = CreateForEach(setUp, tearDown)

func CreateForEach(setUp func(), tearDown func()) func(func()) {
	return func(testFunc func()) {
		setUp()
		testFunc()
		tearDown()
	}
}

func setUp() {
	var testFixture *FixtureFactory = &FixtureFactory{}
	fixture = testFixture.Create()
}

func tearDown() {
	Cleanup(fixture)
}
