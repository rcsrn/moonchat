package server

type room struct {
	roomName string
	users *set
	invitedUsers *set
}

func GetRoomInstance(roomName string) *room {
	room := room{roomName, newSet(), newSet()}
	return &room
}

func (room *room) VerifyRoomMember(userName string) (bool) {
	return room.users.contains(userName)
}

func (room *room) GetMemberList() ([]string) {
	memberList := make([]string, len(room.users.elements))
	i := 0
	for userName, _ := range(room.users.elements) {
		memberList[i] = userName
		i++
	}
	return memberList
}

func (room *room) VerifyInvitedUser(userName string) (bool) {
	return room.invitedUsers.contains(userName) 
}  

func (room *room) AddUser(userName string) {
	room.users.add(userName)
}

func (room *room) AddInvitedUser(userName string) {
	room.invitedUsers.add(userName)
}

func (room *room) RemoveUser(userName string) {
	room.users.remove(userName)
}

func (room *room) RemoveInvitedUser(userName string) {
	room.invitedUsers.remove(userName)
}

func (room *room) IsEmpty() (bool) {
	return room.users.isEmpty()
}

