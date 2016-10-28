import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
  name: 'xmlparse',
  pure: true
})

export class XmlParsePipe  implements PipeTransform{
  transform(value){
    if(value < 10){
      value = '0'+value;
    }
    return value;
  }
}
