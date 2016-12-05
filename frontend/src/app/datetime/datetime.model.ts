export class DateTime {
    public day: string;
    public year: string;
    public month: string;
    public minutes: string;
    public hour: string;

    constructor(date: Date) {
        this.minutes = this.format(date.getMinutes().toString());
        this.hour = this.format(date.getHours().toString());

        this.day = this.format(date.toLocaleString('fr', { 'day': '2-digit' }));
        this.year = date.toLocaleString('fr', { 'year': 'numeric' });
        this.month = date.toLocaleString('fr', { 'month': 'long' });
    }

    /**
     * Add a 0 as a prefix if the digit in parameter has only one digit
     */
    private format(digit: string) {
        return ('0' + digit).slice(-2);
    }
}