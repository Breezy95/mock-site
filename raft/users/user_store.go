package users

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
)

//implementing Raft FSM
type User struct {
	username string
	email string
	pass string
}

type UserStore struct{
	mu sync.Mutex
	logger hclog.Logger
	usermap map[string]User
}

func newUserStore(logger hclog.Logger) *UserStore {
	newUserStore := UserStore {logger: logger, usermap: map[string]User{},
}
	return &newUserStore
}

func (st *UserStore) Persist(sink raft.SnapshotSink) error{
	b, err := json.Marshal(st.usermap)
	if err != nil {
		sink.Cancel()
		return fmt.Errorf("error in persisting keystore")
	}
	_,err1 := sink.Write(b)
	if err1 != nil {
		sink.Cancel()
		return fmt.Errorf("error in writing to sink")
		}
	err2 := sink.Close()
	if err2 != nil{
		return fmt.Errorf("error in closing sink")
	}
	return nil
	}

func (st *UserStore) Snapshot() (raft.FSMSnapshot, error) {
	st.mu.Lock()
	defer st.mu.Unlock()
	return st.clone(),nil
}

func (us *UserStore) Release() {  }


func (us_inst *UserStore) clone() *UserStore {
	us_inst.mu.Lock()
	defer us_inst.mu.Unlock()


	newUserStore := UserStore{ logger: us_inst.logger}
	newUserMap := make(map[string]User)
	for username, val := range us_inst.usermap{
		newUserMap[username] = val
	} 
	newUserStore.usermap = newUserMap

	return &newUserStore
}

func (st *UserStore) Restore(snapshot io.ReadCloser) error {

	return nil
}

func (us *UserStore) Apply(log *raft.Log) interface{} {

	return nil
}

func (us *UserStore) GetUser(username string) (User, error){
	us.mu.Lock()
	defer us.mu.Unlock()

	user,succ := us.usermap[username]
	if succ != true{
		nil_user := User{}
		return nil_user, fmt.Errorf("could not retrieve user with the key" + username)
	}
	return user,nil
}

func (us *UserStore) SetUser(key string, user User) {
	us.mu.Lock()
	defer us.mu.Unlock()
	us.usermap[key] = user
}

func (us *UserStore) DeleteUser(username string){
	us.mu.Lock()
	defer us.mu.Unlock()
	delete(us.usermap, username)	
}





