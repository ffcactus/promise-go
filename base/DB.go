package base

// DBInterface is the interface that DB should have.
type DBInterface interface {
	Post(ModelInterface) (bool, ModelInterface, bool, error)
}

// DB is the DB implementation in Promise project.
type DB struct {

}
