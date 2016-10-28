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
    this.texte_1 = "Bonjour";
    this.texte_2 = "Aurélien";
  }

  ngOnInit(){
    console.log('Init Message');
  }

}
