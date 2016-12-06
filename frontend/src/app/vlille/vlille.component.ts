import {Component, OnInit,ElementRef} from "@angular/core";
import {ViewEncapsulation} from '@angular/core';
import {XmlParsePipe} from '../pipes/xmlparse.pipe';
import {VlilleService} from './vlille.service';
import {Observable} from 'rxjs/Rx';

@Component({
  selector: 'vlille',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './vlille.component.html',
  styleUrls: ['./vlille.component.scss']
})

export class VlilleComponent implements OnInit{

  public stationids : any[] = [];
  public stations : any[] = [];
  public interval : number;

  constructor(private vlilleservice: VlilleService) {
    this.stationids = [
      {name:'Rihour',id:10},
      {name:'Cormontaigne',id:36},
      {name:'Mairie de Lille',id:64},
      {name:'Gare Lille Flandres',id:25}
    ];
    this.interval = 30000; // secondes  
  }


  ngOnInit(){
    console.log('ngOnInit');

    this.getBornesData();

    setInterval(() => {
      this.getBornesData();
    }, this.interval);
  }

  getBornesData(){
    this.stations = [];
    Observable
      .from(this.stationids)
      .flatMap((i) => this.vlilleservice.getBorneData(i))
      .subscribe((o) => {
        this.stations.push(o);
      });
  }

}
