$(document).ready(function() {
  $('body').layout({ 
    defaults: {
      fxName: "slide",
      fxSpeed: "slow",
      spacing_closed: 14,
      initClosed: false
    },
    north: {
      size: 50,
      resizable: false,
      closable: false,
    },
    west: {
      applyDefaultStyles: true,
      size: 250,
      resizable: true,
      closable: true,
      //slidable: true,
      resizerDragOpacity: 1,
      resizerCursor: "w-resize",
      sliderCursor: "pointer",
      slideTrigger_open: "click",
      slideTrigger_close: "click",
      togglerTip_open: "Close",
      togglerTip_closed: "Open",
      fxName: "slide",
      fxSpeed: "normal",
      initClosed: false,
    }
  });
});
