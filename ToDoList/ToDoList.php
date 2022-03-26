<?php
include_once "Task.php";

class ToDoList
{
    private $list;
    private $listName;

    function __construct($listName = "")
    {
        $this->listName = $listName;
        $this->list = array();
    }

    function addTask($description, $criticalValue)
    {
        array_push($this->list, new Task($description, $criticalValue));
    }

    function sortByCriticalValue()
    {
        //For array of objects:
        //Sort function will sort the array according to the first field value in the class
        //Beware of: First field value can change if you will inherit some class

        //There is no comparable interface in php, which we can implement as per our use case
        sort($this->list);
    }
}