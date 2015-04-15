import QtQuick 2.0
import GoExtensions 1.0
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.1
ApplicationWindow { 
 id:app
 visible: true
 visibility:"FullScreen"
 title:"Painting Test"
 width:1000
 height:400
 color: "transparent"
 ColumnLayout{
  spacing: 0
  anchors.fill: parent
  GoPlot {
   id:plot1
   name:"plot1"
   color:"#F0000000"
   // x: 10; y: 0;
   // anchors.fill : parent

   Layout.fillWidth:true
   Layout.fillHeight:true
   Layout.alignment:Qt.AlignTop

   // width: parent.width
   height: 150
   yMax:50
   // minimumHeight:100
   
   anchors.topMargin: 20
   anchors.leftMargin: 10
   anchors.bottomMargin: 50
   anchors.rightMargin: 10

   objectName:"myobject"
   npoints:100

   signal mousePressed()
   MouseArea {
    id:plotmousearea      
    anchors.fill: parent
    onWheel: {

      if (wheel.modifiers & Qt.ControlModifier) {
        console.log("The data is ",wheel.angleDelta.y)
         myapp.scroll(wheel.angleDelta.y)
        
      }

    }
  }    

  Component.onCompleted: {
   console.log("Printing from QML npoints  := ", plot1.npoints) 
   console.log("Printing from QML npoints  := ", plot1.yMax) 
   console.log("Printing from QML My Dimensions := ", plot1.height, plot1.width) 
   plotmousearea.clicked.connect(mousePressed)

 }
 // SequentialAnimation on x {
   //      loops: Animation.Infinite
   //      NumberAnimation { from: 0; to: 320; duration: 3000; easing.type: Easing.OutQuad}
   //      NumberAnimation { from: 320; to: 0; duration: 1000; easing.type: Easing.OutQuad}
   //  }  
 }

 Rectangle{  
  color:"#800000FF"
  height:100
  // Layout.alignment: Qt.AlignHCenter
  Layout.fillWidth: true
  // anchors.topMargin: 55
  RowLayout{

    spacing:10
    Text    {
      width:parent.width
      font.italic: true   
      baselineOffset:15
      horizontalAlignment: Text.AlignHCenter
      verticalAlignment: Text.AlignVCenter
      wrapMode: Text.Wrap
      color:"white"; text:"\nFigure showing sine wave. Click on the plot to randomize "
    }
    SpinBox {
      id: spinbox
      value:myapp.sinewaves
      minimumValue:.01
      decimals: 4
      stepSize:0.5
      onEditingFinished:{
        myapp.setScale(spinbox.value)  
      }


    }
    TextField { text: "Hello"; font.capitalization: Font.AllLowercase ;placeholderText:"Type anything here"}
    TextArea {
      // width:parent.width
      font.italic: true   
      baselineOffset:15
      horizontalAlignment: Text.AlignHCenter

      
    }
    Slider {
      id:vSlider
      value: myapp.yoffset
      stepSize:10
      minimumValue:-200
      maximumValue:200
      onValueChanged:{
        myapp.updateOffset(vSlider.value,hSlider.value)  
      }
    }
    Slider {
     id:hSlider
     value: myapp.xoffset
     stepSize:10
     minimumValue:-200
     maximumValue:plot1.height
     onValueChanged:{
      myapp.updateOffset(vSlider.value,hSlider.value)  
    }
  }

}
} 
// Label {

 //   id:figtitle
 //   Layout.alignment: Qt.AlignHCenter
 //   Layout.fillHeight: false
 //   // font.pixelSize: 10
 //   font.bold:true
 //   color: "steelblue" // steelblue
 //   text:"Figure : Plot Rendered by Golang  "

 // }

}


}