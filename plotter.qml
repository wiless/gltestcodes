import QtQuick 2.0
import GoExtensions 1.0
import QtQuick.Controls 1.1

ApplicationWindow { 
 id:app
 visible: true
 visibility:"FullScreen"
 title:"Painting Test"
 width:1000
 height:400
    GoPlot {
     id:plot1
     x: 0; y: 0;
     width: 500 
     height: 150
     yMax:50
  anchors.topMargin: 0
   anchors.leftMargin: 0
    update:true
     
     objectName:"myobject"
     npoints:100

     signal mousePressed()
     MouseArea {
      id:plotmousearea      
      anchors.fill: parent
      }    
    Component.onCompleted: {
     console.log("Printing from QML npoints  := ", plot1.npoints) 
     console.log("Printing from QML npoints  := ", plot1.yMax) 
     plotmousearea.clicked.connect(mousePressed)
   }

    // SequentialAnimation on x {
    //     loops: Animation.Infinite
    //     NumberAnimation { from: 0; to: 320; duration: 3000; easing.type: Easing.OutQuad}
    //     NumberAnimation { from: 320; to: 0; duration: 1000; easing.type: Easing.OutQuad}
    // }
  
}
}