package model

type Income struct {
    Id int
    Source string
    Amount float64
}

func GetIncome(id int) *Income {
    db := getdbconn()
    i := new(Income)
    db.QueryRow("SELECT id, source, amount FROM income WHERE id = $1", id).Scan(&i.Id, &i.Source, &i.Amount)
    return i
}


func AddIncome(i Income) int {
    db := getdbconn()
    var insertid int
    db.QueryRow("INSERT INTO income(source, amount) VALUES ($1, $2) returning id;", i.Source, i.Amount).Scan(&insertid)
    return insertid   
}

func UpdateIncome(i *Income) {
    db := getdbconn()
    db.QueryRow("UPDATE income SET source=$2, amount=$3 WHERE id=$1", i.Id, i.Source, i.Amount)
}

func DeleteIncome(i *Income) {
    db := getdbconn()
    db.QueryRow("DELETE FROM income WHERE id=$1", i.Id)
}

func GetAllIncomes() []*Income {
    db := getdbconn()
    rows, _ := db.Query("SELECT id, source, amount FROM income")
    incms := make([]*Income, 0)
    for rows.Next() {
        incm := new(Income)
        rows.Scan(&incm.Id, &incm.Source, &incm.Amount)
            incms = append(incms, incm)
    }

    return incms
}
