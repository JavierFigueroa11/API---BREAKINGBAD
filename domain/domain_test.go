package domain

import "testing"

func TestValidCharID(t *testing.T) {
	c := CharactersBB {
		CharID: 1,
	}
	v := c.ValidCharID()
	if v != true {
		t.Error("Error in Data CharID...")
		return
	}
}

func TestInvalidCharID(t *testing.T) {
	c := CharactersBB{
		CharID: -1,
	}
	v := c.ValidCharID()
	if v == true {
		t.Fail()
		return
	}
}

func TestValidName(t *testing.T) {
	c := CharactersBB{
		Name: "Name",
	}
	v := c.HasName()
	if v != true {
		t.Error("Error in Data Name...")
		return
	}
}

func TestInvalidName(t *testing.T) {
	c := CharactersBB{
		Name: "",
	}
	v := c.HasName()
	if v == true {
		t.Fail()
		return
	}
}

func TestValidBirthday (t *testing.T) {
	c := CharactersBB{
		Birthday:"DD/MM/YYYY",
	}
	v := c.HasBirthday()
	if v!= true {
		t.Error("Error in Data Birthday...")
		return
	}
}

func TestInvalidBirthday(t *testing.T){
	c := CharactersBB{
		Birthday: "",
	}
	v := c.HasBirthday()
	if v== true {
		t.Fail()
		return
	}
}

func TestValidIMG (t *testing.T) {
	c:= CharactersBB{
		Img: "https/...",
	}
	v:= c.HasIMG()
	if v != true{
		t.Error("Error in IMG")
		return
	}
}

func TestInvalidIMG(t *testing.T){
	c:= CharactersBB{
		Img:"",
	}
	v:= c.HasIMG()
	if v == true {
		t.Fail()
		return
	}
}

func TestValidStatus(t *testing.T) {
	c:= CharactersBB{
		Status: "Status",
	}
	v:= c.HasStatus()
	if v != true {
		t.Error("Error in Data Status")
		return
	}
}

func TestInvalidStatus(t *testing.T) {
	c:= CharactersBB{
		Status: "",
	}
	v:= c.HasStatus()
	if v == true {
		t.Fail()
		return
	}
}

func TestValidNickname(t *testing.T) {
	c := CharactersBB{
		Nickname:"Nickname",
	}
	v:= c.HasNickname()
	if v!= true {
		t.Error("Error in Nickname")
		return
	}
}

func TestInvalidNickname(t *testing.T) {
	c := CharactersBB{
		Nickname:"",
	}
	v:= c.HasNickname()
	if v== true {
		t.Fail()
		return
	}
}

func TestValidPortrayed(t *testing.T) {
	c := CharactersBB{
		Portrayed:"Portrayed",
	}
	v:= c.HasPortrayed()
	if v!= true {
		t.Error("Error in Data Portrayed")
		return
	}
}

func TestInvalidPortrayed(t *testing.T) {
	c := CharactersBB{
		Portrayed:"",
	}
	v:= c.HasPortrayed()
	if v== true {
		t.Fail()
		return
	}
}

func TestValidCategory(t *testing.T) {
	c := CharactersBB{
		Category:"Category",
	}
	v:= c.HasCategory()
	if v!= true {
		t.Error("Error in Data Category")
		return
	}
}

func TestInvalidCategory(t *testing.T) {
	c := CharactersBB{
		Category:"",
	}
	v:= c.HasCategory()
	if v== true {
		t.Fail()
		return
	}
}

