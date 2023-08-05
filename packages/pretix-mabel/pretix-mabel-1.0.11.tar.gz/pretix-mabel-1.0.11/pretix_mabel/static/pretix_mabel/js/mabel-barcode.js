// var PDFRenderer = function PDFRenderer(doc, stickerWidth, stickerHeight) {
//   this.doc = doc;
//   this.offX = 0;
//   this.offY = 0;
//   this.stickerWidth = stickerWidth;
//   this.stickerHeight = stickerHeight;
// };

// PDFRenderer.prototype = Object.create(JsBarcode.Renderer.prototype);

// PDFRenderer.prototype.coffX = function() {
//   return this.offX + this.stickerWidth / 2 - this.width / 2;
// };

// PDFRenderer.prototype.size = function(w, h) {
//   this.width = w;
//   this.height = h;
// };

// PDFRenderer.prototype.rect = function(colour, x, y, w, h) {
//   this.doc
//     .rect(mm(x + this.coffX()), mm(y + this.offY), mm(w), mm(h))
//     .fill(colour);
// };

// PDFRenderer.prototype.text = function(text, x, y, size, font, align) {
//   if (align === "right") {
//     x -= this.width;
//   } else if (align === "center") {
//     x -= this.width / 2;
//   }
//   this.doc
//     .font(font)
//     .fontSize(size)
//     .text(text, mm(x + this.coffX()), mm(y - 8 + this.offY), {
//       width: mm(this.width),
//       align: align
//     });
// };

function mm(n) {
  return n * 2.8346456692895527;
}

function drawStickerBase(doc) {
  console.log("Drawing base");
  // modelled after https://www.a4labels.com/products/transparent-gloss-labels-63-x-38mm/24696
  for (var y = 0; y < 7; y++) {
    for (var x = 0; x < 3; x += 1) {
      doc.roundedRect(
        mm(7.75 + (63.5 + 2) * x),
        mm(15.5 + 38.1 * y),
        mm(63.5),
        mm(38.1),
        mm(1.5)
      );
    }
  }
  doc.stroke("#000000");
}

function drawStickerCode(offX, offY, doc, code, stickerWidth) {
  var qrcode = new QrCode(code);
  var matrix = qrcode.getData();
  var pixelSize = 0.58;
  for (var y = 0; y < matrix.length; y++) {
    for (var x = 0; x < matrix[y].length; x++) {
      if (matrix[y][x] === 1) {
        doc.rect(
          mm(stickerWidth / 2 + offX + pixelSize * (x - matrix[y].length / 2)),
          mm(offY + y * pixelSize),
          mm(pixelSize),
          mm(pixelSize)
        );
      }
    }
  }
}

function setUpButton(buttonId, linkId) {
  var tickets = JSON.parse(document.getElementById("ticket-data").innerHTML);

  document
    .getElementById("pdf-generator")
    .addEventListener("click", function() {
      makePDF(linkId, tickets);
    });
}

function makePDF(linkId, tickets) {
  const link = document.getElementById(linkId);
  const button = link.getElementsByTagName("button")[0];

  button.disabled = true;
  button.innerText = "Loading...";
  var totalProgress = tickets.length * 4;

  // create a document and pipe to a blob
  var doc = new PDFDocument({
    size: "a4",
    margin: 0
  });

  var stream = doc.pipe(blobStream());

  var showGrid = false;

  var cols = 3,
    rows = 7;

  var w = 35,
    textHeight = 13,
    offX = 7.75,
    offY = 32,
    borderX = 2,
    borderY = 25.1,
    stickerWidth = 63.5;

  // proceed page by page
  const pages = Math.ceil(tickets.length / (cols * rows));
  let page = 0;
  function forEachTicket(page, cb) {
    for (var i = page * cols * rows; i < (page + 1) * cols * rows; i++) {
      if (i < tickets.length) {
        var x = i % cols;
        var y = Math.floor(i / cols) % rows;
        var ticket = tickets[i];
        cb(x, y, ticket);
      }
    }
  }

  var progress = 0;
  const progressCounter = setInterval(function() {
    button.innerText =
      "Loading... (" + Math.round(100 * progress / totalProgress) + "%)";
  }, 100);

  while (page < pages) {
    // sticker base
    if (showGrid) drawStickerBase(doc);

    // Name labels
    doc.font("Helvetica").fontSize(14);
    forEachTicket(page, function(x, y, ticket) {
      progress++;
      doc.text(
        ticket.guest_name,
        mm(offX + x * (stickerWidth + borderX)),
        mm(offY + y * (textHeight + borderY) - 12),
        {
          width: mm(stickerWidth),
          align: "center"
        }
      );
    });

    // Ticket type
    doc.font("Helvetica-Bold").fontSize(10);
    forEachTicket(page, function(x, y, ticket) {
      progress++;
      doc.text(
        ticket.ticket_type,
        mm(offX + x * (stickerWidth + borderX)),
        mm(offY + y * (textHeight + borderY) - 6),
        {
          width: mm(stickerWidth),
          align: "center"
        }
      );
    });

    // QR codes
    forEachTicket(page, function(x, y, ticket) {
      progress++;
      drawStickerCode(
        offX + x * (stickerWidth + borderX),
        offY + y * (textHeight + borderY) - 1,
        doc,
        ticket.code,
        stickerWidth
      );
    });
    doc.fill("#101233");

    doc.font("Helvetica").fontSize(8);
    forEachTicket(page, function(x, y, ticket) {
      progress++;
      doc.text(
        ticket.code,
        mm(offX + x * (stickerWidth + borderX)),
        mm(offY + y * (textHeight + borderY) + 18),
        {
          width: mm(stickerWidth),
          align: "center"
        }
      );
    });

    // new page
    page += 1;
    if (page < pages) {
      doc.addPage();
    }
  }

  doc.end();
  clearInterval(progressCounter);

  stream.on("finish", function() {
    var url = stream.toBlobURL("application/pdf");
    console.log(url);
    button.innerText = "Click here to download";
    button.disabled = false;
    link.href = url;

    console.log("Document writton");
  });

  stream.on("pipe", function() {
    console.log("Got data");
  });
}

console.log("Getting ticket data...");
// APICaller.get("ticket/getAllDetailed", {}, function(err, data) {
//   if (err) return alert(err);
//   if (data.length < 1) return alert("No tickets found");
//   console.log("Ticket data obtained.");
//   vm.allTickets = _.filter(data, function(t) {
//     return t.status_id === 2;
//   });
//   allTickets = vm.allTickets;
// });

setUpButton("pdf-generator", "pdf");
