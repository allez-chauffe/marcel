import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'firstletterupper',
  pure: true
})

export class FirstletterupperPipe implements PipeTransform {
  transform(value:string):any {
    return value.charAt(0).toUpperCase() + value.slice(1);
  }
}
