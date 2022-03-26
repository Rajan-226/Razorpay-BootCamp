<?php
include_once "ToDoList.php";
include_once "User.php";

class StandardUser extends User
{
    //maximum number of lists that can be added
    private $maxListCount;

    function __construct($name = "")
    {
        parent::__construct($name);
        $this->maxListCount = 10;
    }

    public function createNewList(): bool
    {
        //if after adding,limit will get exceed
        if (count($this->lists) == $this->maxListCount) {
            print("List limit exceeded, cannot add more todo lists for user {$this->name}\n\n");
            return false;
        }

        array_push($this->lists, new ToDoList());

        $listNumber = count($this->lists);
        print "List {$listNumber} added successfully for Standard User {$this->name}\n\n";
        return true;
    }
}