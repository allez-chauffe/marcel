var items = document.getElementsByTagName('df-schedule-item');
var talks = [];

for (var i in items ) {
    var talk = {};
    if(items[i].title != undefined){
        var description = items[i].getElementsByClassName('schedule-item__description');
        if (description[0] != undefined) {
            talk.description = description[0].getElementsByTagName('p')[0].innerHTML;
        }
        talk.time = items[i].getElementsByTagName('time')[0].title;
        talk.title = items[i].title;
        talk.name = items[i].name;
        talk.company = items[i].company;
        talk.abstract = items[i].abstract;
        talk.duration = items[i].duration;
        talk.type = items[i].type;
        if (talks.length < 12) {
            talk.room = "Amphi";
        } else if (talks.length < 22) {
            talk.room = "Talk room";
        } else {
            talk.room = "Codelab room";
        }
        talks.push(talk);
    }
    if(talks.length == 26) {
        break;
    }
}

console.log(JSON.stringify(talks));

// Copy and paste this script into firefox web console and then copy the result into a JSON file.