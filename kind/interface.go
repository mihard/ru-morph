package kind

const GEN_MALE = "m"
const GEN_FEMALE = "f"
const GEN_MIDDLE = "n"

const ENFORCED_ORIGINAL = "O"

type IWord interface {
	Gen() string
	Word() string
	Locative(forceGen string) (string, bool)
}
