<?php

interface StackInterface
{
    public function push(mixed $val);
    public function pop(): mixed;
    public function isEmpty(): bool;
    public function top(): mixed;
    public function flush();
}
