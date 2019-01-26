/* eslint-disable */

import { paper, Point, Path } from 'paper'

export class Board {
    constructor (canvasId) {
        this.canvasElem = document.getElementById(canvasId)

        // total canvas width and height
        this.cWidth = this.canvasElem.scrollWidth
        this.cHeight = this.canvasElem.scrollHeight

        // map edges
        this.mapLeftEdge = 0.25 * this.cWidth
        this.mapRightEdge = 0.75 * this.cWidth
        this.mapTopEdge = 0.1 * this.cHeight
        this.mapBotEdge = 0.7 * this.cHeight

        paper.setup(canvasId)
    }
}

Board.prototype.createCountry = function () {
    var myCircle = new Path.Circle(new Point(this.mapLeftEdge, 70), 5);
    myCircle.fillColor = 'black';
    paper.view.draw()
}

Board.prototype.setupCountryPoints = function () {
    var mapWidth = this.mapRightEdge - this.mapLeftEdge
    var dotSpacingX = mapWidth / 6
    var mapHeight = this.mapBotEdge - this.mapTopEdge
    var dotSpacingY = mapHeight / 6

    console.log(mapWidth)
    console.log(mapHeight)
    console.log(this.cWidth)
    console.log(this.cHeight)

    // top row
    var xPos = this.mapLeftEdge + dotSpacingX
    var yPos = this.mapTopEdge + dotSpacingY

    for (var i = 0; i < 3; i++) {
        var myCircle = new Path.Circle(new Point(xPos, yPos), 5);
        myCircle.fillColor = 'black'
        xPos += (2*dotSpacingX)
    }

    // bot row
    var xPos = this.mapLeftEdge + dotSpacingX
    var yPos = this.mapTopEdge + (dotSpacingY*3)
    for (var i = 0; i < 3; i++) {
        var myCircle = new Path.Circle(new Point(xPos, yPos), 5);
        myCircle.fillColor = 'black'
        xPos += (2*dotSpacingX)
    }   
}