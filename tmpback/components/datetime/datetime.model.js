class DateTimeModel {
  constructor (date) {
      this.minutes = this.format(date.getMinutes().toString());
      this.hour = this.format(date.getHours().toString());

      this.day = this.format(date.toLocaleString('fr', { 'day': '2-digit' }));
      this.year = date.toLocaleString('fr', { 'year': 'numeric' });
      this.month = date.toLocaleString('fr', { 'month': 'long' });
  }

  /**
   * Add a 0 as a prefix if the digit in parameter has only one digit
   */
  format (digit) {
      return ('0' + digit).slice(-2);
  }
}