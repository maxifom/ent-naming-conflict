package message


type MessageType string 

func (m MessageType) Values() []string {
	return []string{
		"type1",
		"type2",
	}
}