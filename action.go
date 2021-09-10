package adaptivecards

const (
	TypeActionOpenUrl = "Action.OpenUrl"
)

type Action interface {
	Action() string
}
