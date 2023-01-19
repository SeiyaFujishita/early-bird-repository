package wakeup

import (

"fmt"
"backend/domain/model"
)
func (r *WakeUpRepository) CreateWakeUp(WakeUp *model.WakeUp) (err error) {
	fmt.Println("------------------------")
	fmt.Printf("%+v\n",&WakeUp)
	fmt.Println("------------------------")
	err = r.DB.Create(&WakeUp).Error
	fmt.Println("------------------------")
	fmt.Printf("%+v\n",err)
	fmt.Println("------------------------")
	if err != nil {
		return err
	}

	return nil
}
