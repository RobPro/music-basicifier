let num = 0
let isFirstTime = false
let tempo = 0
input.buttonA.onEvent(ButtonEvent.Click, function () {
    pause(3000)
    num = 0
    network.infraredSendNumber(num)
    if (isFirstTime) {
        pause(2000)
        num = 1
        music.setVolume(152)
        for (let i = 0; i < 3; i++) {
            music.playTone(262, music.beat(BeatFraction.Double))
            music.playTone(392, music.beat(BeatFraction.Breve))
            pause(2000)
        }
        network.infraredSendNumber(num)
        pause(4000)
        tempo = 290
        playFirstBit()
        playFirstBit()
    }
    isFirstTime = true
})
function playFirstBit () {
    music.playMelody("- - - - C C C C ", tempo)
    music.playMelody("F F G G G A A A ", tempo)
    music.playMelody("G G G G E E E E ", tempo)
    music.playMelody("- - - - C C C C ", tempo)
    music.playMelody("F F G G G A A A ", tempo)
    music.playMelody("G G G G G G G G ", tempo)
    music.playMelody("- - - - C C C C ", tempo)
    music.playMelody("F F G G G A A A ", tempo)
    music.playMelody("G G G G E E E E ", tempo)
    music.playMelody("- - - - E E E E ", tempo)
    music.playMelody("F F F E E C C - ", tempo)
    music.playMelody("C C C C C C C C ", tempo)
}
network.onInfraredReceivedNumber(function (num) {
    music.setVolume(255)
    while (true) {
        num = 213
        music.playTone(131, num)
    }
})
input.onGesture(Gesture.Shake, function () {
    num += 1
})
input.buttonB.onEvent(ButtonEvent.Click, function () {
    num += -1
})
