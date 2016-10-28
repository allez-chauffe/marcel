import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'addzero',
  pure: true
})

export class AddZeroPipe implements PipeTransform {

  transform(value){
    if(value < 10){
      value = '0'+value;
    }
    return value;
  }

}
