package server

type room struct {
	roomName string
	users *set
	invitedUsers *set
}

func getRoomInstance(roomName string) *room {
	room := room{roomName, newSet(), newSet()}
	return &room
}

func (room *room) verifyRoomMember(userName string) (bool) {
	return room.users.contains(userName)
}

func (room *room) getMemberList() ([]string) {
	memberList := make([]string, len(room.users.elements))
	i := 0
	for userName, _ := range(room.users.elements) {
		memberList[i] = userName
		i++
	}
	return memberList
}

func (room *room) verifyInvitedUser(userName string) (bool) {
	return room.invitedUsers.contains(userName) 
}  

func (room *room) addUser(userName string) {
	room.users.add(userName)
}

func (room *room) addInvitedUser(userName string) {
	room.invitedUsers.add(userName)
}

func (room *room) removeUser(userName string) {
	room.users.remove(userName)
}

func (room *room) removeInvitedUser(userName string) {
	room.invitedUsers.remove(userName)
}

func (room *room) isEmpty() (bool) {
	return room.users.isEmpty()
}

