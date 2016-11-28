import { Component, OnInit } from '@angular/core';
import { ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'message',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './message.component.html',
  styleUrls: ['./message.component.scss']
})

export class MessageComponent implements OnInit {

  public texte_1: string;
  public texte_2: string;

  constructor() {
    this.texte_1 = "Bienvenue";
    this.texte_2 = "à Zenika Lille";
  }

  ngOnInit(){
    console.log('Init Message');
  }

}
