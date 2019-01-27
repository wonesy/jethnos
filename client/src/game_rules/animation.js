/* eslint-disable */

import { paper, Point, Path, view, Tool, hitOptions, project } from 'paper'

/* globals */
class Measurements {
    constructor (canvasId) {
        this.canvasElem = document.getElementById(canvasId)

        // total canvas width and height
        this.cWidth = this.canvasElem.scrollWidth
        this.cHeight = this.canvasElem.scrollHeight
        
        // map edges
        this.mapLeftEdge = 0.2 * this.cWidth
        this.mapRightEdge = 0.8 * this.cWidth
        this.mapTopEdge = 0.1 * this.cHeight
        this.mapBotEdge = 0.7 * this.cHeight
        
        // realm centers and offsets
        this.mapWidth = this.mapRightEdge - this.mapLeftEdge
        this.dotSpacingX = this.mapWidth / 6
        this.mapHeight = this.mapBotEdge - this.mapTopEdge
        this.dotSpacingY = this.mapHeight / 6
    }
}

var tool
var measurements

export class Board {
    constructor (canvasId) {
        this.m = new Measurements(canvasId)
        paper.setup(canvasId)
        tool = new Tool()
        tool.onMouseDown = onMouseDown
        tool.onMouseDrag = onMouseDrag
        tool.onMouseMove = onMouseMove
    }
}

export var RealmPositions = {
    NORTHWEST: 0,
    NORTHMID: 1,
    NORTHEAST: 2,
    SOUTHWEST: 3,
    SOUTHMID: 4,
    SOUTHEAST: 5
}
Object.freeze(RealmPositions)

export class Realm {
    constructor (name, position) {
        this.name = name
        this.position = position
    }
}

Board.prototype.createRealms = function () {
    var nwCenter = new Point(
        this.m.mapLeftEdge + this.m.dotSpacingX,
        this.m.mapTopEdge + this.m.dotSpacingY
    )
    var nwPath = drawRealm(nwCenter, this.m)

    var nmCenter = new Point(
        this.m.mapLeftEdge + (this.m.dotSpacingX*3),
        this.m.mapTopEdge + this.m.dotSpacingY
    )
    var nmPath = drawRealm(nmCenter, this.m)
}

function dotJitter(range) {
    // var neg = Math.round(Math.random() + 1)
    return Math.random() * range
}

function makePointsStaticY(points, staticY, startX, endX) {
    var pointsList = []
    var x, y
    var span = endX - startX

    for (var i = 0; i < points; i++) {
        // add a little jitter randomness to the height
        y = staticY + dotJitter(20)
        x = startX + (span * (i / points))
        pointsList.push(new Point(x,y))
    }
    return pointsList
}

function makePointsStaticX(points, staticX, startY, endY) {
    var pointsList = []
    var x, y
    var span = endY - startY

    for (var i = 0; i < points; i++) {
        // add a little jitter randomness to the height
        x = staticX + dotJitter(20)
        y = startY + (span * (i / points))
        pointsList.push(new Point(x,y))
    }
    return pointsList
}

var drawRealm = function (center, meas) {
    console.log('realm')
    var path = new Path()
    path.closed = true
    path.strokeColor = 'red'
    path.strokeWidth = 5
    path.fillColor = 'pink'

    var center = new Point(
        meas.mapLeftEdge + meas.dotSpacingX,
        meas.mapTopEdge + meas.dotSpacingY
    )

    //
    // left edge, bot to top
    //
    var leftStart = center.y + meas.dotSpacingY
    var leftEnd = center.y - meas.dotSpacingY
    var leftPts = makePointsStaticX(10, center.x-meas.dotSpacingX, leftStart, leftEnd)
    for (var i = 0; i < leftPts.length; i++) {
        path.add(leftPts[i])
    }

    //
    // top edge, left to right
    //
    var topStart = center.x - meas.dotSpacingX
    var topEnd = center.x + meas.dotSpacingX
    var topPts = makePointsStaticY(10, center.y-meas.dotSpacingY, topStart, topEnd)
    for (var i = 0; i < topPts.length; i++) {
        path.add(topPts[i])
    }

    //
    // right edge, top to bot
    //
    var rightStart = center.y - meas.dotSpacingY
    var rightEnd = center.y + meas.dotSpacingY
    var rightPts = makePointsStaticX(10, center.x+meas.dotSpacingX, rightStart, rightEnd)
    for (var i = 0; i < rightPts.length; i++) {
        path.add(rightPts[i])
    }

    //
    // bot edge, right to left
    //
    var botStart = center.x + meas.dotSpacingX
    var botEnd = center.x - meas.dotSpacingX
    var botPts = makePointsStaticY(10, center.y+meas.dotSpacingY, botStart, botEnd)
    for (var i = 0; i < botPts.length; i++) {
        path.add(botPts[i])
    }

    path.smooth()
    
    return path
}

Board.prototype.randomShape = function () {
    var myPath = new Path();
    myPath.closed = true
    myPath.strokeColor = 'black';
    myPath.fillColor = 'orange'

    myPath.add(new Point());
    myPath.add(new Point(Math.random() * view.size.height));
    myPath.add(new Point(100, Math.random() * view.size.height * 0.9));
    myPath.data.remember = "yo"
}

Board.prototype.randomRegion = function () {
    var center = new Point(500,500)
    var myPath = new Path();
    myPath.closed = true
    myPath.strokeColor = 'black';
    myPath.fillColor = 'green'

    // left edge
    var x = center.x - 100
    var y = center.y - 25
    var p = new Point(x,y)
    myPath.add(p)

    var x = center.x - 100
    var y = center.y
    p = new Point(x,y)
    myPath.add(p)

    var x = center.x - 100
    var y = center.y + 25
    p = new Point(x,y)
    myPath.add(p)

    // right edge
    var x = center.x + 100
    var y = center.y + 25
    p = new Point(x,y)
    myPath.add(p)

    var x = center.x + 100
    var y = center.y
    p = new Point(x,y)
    myPath.add(p)

    var x = center.x + 100
    var y = center.y - 25
    p = new Point(x,y)
    myPath.add(p)

    myPath.data.remember = "yo"
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

    console.log(view.size)
}

var segment, path;
var movePath = false;
function onMouseDown(event) {
	segment = path = null;
	var hitResult = project.hitTest(event.point, hitOptions);
	if (!hitResult) {
        return
    }

    console.log(hitResult)
    console.log(hitResult.item)
    console.log(hitResult.item.data)
        
	if (event.modifiers.shift) {
		if (hitResult.type == 'segment') {
			hitResult.segment.remove();
		};
		return;
	}
    path = hitResult.item;
    if (hitResult.type == 'segment') {
        segment = hitResult.segment;
    }
	movePath = hitResult.type == 'fill';
	if (movePath)
		project.activeLayer.addChild(hitResult.item);
}

function onMouseMove(event) {
    project.activeLayer.selected = false;
	if (event.item) {
        event.item.selected = true;
    }
}

function onMouseDrag(event) {
	if (path) {
        path.position.x += event.delta.x
        path.position.y += event.delta.y 
        paper.view.draw()      
	}
}
