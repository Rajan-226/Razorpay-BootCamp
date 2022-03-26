<?php
include_once "ToDoList.php";
include_once "Task.php";

abstract class User
{
    //name of user
    protected $name;

    //array of toDolist
    protected $lists;

    function __construct($name)
    {
        $this->name = $name;
        $this->lists = array();
    }

    function setName(string $name): void
    {
        $this->name = $name;
    }

    function getName(): string
    {
        return $this->name;
    }

    function addTask($listNumber, $description, $importance = 1)
    {
        $this->lists[$listNumber]->addTask($description, $importance);
    }

    function displayALlLists()
    {
        foreach ($this->lists as $list) {
            print_r($list);
        }
        print "\n";
    }

    private function isPresent($id)
    {
        if ($id >= count($this->lists)) {
            print "List does not exists\n";
            return false;
        }
        return true;
    }

    function displaySingleList($id)
    {
        if ($this->isPresent($id)) {
            return $this->lists[$id];
        }
    }

    function sortByImportance($id)
    {
        if ($this->isPresent($id)) {
            return $this->lists[$id]->sortByCriticalValue();
        }
    }

    abstract protected function createNewList(): bool;
}