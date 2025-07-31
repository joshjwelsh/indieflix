import 'package:flutter/material.dart';
import 'dart:async';
import 'dart:convert';


void main() {
  runApp(
    MaterialApp(
      home: App(),
    )
  );
}

class App extends StatelessWidget{
  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Scaffold(
        appBar: AppBar(
          title: Text('hello world!'),
          centerTitle: true,
        ),
    );
  }
}

class Server{
  final String 
}