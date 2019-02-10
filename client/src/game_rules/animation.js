/* eslint-disable */

import { paper, Point, Path, view, Tool, hitOptions, project, PointText } from 'paper'
import { Realm } from './realms'
import { GameState } from './gamestate'

var globalTool
var globalMapSrc
var globalCanvasElem

/* globals */
window.addEventListener('resize', function () {
    paper.clear()
    drawMap()
})

function toRadians (angle) {
    return angle * (Math.PI / 180);
}

export class Board {
    constructor (canvasId, mapSrc) {
        globalCanvasElem = document.getElementById(canvasId)
        paper.setup(canvasId)
        globalTool = new Tool()
        globalTool.onMouseDown = onMouseDown
        globalTool.onMouseDrag = onMouseDrag
        globalTool.onMouseMove = onMouseMove

        globalMapSrc = mapSrc

        drawMap()
    }
}

function drawMap() {
    var mapWidth = globalCanvasElem.Width * 0.5
    var mapHeight = globalCanvasElem.scrollHeight * 0.4

    var raster = new paper.Raster({
        source: globalMapSrc,
        position: paper.view.center,
        size: new paper.Size(mapWidth, mapHeight - 150),
        data: {
            background: true
        }
    });

    writeName(40, 40, 'green', 'Noremac Leahhcim')
    writeName(35, 50, 'orange', 'Staffaronia')
    writeName(11.5, 910, 'purple', 'Deedubbya')
    writeName(52, 520, 'red', 'Shotsguyville')
    writeName(31, 690, 'blue', 'Steakton')
    writeName(19, 950, 'pink', 'Bigdansylvania')
}

function writeName(angle, pctLength, color, content) {
    var canvasHeight = globalCanvasElem.scrollHeight
    var hypotenuse = canvasHeight / Math.cos(toRadians(angle))
    var length = hypotenuse * pctLength / 100

    var p = new paper.Point({
        angle: angle,
        length: length
    })
    var pC = new paper.Path.Circle(p, 5)
    pC.fillColor = color
    pC.data = {type: 'realm'}

    var textP = new paper.Point(p)

    var text = new paper.PointText(textP)
    text.fillColor = 'black'
    text.content = content
    text.fontFamily = 'Bookman'
    text.fontWeight = 'bold'
    text.fontSize = 15
    text.data = {
        type: 'title'
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


var segment, path;
var movePath = false;
function onMouseDown(event) {
	segment = path = null;
	var hitResult = project.hitTest(event.point, hitOptions);
	if (!hitResult) {
        return
    }

    console.log(hitResult.type)
    console.log(hitResult.item)
    console.log(hitResult.item.data)

    // if (hitResult.item.data.background) {
    //     return
    // }
        
	if (event.modifiers.shift) {
		if (hitResult.type == 'segment') {
			hitResult.segment.remove();
		};
		return;
    }

    if (hitResult.item.data.draggable) {
        path = hitResult.item;
    }

    if (hitResult.type == 'segment') {
        segment = hitResult.segment;
    }
	movePath = hitResult.type == 'fill';
	if (movePath)
		project.activeLayer.addChild(hitResult.item);
}

function onMouseMove(event) {
    project.activeLayer.selected = true;
    document.body.style.cursor = "default";

	if (event.item && event.item.data.type === 'title') {
        document.body.style.cursor = "pointer";
    }
}

function onMouseDrag(event) {
	if (path) {
        path.position.x += event.delta.x
        path.position.y += event.delta.y 
        paper.view.draw()      
	}
}
