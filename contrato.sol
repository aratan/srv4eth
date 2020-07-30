//pragma solidity ^0.4.0;

contract HolaMundo {
    string saludo = "Hola mundo";
    
    function getSaludo() constant returns(string) {
        return saludo;
    }
    
    function setSaludo(string nuevoSaludo) returns(string) {
        saludo = nuevoSaludo;
        return saludo;
    }
}
