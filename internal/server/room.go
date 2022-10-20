package server

type Room struct {
	roomName string
	members *set
	invitedUsers *set
}

func GetRoomInstance(roomName string) *Room {
	room := Room{roomName, newSet(), newSet()}
	return &room
}

func (room *Room) GetMembers() map[string]struct{} {
	return room.members.getElements()
}

func (room *Room) GetInvitedUsers() map[string]struct{} {
	return room.invitedUsers.getElements()
}

func (room *Room) GetRoomName() string {
	return room.roomName
}

func (room *Room) VerifyRoomMember(userName string) (bool) {
	return room.members.contains(userName)
}

func (room *Room) GetMemberList() ([]string) {
	memberList := make([]string, len(room.members.elements))
	i := 0
	for userName, _ := range(room.members.elements) {
		memberList[i] = userName
		i++
	}
	return memberList
}

func (room *Room) VerifyInvitedUser(userName string) (bool) {
	return room.invitedUsers.contains(userName) 
}  

func (room *Room) AddUser(userName string) {
	room.members.add(userName)
}

func (room *Room) AddInvitedUser(userName string) {
	room.invitedUsers.add(userName)
}

func (room *Room) RemoveUser(userName string) {
	room.members.remove(userName)
}

func (room *Room) RemoveInvitedUser(userName string) {
	room.invitedUsers.remove(userName)
}

func (room *Room) IsEmpty() (bool) {
	return room.members.isEmpty()
}

