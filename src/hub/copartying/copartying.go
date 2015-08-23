package copartying

import ara "github.com/diegogub/aranGO"

type PartyDoc struct {
  ara.Document // Must include arango Document in every struct you want to save id, key, rev after saving it
  // Id int
  // DefaultKey string `json:"id"`
  // required tag for strings.
  Name     string `required:"-"`
  // unique tag validate within collection users, if username is unique
  // Username string `unique:"users"`
  // // enum tag checks string value
  // Type     string `enum:"A,M,S"`
  // // Next release I will be implementing some other tags to validate int and arrays
  // Age      int
  // Likes    []string
}

var db_session ara.Session
var db_name string

const COL_PARTIES string = "parties"

func Init(database_session *ara.Session, database_name string) {
  db_session = *database_session
  db_name = database_name
}

func CreateParty(party PartyDoc) string {
  err := db_session.DB(db_name).Col(COL_PARTIES).Save(&party)
  if err != nil {
    panic(err)
  }
  return party.Key
}

func GetParty(key string) PartyDoc {
  var party PartyDoc
  err := db_session.DB(db_name).Col(COL_PARTIES).Get(key, &party)
  if err != nil {
    panic(err)
  }
  return party
}
