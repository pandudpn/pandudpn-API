package sql

func (s *Db) Rollback() error {
	return nil
}

func (s *Db) Commit() error {
	return nil
}

func (s *Db) TxEnd(txFunc func() error) error {
	return nil
}

func (s *Tx) Rollback() error {
	return s.DB.Rollback()
}

func (s *Tx) Commit() error {
	return s.DB.Commit()
}

func (s *Tx) TxEnd(txFunc func() error) error {
	var err error

	defer func() {
		if p := recover(); p != nil {
			s.Rollback()
			panic(p)
		} else if err != nil {
			s.Rollback()
		} else {
			err = s.Commit()
		}
	}()

	err = txFunc()
	return err
}
