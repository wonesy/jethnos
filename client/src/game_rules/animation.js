/* eslint-disable */

import { paper, Point, Path, view, Tool, hitOptions, project, PointText } from 'paper'
import { Realm } from './realms'
import { GameState } from './gamestate'

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

// Raster.prototype.resize = function (width) {
//     this.width = width
//     this.height = width/2
// }

export class Board {
    constructor (canvasId, mapImgSrc) {
        this.m = new Measurements(canvasId)
        paper.setup(canvasId)
        tool = new Tool()
        tool.onMouseDown = onMouseDown
        tool.onMouseDrag = onMouseDrag
        tool.onMouseMove = onMouseMove

        this.background = new Raster({
            source: mapImgSrc,
            position: view.center,
            width: 1000,
            height: 500,
        });
        this.background.data = {
            'type': 'map',
            'draggable': false,
            'clickable': false
        }

        this.gamestate = new GameState()

        // Northwest
        drawRealmName(190, 260, this.gamestate.realms[0])

        // Northmiddle
        drawRealmName(240, 150, this.gamestate.realms[1])

        // Northeast
        drawRealmName(337, 310, this.gamestate.realms[2])

        // Southwest
        drawRealmName(160, 330, this.gamestate.realms[3])

        // South
        drawRealmName(110, 40, this.gamestate.realms[4])

        // Southeast
        drawRealmName(359, 295, this.gamestate.realms[5])
    }
}

function drawRealmName(angle, length, realm) {
    // create a point with an angle and length relative to the center
    var p = new Point({
        angle: angle,
        length: length
    })
    p.x += view.center.x
    p.y += view.center.y

    var text = new PointText(p)
    text.justification = 'center';
    text.fillColor = 'black';
    text.fontFamily = 'Comfortaa'
    text.fontSize = 18
    text.fontWeight = 'bold'
    text.content = realm.name;
    text.data = {
        'type': 'realm',
        'draggable': false,
        'clickable': true,
        'realm': realm
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
    document.getElementById('board').style.cursor = 'default'
	if (event.item && event.item.data.clickable) {
        document.getElementById('board').style.cursor = 'pointer'
    }
}

function onMouseDrag(event) {
	if (path) {
        path.position.x += event.delta.x
        path.position.y += event.delta.y 
        paper.view.draw()      
	}
}
