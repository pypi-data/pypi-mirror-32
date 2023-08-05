/* -*- Mode: Javascript; indent-tabs-mode:nil; js-indent-level: 2 -*- */
/* vim: set ts=2 et sw=2 tw=80: */

/*************************************************************
 *
 *  MathJax/jax/output/SVG/fonts/TeX/fontdata.js
 *  
 *  Initializes the SVG OutputJax to use the MathJax TeX fonts
 *  for displaying mathematics.
 *
 *  ---------------------------------------------------------------------
 *  
 *  Copyright (c) 2011-2015 The MathJax Consortium
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

(function (SVG,MML,AJAX,HUB) {
  var VERSION = "2.6.0";
  
  var MAIN   = "MathJax_Main",
      BOLD   = "MathJax_Main-bold",
      ITALIC = "MathJax_Math-italic",
      AMS    = "MathJax_AMS",
      SIZE1  = "MathJax_Size1",
      SIZE2  = "MathJax_Size2",
      SIZE3  = "MathJax_Size3",
      SIZE4  = "MathJax_Size4";
  var H = "H", V = "V", EXTRAH = {load:"extra", dir:H}, EXTRAV = {load:"extra", dir:V};
  var STDHW = [[1000,MAIN],[1200,SIZE1],[1800,SIZE2],[2400,SIZE3],[3000,SIZE4]];
  var ARROWREP = [0x2212,MAIN,0,0,0,0,.1];   // add depth for arrow extender
  var DARROWREP = [0x3D,MAIN,0,0,0,0,.1];    // add depth for arrow extender

  SVG.Augment({
    FONTDATA: {
      version: VERSION,
      
      baselineskip: 1200,
      lineH: 800, lineD: 200,
      
      FONTS: {
        "MathJax_Main":             "Main/Regular/Main.js",
        "MathJax_Main-bold":        "Main/Bold/Main.js",
        "MathJax_Main-italic":      "Main/Italic/Main.js",
        "MathJax_Math-italic":      "Math/Italic/Main.js",
        "MathJax_Math-bold-italic": "Math/BoldItalic/Main.js",
        "MathJax_Caligraphic":      "Caligraphic/Regular/Main.js",
        "MathJax_Size1":            "Size1/Regular/Main.js",
        "MathJax_Size2":            "Size2/Regular/Main.js",
        "MathJax_Size3":            "Size3/Regular/Main.js",
        "MathJax_Size4":            "Size4/Regular/Main.js",
        "MathJax_AMS":              "AMS/Regular/Main.js",
        "MathJax_Fraktur":          "Fraktur/Regular/Main.js",
        "MathJax_Fraktur-bold":     "Fraktur/Bold/Main.js",
        "MathJax_SansSerif":        "SansSerif/Regular/Main.js",
        "MathJax_SansSerif-bold":   "SansSerif/Bold/Main.js",
        "MathJax_SansSerif-italic": "SansSerif/Italic/Main.js",
        "MathJax_Script":           "Script/Regular/Main.js",
        "MathJax_Typewriter":       "Typewriter/Regular/Main.js",
        "MathJax_Caligraphic-bold": "Caligraphic/Bold/Main.js"
      },
      
      VARIANT: {
        "normal": {fonts:[MAIN,SIZE1,AMS],
                   offsetG: 0x03B1, variantG: "italic",
                   remap: {0x391:0x41, 0x392:0x42, 0x395:0x45, 0x396:0x5A, 0x397:0x48,
                           0x399:0x49, 0x39A:0x4B, 0x39C:0x4D, 0x39D:0x4E, 0x39F:0x4F,
                           0x3A1:0x50, 0x3A4:0x54, 0x3A7:0x58,
                           0x2216:[0x2216,"-TeX-variant"],  // \smallsetminus
                           0x210F:[0x210F,"-TeX-variant"],  // \hbar
                           0x2032:[0x27,"sans-serif-italic"],  // HACK: a smaller prime
                           0x29F8:[0x002F,MML.VARIANT.ITALIC]}},
        "bold":   {fonts:[BOLD,SIZE1,AMS], bold:true,
                   offsetG: 0x03B1, variantG: "bold-italic",
                   remap: {0x391:0x41, 0x392:0x42, 0x395:0x45, 0x396:0x5A, 0x397:0x48,
                           0x399:0x49, 0x39A:0x4B, 0x39C:0x4D, 0x39D:0x4E, 0x39F:0x4F,
                           0x3A1:0x50, 0x3A4:0x54, 0x3A7:0x58, 0x29F8:[0x002F,"bold-italic"],
                           0x219A:"\u2190\u0338", 0x219B:"\u2192\u0338", 0x21AE:"\u2194\u0338",
                           0x21CD:"\u21D0\u0338", 0x21CE:"\u21D4\u0338", 0x21CF:"\u21D2\u0338",
                           0x2204:"\u2203\u0338", 0x2224:"\u2223\u0338", 0x2226:"\u2225\u0338",
                           0x2241:"\u223C\u0338", 0x2247:"\u2245\u0338", 
                           0x226E:"<\u0338", 0x226F:">\u0338",
                           0x2270:"\u2264\u0338", 0x2271:"\u2265\u0338",
                           0x2280:"\u227A\u0338", 0x2281:"\u227B\u0338",
                           0x2288:"\u2286\u0338", 0x2289:"\u2287\u0338",
                           0x22AC:"\u22A2\u0338", 0x22AD:"\u22A8\u0338",
//                         0x22AE:"\u22A9\u0338", 0x22AF:"\u22AB\u0338",
                           0x22E0:"\u227C\u0338", 0x22E1:"\u227D\u0338"//,
//                         0x22EA:"\u22B2\u0338", 0x22EB:"\u22B3\u0338",
//                         0x22EC:"\u22B4\u0338", 0x22ED:"\u22B5\u0338"
                  }},
        "italic": {fonts:[ITALIC,"MathJax_Main-italic",MAIN,SIZE1,AMS], italic:true,
                   remap: {0x391:0x41, 0x392:0x42, 0x395:0x45, 0x396:0x5A, 0x397:0x48,
                           0x399:0x49, 0x39A:0x4B, 0x39C:0x4D, 0x39D:0x4E, 0x39F:0x4F,
                           0x3A1:0x50, 0x3A4:0x54, 0x3A7:0x58}},
        "bold-italic": {fonts:["MathJax_Math-bold-italic",BOLD,SIZE1,AMS], bold:true, italic:true,
                   remap: {0x391:0x41, 0x392:0x42, 0x395:0x45, 0x396:0x5A, 0x397:0x48,
                           0x399:0x49, 0x39A:0x4B, 0x39C:0x4D, 0x39D:0x4E, 0x39F:0x4F,
                           0x3A1:0x50, 0x3A4:0x54, 0x3A7:0x58}},
        "double-struck": {fonts:[AMS, MAIN]},
        "fraktur": {fonts:["MathJax_Fraktur",MAIN,SIZE1,AMS]},
        "bold-fraktur": {fonts:["MathJax_Fraktur-bold",BOLD,SIZE1,AMS], bold:true},
        "script": {fonts:["MathJax_Script",MAIN,SIZE1,AMS]},
        "bold-script": {fonts:["MathJax_Script",BOLD,SIZE1,AMS], bold:true},
        "sans-serif": {fonts:["MathJax_SansSerif",MAIN,SIZE1,AMS]},
        "bold-sans-serif": {fonts:["MathJax_SansSerif-bold",BOLD,SIZE1,AMS], bold:true},
        "sans-serif-italic": {fonts:["MathJax_SansSerif-italic","MathJax_Main-italic",SIZE1,AMS], italic:true},
        "sans-serif-bold-italic": {fonts:["MathJax_SansSerif-italic","MathJax_Main-italic",SIZE1,AMS], bold:true, italic:true},
        "monospace": {fonts:["MathJax_Typewriter",MAIN,SIZE1,AMS]},
        "-tex-caligraphic": {fonts:["MathJax_Caligraphic",MAIN], offsetA: 0x41, variantA: "italic"},
        "-tex-oldstyle": {fonts:["MathJax_Caligraphic",MAIN]},
        "-tex-mathit": {fonts:["MathJax_Main-italic",ITALIC,MAIN,SIZE1,AMS], italic:true, noIC: true,
                   remap: {0x391:0x41, 0x392:0x42, 0x395:0x45, 0x396:0x5A, 0x397:0x48,
                           0x399:0x49, 0x39A:0x4B, 0x39C:0x4D, 0x39D:0x4E, 0x39F:0x4F,
                           0x3A1:0x50, 0x3A4:0x54, 0x3A7:0x58}},
        "-TeX-variant": {fonts:[AMS,MAIN,SIZE1],   // HACK: to get larger prime for \prime
                   remap: {
                     0x2268: 0xE00C, 0x2269: 0xE00D, 0x2270: 0xE011, 0x2271: 0xE00E,
                     0x2A87: 0xE010, 0x2A88: 0xE00F, 0x2224: 0xE006, 0x2226: 0xE007,
                     0x2288: 0xE016, 0x2289: 0xE018, 0x228A: 0xE01A, 0x228B: 0xE01B,
                     0x2ACB: 0xE017, 0x2ACC: 0xE019, 0x03DC: 0xE008, 0x03F0: 0xE009,
                     0x2216:[0x2216,MML.VARIANT.NORMAL], // \setminus
                     0x210F:[0x210F,MML.VARIANT.NORMAL]  // \hslash
                   }},
        "-largeOp": {fonts:[SIZE2,SIZE1,MAIN]},
        "-smallOp": {fonts:[SIZE1,MAIN]},
        "-tex-caligraphic-bold": {fonts:["MathJax_Caligraphic-bold","MathJax_Main-bold","MathJax_Main","MathJax_Math","MathJax_Size1"], bold:true,
                                  offsetA: 0x41, variantA: "bold-italic"},
        "-tex-oldstyle-bold": {fonts:["MathJax_Caligraphic-bold","MathJax_Main-bold","MathJax_Main","MathJax_Math","MathJax_Size1"], bold:true}
      },
      
      RANGES: [
        {name: "alpha", low: 0x61, high: 0x7A, offset: "A", add: 32},
        {name: "number", low: 0x30, high: 0x39, offset: "N"},
        {name: "greek", low: 0x03B1, high: 0x03F6, offset: "G"}
      ],
      
      RULECHAR: 0x2212,
      
      REMAP: {
        0x00A0: 0x20,                   // non-breaking space
        0x203E: 0x2C9,                  // overline
        0x20D0: 0x21BC, 0x20D1: 0x21C0, // combining left and right harpoons
        0x20D6: 0x2190, 0x20E1: 0x2194, // combining left arrow and lef-right arrow
        0x20EC: 0x21C1, 0x20ED: 0x21BD, // combining low right and left harpoons
        0x20EE: 0x2190, 0x20EF: 0x2192, // combining low left and right arrows
        0x20F0: 0x2A,                   // combining asterisk
        0xFE37: 0x23DE, 0xFE38: 0x23DF, // OverBrace, UnderBrace

        0xB7: 0x22C5,                   // center dot
        0x2B9: 0x2032,                  // prime,
        0x3D2: 0x3A5,                   // Upsilon
        0x2015: 0x2014, 0x2017: 0x5F,   // horizontal bars
        0x2022: 0x2219, 0x2044: 0x2F,   // bullet, fraction slash
        0x2305: 0x22BC, 0x2306: 0x2A5E, // barwedge, doublebarwedge
        0x25AA: 0x25A0, 0x25B4: 0x25B2, // blacksquare, blacktriangle
        0x25B5: 0x25B3, 0x25B8: 0x25B6, // triangle, blacktriangleright
        0x25BE: 0x25BC, 0x25BF: 0x25BD, // blacktriangledown, triangledown
        0x25C2: 0x25C0,                 // blacktriangleleft
        0x2329: 0x27E8, 0x232A: 0x27E9, // langle, rangle
        0x3008: 0x27E8, 0x3009: 0x27E9, // langle, rangle
        0x2758: 0x2223,                 // VerticalSeparator
        0x2A2F: 0xD7,                   // cross product

        0x25FB: 0x25A1, 0x25FC: 0x25A0, // square, blacksquare

        //
        //  Letter-like symbols (that appear elsewhere)
        //
        0x2102: [0x0043,MML.VARIANT.DOUBLESTRUCK],
//      0x210A: [0x0067,MML.VARIANT.SCRIPT],
        0x210B: [0x0048,MML.VARIANT.SCRIPT],
        0x210C: [0x0048,MML.VARIANT.FRAKTUR],
        0x210D: [0x0048,MML.VARIANT.DOUBLESTRUCK],
        0x210E: [0x0068,MML.VARIANT.ITALIC],
        0x2110: [0x004A,MML.VARIANT.SCRIPT],
        0x2111: [0x0049,MML.VARIANT.FRAKTUR],
        0x2112: [0x004C,MML.VARIANT.SCRIPT],
        0x2115: [0x004E,MML.VARIANT.DOUBLESTRUCK],
        0x2119: [0x0050,MML.VARIANT.DOUBLESTRUCK],
        0x211A: [0x0051,MML.VARIANT.DOUBLESTRUCK],
        0x211B: [0x0052,MML.VARIANT.SCRIPT],
        0x211C: [0x0052,MML.VARIANT.FRAKTUR],
        0x211D: [0x0052,MML.VARIANT.DOUBLESTRUCK],
        0x2124: [0x005A,MML.VARIANT.DOUBLESTRUCK],
        0x2126: [0x03A9,MML.VARIANT.NORMAL],
        0x2128: [0x005A,MML.VARIANT.FRAKTUR],
        0x212C: [0x0042,MML.VARIANT.SCRIPT],
        0x212D: [0x0043,MML.VARIANT.FRAKTUR],
//      0x212F: [0x0065,MML.VARIANT.SCRIPT],
        0x2130: [0x0045,MML.VARIANT.SCRIPT],
        0x2131: [0x0046,MML.VARIANT.SCRIPT],
        0x2133: [0x004D,MML.VARIANT.SCRIPT],
//      0x2134: [0x006F,MML.VARIANT.SCRIPT],

        0x2247: 0x2246,                 // wrong placement of this character
        0x231C: 0x250C, 0x231D:0x2510,  // wrong placement of \ulcorner, \urcorner
        0x231E: 0x2514, 0x231F:0x2518,  // wrong placement of \llcorner, \lrcorner

        //
        //  compound symbols not in these fonts
        //  
        0x2204: "\u2203\u0338",    // \not\exists
        0x220C: "\u220B\u0338",    // \not\ni
        0x2244: "\u2243\u0338",    // \not\simeq
        0x2249: "\u2248\u0338",    // \not\approx
        0x2262: "\u2261\u0338",    // \not\equiv
        0x226D: "\u224D\u0338",    // \not\asymp
        0x2274: "\u2272\u0338",    // \not\lesssim
        0x2275: "\u2273\u0338",    // \not\gtrsim
        0x2278: "\u2276\u0338",    // \not\lessgtr
        0x2279: "\u2277\u0338",    // \not\gtrless
        0x2284: "\u2282\u0338",    // \not\subset
        0x2285: "\u2283\u0338",    // \not\supset
        0x22E2: "\u2291\u0338",    // \not\sqsubseteq
        0x22E3: "\u2292\u0338",    // \not\sqsupseteq

        0x2A0C: "\u222C\u222C",    // quadruple integral

        0x2033: "\u2032\u2032",        // double prime
        0x2034: "\u2032\u2032\u2032",  // triple prime
        0x2036: "\u2035\u2035",        // double back prime
        0x2037: "\u2035\u2035\u2035",  // trile back prime
        0x2057: "\u2032\u2032\u2032\u2032",  // quadruple prime
        0x20DB: "...",                 // combining three dots above (only works with mover/under)
        0x20DC: "...."                 // combining four dots above (only works with mover/under)
      },
      
      REMAPACCENT: {
        "\u2192":"\u20D7",
        "\u2032":"'",
        "\u2035":"`"
      },
      REMAPACCENTUNDER: {
      },
      
      PLANE1MAP: [
        [0x1D400,0x1D419, 0x41, MML.VARIANT.BOLD],
        [0x1D41A,0x1D433, 0x61, MML.VARIANT.BOLD],
        [0x1D434,0x1D44D, 0x41, MML.VARIANT.ITALIC],
        [0x1D44E,0x1D467, 0x61, MML.VARIANT.ITALIC],
        [0x1D468,0x1D481, 0x41, MML.VARIANT.BOLDITALIC],
        [0x1D482,0x1D49B, 0x61, MML.VARIANT.BOLDITALIC],
        [0x1D49C,0x1D4B5, 0x41, MML.VARIANT.SCRIPT],
//      [0x1D4B6,0x1D4CF, 0x61, MML.VARIANT.SCRIPT],
//      [0x1D4D0,0x1D4E9, 0x41, MML.VARIANT.BOLDSCRIPT],
//      [0x1D4EA,0x1D503, 0x61, MML.VARIANT.BOLDSCRIPT],
        [0x1D504,0x1D51D, 0x41, MML.VARIANT.FRAKTUR],
        [0x1D51E,0x1D537, 0x61, MML.VARIANT.FRAKTUR],
        [0x1D538,0x1D551, 0x41, MML.VARIANT.DOUBLESTRUCK],
//      [0x1D552,0x1D56B, 0x61, MML.VARIANT.DOUBLESTRUCK],
        [0x1D56C,0x1D585, 0x41, MML.VARIANT.BOLDFRAKTUR],
        [0x1D586,0x1D59F, 0x61, MML.VARIANT.BOLDFRAKTUR],
        [0x1D5A0,0x1D5B9, 0x41, MML.VARIANT.SANSSERIF],
        [0x1D5BA,0x1D5D3, 0x61, MML.VARIANT.SANSSERIF],
        [0x1D5D4,0x1D5ED, 0x41, MML.VARIANT.BOLDSANSSERIF],
        [0x1D5EE,0x1D607, 0x61, MML.VARIANT.BOLDSANSSERIF],
        [0x1D608,0x1D621, 0x41, MML.VARIANT.SANSSERIFITALIC],
        [0x1D622,0x1D63B, 0x61, MML.VARIANT.SANSSERIFITALIC],
//      [0x1D63C,0x1D655, 0x41, MML.VARIANT.SANSSERIFBOLDITALIC],
//      [0x1D656,0x1D66F, 0x61, MML.VARIANT.SANSSERIFBOLDITALIC],
        [0x1D670,0x1D689, 0x41, MML.VARIANT.MONOSPACE],
        [0x1D68A,0x1D6A3, 0x61, MML.VARIANT.MONOSPACE],
        
        [0x1D6A8,0x1D6C1, 0x391, MML.VARIANT.BOLD],
//      [0x1D6C2,0x1D6E1, 0x3B1, MML.VARIANT.BOLD],
        [0x1D6E2,0x1D6FA, 0x391, MML.VARIANT.ITALIC],
        [0x1D6FC,0x1D71B, 0x3B1, MML.VARIANT.ITALIC],
        [0x1D71C,0x1D734, 0x391, MML.VARIANT.BOLDITALIC],
        [0x1D736,0x1D755, 0x3B1, MML.VARIANT.BOLDITALIC],
        [0x1D756,0x1D76E, 0x391, MML.VARIANT.BOLDSANSSERIF],
//      [0x1D770,0x1D78F, 0x3B1, MML.VARIANT.BOLDSANSSERIF],
        [0x1D790,0x1D7A8, 0x391, MML.VARIANT.SANSSERIFBOLDITALIC],
//      [0x1D7AA,0x1D7C9, 0x3B1, MML.VARIANT.SANSSERIFBOLDITALIC],
        
        [0x1D7CE,0x1D7D7, 0x30, MML.VARIANT.BOLD],
//      [0x1D7D8,0x1D7E1, 0x30, MML.VARIANT.DOUBLESTRUCK],
        [0x1D7E2,0x1D7EB, 0x30, MML.VARIANT.SANSSERIF],
        [0x1D7EC,0x1D7F5, 0x30, MML.VARIANT.BOLDSANSSERIF],
        [0x1D7F6,0x1D7FF, 0x30, MML.VARIANT.MONOSPACE]
      ],

      REMAPGREEK: {
        0x391: 0x41, 0x392: 0x42, 0x395: 0x45, 0x396: 0x5A,
        0x397: 0x48, 0x399: 0x49, 0x39A: 0x4B, 0x39C: 0x4D,
        0x39D: 0x4E, 0x39F: 0x4F, 0x3A1: 0x50, 0x3A2: 0x398,
        0x3A4: 0x54, 0x3A7: 0x58, 0x3AA: 0x2207,
        0x3CA: 0x2202, 0x3CB: 0x3F5, 0x3CC: 0x3D1, 0x3CD: 0x3F0,
        0x3CE: 0x3D5, 0x3CF: 0x3F1, 0x3D0: 0x3D6
      },
      
      RemapPlane1: function (n,variant) {
        for (var i = 0, m = this.PLANE1MAP.length; i < m; i++) {
          if (n < this.PLANE1MAP[i][0]) break;
          if (n <= this.PLANE1MAP[i][1]) {
            n = n - this.PLANE1MAP[i][0] + this.PLANE1MAP[i][2];
            if (this.REMAPGREEK[n]) {n = this.REMAPGREEK[n]}
            variant = this.VARIANT[this.PLANE1MAP[i][3]];
            break;
          }
        }
        return {n: n, variant: variant};
      },
      
      DELIMITERS: {
        0x0028: // (
        {
          dir: V, HW: STDHW,
          stretch: {top: [0x239B,SIZE4], ext: [0x239C,SIZE4], bot: [0x239D,SIZE4]}
        },
        0x0029: // )
        {
          dir: V, HW: STDHW,
          stretch: {top:[0x239E,SIZE4], ext:[0x239F,SIZE4], bot:[0x23A0,SIZE4]}
        },
        0x002F: // /
        {
          dir: V, HW: STDHW
        },
        0x005B: // [
        {
          dir: V, HW: STDHW,
          stretch: {top:[0x23A1,SIZE4], ext:[0x23A2,SIZE4], bot:[0x23A3,SIZE4]}
        },
        0x005C: // \
        {
          dir: V, HW: STDHW
        },
        0x005D: // ]
        {
          dir: V, HW: STDHW,
          stretch: {top:[0x23A4,SIZE4], ext:[0x23A5,SIZE4], bot:[0x23A6,SIZE4]}
        },
        0x007B: // {
        {
          dir: V, HW: STDHW,
          stretch: {top:[0x23A7,SIZE4], mid:[0x23A8,SIZE4], bot:[0x23A9,SIZE4], ext:[0x23AA,SIZE4]}
        },
        0x007C: // |
        {
          dir: V, HW: [[1000,MAIN]], stretch: {ext:[0x2223,MAIN]}
        },
        0x007D: // }
        {
          dir: V, HW: STDHW,
          stretch: {top: [0x23AB,SIZE4], mid:[0x23AC,SIZE4], bot: [0x23AD,SIZE4], ext: [0x23AA,SIZE4]}
        },
        0x00AF: // macron
        {
          dir: H, HW: [[.59,MAIN]], stretch: {rep:[0xAF,MAIN]}
        },
        0x02C6: // wide hat
        {
          dir: H, HW: [[267+250,MAIN],[567+250,SIZE1],[1005+330,SIZE2],[1447+330,SIZE3],[1909,SIZE4]]
        },
        0x02DC: // wide tilde
        {
          dir: H, HW: [[333+250,MAIN],[555+250,SIZE1],[1000+330,SIZE2],[1443+330,SIZE3],[1887,SIZE4]]
        },
        0x2016: // vertical arrow extension
        {
          dir: V, HW: [[602,SIZE1],[1000,MAIN,null,0x2225]], stretch: {ext:[0x2225,MAIN]}
        },
        0x2190: // left arrow
        {
          dir: H, HW: [[1000,MAIN]], stretch: {left:[0x2190,MAIN], rep:ARROWREP, fuzz:300}
        },
        0x2191: // \uparrow
        {
          dir: V, HW: [[888,MAIN]], stretch: {top:[0x2191,SIZE1], ext:[0x23D0,SIZE1]}
        },
        0x2192: // right arrow
        {
          dir: H, HW: [[1000,MAIN]], stretch: {rep:ARROWREP, right:[0x2192,MAIN], fuzz:300}
        },
        0x2193: // \downarrow
        {
          dir: V, HW: [[888,MAIN]], stretch: {ext:[0x23D0,SIZE1], bot:[0x2193,SIZE1]}
        },
        0x2194: // left-right arrow
        {
          dir: H, HW: [[1000,MAIN]],
          stretch: {left:[0x2190,MAIN], rep:ARROWREP, right:[0x2192,MAIN], fuzz:300}
        },
        0x2195: // \updownarrow
        {
          dir: V, HW: [[1044,MAIN]],
          stretch: {top:[0x2191,SIZE1], ext:[0x23D0,SIZE1], bot:[0x2193,SIZE1]}
        },
        0x21D0: // left double arrow
        {
          dir: H, HW: [[1000,MAIN]], stretch: {left:[0x21D0,MAIN], rep:DARROWREP, fuzz:300}
        },
        0x21D1: // \Uparrow
        {
          dir: V, HW: [[888,MAIN]], stretch: {top:[0x21D1,SIZE1], ext:[0x2016,SIZE1]}
        },
        0x21D2: // right double arrow
        {
          dir: H, HW: [[1000,MAIN]], stretch: {rep:DARROWREP, right:[0x21D2,MAIN], fuzz:300}
        },
        0x21D3: // \Downarrow
        {
          dir: V, HW: [[888,MAIN]], stretch: {ext:[0x2016,SIZE1], bot:[0x21D3,SIZE1]}
        },
        0x21D4: // left-right double arrow
        {
          dir: H, HW: [[1000,MAIN]],
          stretch: {left:[0x21D0,MAIN], rep:DARROWREP, right:[0x21D2,MAIN], fuzz:300}
        },
        0x21D5: // \Updownarrow
        {
          dir: V, HW: [[1044,MAIN]],
          stretch: {top:[0x21D1,SIZE1], ext:[0x2016,SIZE1], bot:[0x21D3,SIZE1]}
        },
        0x2212: // horizontal line
        {
          dir: H, HW: [[778,MAIN]], stretch: {rep:[0x2212,MAIN], fuzz:300}
        },
        0x221A: // \surd
        {
          dir: V, HW: STDHW,
          stretch: {top:[0xE001,SIZE4], ext:[0xE000,SIZE4], bot:[0x23B7,SIZE4], fullExtenders:true}
        },
        0x2223: // \vert
        {
          dir: V, HW: [[1000,MAIN]], stretch: {ext:[0x2223,MAIN]}
        },
        0x2225: // \Vert
        {
          dir: V, HW: [[1000,MAIN]], stretch: {ext:[0x2225,MAIN]}
        },
        0x2308: // \lceil
        {
          dir: V, HW: STDHW, stretch: {top:[0x23A1,SIZE4], ext:[0x23A2,SIZE4]}
        },
        0x2309: // \rceil
        {
          dir: V, HW: STDHW, stretch: {top:[0x23A4,SIZE4], ext:[0x23A5,SIZE4]}
        },
        0x230A: // \lfloor
        {
          dir: V, HW: STDHW, stretch: {ext:[0x23A2,SIZE4], bot:[0x23A3,SIZE4]}
        },
        0x230B: // \rfloor
        {
          dir: V, HW: STDHW, stretch: {ext:[0x23A5,SIZE4], bot:[0x23A6,SIZE4]}
        },
        0x23AA: // \bracevert
        {
          dir: V, HW: [[320,SIZE4]],
          stretch: {top:[0x23AA,SIZE4], ext:[0x23AA,SIZE4], bot:[0x23AA,SIZE4]}
        },
        0x23B0: // \lmoustache
        {
          dir: V, HW: [[989,MAIN]],
          stretch: {top:[0x23A7,SIZE4], ext:[0x23AA,SIZE4], bot:[0x23AD,SIZE4]}
        },
        0x23B1: // \rmoustache
        {
          dir: V, HW: [[989,MAIN]],
          stretch: {top:[0x23AB,SIZE4], ext:[0x23AA,SIZE4], bot:[0x23A9,SIZE4]}
        },
        0x23D0: // vertical line extension
        {
          dir: V, HW: [[602,SIZE1],[1000,MAIN,null,0x2223]], stretch: {ext:[0x2223,MAIN]}
        },
        0x23DE: // horizontal brace down
        {
          dir: H, HW: [],
          stretch: {min:.9, left:[0xE150,SIZE4], mid:[[0xE153,0xE152],SIZE4], right:[0xE151,SIZE4], rep:[0xE154,SIZE4]}
        },
        0x23DF: // horizontal brace up
        {
          dir: H, HW: [],
          stretch: {min:.9, left:[0xE152,SIZE4], mid:[[0xE151,0xE150],SIZE4], right:[0xE153,SIZE4], rep:[0xE154,SIZE4]}
        },
        0x27E8: // \langle
        {
          dir: V, HW: STDHW
        },
        0x27E9: // \rangle
        {
          dir: V, HW: STDHW
        },
        0x27EE: // \lgroup
        {
          dir: V, HW: [[989,MAIN]],
          stretch: {top:[0x23A7,SIZE4], ext:[0x23AA,SIZE4], bot:[0x23A9,SIZE4]}
        },
        0x27EF: // \rgroup
        {
          dir: V, HW: [[989,MAIN]],
          stretch: {top:[0x23AB,SIZE4], ext:[0x23AA,SIZE4], bot:[0x23AD,SIZE4]}
        },
        0x002D: {alias: 0x2212, dir:H}, // minus
        0x005E: {alias: 0x02C6, dir:H}, // wide hat
        0x005F: {alias: 0x2212, dir:H}, // low line
        0x007E: {alias: 0x02DC, dir:H}, // wide tilde
        0x02C9: {alias: 0x00AF, dir:H}, // macron
        0x0302: {alias: 0x02C6, dir:H}, // wide hat
        0x0303: {alias: 0x02DC, dir:H}, // wide tilde
        0x030C: {alias: 0x02C7, dir:H}, // wide caron
        0x0332: {alias: 0x2212, dir:H}, // combining low line
        0x2015: {alias: 0x2212, dir:H}, // horizontal line
        0x2017: {alias: 0x2212, dir:H}, // horizontal line
        0x203E: {alias: 0x00AF, dir:H}, // over line
        0x2215: {alias: 0x002F, dir:V}, // division slash
        0x2329: {alias: 0x27E8, dir:V}, // langle
        0x232A: {alias: 0x27E9, dir:V}, // rangle
        0x23AF: {alias: 0x2212, dir:H}, // horizontal line extension
        0x2500: {alias: 0x2212, dir:H}, // horizontal line
        0x2758: {alias: 0x2223, dir:V}, // vertical separator
        0x3008: {alias: 0x27E8, dir:V}, // langle
        0x3009: {alias: 0x27E9, dir:V}, // rangle
        0xFE37: {alias: 0x23DE, dir:H}, // horizontal brace down
        0xFE38: {alias: 0x23DF, dir:H},  // horizontal brace up

        0x003D: EXTRAH, // equal sign
        0x219E: EXTRAH, // left two-headed arrow
        0x21A0: EXTRAH, // right two-headed arrow
        0x21A4: EXTRAH, // left arrow from bar
        0x21A5: EXTRAV, // up arrow from bar
        0x21A6: EXTRAH, // right arrow from bar
        0x21A7: EXTRAV, // down arrow from bar
        0x21B0: EXTRAV, // up arrow with top leftwards
        0x21B1: EXTRAV, // up arrow with top right
        0x21BC: EXTRAH, // left harpoon with barb up
        0x21BD: EXTRAH, // left harpoon with barb down
        0x21BE: EXTRAV, // up harpoon with barb right
        0x21BF: EXTRAV, // up harpoon with barb left
        0x21C0: EXTRAH, // right harpoon with barb up
        0x21C1: EXTRAH, // right harpoon with barb down
        0x21C2: EXTRAV, // down harpoon with barb right
        0x21C3: EXTRAV, // down harpoon with barb left
        0x21DA: EXTRAH, // left triple arrow
        0x21DB: EXTRAH, // right triple arrow
        0x23B4: EXTRAH, // top square bracket
        0x23B5: EXTRAH, // bottom square bracket
        0x23DC: EXTRAH, // top paren
        0x23DD: EXTRAH, // bottom paren
        0x23E0: EXTRAH, // top tortoise shell
        0x23E1: EXTRAH, // bottom tortoise shell
        0x2906: EXTRAH, // leftwards double arrow from bar
        0x2907: EXTRAH, // rightwards double arrow from bar
        0x294E: EXTRAH, // left barb up right barb up harpoon
        0x294F: EXTRAV, // up barb right down barb right harpoon
        0x2950: EXTRAH, // left barb dow right barb down harpoon
        0x2951: EXTRAV, // up barb left down barb left harpoon
        0x295A: EXTRAH, // leftwards harpoon with barb up from bar
        0x295B: EXTRAH, // rightwards harpoon with barb up from bar
        0x295C: EXTRAV, // up harpoon with barb right from bar
        0x295D: EXTRAV, // down harpoon with barb right from bar
        0x295E: EXTRAH, // leftwards harpoon with barb down from bar
        0x295F: EXTRAH, // rightwards harpoon with barb down from bar
        0x2960: EXTRAV, // up harpoon with barb left from bar
        0x2961: EXTRAV, // down harpoon with barb left from bar
        0x2312: {alias: 0x23DC, dir:H}, // arc
        0x2322: {alias: 0x23DC, dir:H}, // frown
        0x2323: {alias: 0x23DD, dir:H}, // smile
        0x27F5: {alias: 0x2190, dir:H}, // long left arrow
        0x27F6: {alias: 0x2192, dir:H}, // long right arrow
        0x27F7: {alias: 0x2194, dir:H}, // long left-right arrow
        0x27F8: {alias: 0x21D0, dir:H}, // long left double arrow
        0x27F9: {alias: 0x21D2, dir:H}, // long right double arrow
        0x27FA: {alias: 0x21D4, dir:H}, // long left-right double arrow
        0x27FB: {alias: 0x21A4, dir:H}, // long left arrow from bar
        0x27FC: {alias: 0x21A6, dir:H}, // long right arrow from bar
        0x27FD: {alias: 0x2906, dir:H}, // long left double arrow from bar
        0x27FE: {alias: 0x2907, dir:H}  // long right double arrow from bar
      }
    }
  });

  
  SVG.FONTDATA.FONTS['MathJax_Main'] = {
    directory: 'Main/Regular',
    family: 'MathJax_Main',
    id: 'MJMAIN',
    skew: {
      0x131: 0.0278,
      0x237: 0.0833,
      0x2113: 0.111,
      0x2118: 0.111,
      0x2202: 0.0833
    },
    Ranges: [
      [0x20,0x7F,"BasicLatin"],
      [0x100,0x17F,"LatinExtendedA"],
      [0x180,0x24F,"LatinExtendedB"],
      [0x2B0,0x2FF,"SpacingModLetters"],
      [0x300,0x36F,"CombDiacritMarks"],
      [0x370,0x3FF,"GreekAndCoptic"],
      [0x2100,0x214F,"LetterlikeSymbols"],
      [0x2200,0x22FF,"MathOperators"],
      [0x25A0,0x25FF,"GeometricShapes"],
      [0x2600,0x26FF,"MiscSymbols"],
      [0x2A00,0x2AFF,"SuppMathOperators"]
    ],
  
      // SPACE
      0x20: [0,0,250,0,0,''],
  
      // LEFT PARENTHESIS
      0x28: [750,250,389,94,333,'94 250Q94 319 104 381T127 488T164 576T202 643T244 695T277 729T302 750H315H319Q333 750 333 741Q333 738 316 720T275 667T226 581T184 443T167 250T184 58T225 -81T274 -167T316 -220T333 -241Q333 -250 318 -250H315H302L274 -226Q180 -141 137 -14T94 250'],
  
      // RIGHT PARENTHESIS
      0x29: [750,250,389,55,294,'60 749L64 750Q69 750 74 750H86L114 726Q208 641 251 514T294 250Q294 182 284 119T261 12T224 -76T186 -143T145 -194T113 -227T90 -246Q87 -249 86 -250H74Q66 -250 63 -250T58 -247T55 -238Q56 -237 66 -225Q221 -64 221 250T66 725Q56 737 55 738Q55 746 60 749'],
  
      // PLUS SIGN
      0x2B: [583,82,778,56,722,'56 237T56 250T70 270H369V420L370 570Q380 583 389 583Q402 583 409 568V270H707Q722 262 722 250T707 230H409V-68Q401 -82 391 -82H389H387Q375 -82 369 -68V230H70Q56 237 56 250'],
  
      // COMMA
      0x2C: [121,195,278,78,210,'78 35T78 60T94 103T137 121Q165 121 187 96T210 8Q210 -27 201 -60T180 -117T154 -158T130 -185T117 -194Q113 -194 104 -185T95 -172Q95 -168 106 -156T131 -126T157 -76T173 -3V9L172 8Q170 7 167 6T161 3T152 1T140 0Q113 0 96 17'],
  
      // FULL STOP
      0x2E: [120,0,278,78,199,'78 60Q78 84 95 102T138 120Q162 120 180 104T199 61Q199 36 182 18T139 0T96 17T78 60'],
  
      // SOLIDUS
      0x2F: [750,250,500,56,444,'423 750Q432 750 438 744T444 730Q444 725 271 248T92 -240Q85 -250 75 -250Q68 -250 62 -245T56 -231Q56 -221 230 257T407 740Q411 750 423 750'],
  
      // DIGIT ZERO
      0x30: [666,22,500,39,460,'96 585Q152 666 249 666Q297 666 345 640T423 548Q460 465 460 320Q460 165 417 83Q397 41 362 16T301 -15T250 -22Q224 -22 198 -16T137 16T82 83Q39 165 39 320Q39 494 96 585ZM321 597Q291 629 250 629Q208 629 178 597Q153 571 145 525T137 333Q137 175 145 125T181 46Q209 16 250 16Q290 16 318 46Q347 76 354 130T362 333Q362 478 354 524T321 597'],
  
      // DIGIT ONE
      0x31: [666,0,500,83,427,'213 578L200 573Q186 568 160 563T102 556H83V602H102Q149 604 189 617T245 641T273 663Q275 666 285 666Q294 666 302 660V361L303 61Q310 54 315 52T339 48T401 46H427V0H416Q395 3 257 3Q121 3 100 0H88V46H114Q136 46 152 46T177 47T193 50T201 52T207 57T213 61V578'],
  
      // DIGIT TWO
      0x32: [666,0,500,50,449,'109 429Q82 429 66 447T50 491Q50 562 103 614T235 666Q326 666 387 610T449 465Q449 422 429 383T381 315T301 241Q265 210 201 149L142 93L218 92Q375 92 385 97Q392 99 409 186V189H449V186Q448 183 436 95T421 3V0H50V19V31Q50 38 56 46T86 81Q115 113 136 137Q145 147 170 174T204 211T233 244T261 278T284 308T305 340T320 369T333 401T340 431T343 464Q343 527 309 573T212 619Q179 619 154 602T119 569T109 550Q109 549 114 549Q132 549 151 535T170 489Q170 464 154 447T109 429'],
  
      // DIGIT THREE
      0x33: [665,22,500,42,457,'127 463Q100 463 85 480T69 524Q69 579 117 622T233 665Q268 665 277 664Q351 652 390 611T430 522Q430 470 396 421T302 350L299 348Q299 347 308 345T337 336T375 315Q457 262 457 175Q457 96 395 37T238 -22Q158 -22 100 21T42 130Q42 158 60 175T105 193Q133 193 151 175T169 130Q169 119 166 110T159 94T148 82T136 74T126 70T118 67L114 66Q165 21 238 21Q293 21 321 74Q338 107 338 175V195Q338 290 274 322Q259 328 213 329L171 330L168 332Q166 335 166 348Q166 366 174 366Q202 366 232 371Q266 376 294 413T322 525V533Q322 590 287 612Q265 626 240 626Q208 626 181 615T143 592T132 580H135Q138 579 143 578T153 573T165 566T175 555T183 540T186 520Q186 498 172 481T127 463'],
  
      // DIGIT FOUR
      0x34: [677,0,500,28,471,'462 0Q444 3 333 3Q217 3 199 0H190V46H221Q241 46 248 46T265 48T279 53T286 61Q287 63 287 115V165H28V211L179 442Q332 674 334 675Q336 677 355 677H373L379 671V211H471V165H379V114Q379 73 379 66T385 54Q393 47 442 46H471V0H462ZM293 211V545L74 212L183 211H293'],
  
      // DIGIT FIVE
      0x35: [666,22,500,50,449,'164 157Q164 133 148 117T109 101H102Q148 22 224 22Q294 22 326 82Q345 115 345 210Q345 313 318 349Q292 382 260 382H254Q176 382 136 314Q132 307 129 306T114 304Q97 304 95 310Q93 314 93 485V614Q93 664 98 664Q100 666 102 666Q103 666 123 658T178 642T253 634Q324 634 389 662Q397 666 402 666Q410 666 410 648V635Q328 538 205 538Q174 538 149 544L139 546V374Q158 388 169 396T205 412T256 420Q337 420 393 355T449 201Q449 109 385 44T229 -22Q148 -22 99 32T50 154Q50 178 61 192T84 210T107 214Q132 214 148 197T164 157'],
  
      // DIGIT SIX
      0x36: [666,22,500,41,456,'42 313Q42 476 123 571T303 666Q372 666 402 630T432 550Q432 525 418 510T379 495Q356 495 341 509T326 548Q326 592 373 601Q351 623 311 626Q240 626 194 566Q147 500 147 364L148 360Q153 366 156 373Q197 433 263 433H267Q313 433 348 414Q372 400 396 374T435 317Q456 268 456 210V192Q456 169 451 149Q440 90 387 34T253 -22Q225 -22 199 -14T143 16T92 75T56 172T42 313ZM257 397Q227 397 205 380T171 335T154 278T148 216Q148 133 160 97T198 39Q222 21 251 21Q302 21 329 59Q342 77 347 104T352 209Q352 289 347 316T329 361Q302 397 257 397'],
  
      // DIGIT SEVEN
      0x37: [676,22,500,55,485,'55 458Q56 460 72 567L88 674Q88 676 108 676H128V672Q128 662 143 655T195 646T364 644H485V605L417 512Q408 500 387 472T360 435T339 403T319 367T305 330T292 284T284 230T278 162T275 80Q275 66 275 52T274 28V19Q270 2 255 -10T221 -22Q210 -22 200 -19T179 0T168 40Q168 198 265 368Q285 400 349 489L395 552H302Q128 552 119 546Q113 543 108 522T98 479L95 458V455H55V458'],
  
      // DIGIT EIGHT
      0x38: [666,22,500,43,457,'70 417T70 494T124 618T248 666Q319 666 374 624T429 515Q429 485 418 459T392 417T361 389T335 371T324 363L338 354Q352 344 366 334T382 323Q457 264 457 174Q457 95 399 37T249 -22Q159 -22 101 29T43 155Q43 263 172 335L154 348Q133 361 127 368Q70 417 70 494ZM286 386L292 390Q298 394 301 396T311 403T323 413T334 425T345 438T355 454T364 471T369 491T371 513Q371 556 342 586T275 624Q268 625 242 625Q201 625 165 599T128 534Q128 511 141 492T167 463T217 431Q224 426 228 424L286 386ZM250 21Q308 21 350 55T392 137Q392 154 387 169T375 194T353 216T330 234T301 253T274 270Q260 279 244 289T218 306L210 311Q204 311 181 294T133 239T107 157Q107 98 150 60T250 21'],
  
      // DIGIT NINE
      0x39: [666,22,500,42,456,'352 287Q304 211 232 211Q154 211 104 270T44 396Q42 412 42 436V444Q42 537 111 606Q171 666 243 666Q245 666 249 666T257 665H261Q273 665 286 663T323 651T370 619T413 560Q456 472 456 334Q456 194 396 97Q361 41 312 10T208 -22Q147 -22 108 7T68 93T121 149Q143 149 158 135T173 96Q173 78 164 65T148 49T135 44L131 43Q131 41 138 37T164 27T206 22H212Q272 22 313 86Q352 142 352 280V287ZM244 248Q292 248 321 297T351 430Q351 508 343 542Q341 552 337 562T323 588T293 615T246 625Q208 625 181 598Q160 576 154 546T147 441Q147 358 152 329T172 282Q197 248 244 248'],
  
      // COLON
      0x3A: [430,0,278,78,199,'78 370Q78 394 95 412T138 430Q162 430 180 414T199 371Q199 346 182 328T139 310T96 327T78 370ZM78 60Q78 84 95 102T138 120Q162 120 180 104T199 61Q199 36 182 18T139 0T96 17T78 60'],
  
      // SEMICOLON
      0x3B: [430,194,278,78,202,'78 370Q78 394 95 412T138 430Q162 430 180 414T199 371Q199 346 182 328T139 310T96 327T78 370ZM78 60Q78 85 94 103T137 121Q202 121 202 8Q202 -44 183 -94T144 -169T118 -194Q115 -194 106 -186T95 -174Q94 -171 107 -155T137 -107T160 -38Q161 -32 162 -22T165 -4T165 4Q165 5 161 4T142 0Q110 0 94 18T78 60'],
  
      // LESS-THAN SIGN
      0x3C: [540,40,778,83,695,'694 -11T694 -19T688 -33T678 -40Q671 -40 524 29T234 166L90 235Q83 240 83 250Q83 261 91 266Q664 540 678 540Q681 540 687 534T694 519T687 505Q686 504 417 376L151 250L417 124Q686 -4 687 -5Q694 -11 694 -19'],
  
      // EQUALS SIGN
      0x3D: [367,-133,778,56,722,'56 347Q56 360 70 367H707Q722 359 722 347Q722 336 708 328L390 327H72Q56 332 56 347ZM56 153Q56 168 72 173H708Q722 163 722 153Q722 140 707 133H70Q56 140 56 153'],
  
      // GREATER-THAN SIGN
      0x3E: [540,40,778,82,694,'84 520Q84 528 88 533T96 539L99 540Q106 540 253 471T544 334L687 265Q694 260 694 250T687 235Q685 233 395 96L107 -40H101Q83 -38 83 -20Q83 -19 83 -17Q82 -10 98 -1Q117 9 248 71Q326 108 378 132L626 250L378 368Q90 504 86 509Q84 513 84 520'],
  
      // LEFT SQUARE BRACKET
      0x5B: [750,250,278,118,255,'118 -250V750H255V710H158V-210H255V-250H118'],
  
      // REVERSE SOLIDUS
      0x5C: [750,250,500,56,444,'56 731Q56 740 62 745T75 750Q85 750 92 740Q96 733 270 255T444 -231Q444 -239 438 -244T424 -250Q414 -250 407 -240Q404 -236 230 242T56 731'],
  
      // RIGHT SQUARE BRACKET
      0x5D: [750,250,278,22,159,'22 710V750H159V-250H22V-210H119V710H22'],
  
      // CIRCUMFLEX ACCENT
      0x5E: [694,-531,500,112,387,'112 560L249 694L257 686Q387 562 387 560L361 531Q359 532 303 581L250 627L195 580Q182 569 169 557T148 538L140 532Q138 530 125 546L112 560'],
  
      // LATIN SMALL LETTER A
      0x61: [448,11,500,34,493,'137 305T115 305T78 320T63 359Q63 394 97 421T218 448Q291 448 336 416T396 340Q401 326 401 309T402 194V124Q402 76 407 58T428 40Q443 40 448 56T453 109V145H493V106Q492 66 490 59Q481 29 455 12T400 -6T353 12T329 54V58L327 55Q325 52 322 49T314 40T302 29T287 17T269 6T247 -2T221 -8T190 -11Q130 -11 82 20T34 107Q34 128 41 147T68 188T116 225T194 253T304 268H318V290Q318 324 312 340Q290 411 215 411Q197 411 181 410T156 406T148 403Q170 388 170 359Q170 334 154 320ZM126 106Q126 75 150 51T209 26Q247 26 276 49T315 109Q317 116 318 175Q318 233 317 233Q309 233 296 232T251 223T193 203T147 166T126 106'],
  
      // LATIN SMALL LETTER B
      0x62: [695,11,556,20,522,'307 -11Q234 -11 168 55L158 37Q156 34 153 28T147 17T143 10L138 1L118 0H98V298Q98 599 97 603Q94 622 83 628T38 637H20V660Q20 683 22 683L32 684Q42 685 61 686T98 688Q115 689 135 690T165 693T176 694H179V543Q179 391 180 391L183 394Q186 397 192 401T207 411T228 421T254 431T286 439T323 442Q401 442 461 379T522 216Q522 115 458 52T307 -11ZM182 98Q182 97 187 90T196 79T206 67T218 55T233 44T250 35T271 29T295 26Q330 26 363 46T412 113Q424 148 424 212Q424 287 412 323Q385 405 300 405Q270 405 239 390T188 347L182 339V98'],
  
      // LATIN SMALL LETTER C
      0x63: [448,12,444,34,415,'370 305T349 305T313 320T297 358Q297 381 312 396Q317 401 317 402T307 404Q281 408 258 408Q209 408 178 376Q131 329 131 219Q131 137 162 90Q203 29 272 29Q313 29 338 55T374 117Q376 125 379 127T395 129H409Q415 123 415 120Q415 116 411 104T395 71T366 33T318 2T249 -11Q163 -11 99 53T34 214Q34 318 99 383T250 448T370 421T404 357Q404 334 387 320'],
  
      // LATIN SMALL LETTER D
      0x64: [695,11,556,34,535,'376 495Q376 511 376 535T377 568Q377 613 367 624T316 637H298V660Q298 683 300 683L310 684Q320 685 339 686T376 688Q393 689 413 690T443 693T454 694H457V390Q457 84 458 81Q461 61 472 55T517 46H535V0Q533 0 459 -5T380 -11H373V44L365 37Q307 -11 235 -11Q158 -11 96 50T34 215Q34 315 97 378T244 442Q319 442 376 393V495ZM373 342Q328 405 260 405Q211 405 173 369Q146 341 139 305T131 211Q131 155 138 120T173 59Q203 26 251 26Q322 26 373 103V342'],
  
      // LATIN SMALL LETTER E
      0x65: [448,11,444,28,415,'28 218Q28 273 48 318T98 391T163 433T229 448Q282 448 320 430T378 380T406 316T415 245Q415 238 408 231H126V216Q126 68 226 36Q246 30 270 30Q312 30 342 62Q359 79 369 104L379 128Q382 131 395 131H398Q415 131 415 121Q415 117 412 108Q393 53 349 21T250 -11Q155 -11 92 58T28 218ZM333 275Q322 403 238 411H236Q228 411 220 410T195 402T166 381T143 340T127 274V267H333V275'],
  
      // LATIN SMALL LETTER F
      0x66: [705,0,306,26,372,'273 0Q255 3 146 3Q43 3 34 0H26V46H42Q70 46 91 49Q99 52 103 60Q104 62 104 224V385H33V431H104V497L105 564L107 574Q126 639 171 668T266 704Q267 704 275 704T289 705Q330 702 351 679T372 627Q372 604 358 590T321 576T284 590T270 627Q270 647 288 667H284Q280 668 273 668Q245 668 223 647T189 592Q183 572 182 497V431H293V385H185V225Q185 63 186 61T189 57T194 54T199 51T206 49T213 48T222 47T231 47T241 46T251 46H282V0H273'],
  
      // LATIN SMALL LETTER G
      0x67: [453,206,500,29,485,'329 409Q373 453 429 453Q459 453 472 434T485 396Q485 382 476 371T449 360Q416 360 412 390Q410 404 415 411Q415 412 416 414V415Q388 412 363 393Q355 388 355 386Q355 385 359 381T368 369T379 351T388 325T392 292Q392 230 343 187T222 143Q172 143 123 171Q112 153 112 133Q112 98 138 81Q147 75 155 75T227 73Q311 72 335 67Q396 58 431 26Q470 -13 470 -72Q470 -139 392 -175Q332 -206 250 -206Q167 -206 107 -175Q29 -140 29 -75Q29 -39 50 -15T92 18L103 24Q67 55 67 108Q67 155 96 193Q52 237 52 292Q52 355 102 398T223 442Q274 442 318 416L329 409ZM299 343Q294 371 273 387T221 404Q192 404 171 388T145 343Q142 326 142 292Q142 248 149 227T179 192Q196 182 222 182Q244 182 260 189T283 207T294 227T299 242Q302 258 302 292T299 343ZM403 -75Q403 -50 389 -34T348 -11T299 -2T245 0H218Q151 0 138 -6Q118 -15 107 -34T95 -74Q95 -84 101 -97T122 -127T170 -155T250 -167Q319 -167 361 -139T403 -75'],
  
      // LATIN SMALL LETTER H
      0x68: [695,0,556,25,542,'41 46H55Q94 46 102 60V68Q102 77 102 91T102 124T102 167T103 217T103 272T103 329Q103 366 103 407T103 482T102 542T102 586T102 603Q99 622 88 628T43 637H25V660Q25 683 27 683L37 684Q47 685 66 686T103 688Q120 689 140 690T170 693T181 694H184V367Q244 442 328 442Q451 442 463 329Q464 322 464 190V104Q464 66 466 59T477 49Q498 46 526 46H542V0H534L510 1Q487 2 460 2T422 3Q319 3 310 0H302V46H318Q379 46 379 62Q380 64 380 200Q379 335 378 343Q372 371 358 385T334 402T308 404Q263 404 229 370Q202 343 195 315T187 232V168V108Q187 78 188 68T191 55T200 49Q221 46 249 46H265V0H257L234 1Q210 2 183 2T145 3Q42 3 33 0H25V46H41'],
  
      // LATIN SMALL LETTER I
      0x69: [669,0,278,26,255,'69 609Q69 637 87 653T131 669Q154 667 171 652T188 609Q188 579 171 564T129 549Q104 549 87 564T69 609ZM247 0Q232 3 143 3Q132 3 106 3T56 1L34 0H26V46H42Q70 46 91 49Q100 53 102 60T104 102V205V293Q104 345 102 359T88 378Q74 385 41 385H30V408Q30 431 32 431L42 432Q52 433 70 434T106 436Q123 437 142 438T171 441T182 442H185V62Q190 52 197 50T232 46H255V0H247'],
  
      // LATIN SMALL LETTER J
      0x6A: [669,205,306,-55,218,'98 609Q98 637 116 653T160 669Q183 667 200 652T217 609Q217 579 200 564T158 549Q133 549 116 564T98 609ZM28 -163Q58 -168 64 -168Q124 -168 135 -77Q137 -65 137 141T136 353Q132 371 120 377T72 385H52V408Q52 431 54 431L58 432Q62 432 70 432T87 433T108 434T133 436Q151 437 171 438T202 441T214 442H218V184Q217 -36 217 -59T211 -98Q195 -145 153 -175T58 -205Q9 -205 -23 -179T-55 -117Q-55 -94 -40 -79T-2 -64T36 -79T52 -118Q52 -143 28 -163'],
  
      // LATIN SMALL LETTER K
      0x6B: [695,0,528,20,511,'36 46H50Q89 46 97 60V68Q97 77 97 91T97 124T98 167T98 217T98 272T98 329Q98 366 98 407T98 482T98 542T97 586T97 603Q94 622 83 628T38 637H20V660Q20 683 22 683L32 684Q42 685 61 686T98 688Q115 689 135 690T165 693T176 694H179V463L180 233L240 287Q300 341 304 347Q310 356 310 364Q310 383 289 385H284V431H293Q308 428 412 428Q475 428 484 431H489V385H476Q407 380 360 341Q286 278 286 274Q286 273 349 181T420 79Q434 60 451 53T500 46H511V0H505Q496 3 418 3Q322 3 307 0H299V46H306Q330 48 330 65Q330 72 326 79Q323 84 276 153T228 222L176 176V120V84Q176 65 178 59T189 49Q210 46 238 46H254V0H246Q231 3 137 3T28 0H20V46H36'],
  
      // LATIN SMALL LETTER L
      0x6C: [695,0,278,26,263,'42 46H56Q95 46 103 60V68Q103 77 103 91T103 124T104 167T104 217T104 272T104 329Q104 366 104 407T104 482T104 542T103 586T103 603Q100 622 89 628T44 637H26V660Q26 683 28 683L38 684Q48 685 67 686T104 688Q121 689 141 690T171 693T182 694H185V379Q185 62 186 60Q190 52 198 49Q219 46 247 46H263V0H255L232 1Q209 2 183 2T145 3T107 3T57 1L34 0H26V46H42'],
  
      // LATIN SMALL LETTER M
      0x6D: [443,0,833,25,819,'41 46H55Q94 46 102 60V68Q102 77 102 91T102 122T103 161T103 203Q103 234 103 269T102 328V351Q99 370 88 376T43 385H25V408Q25 431 27 431L37 432Q47 433 65 434T102 436Q119 437 138 438T167 441T178 442H181V402Q181 364 182 364T187 369T199 384T218 402T247 421T285 437Q305 442 336 442Q351 442 364 440T387 434T406 426T421 417T432 406T441 395T448 384T452 374T455 366L457 361L460 365Q463 369 466 373T475 384T488 397T503 410T523 422T546 432T572 439T603 442Q729 442 740 329Q741 322 741 190V104Q741 66 743 59T754 49Q775 46 803 46H819V0H811L788 1Q764 2 737 2T699 3Q596 3 587 0H579V46H595Q656 46 656 62Q657 64 657 200Q656 335 655 343Q649 371 635 385T611 402T585 404Q540 404 506 370Q479 343 472 315T464 232V168V108Q464 78 465 68T468 55T477 49Q498 46 526 46H542V0H534L510 1Q487 2 460 2T422 3Q319 3 310 0H302V46H318Q379 46 379 62Q380 64 380 200Q379 335 378 343Q372 371 358 385T334 402T308 404Q263 404 229 370Q202 343 195 315T187 232V168V108Q187 78 188 68T191 55T200 49Q221 46 249 46H265V0H257L234 1Q210 2 183 2T145 3Q42 3 33 0H25V46H41'],
  
      // LATIN SMALL LETTER N
      0x6E: [443,0,556,25,542,'41 46H55Q94 46 102 60V68Q102 77 102 91T102 122T103 161T103 203Q103 234 103 269T102 328V351Q99 370 88 376T43 385H25V408Q25 431 27 431L37 432Q47 433 65 434T102 436Q119 437 138 438T167 441T178 442H181V402Q181 364 182 364T187 369T199 384T218 402T247 421T285 437Q305 442 336 442Q450 438 463 329Q464 322 464 190V104Q464 66 466 59T477 49Q498 46 526 46H542V0H534L510 1Q487 2 460 2T422 3Q319 3 310 0H302V46H318Q379 46 379 62Q380 64 380 200Q379 335 378 343Q372 371 358 385T334 402T308 404Q263 404 229 370Q202 343 195 315T187 232V168V108Q187 78 188 68T191 55T200 49Q221 46 249 46H265V0H257L234 1Q210 2 183 2T145 3Q42 3 33 0H25V46H41'],
  
      // LATIN SMALL LETTER O
      0x6F: [448,10,500,28,471,'28 214Q28 309 93 378T250 448Q340 448 405 380T471 215Q471 120 407 55T250 -10Q153 -10 91 57T28 214ZM250 30Q372 30 372 193V225V250Q372 272 371 288T364 326T348 362T317 390T268 410Q263 411 252 411Q222 411 195 399Q152 377 139 338T126 246V226Q126 130 145 91Q177 30 250 30'],
  
      // LATIN SMALL LETTER P
      0x70: [443,194,556,20,522,'36 -148H50Q89 -148 97 -134V-126Q97 -119 97 -107T97 -77T98 -38T98 6T98 55T98 106Q98 140 98 177T98 243T98 296T97 335T97 351Q94 370 83 376T38 385H20V408Q20 431 22 431L32 432Q42 433 61 434T98 436Q115 437 135 438T165 441T176 442H179V416L180 390L188 397Q247 441 326 441Q407 441 464 377T522 216Q522 115 457 52T310 -11Q242 -11 190 33L182 40V-45V-101Q182 -128 184 -134T195 -145Q216 -148 244 -148H260V-194H252L228 -193Q205 -192 178 -192T140 -191Q37 -191 28 -194H20V-148H36ZM424 218Q424 292 390 347T305 402Q234 402 182 337V98Q222 26 294 26Q345 26 384 80T424 218'],
  
      // LATIN SMALL LETTER Q
      0x71: [442,194,528,33,535,'33 218Q33 308 95 374T236 441H246Q330 441 381 372L387 364Q388 364 404 403L420 442H457V156Q457 -132 458 -134Q462 -142 470 -145Q491 -148 519 -148H535V-194H527L504 -193Q480 -192 453 -192T415 -191Q312 -191 303 -194H295V-148H311Q339 -148 360 -145Q369 -141 371 -135T373 -106V-41V49Q313 -11 236 -11Q154 -11 94 53T33 218ZM376 300Q346 389 278 401Q275 401 269 401T261 402Q211 400 171 350T131 214Q131 137 165 82T253 27Q296 27 328 54T376 118V300'],
  
      // LATIN SMALL LETTER R
      0x72: [443,0,392,20,364,'36 46H50Q89 46 97 60V68Q97 77 97 91T98 122T98 161T98 203Q98 234 98 269T98 328L97 351Q94 370 83 376T38 385H20V408Q20 431 22 431L32 432Q42 433 60 434T96 436Q112 437 131 438T160 441T171 442H174V373Q213 441 271 441H277Q322 441 343 419T364 373Q364 352 351 337T313 322Q288 322 276 338T263 372Q263 381 265 388T270 400T273 405Q271 407 250 401Q234 393 226 386Q179 341 179 207V154Q179 141 179 127T179 101T180 81T180 66V61Q181 59 183 57T188 54T193 51T200 49T207 48T216 47T225 47T235 46T245 46H276V0H267Q249 3 140 3Q37 3 28 0H20V46H36'],
  
      // LATIN SMALL LETTER S
      0x73: [448,11,394,33,359,'295 316Q295 356 268 385T190 414Q154 414 128 401Q98 382 98 349Q97 344 98 336T114 312T157 287Q175 282 201 278T245 269T277 256Q294 248 310 236T342 195T359 133Q359 71 321 31T198 -10H190Q138 -10 94 26L86 19L77 10Q71 4 65 -1L54 -11H46H42Q39 -11 33 -5V74V132Q33 153 35 157T45 162H54Q66 162 70 158T75 146T82 119T101 77Q136 26 198 26Q295 26 295 104Q295 133 277 151Q257 175 194 187T111 210Q75 227 54 256T33 318Q33 357 50 384T93 424T143 442T187 447H198Q238 447 268 432L283 424L292 431Q302 440 314 448H322H326Q329 448 335 442V310L329 304H301Q295 310 295 316'],
  
      // LATIN SMALL LETTER T
      0x74: [615,10,389,18,333,'27 422Q80 426 109 478T141 600V615H181V431H316V385H181V241Q182 116 182 100T189 68Q203 29 238 29Q282 29 292 100Q293 108 293 146V181H333V146V134Q333 57 291 17Q264 -10 221 -10Q187 -10 162 2T124 33T105 68T98 100Q97 107 97 248V385H18V422H27'],
  
      // LATIN SMALL LETTER U
      0x75: [443,11,556,25,542,'383 58Q327 -10 256 -10H249Q124 -10 105 89Q104 96 103 226Q102 335 102 348T96 369Q86 385 36 385H25V408Q25 431 27 431L38 432Q48 433 67 434T105 436Q122 437 142 438T172 441T184 442H187V261Q188 77 190 64Q193 49 204 40Q224 26 264 26Q290 26 311 35T343 58T363 90T375 120T379 144Q379 145 379 161T380 201T380 248V315Q380 361 370 372T320 385H302V431Q304 431 378 436T457 442H464V264Q464 84 465 81Q468 61 479 55T524 46H542V0Q540 0 467 -5T390 -11H383V58'],
  
      // LATIN SMALL LETTER V
      0x76: [431,11,528,19,508,'338 431Q344 429 422 429Q479 429 503 431H508V385H497Q439 381 423 345Q421 341 356 172T288 -2Q283 -11 263 -11Q244 -11 239 -2Q99 359 98 364Q93 378 82 381T43 385H19V431H25L33 430Q41 430 53 430T79 430T104 429T122 428Q217 428 232 431H240V385H226Q187 384 184 370Q184 366 235 234L286 102L377 341V349Q377 363 367 372T349 383T335 385H331V431H338'],
  
      // LATIN SMALL LETTER W
      0x77: [431,11,722,18,703,'90 368Q84 378 76 380T40 385H18V431H24L43 430Q62 430 84 429T116 428Q206 428 221 431H229V385H215Q177 383 177 368Q177 367 221 239L265 113L339 328L333 345Q323 374 316 379Q308 384 278 385H258V431H264Q270 428 348 428Q439 428 454 431H461V385H452Q404 385 404 369Q404 366 418 324T449 234T481 143L496 100L537 219Q579 341 579 347Q579 363 564 373T530 385H522V431H529Q541 428 624 428Q692 428 698 431H703V385H697Q696 385 691 385T682 384Q635 377 619 334L559 161Q546 124 528 71Q508 12 503 1T487 -11H479Q460 -11 456 -4Q455 -3 407 133L361 267Q359 263 266 -4Q261 -11 243 -11H238Q225 -11 220 -3L90 368'],
  
      // LATIN SMALL LETTER X
      0x78: [431,0,528,11,516,'201 0Q189 3 102 3Q26 3 17 0H11V46H25Q48 47 67 52T96 61T121 78T139 96T160 122T180 150L226 210L168 288Q159 301 149 315T133 336T122 351T113 363T107 370T100 376T94 379T88 381T80 383Q74 383 44 385H16V431H23Q59 429 126 429Q219 429 229 431H237V385Q201 381 201 369Q201 367 211 353T239 315T268 274L272 270L297 304Q329 345 329 358Q329 364 327 369T322 376T317 380T310 384L307 385H302V431H309Q324 428 408 428Q487 428 493 431H499V385H492Q443 385 411 368Q394 360 377 341T312 257L296 236L358 151Q424 61 429 57T446 50Q464 46 499 46H516V0H510H502Q494 1 482 1T457 2T432 2T414 3Q403 3 377 3T327 1L304 0H295V46H298Q309 46 320 51T331 63Q331 65 291 120L250 175Q249 174 219 133T185 88Q181 83 181 74Q181 63 188 55T206 46Q208 46 208 23V0H201'],
  
      // LATIN SMALL LETTER Y
      0x79: [431,204,528,19,508,'69 -66Q91 -66 104 -80T118 -116Q118 -134 109 -145T91 -160Q84 -163 97 -166Q104 -168 111 -168Q131 -168 148 -159T175 -138T197 -106T213 -75T225 -43L242 0L170 183Q150 233 125 297Q101 358 96 368T80 381Q79 382 78 382Q66 385 34 385H19V431H26L46 430Q65 430 88 429T122 428Q129 428 142 428T171 429T200 430T224 430L233 431H241V385H232Q183 385 185 366L286 112Q286 113 332 227L376 341V350Q376 365 366 373T348 383T334 385H331V431H337H344Q351 431 361 431T382 430T405 429T422 429Q477 429 503 431H508V385H497Q441 380 422 345Q420 343 378 235T289 9T227 -131Q180 -204 113 -204Q69 -204 44 -177T19 -116Q19 -89 35 -78T69 -66'],
  
      // LATIN SMALL LETTER Z
      0x7A: [431,0,444,28,401,'42 263Q44 270 48 345T53 423V431H393Q399 425 399 415Q399 403 398 402L381 378Q364 355 331 309T265 220L134 41L182 40H206Q254 40 283 46T331 77Q352 105 359 185L361 201Q361 202 381 202H401V196Q401 195 393 103T384 6V0H209L34 1L31 3Q28 8 28 17Q28 30 29 31T160 210T294 394H236Q169 393 152 388Q127 382 113 367Q89 344 82 264V255H42V263'],
  
      // LEFT CURLY BRACKET
      0x7B: [750,250,500,65,434,'434 -231Q434 -244 428 -250H410Q281 -250 230 -184Q225 -177 222 -172T217 -161T213 -148T211 -133T210 -111T209 -84T209 -47T209 0Q209 21 209 53Q208 142 204 153Q203 154 203 155Q189 191 153 211T82 231Q71 231 68 234T65 250T68 266T82 269Q116 269 152 289T203 345Q208 356 208 377T209 529V579Q209 634 215 656T244 698Q270 724 324 740Q361 748 377 749Q379 749 390 749T408 750H428Q434 744 434 732Q434 719 431 716Q429 713 415 713Q362 710 332 689T296 647Q291 634 291 499V417Q291 370 288 353T271 314Q240 271 184 255L170 250L184 245Q202 239 220 230T262 196T290 137Q291 131 291 1Q291 -134 296 -147Q306 -174 339 -192T415 -213Q429 -213 431 -216Q434 -219 434 -231'],
  
      // VERTICAL LINE
      0x7C: [750,249,278,119,159,'139 -249H137Q125 -249 119 -235V251L120 737Q130 750 139 750Q152 750 159 735V-235Q151 -249 141 -249H139'],
  
      // RIGHT CURLY BRACKET
      0x7D: [750,250,500,65,434,'65 731Q65 745 68 747T88 750Q171 750 216 725T279 670Q288 649 289 635T291 501Q292 362 293 357Q306 312 345 291T417 269Q428 269 431 266T434 250T431 234T417 231Q380 231 345 210T298 157Q293 143 292 121T291 -28V-79Q291 -134 285 -156T256 -198Q202 -250 89 -250Q71 -250 68 -247T65 -230Q65 -224 65 -223T66 -218T69 -214T77 -213Q91 -213 108 -210T146 -200T183 -177T207 -139Q208 -134 209 3L210 139Q223 196 280 230Q315 247 330 250Q305 257 280 270Q225 304 212 352L210 362L209 498Q208 635 207 640Q195 680 154 696T77 713Q68 713 67 716T65 731'],
  
      // DIAERESIS
      0xA8: [669,-554,500,95,405,'95 612Q95 633 112 651T153 669T193 652T210 612Q210 588 194 571T152 554L127 560Q95 577 95 612ZM289 611Q289 634 304 649T335 668Q336 668 340 668T346 669Q369 669 386 652T404 612T387 572T346 554Q323 554 306 570T289 611'],
  
      // NOT SIGN
      0xAC: [356,-89,667,56,611,'56 323T56 336T70 356H596Q603 353 611 343V102Q598 89 591 89Q587 89 584 90T579 94T575 98T572 102L571 209V316H70Q56 323 56 336'],
  
      // MACRON
      0xAF: [590,-544,500,69,430,'69 544V590H430V544H69'],
  
      // DEGREE SIGN
      0xB0: [715,-542,500,147,352,'147 628Q147 669 179 692T244 715Q298 715 325 689T352 629Q352 592 323 567T249 542Q202 542 175 567T147 628ZM313 628Q313 660 300 669T259 678H253Q248 678 242 678T234 679Q217 679 207 674T192 659T188 644T187 629Q187 600 198 590Q210 579 250 579H265Q279 579 288 581T305 595T313 628'],
  
      // PLUS-MINUS SIGN
      0xB1: [666,0,778,56,722,'56 320T56 333T70 353H369V502Q369 651 371 655Q376 666 388 666Q402 666 405 654T409 596V500V353H707Q722 345 722 333Q722 320 707 313H409V40H707Q722 32 722 20T707 0H70Q56 7 56 20T70 40H369V313H70Q56 320 56 333'],
  
      // ACUTE ACCENT
      0xB4: [699,-505,500,203,393,'349 699Q367 699 380 686T393 656Q393 651 392 647T387 637T380 627T367 616T351 602T330 585T303 563L232 505L217 519Q203 533 204 533Q204 534 229 567T282 636T313 678L316 681Q318 684 321 686T328 692T337 697T349 699'],
  
      // MULTIPLICATION SIGN
      0xD7: [491,-9,778,147,630,'630 29Q630 9 609 9Q604 9 587 25T493 118L389 222L284 117Q178 13 175 11Q171 9 168 9Q160 9 154 15T147 29Q147 36 161 51T255 146L359 250L255 354Q174 435 161 449T147 471Q147 480 153 485T168 490Q173 490 175 489Q178 487 284 383L389 278L493 382Q570 459 587 475T609 491Q630 491 630 471Q630 464 620 453T522 355L418 250L522 145Q606 61 618 48T630 29'],
  
      // DIVISION SIGN
      0xF7: [537,36,778,56,721,'318 466Q318 500 339 518T386 537Q418 537 438 517T458 466Q458 438 440 417T388 396Q355 396 337 417T318 466ZM56 237T56 250T70 270H706Q721 262 721 250T706 230H70Q56 237 56 250ZM318 34Q318 68 339 86T386 105Q418 105 438 85T458 34Q458 6 440 -15T388 -36Q355 -36 337 -15T318 34'],
  
      // MODIFIER LETTER CIRCUMFLEX ACCENT
      0x2C6: [694,-531,500,112,387,'112 560L249 694L257 686Q387 562 387 560L361 531Q359 532 303 581L250 627L195 580Q182 569 169 557T148 538L140 532Q138 530 125 546L112 560'],
  
      // CARON
      0x2C7: [644,-513,500,114,385,'114 611L127 630L136 644Q138 644 193 612Q248 581 250 581L306 612Q361 644 363 644L385 611L318 562L249 513L114 611'],
  
      // MODIFIER LETTER MACRON
      0x2C9: [590,-544,500,69,430,'69 544V590H430V544H69'],
  
      // MODIFIER LETTER ACUTE ACCENT
      0x2CA: [699,-505,500,203,393,'349 699Q367 699 380 686T393 656Q393 651 392 647T387 637T380 627T367 616T351 602T330 585T303 563L232 505L217 519Q203 533 204 533Q204 534 229 567T282 636T313 678L316 681Q318 684 321 686T328 692T337 697T349 699'],
  
      // MODIFIER LETTER GRAVE ACCENT
      0x2CB: [699,-505,500,106,296,'106 655Q106 671 119 685T150 699Q166 699 177 688Q190 671 222 629T275 561T295 533T282 519L267 505L196 563Q119 626 113 634Q106 643 106 655'],
  
      // BREVE
      0x2D8: [694,-515,500,92,407,'250 515Q179 515 138 565T92 683V694H129V689Q129 688 129 683T130 675Q137 631 169 599T248 567Q304 567 337 608T370 689V694H407V683Q403 617 361 566T250 515'],
  
      // DOT ABOVE
      0x2D9: [669,-549,500,190,309,'190 609Q190 637 208 653T252 669Q275 667 292 652T309 609Q309 579 292 564T250 549Q225 549 208 564T190 609'],
  
      // SMALL TILDE
      0x2DC: [668,-565,500,83,416,'179 601Q164 601 151 595T131 584T111 565L97 577L83 588Q83 589 95 603T121 633T142 654Q165 668 187 668T253 650T320 632Q335 632 348 638T368 649T388 668L402 656L416 645Q375 586 344 572Q330 565 313 565Q292 565 248 583T179 601'],
  
      // EN DASH
      0x2013: [285,-248,500,0,499,'0 248V285H499V248H0'],
  
      // EM DASH
      0x2014: [285,-248,1000,0,999,'0 248V285H999V248H0'],
  
      // LEFT SINGLE QUOTATION MARK
      0x2018: [694,-379,278,64,199,'64 494Q64 548 86 597T131 670T160 694Q163 694 172 685T182 672Q182 669 170 656T144 625T116 573T101 501Q101 489 102 489T107 491T120 497T138 500Q163 500 180 483T198 440T181 397T139 379Q110 379 87 405T64 494'],
  
      // RIGHT SINGLE QUOTATION MARK
      0x2019: [694,-379,278,78,212,'78 634Q78 659 95 676T138 694Q166 694 189 668T212 579Q212 525 190 476T146 403T118 379Q114 379 105 388T95 401Q95 404 107 417T133 448T161 500T176 572Q176 584 175 584T170 581T157 576T139 573Q114 573 96 590T78 634'],
  
      // LEFT DOUBLE QUOTATION MARK
      0x201C: [694,-379,500,128,466,'128 494Q128 528 137 560T158 616T185 658T209 685T223 694T236 685T245 670Q244 668 231 654T204 622T178 571T164 501Q164 489 165 489T170 491T183 497T201 500Q226 500 244 483T262 440T245 397T202 379Q173 379 151 405T128 494ZM332 494Q332 528 341 560T362 616T389 658T413 685T427 694T439 685T449 672Q449 669 437 656T411 625T383 573T368 501Q368 489 369 489T374 491T387 497T405 500Q430 500 448 483T466 440T449 397T406 379Q377 379 355 405T332 494'],
  
      // RIGHT DOUBLE QUOTATION MARK
      0x201D: [694,-379,500,34,372,'34 634Q34 659 50 676T93 694Q121 694 144 668T168 579Q168 525 146 476T101 403T73 379Q69 379 60 388T50 401Q50 404 62 417T88 448T116 500T131 572Q131 584 130 584T125 581T112 576T94 573Q69 573 52 590T34 634ZM238 634Q238 659 254 676T297 694Q325 694 348 668T372 579Q372 525 350 476T305 403T277 379Q273 379 264 388T254 401Q254 404 266 417T292 448T320 500T335 572Q335 584 334 584T329 581T316 576T298 573Q273 573 256 590T238 634'],
  
      // DAGGER
      0x2020: [705,216,444,54,389,'182 675Q195 705 222 705Q234 705 243 700T253 691T263 675L262 655Q262 620 252 549T240 454V449Q250 451 288 461T346 472T377 461T389 431Q389 417 379 404T346 390Q327 390 288 401T243 412H240V405Q245 367 250 339T258 301T261 274T263 225Q263 124 255 -41T239 -213Q236 -216 222 -216H217Q206 -216 204 -212T200 -186Q199 -175 199 -168Q181 38 181 225Q181 265 182 280T191 327T204 405V412H201Q196 412 157 401T98 390Q76 390 66 403T55 431T65 458T98 472Q116 472 155 462T205 449Q204 452 204 460T201 490T193 547Q182 619 182 655V675'],
  
      // DOUBLE DAGGER
      0x2021: [705,205,444,54,389,'181 658Q181 705 222 705T263 658Q263 633 252 572T240 497Q240 496 241 496Q243 496 285 507T345 519Q365 519 376 508T388 478Q388 466 384 458T375 447T361 438H344Q318 438 282 448T241 459Q240 458 240 456Q240 449 251 384T263 297Q263 278 255 267T238 253T222 250T206 252T190 266T181 297Q181 323 192 383T204 458Q204 459 203 459Q198 459 162 449T101 438H84Q74 443 70 446T61 457T56 478Q56 497 67 508T99 519Q117 519 159 508T203 496Q204 496 204 499Q204 507 193 572T181 658ZM181 202Q181 249 222 249T263 202Q263 185 259 161T249 103T240 48V41H243Q248 41 287 52T346 63T377 52T389 22Q389 8 379 -5T346 -19Q327 -19 288 -8T243 3H240V-4Q243 -24 249 -58T259 -117T263 -158Q263 -177 255 -188T238 -202T222 -205T206 -203T190 -189T181 -158Q181 -141 185 -117T195 -59T204 -4V3H201Q196 3 157 -8T98 -19Q76 -19 66 -6T55 22T65 49T98 63Q117 63 156 52T201 41H204V48Q201 68 195 102T185 161T181 202'],
  
      // HORIZONTAL ELLIPSIS
      0x2026: [120,0,1172,78,1093,'78 60Q78 84 95 102T138 120Q162 120 180 104T199 61Q199 36 182 18T139 0T96 17T78 60ZM525 60Q525 84 542 102T585 120Q609 120 627 104T646 61Q646 36 629 18T586 0T543 17T525 60ZM972 60Q972 84 989 102T1032 120Q1056 120 1074 104T1093 61Q1093 36 1076 18T1033 0T990 17T972 60'],
  
      // PRIME
      0x2032: [560,-43,275,30,262,'79 43Q73 43 52 49T30 61Q30 68 85 293T146 528Q161 560 198 560Q218 560 240 545T262 501Q262 496 260 486Q259 479 173 263T84 45T79 43'],
  
      // COMBINING RIGHT ARROW ABOVE
      0x20D7: [714,-516,0,-471,-29,'-123 694Q-123 702 -118 708T-103 714Q-93 714 -88 706T-80 687T-67 660T-40 633Q-29 626 -29 615Q-29 606 -36 600T-53 590T-83 571T-121 531Q-135 516 -143 516T-157 522T-163 536T-152 559T-129 584T-116 595H-287L-458 596Q-459 597 -461 599T-466 602T-469 607T-471 615Q-471 622 -458 635H-99Q-123 673 -123 694'],
  
      // LEFTWARDS ARROW
      0x2190: [511,11,1000,55,944,'944 261T944 250T929 230H165Q167 228 182 216T211 189T244 152T277 96T303 25Q308 7 308 0Q308 -11 288 -11Q281 -11 278 -11T272 -7T267 2T263 21Q245 94 195 151T73 236Q58 242 55 247Q55 254 59 257T73 264Q121 283 158 314T215 375T247 434T264 480L267 497Q269 503 270 505T275 509T288 511Q308 511 308 500Q308 493 303 475Q293 438 278 406T246 352T215 315T185 287T165 270H929Q944 261 944 250'],
  
      // UPWARDS ARROW
      0x2191: [694,193,500,17,483,'27 414Q17 414 17 433Q17 437 17 439T17 444T19 447T20 450T22 452T26 453T30 454T36 456Q80 467 120 494T180 549Q227 607 238 678Q240 694 251 694Q259 694 261 684Q261 677 265 659T284 608T320 549Q340 525 363 507T405 479T440 463T467 455T479 451Q483 447 483 433Q483 413 472 413Q467 413 458 416Q342 448 277 545L270 555V-179Q262 -193 252 -193H250H248Q236 -193 230 -179V555L223 545Q192 499 146 467T70 424T27 414'],
  
      // RIGHTWARDS ARROW
      0x2192: [511,11,1000,56,944,'56 237T56 250T70 270H835Q719 357 692 493Q692 494 692 496T691 499Q691 511 708 511H711Q720 511 723 510T729 506T732 497T735 481T743 456Q765 389 816 336T935 261Q944 258 944 250Q944 244 939 241T915 231T877 212Q836 186 806 152T761 85T740 35T732 4Q730 -6 727 -8T711 -11Q691 -11 691 0Q691 7 696 25Q728 151 835 230H70Q56 237 56 250'],
  
      // DOWNWARDS ARROW
      0x2193: [694,194,500,17,483,'473 86Q483 86 483 67Q483 63 483 61T483 56T481 53T480 50T478 48T474 47T470 46T464 44Q428 35 391 14T316 -55T264 -168Q264 -170 263 -173T262 -180T261 -184Q259 -194 251 -194Q242 -194 238 -176T221 -121T180 -49Q169 -34 155 -21T125 2T95 20T67 33T44 42T27 47L21 49Q17 53 17 67Q17 87 28 87Q33 87 42 84Q158 52 223 -45L230 -55V312Q230 391 230 482T229 591Q229 662 231 676T243 693Q244 694 251 694Q264 692 270 679V-55L277 -45Q307 1 353 33T430 76T473 86'],
  
      // LEFT RIGHT ARROW
      0x2194: [511,11,1000,55,944,'263 479Q267 501 271 506T288 511Q308 511 308 500Q308 493 303 475Q293 438 278 406T246 352T215 315T185 287T165 270H835Q729 349 696 475Q691 493 691 500Q691 511 711 511Q720 511 723 510T729 506T732 497T735 481T743 456Q765 389 816 336T935 261Q944 258 944 250Q944 244 939 241T915 231T877 212Q836 186 806 152T761 85T740 35T732 4Q730 -6 727 -8T711 -11Q691 -11 691 0Q691 7 696 25Q728 151 835 230H165Q167 228 182 216T211 189T244 152T277 96T303 25Q308 7 308 0Q308 -11 288 -11Q281 -11 278 -11T272 -7T267 2T263 21Q245 94 195 151T73 236Q58 242 55 247Q55 254 59 257T73 264Q144 292 194 349T263 479'],
  
      // UP DOWN ARROW
      0x2195: [772,272,500,17,483,'27 492Q17 492 17 511Q17 515 17 517T17 522T19 525T20 528T22 530T26 531T30 532T36 534Q80 545 120 572T180 627Q210 664 223 701T238 755T250 772T261 762Q261 757 264 741T282 691T319 628Q352 589 390 566T454 536L479 529Q483 525 483 511Q483 491 472 491Q467 491 458 494Q342 526 277 623L270 633V-133L277 -123Q307 -77 353 -45T430 -2T473 8Q483 8 483 -11Q483 -15 483 -17T483 -22T481 -25T480 -28T478 -30T474 -31T470 -32T464 -34Q407 -49 364 -84T300 -157T270 -223T261 -262Q259 -272 250 -272Q242 -272 239 -255T223 -201T180 -127Q169 -112 155 -99T125 -76T95 -58T67 -45T44 -36T27 -31L21 -29Q17 -25 17 -11Q17 9 28 9Q33 9 42 6Q158 -26 223 -123L230 -133V633L223 623Q192 577 146 545T70 502T27 492'],
  
      // NORTH WEST ARROW
      0x2196: [720,195,1000,29,944,'204 662Q257 662 301 676T369 705T394 720Q398 720 407 711T417 697Q417 688 389 671T310 639T212 623Q176 623 153 628Q151 628 221 557T546 232Q942 -164 943 -168Q944 -170 944 -174Q944 -182 938 -188T924 -195Q922 -195 916 -193Q912 -191 517 204Q440 281 326 394T166 553L121 598Q126 589 126 541Q126 438 70 349Q59 332 52 332Q48 332 39 341T29 355Q29 358 38 372T57 407T77 464T86 545Q86 583 78 614T63 663T55 683Q55 693 65 693Q73 693 82 688Q136 662 204 662'],
  
      // NORTH EAST ARROW
      0x2197: [720,195,1000,55,971,'582 697Q582 701 591 710T605 720Q607 720 630 706T697 677T795 662Q830 662 863 670T914 686T934 694Q942 694 944 685Q944 680 936 663T921 615T913 545Q913 490 927 446T956 379T970 355Q970 351 961 342T947 332Q940 332 929 349Q874 436 874 541Q874 590 878 598L832 553Q787 508 673 395T482 204Q87 -191 83 -193Q77 -195 75 -195Q67 -195 61 -189T55 -174Q55 -170 56 -168Q58 -164 453 232Q707 487 777 557T847 628Q824 623 787 623Q689 623 599 679Q582 690 582 697'],
  
      // SOUTH EAST ARROW
      0x2198: [695,220,1000,55,970,'55 675Q55 683 60 689T75 695Q77 695 83 693Q87 691 482 296Q532 246 605 174T717 62T799 -20T859 -80T878 -97Q874 -93 874 -41Q874 64 929 151Q940 168 947 168Q951 168 960 159T970 145Q970 143 956 121T928 54T913 -45Q913 -83 920 -114T936 -163T944 -185Q942 -194 934 -194Q932 -194 914 -186T864 -170T795 -162Q743 -162 698 -176T630 -205T605 -220Q601 -220 592 -211T582 -197Q582 -187 611 -170T691 -138T787 -123Q824 -123 847 -128Q848 -128 778 -57T453 268Q58 664 56 668Q55 670 55 675'],
  
      // SOUTH WEST ARROW
      0x2199: [695,220,1000,29,944,'126 -41Q126 -92 121 -97Q121 -98 139 -80T200 -20T281 61T394 173T517 296Q909 690 916 693Q922 695 924 695Q932 695 938 689T944 674Q944 670 943 668Q942 664 546 268Q292 13 222 -57T153 -128Q176 -123 212 -123Q310 -123 400 -179Q417 -190 417 -197Q417 -201 408 -210T394 -220Q392 -220 369 -206T302 -177T204 -162Q131 -162 67 -194Q63 -195 59 -192T55 -183Q55 -180 62 -163T78 -115T86 -45Q86 10 72 54T44 120T29 145Q29 149 38 158T52 168Q59 168 70 151Q126 62 126 -41'],
  
      // RIGHTWARDS ARROW FROM BAR
      0x21A6: [511,11,1000,54,944,'95 155V109Q95 83 92 73T75 63Q61 63 58 74T54 130Q54 140 54 180T55 250Q55 421 57 425Q61 437 75 437Q88 437 91 428T95 393V345V270H835Q719 357 692 493Q692 494 692 496T691 499Q691 511 708 511H711Q720 511 723 510T729 506T732 497T735 481T743 456Q765 389 816 336T935 261Q944 258 944 250Q944 244 939 241T915 231T877 212Q836 186 806 152T761 85T740 35T732 4Q730 -6 727 -8T711 -11Q691 -11 691 0Q691 7 696 25Q728 151 835 230H95V155'],
  
      // LEFTWARDS ARROW WITH HOOK
      0x21A9: [511,11,1126,55,1070,'903 424T903 444T929 464Q976 464 1023 434T1070 347Q1070 316 1055 292T1016 256T971 237T929 230H165Q167 228 182 216T211 189T244 152T277 96T303 25Q308 7 308 0Q308 -11 288 -11Q281 -11 278 -11T272 -7T267 2T263 21Q245 94 195 151T73 236Q58 242 55 247Q55 254 59 257T73 264Q121 283 158 314T215 375T247 434T264 480L267 497Q269 503 270 505T275 509T288 511Q308 511 308 500Q308 493 303 475Q293 438 278 406T246 352T215 315T185 287T165 270H926Q929 270 941 271T960 275T978 280T998 290T1015 307Q1030 325 1030 347Q1030 355 1027 364T1014 387T983 411T929 424H928Q903 424 903 444'],
  
      // RIGHTWARDS ARROW WITH HOOK
      0x21AA: [511,11,1126,55,1070,'55 347Q55 380 72 404T113 441T159 458T197 464Q222 464 222 444Q222 429 204 426T157 417T110 387Q95 369 95 347Q95 339 98 330T111 307T142 283T196 270H961Q845 357 818 493Q818 494 818 496T817 499Q817 511 834 511H837Q846 511 849 510T855 506T858 497T861 481T869 456Q891 389 942 336T1061 261Q1070 258 1070 250Q1070 244 1065 241T1041 231T1003 212Q962 186 932 152T887 85T866 35T858 4Q856 -6 853 -8T837 -11Q817 -11 817 0Q817 7 822 25Q854 151 961 230H196Q149 230 102 260T55 347'],
  
      // LEFTWARDS HARPOON WITH BARB UPWARDS
      0x21BC: [511,-230,1000,55,944,'62 230Q56 236 55 244Q55 252 57 255T69 265Q114 292 151 326T208 391T243 448T265 491T273 509Q276 511 288 511Q304 511 306 505Q309 501 303 484Q293 456 279 430T251 383T223 344T196 313T173 291T156 276L148 270H929Q944 261 944 250T929 230H62'],
  
      // LEFTWARDS HARPOON WITH BARB DOWNWARDS
      0x21BD: [270,11,1000,55,944,'55 256Q56 264 62 270H929Q944 261 944 250T929 230H148Q149 229 165 215T196 185T231 145T270 87T303 16Q309 -1 306 -5Q304 -11 288 -11Q279 -11 276 -10T269 -4T264 10T253 36T231 75Q172 173 69 235Q59 242 57 245T55 256'],
  
      // RIGHTWARDS HARPOON WITH BARB UPWARDS
      0x21C0: [511,-230,1000,56,945,'691 500Q691 511 711 511Q720 511 723 510T730 504T735 490T746 464T768 425Q796 378 835 339T897 285T933 263Q941 258 942 256T944 245T937 230H70Q56 237 56 250T70 270H852Q802 308 762 364T707 455T691 500'],
  
      // RIGHTWARDS HARPOON WITH BARB DOWNWARDS
      0x21C1: [270,11,1000,56,944,'56 237T56 250T70 270H937Q944 263 944 256Q944 251 944 250T943 246T940 242T933 238Q794 153 734 7Q729 -7 726 -9T711 -11Q695 -11 693 -5Q690 -1 696 16Q721 84 763 139T852 230H70Q56 237 56 250'],
  
      // RIGHTWARDS HARPOON OVER LEFTWARDS HARPOON
      0x21CC: [671,11,1000,55,945,'691 660Q691 671 711 671Q720 671 723 670T730 664T735 650T746 624T768 585Q797 538 836 499T897 445T933 423Q941 418 942 416T944 405T937 390H70Q56 397 56 410T70 430H852Q802 468 762 524T707 615T691 660ZM55 256Q56 264 62 270H929Q944 261 944 250T929 230H148Q149 229 165 215T196 185T231 145T270 87T303 16Q309 -1 306 -5Q304 -11 288 -11Q279 -11 276 -10T269 -4T264 10T253 36T231 75Q172 173 69 235Q59 242 57 245T55 256'],
  
      // LEFTWARDS DOUBLE ARROW
      0x21D0: [525,24,1000,56,945,'944 153Q944 140 929 133H318L328 123Q379 69 414 0Q419 -13 419 -17Q419 -24 399 -24Q388 -24 385 -23T377 -12Q332 77 253 144T72 237Q62 240 59 242T56 250T59 257T70 262T89 268T119 278T160 296Q303 366 377 512Q382 522 385 523T401 525Q419 524 419 515Q419 510 414 500Q379 431 328 377L318 367H929Q944 359 944 347Q944 336 930 328L602 327H274L264 319Q225 289 147 250Q148 249 165 241T210 217T264 181L274 173H930Q931 172 933 171T936 169T938 167T941 164T942 162T943 158T944 153'],
  
      // UPWARDS DOUBLE ARROW
      0x21D1: [694,194,611,31,579,'228 -179Q227 -180 226 -182T223 -186T221 -189T218 -192T214 -193T208 -194Q196 -194 189 -181L188 125V430L176 419Q122 369 59 338Q46 330 40 330Q38 330 31 337V350Q31 362 33 365T46 374Q60 381 77 390T128 426T190 484T247 567T292 677Q295 688 298 692Q302 694 305 694Q313 694 318 677Q334 619 363 568T420 485T481 427T532 391T564 374Q575 368 577 365T579 350V337Q572 330 570 330Q564 330 551 338Q487 370 435 419L423 430L422 125V-181Q409 -194 401 -194Q397 -194 394 -193T388 -189T385 -184T382 -180V-177V475L373 487Q331 541 305 602Q304 601 300 591T290 571T278 548T260 519T238 488L229 476L228 148V-179'],
  
      // RIGHTWARDS DOUBLE ARROW
      0x21D2: [525,24,1000,56,944,'580 514Q580 525 596 525Q601 525 604 525T609 525T613 524T615 523T617 520T619 517T622 512Q659 438 720 381T831 300T927 263Q944 258 944 250T935 239T898 228T840 204Q696 134 622 -12Q618 -21 615 -22T600 -24Q580 -24 580 -17Q580 -13 585 0Q620 69 671 123L681 133H70Q56 140 56 153Q56 168 72 173H725L735 181Q774 211 852 250Q851 251 834 259T789 283T735 319L725 327H72Q56 332 56 347Q56 360 70 367H681L671 377Q638 412 609 458T580 514'],
  
      // DOWNWARDS DOUBLE ARROW
      0x21D3: [694,194,611,31,579,'401 694Q412 694 422 681V375L423 70L435 81Q487 130 551 162Q564 170 570 170Q572 170 579 163V150Q579 138 577 135T564 126Q541 114 518 99T453 48T374 -46T318 -177Q313 -194 305 -194T293 -178T272 -119T225 -31Q158 70 46 126Q35 132 33 135T31 150V163Q38 170 40 170Q46 170 59 162Q122 131 176 81L188 70V375L189 681Q199 694 208 694Q219 694 228 680V352L229 25L238 12Q279 -42 305 -102Q344 -23 373 13L382 25V678Q387 692 401 694'],
  
      // LEFT RIGHT DOUBLE ARROW
      0x21D4: [526,25,1000,33,966,'308 524Q318 526 323 526Q340 526 340 514Q340 507 336 499Q326 476 314 454T292 417T274 391T260 374L255 368Q255 367 500 367Q744 367 744 368L739 374Q734 379 726 390T707 416T685 453T663 499Q658 511 658 515Q658 525 680 525Q687 524 690 523T695 519T701 507Q766 359 902 287Q921 276 939 269T961 259T966 250Q966 246 965 244T960 240T949 236T930 228T902 213Q763 137 701 -7Q697 -16 695 -19T690 -23T680 -25Q658 -25 658 -15Q658 -11 663 1Q673 24 685 46T707 83T725 109T739 126L744 132Q744 133 500 133Q255 133 255 132L260 126Q265 121 273 110T292 84T314 47T336 1Q341 -11 341 -15Q341 -25 319 -25Q312 -24 309 -23T304 -19T298 -7Q233 141 97 213Q83 221 70 227T51 235T41 239T35 243T34 250T35 256T40 261T51 265T70 273T97 287Q235 363 299 509Q305 522 308 524ZM792 319L783 327H216Q183 294 120 256L110 250L120 244Q173 212 207 181L216 173H783L792 181Q826 212 879 244L889 250L879 256Q826 288 792 319'],
  
      // UP DOWN DOUBLE ARROW
      0x21D5: [772,272,611,31,579,'290 755Q298 772 305 772T318 757T343 706T393 633Q431 588 473 558T545 515T579 497V484Q579 464 570 464Q564 464 550 470Q485 497 423 550L422 400V100L423 -50Q485 3 550 30Q565 36 570 36Q579 36 579 16V3Q575 -1 549 -12T480 -53T393 -132Q361 -172 342 -208T318 -258T305 -272T293 -258T268 -208T217 -132Q170 -80 128 -51T61 -12T31 3V16Q31 36 40 36Q46 36 61 30Q86 19 109 6T146 -18T173 -38T188 -50V550Q186 549 173 539T147 519T110 495T61 470Q46 464 40 464Q31 464 31 484V497Q34 500 63 513T135 557T217 633Q267 692 290 755ZM374 598Q363 610 351 625T332 651T316 676T305 695L294 676Q282 657 267 636T236 598L228 589V-89L236 -98Q247 -110 259 -125T278 -151T294 -176T305 -195L316 -176Q328 -157 343 -136T374 -98L382 -89V589L374 598'],
  
      // FOR ALL
      0x2200: [694,22,556,0,556,'0 673Q0 684 7 689T20 694Q32 694 38 680T82 567L126 451H430L473 566Q483 593 494 622T512 668T519 685Q524 694 538 694Q556 692 556 674Q556 670 426 329T293 -15Q288 -22 278 -22T263 -15Q260 -11 131 328T0 673ZM414 410Q414 411 278 411T142 410L278 55L414 410'],
  
      // PARTIAL DIFFERENTIAL
      0x2202: [715,22,531,42,567,'202 508Q179 508 169 520T158 547Q158 557 164 577T185 624T230 675T301 710L333 715H345Q378 715 384 714Q447 703 489 661T549 568T566 457Q566 362 519 240T402 53Q321 -22 223 -22Q123 -22 73 56Q42 102 42 148V159Q42 276 129 370T322 465Q383 465 414 434T455 367L458 378Q478 461 478 515Q478 603 437 639T344 676Q266 676 223 612Q264 606 264 572Q264 547 246 528T202 508ZM430 306Q430 372 401 400T333 428Q270 428 222 382Q197 354 183 323T150 221Q132 149 132 116Q132 21 232 21Q244 21 250 22Q327 35 374 112Q389 137 409 196T430 306'],
  
      // THERE EXISTS
      0x2203: [694,0,556,56,500,'56 661T56 674T70 694H487Q497 686 500 679V15Q497 10 487 1L279 0H70Q56 7 56 20T70 40H460V327H84Q70 334 70 347T84 367H460V654H70Q56 661 56 674'],
  
      // EMPTY SET
      0x2205: [772,78,500,39,460,'331 696Q335 708 339 722T345 744T350 759T357 769T367 772Q374 772 381 767T388 754Q388 746 377 712L366 673L378 661Q460 575 460 344Q460 281 456 234T432 126T373 27Q319 -22 250 -22Q214 -22 180 -7Q168 -3 168 -4L159 -33Q148 -71 142 -75Q138 -78 132 -78Q124 -78 118 -72T111 -60Q111 -52 122 -18L133 21L125 29Q39 111 39 344Q39 596 137 675Q187 716 251 716Q265 716 278 714T296 710T315 703T331 696ZM276 676Q264 679 246 679Q196 679 159 631Q134 597 128 536T121 356Q121 234 127 174T151 80L234 366Q253 430 275 506T308 618L318 654Q318 656 294 669L276 676ZM181 42Q207 16 250 16Q291 16 324 47Q354 78 366 136T378 356Q378 470 372 528T349 616L348 613Q348 611 264 326L181 42'],
  
      // NABLA
      0x2207: [683,33,833,46,786,'46 676Q46 679 51 683H781Q786 679 786 676Q786 674 617 326T444 -26Q439 -33 416 -33T388 -26Q385 -22 216 326T46 676ZM697 596Q697 597 445 597T193 596Q195 591 319 336T445 80L697 596'],
  
      // ELEMENT OF
      0x2208: [541,41,667,84,583,'84 250Q84 372 166 450T360 539Q361 539 377 539T419 540T469 540H568Q583 532 583 520Q583 511 570 501L466 500Q355 499 329 494Q280 482 242 458T183 409T147 354T129 306T124 272V270H568Q583 262 583 250T568 230H124V228Q124 207 134 177T167 112T231 48T328 7Q355 1 466 0H570Q583 -10 583 -20Q583 -32 568 -40H471Q464 -40 446 -40T417 -41Q262 -41 172 45Q84 127 84 250'],
  
      // stix-negated (vert) set membership, variant
      0x2209: [716,215,667,84,584,'196 25Q84 109 84 250Q84 372 166 450T360 539Q361 539 375 539T413 540T460 540L547 707Q550 716 563 716Q570 716 575 712T581 703T583 696T505 540H568Q583 532 583 520Q583 511 570 501L484 500L366 270H568Q583 262 583 250T568 230H346L247 38Q284 16 328 7Q355 1 466 0H570Q583 -10 583 -20Q583 -32 568 -40H471Q464 -40 447 -40T419 -41Q304 -41 228 3Q117 -211 115 -212Q111 -215 104 -215T92 -212T86 -204T84 -197Q84 -190 89 -183L196 25ZM214 61L301 230H124V228Q124 196 147 147T214 61ZM321 270L440 500Q353 499 329 494Q280 482 242 458T183 409T147 354T129 306T124 272V270H321'],
  
      // CONTAINS AS MEMBER
      0x220B: [541,40,667,83,582,'83 520Q83 532 98 540H195Q202 540 220 540T249 541Q404 541 494 455Q582 374 582 250Q582 165 539 99T434 0T304 -39Q297 -40 195 -40H98Q83 -32 83 -20Q83 -10 96 0H200Q311 1 337 6Q369 14 401 28Q422 39 445 55Q484 85 508 127T537 191T542 228V230H98Q84 237 84 250T98 270H542V272Q542 280 539 295T527 336T497 391T445 445Q422 461 401 472Q386 479 374 483T347 491T325 495T298 498T273 499T239 500T200 500L96 501Q83 511 83 520'],
  
      // MINUS SIGN
      0x2212: [270,-230,778,84,694,'84 237T84 250T98 270H679Q694 262 694 250T679 230H98Q84 237 84 250'],
  
      // MINUS-OR-PLUS SIGN
      0x2213: [500,166,778,56,722,'56 467T56 480T70 500H707Q722 492 722 480T707 460H409V187H707Q722 179 722 167Q722 154 707 147H409V0V-93Q409 -144 406 -155T389 -166Q376 -166 372 -155T368 -105Q368 -96 368 -62T369 -2V147H70Q56 154 56 167T70 187H369V460H70Q56 467 56 480'],
  
      // DIVISION SLASH
      0x2215: [750,250,500,56,444,'423 750Q432 750 438 744T444 730Q444 725 271 248T92 -240Q85 -250 75 -250Q68 -250 62 -245T56 -231Q56 -221 230 257T407 740Q411 750 423 750'],
  
      // SET MINUS
      0x2216: [750,250,500,56,444,'56 731Q56 740 62 745T75 750Q85 750 92 740Q96 733 270 255T444 -231Q444 -239 438 -244T424 -250Q414 -250 407 -240Q404 -236 230 242T56 731'],
  
      // ASTERISK OPERATOR
      0x2217: [465,-35,500,64,435,'229 286Q216 420 216 436Q216 454 240 464Q241 464 245 464T251 465Q263 464 273 456T283 436Q283 419 277 356T270 286L328 328Q384 369 389 372T399 375Q412 375 423 365T435 338Q435 325 425 315Q420 312 357 282T289 250L355 219L425 184Q434 175 434 161Q434 146 425 136T401 125Q393 125 383 131T328 171L270 213Q283 79 283 63Q283 53 276 44T250 35Q231 35 224 44T216 63Q216 80 222 143T229 213L171 171Q115 130 110 127Q106 124 100 124Q87 124 76 134T64 161Q64 166 64 169T67 175T72 181T81 188T94 195T113 204T138 215T170 230T210 250L74 315Q65 324 65 338Q65 353 74 363T98 374Q106 374 116 368T171 328L229 286'],
  
      // RING OPERATOR
      0x2218: [444,-55,500,55,444,'55 251Q55 328 112 386T249 444T386 388T444 249Q444 171 388 113T250 55Q170 55 113 112T55 251ZM245 403Q188 403 142 361T96 250Q96 183 141 140T250 96Q284 96 313 109T354 135T375 160Q403 197 403 250Q403 313 360 358T245 403'],
  
      // BULLET OPERATOR
      0x2219: [444,-55,500,55,444,'55 251Q55 328 112 386T249 444T386 388T444 249Q444 171 388 113T250 55Q170 55 113 112T55 251'],
  
      // SQUARE ROOT
      0x221A: [800,200,833,71,853,'95 178Q89 178 81 186T72 200T103 230T169 280T207 309Q209 311 212 311H213Q219 311 227 294T281 177Q300 134 312 108L397 -77Q398 -77 501 136T707 565T814 786Q820 800 834 800Q841 800 846 794T853 782V776L620 293L385 -193Q381 -200 366 -200Q357 -200 354 -197Q352 -195 256 15L160 225L144 214Q129 202 113 190T95 178'],
  
      // PROPORTIONAL TO
      0x221D: [442,11,778,56,722,'56 124T56 216T107 375T238 442Q260 442 280 438T319 425T352 407T382 385T406 361T427 336T442 315T455 297T462 285L469 297Q555 442 679 442Q687 442 722 437V398H718Q710 400 694 400Q657 400 623 383T567 343T527 294T503 253T495 235Q495 231 520 192T554 143Q625 44 696 44Q717 44 719 46H722V-5Q695 -11 678 -11Q552 -11 457 141Q455 145 454 146L447 134Q362 -11 235 -11Q157 -11 107 56ZM93 213Q93 143 126 87T220 31Q258 31 292 48T349 88T389 137T413 178T421 196Q421 200 396 239T362 288Q322 345 288 366T213 387Q163 387 128 337T93 213'],
  
      // INFINITY
      0x221E: [442,11,1000,55,944,'55 217Q55 305 111 373T254 442Q342 442 419 381Q457 350 493 303L507 284L514 294Q618 442 747 442Q833 442 888 374T944 214Q944 128 889 59T743 -11Q657 -11 580 50Q542 81 506 128L492 147L485 137Q381 -11 252 -11Q166 -11 111 57T55 217ZM907 217Q907 285 869 341T761 397Q740 397 720 392T682 378T648 359T619 335T594 310T574 285T559 263T548 246L543 238L574 198Q605 158 622 138T664 94T714 61T765 51Q827 51 867 100T907 217ZM92 214Q92 145 131 89T239 33Q357 33 456 193L425 233Q364 312 334 337Q285 380 233 380Q171 380 132 331T92 214'],
  
      // ANGLE
      0x2220: [694,0,722,55,666,'71 0L68 2Q65 3 63 5T58 11T55 20Q55 22 57 28Q67 43 346 361Q397 420 474 508Q595 648 616 671T647 694T661 688T666 674Q666 668 663 663Q662 662 627 622T524 503T390 350L120 41L386 40H653Q666 30 666 20Q666 8 651 0H71'],
  
      // DIVIDES
      0x2223: [750,249,278,119,159,'139 -249H137Q125 -249 119 -235V251L120 737Q130 750 139 750Q152 750 159 735V-235Q151 -249 141 -249H139'],
  
      // PARALLEL TO
      0x2225: [750,250,500,132,368,'133 736Q138 750 153 750Q164 750 170 739Q172 735 172 250T170 -239Q164 -250 152 -250Q144 -250 138 -244L137 -243Q133 -241 133 -179T132 250Q132 731 133 736ZM329 739Q334 750 346 750Q353 750 361 744L362 743Q366 741 366 679T367 250T367 -178T362 -243L361 -244Q355 -250 347 -250Q335 -250 329 -239Q327 -235 327 250T329 739'],
  
      // LOGICAL AND
      0x2227: [598,22,667,55,611,'318 591Q325 598 333 598Q344 598 348 591Q349 590 414 445T545 151T611 -4Q609 -22 591 -22Q588 -22 586 -21T581 -20T577 -17T575 -13T572 -9T570 -4L333 528L96 -4Q87 -20 80 -21Q78 -22 75 -22Q57 -22 55 -4Q55 2 120 150T251 444T318 591'],
  
      // LOGICAL OR
      0x2228: [598,22,667,55,611,'55 580Q56 587 61 592T75 598Q86 598 96 580L333 48L570 580Q579 596 586 597Q588 598 591 598Q609 598 611 580Q611 574 546 426T415 132T348 -15Q343 -22 333 -22T318 -15Q317 -14 252 131T121 425T55 580'],
  
      // stix-intersection, serifs
      0x2229: [598,22,667,55,611,'88 -21T75 -21T55 -7V200Q55 231 55 280Q56 414 60 428Q61 430 61 431Q77 500 152 549T332 598Q443 598 522 544T610 405Q611 399 611 194V-7Q604 -22 591 -22Q582 -22 572 -9L570 405Q563 433 556 449T529 485Q498 519 445 538T334 558Q251 558 179 518T96 401Q95 396 95 193V-7Q88 -21 75 -21'],
  
      // stix-union, serifs
      0x222A: [598,22,667,55,611,'591 598H592Q604 598 611 583V376Q611 345 611 296Q610 162 606 148Q605 146 605 145Q586 68 507 23T333 -22Q268 -22 209 -1T106 66T56 173Q55 180 55 384L56 585Q66 598 75 598Q85 598 95 585V378L96 172L98 162Q112 95 181 57T332 18Q415 18 487 58T570 175Q571 180 571 383V583Q579 598 591 598'],
  
      // INTEGRAL
      0x222B: [716,216,417,55,472,'151 -112Q151 -150 106 -161Q106 -165 114 -172T134 -179Q155 -179 170 -146Q181 -120 188 -64T206 101T232 310Q256 472 277 567Q308 716 392 716Q434 716 453 681T472 613Q472 590 458 577T424 564Q404 564 390 578T376 612Q376 650 421 661Q421 663 418 667T407 675T393 679Q387 679 380 675Q360 665 350 619T326 438Q302 190 253 -57Q235 -147 201 -186Q174 -213 138 -216Q93 -216 74 -181T55 -113Q55 -91 69 -78T103 -64Q123 -64 137 -78T151 -112'],
  
      // TILDE OPERATOR
      0x223C: [367,-133,778,55,722,'55 166Q55 241 101 304T222 367Q260 367 296 349T362 304T421 252T484 208T554 189Q616 189 655 236T694 338Q694 350 698 358T708 367Q722 367 722 334Q722 260 677 197T562 134H554Q517 134 481 152T414 196T355 248T292 293T223 311Q179 311 145 286Q109 257 96 218T80 156T69 133Q55 133 55 166'],
  
      // WREATH PRODUCT
      0x2240: [583,83,278,55,222,'55 569Q55 583 83 583Q122 583 151 565T194 519T215 464T222 411Q222 360 194 304T139 193T111 89Q111 38 134 -7T195 -55Q222 -57 222 -69Q222 -83 189 -83Q130 -83 93 -33T55 90Q55 130 72 174T110 252T148 328T166 411Q166 462 144 507T83 555Q55 556 55 569'],
  
      // ASYMPTOTICALLY EQUAL TO
      0x2243: [464,-36,778,55,722,'55 283Q55 356 103 409T217 463Q262 463 297 447T395 382Q431 355 446 344T493 320T554 307H558Q613 307 652 344T694 433Q694 464 708 464T722 432Q722 356 673 304T564 251H554Q510 251 465 275T387 329T310 382T223 407H219Q164 407 122 367Q91 333 85 295T76 253T69 250Q55 250 55 283ZM56 56Q56 71 72 76H706Q722 70 722 56Q722 44 707 36H70Q56 43 56 56'],
  
      // APPROXIMATELY EQUAL TO
      0x2245: [589,-22,1000,55,722,'55 388Q55 463 101 526T222 589Q260 589 296 571T362 526T421 474T484 430T554 411Q616 411 655 458T694 560Q694 572 698 580T708 589Q722 589 722 556Q722 482 677 419T562 356H554Q517 356 481 374T414 418T355 471T292 515T223 533Q179 533 145 508Q109 479 96 440T80 378T69 355Q55 355 55 388ZM56 236Q56 249 70 256H707Q722 248 722 236Q722 225 708 217L390 216H72Q56 221 56 236ZM56 42Q56 57 72 62H708Q722 52 722 42Q722 30 707 22H70Q56 29 56 42'],
  
      // ALMOST EQUAL TO
      0x2248: [483,-55,778,55,722,'55 319Q55 360 72 393T114 444T163 472T205 482Q207 482 213 482T223 483Q262 483 296 468T393 413L443 381Q502 346 553 346Q609 346 649 375T694 454Q694 465 698 474T708 483Q722 483 722 452Q722 386 675 338T555 289Q514 289 468 310T388 357T308 404T224 426Q164 426 125 393T83 318Q81 289 69 289Q55 289 55 319ZM55 85Q55 126 72 159T114 210T163 238T205 248Q207 248 213 248T223 249Q262 249 296 234T393 179L443 147Q502 112 553 112Q609 112 649 141T694 220Q694 249 708 249T722 217Q722 153 675 104T555 55Q514 55 468 76T388 123T308 170T224 192Q164 192 125 159T83 84Q80 55 69 55Q55 55 55 85'],
  
      // EQUIVALENT TO
      0x224D: [484,-16,778,55,722,'55 464Q55 471 60 477T74 484Q80 484 108 464T172 420T268 376T389 356Q436 356 483 368T566 399T630 436T675 467T695 482Q701 484 703 484Q711 484 716 478T722 464Q722 454 707 442Q550 316 389 316Q338 316 286 329T195 362T124 402T76 437T57 456Q55 462 55 464ZM57 45Q66 58 109 88T230 151T381 183Q438 183 494 168T587 135T658 94T703 61T720 45Q722 39 722 36Q722 28 717 22T703 16Q697 16 669 36T606 80T510 124T389 144Q341 144 294 132T211 101T147 64T102 33T82 18Q76 16 74 16Q66 16 61 22T55 36Q55 39 57 45'],
  
      // APPROACHES THE LIMIT
      0x2250: [670,-133,778,56,722,'56 347Q56 360 70 367H707Q722 359 722 347Q722 336 708 328L390 327H72Q56 332 56 347ZM56 153Q56 168 72 173H708Q722 163 722 153Q722 140 707 133H70Q56 140 56 153ZM329 610Q329 634 346 652T389 670Q413 670 431 654T450 611Q450 586 433 568T390 550T347 567T329 610'],
  
      // stix-not (vert) equals
      0x2260: [716,215,778,56,722,'166 -215T159 -215T147 -212T141 -204T139 -197Q139 -190 144 -183L306 133H70Q56 140 56 153Q56 168 72 173H327L406 327H72Q56 332 56 347Q56 360 70 367H426Q597 702 602 707Q605 716 618 716Q625 716 630 712T636 703T638 696Q638 692 471 367H707Q722 359 722 347Q722 336 708 328L451 327L371 173H708Q722 163 722 153Q722 140 707 133H351Q175 -210 170 -212Q166 -215 159 -215'],
  
      // IDENTICAL TO
      0x2261: [464,-36,778,56,722,'56 444Q56 457 70 464H707Q722 456 722 444Q722 430 706 424H72Q56 429 56 444ZM56 237T56 250T70 270H707Q722 262 722 250T707 230H70Q56 237 56 250ZM56 56Q56 71 72 76H706Q722 70 722 56Q722 44 707 36H70Q56 43 56 56'],
  
      // LESS-THAN OR EQUAL TO
      0x2264: [636,138,778,83,694,'674 636Q682 636 688 630T694 615T687 601Q686 600 417 472L151 346L399 228Q687 92 691 87Q694 81 694 76Q694 58 676 56H670L382 192Q92 329 90 331Q83 336 83 348Q84 359 96 365Q104 369 382 500T665 634Q669 636 674 636ZM84 -118Q84 -108 99 -98H678Q694 -104 694 -118Q694 -130 679 -138H98Q84 -131 84 -118'],
  
      // GREATER-THAN OR EQUAL TO
      0x2265: [636,138,778,82,694,'83 616Q83 624 89 630T99 636Q107 636 253 568T543 431T687 361Q694 356 694 346T687 331Q685 329 395 192L107 56H101Q83 58 83 76Q83 77 83 79Q82 86 98 95Q117 105 248 167Q326 204 378 228L626 346L360 472Q291 505 200 548Q112 589 98 597T83 616ZM84 -118Q84 -108 99 -98H678Q694 -104 694 -118Q694 -130 679 -138H98Q84 -131 84 -118'],
  
      // MUCH LESS-THAN
      0x226A: [568,67,1000,56,944,'639 -48Q639 -54 634 -60T619 -67H618Q612 -67 536 -26Q430 33 329 88Q61 235 59 239Q56 243 56 250T59 261Q62 266 336 415T615 567L619 568Q622 567 625 567Q639 562 639 548Q639 540 633 534Q632 532 374 391L117 250L374 109Q632 -32 633 -34Q639 -40 639 -48ZM944 -48Q944 -54 939 -60T924 -67H923Q917 -67 841 -26Q735 33 634 88Q366 235 364 239Q361 243 361 250T364 261Q367 266 641 415T920 567L924 568Q927 567 930 567Q944 562 944 548Q944 540 938 534Q937 532 679 391L422 250L679 109Q937 -32 938 -34Q944 -40 944 -48'],
  
      // MUCH GREATER-THAN
      0x226B: [567,67,1000,55,944,'55 539T55 547T60 561T74 567Q81 567 207 498Q297 449 365 412Q633 265 636 261Q639 255 639 250Q639 241 626 232Q614 224 365 88Q83 -65 79 -66Q76 -67 73 -67Q65 -67 60 -61T55 -47Q55 -39 61 -33Q62 -33 95 -15T193 39T320 109L321 110H322L323 111H324L325 112L326 113H327L329 114H330L331 115H332L333 116L334 117H335L336 118H337L338 119H339L340 120L341 121H342L343 122H344L345 123H346L347 124L348 125H349L351 126H352L353 127H354L355 128L356 129H357L358 130H359L360 131H361L362 132L363 133H364L365 134H366L367 135H368L369 136H370L371 137L372 138H373L374 139H375L376 140L378 141L576 251Q63 530 62 533Q55 539 55 547ZM360 539T360 547T365 561T379 567Q386 567 512 498Q602 449 670 412Q938 265 941 261Q944 255 944 250Q944 241 931 232Q919 224 670 88Q388 -65 384 -66Q381 -67 378 -67Q370 -67 365 -61T360 -47Q360 -39 366 -33Q367 -33 400 -15T498 39T625 109L626 110H627L628 111H629L630 112L631 113H632L634 114H635L636 115H637L638 116L639 117H640L641 118H642L643 119H644L645 120L646 121H647L648 122H649L650 123H651L652 124L653 125H654L656 126H657L658 127H659L660 128L661 129H662L663 130H664L665 131H666L667 132L668 133H669L670 134H671L672 135H673L674 136H675L676 137L677 138H678L679 139H680L681 140L683 141L881 251Q368 530 367 533Q360 539 360 547'],
  
      // PRECEDES
      0x227A: [539,41,778,84,694,'84 249Q84 262 91 266T117 270Q120 270 126 270T137 269Q388 273 512 333T653 512Q657 539 676 539Q685 538 689 532T694 520V515Q689 469 672 431T626 366T569 320T500 286T435 265T373 249Q379 248 404 242T440 233T477 221T533 199Q681 124 694 -17Q694 -41 674 -41Q658 -41 653 -17Q646 41 613 84T533 154T418 197T284 220T137 229H114Q104 229 98 230T88 235T84 249'],
  
      // SUCCEEDS
      0x227B: [539,41,778,83,694,'84 517Q84 539 102 539Q115 539 119 529T125 503T137 459T171 404Q277 275 640 269H661Q694 269 694 249T661 229H640Q526 227 439 214T283 173T173 98T124 -17Q118 -41 103 -41Q83 -41 83 -17Q88 29 105 67T151 132T208 178T277 212T342 233T404 249Q401 250 380 254T345 263T302 276T245 299Q125 358 92 468Q84 502 84 517'],
  
      // SUBSET OF
      0x2282: [541,41,778,84,694,'84 250Q84 372 166 450T360 539Q361 539 370 539T395 539T430 540T475 540T524 540H679Q694 532 694 520Q694 511 681 501L522 500H470H441Q366 500 338 496T266 472Q244 461 224 446T179 404T139 337T124 250V245Q124 157 185 89Q244 25 328 7Q348 2 366 2T522 0H681Q694 -10 694 -20Q694 -32 679 -40H526Q510 -40 480 -40T434 -41Q350 -41 289 -25T172 45Q84 127 84 250'],
  
      // SUPERSET OF
      0x2283: [541,40,778,83,693,'83 520Q83 532 98 540H251Q267 540 297 540T343 541Q427 541 488 525T605 455Q693 374 693 250Q693 165 650 99T545 0T415 -39Q407 -40 251 -40H98Q83 -32 83 -20Q83 -10 96 0H255H308H337Q412 0 439 4T512 28Q533 39 553 54T599 96T639 163T654 250Q654 341 592 411Q557 449 512 472Q468 491 439 495T335 500H306H255L96 501Q83 511 83 520'],
  
      // SUBSET OF OR EQUAL TO
      0x2286: [637,138,778,84,694,'84 346Q84 468 166 546T360 635Q361 635 370 635T395 635T430 636T475 636T524 636H679Q694 628 694 616Q694 607 681 597L522 596H470H441Q366 596 338 592T266 568Q244 557 224 542T179 500T139 433T124 346V341Q124 253 185 185Q244 121 328 103Q348 98 366 98T522 96H681Q694 86 694 76Q694 64 679 56H526Q510 56 480 56T434 55Q350 55 289 71T172 141Q84 223 84 346ZM104 -131T104 -118T118 -98H679Q694 -106 694 -118T679 -138H118Q104 -131 104 -118'],
  
      // SUPERSET OF OR EQUAL TO
      0x2287: [637,138,778,83,693,'83 616Q83 628 98 636H251Q267 636 297 636T343 637Q427 637 488 621T605 551Q693 470 693 346Q693 261 650 195T545 96T415 57Q407 56 251 56H98Q83 64 83 76Q83 86 96 96H255H308H337Q412 96 439 100T512 124Q533 135 553 150T599 192T639 259T654 346Q654 437 592 507Q557 545 512 568Q468 587 439 591T335 596H306H255L96 597Q83 607 83 616ZM84 -131T84 -118T98 -98H659Q674 -106 674 -118T659 -138H98Q84 -131 84 -118'],
  
      // MULTISET UNION
      0x228E: [598,22,667,55,611,'591 598H592Q604 598 611 583V376Q611 345 611 296Q610 162 606 148Q605 146 605 145Q586 68 507 23T333 -22Q268 -22 209 -1T106 66T56 173Q55 180 55 384L56 585Q66 598 75 598Q85 598 95 585V378L96 172L98 162Q112 95 181 57T332 18Q415 18 487 58T570 175Q571 180 571 383V583Q579 598 591 598ZM313 406Q313 417 313 435T312 459Q312 483 316 493T333 503T349 494T353 461V406V325H515Q516 325 519 323T527 316T531 305T527 294T520 287T515 285H353V204V152Q353 127 350 117T333 107T316 117T312 152Q312 158 312 175T313 204V285H151Q150 285 147 287T139 294T135 305T139 316T146 323T151 325H313V406'],
  
      // SQUARE IMAGE OF OR EQUAL TO
      0x2291: [636,138,778,84,714,'94 620Q98 632 110 636H699Q714 628 714 616T699 596H134V96H698Q714 90 714 76Q714 64 699 56H109Q104 59 95 69L94 344V620ZM84 -118Q84 -103 100 -98H698Q714 -104 714 -118Q714 -130 699 -138H98Q84 -131 84 -118'],
  
      // SQUARE ORIGINAL OF OR EQUAL TO
      0x2292: [636,138,778,64,694,'64 603T64 616T78 636H668Q675 633 683 623V69Q675 59 668 56H78Q64 63 64 76Q64 91 80 96H643V596H78Q64 603 64 616ZM64 -118Q64 -108 79 -98H678Q694 -104 694 -118Q694 -130 679 -138H78Q64 -131 64 -118'],
  
      // stix-square intersection, serifs
      0x2293: [598,0,667,61,605,'83 0Q79 0 76 1T71 3T67 6T65 9T63 13T61 16V301L62 585Q70 595 76 598H592Q602 590 605 583V15Q598 2 587 0Q583 0 580 1T575 3T571 6T569 9T567 13T565 16V558H101V15Q94 2 83 0'],
  
      // stix-square union, serifs
      0x2294: [598,0,667,61,605,'77 0Q65 4 61 16V301L62 585Q72 598 81 598Q94 598 101 583V40H565V583Q573 598 585 598Q598 598 605 583V15Q602 10 592 1L335 0H77'],
  
      // stix-circled plus (with rim)
      0x2295: [583,83,778,56,722,'56 250Q56 394 156 488T384 583Q530 583 626 485T722 250Q722 110 625 14T390 -83Q249 -83 153 14T56 250ZM364 542Q308 539 251 509T148 418T96 278V270H369V542H364ZM681 278Q675 338 650 386T592 462T522 509T458 535T412 542H409V270H681V278ZM96 222Q104 150 139 95T219 12T302 -29T366 -42H369V230H96V222ZM681 222V230H409V-42H412Q429 -42 456 -36T521 -10T590 37T649 113T681 222'],
  
      // CIRCLED MINUS
      0x2296: [583,83,778,56,722,'56 250Q56 394 156 488T384 583Q530 583 626 485T722 250Q722 110 625 14T390 -83Q249 -83 153 14T56 250ZM681 278Q669 385 591 463T381 542Q283 542 196 471T96 278V270H681V278ZM275 -42T388 -42T585 32T681 222V230H96V222Q108 107 191 33'],
  
      // stix-circled times (with rim)
      0x2297: [583,83,778,56,722,'56 250Q56 394 156 488T384 583Q530 583 626 485T722 250Q722 110 625 14T390 -83Q249 -83 153 14T56 250ZM582 471Q531 510 496 523Q446 542 381 542Q324 542 272 519T196 471L389 278L485 375L582 471ZM167 442Q95 362 95 250Q95 137 167 58L359 250L167 442ZM610 58Q682 138 682 250Q682 363 610 442L418 250L610 58ZM196 29Q209 16 230 2T295 -27T388 -42Q409 -42 429 -40T465 -33T496 -23T522 -11T544 1T561 13T574 22T582 29L388 222L196 29'],
  
      // CIRCLED DIVISION SLASH
      0x2298: [583,83,778,56,722,'56 250Q56 394 156 488T384 583Q530 583 626 485T722 250Q722 110 625 14T390 -83Q249 -83 153 14T56 250ZM582 471Q581 472 571 480T556 491T539 502T517 514T491 525T460 534T424 539T381 542Q272 542 184 460T95 251Q95 198 113 150T149 80L167 58L582 471ZM388 -42Q513 -42 597 44T682 250Q682 363 610 442L196 29Q209 16 229 2T295 -27T388 -42'],
  
      // CIRCLED DOT OPERATOR
      0x2299: [583,83,778,56,722,'56 250Q56 394 156 488T384 583Q530 583 626 485T722 250Q722 110 625 14T390 -83Q249 -83 153 14T56 250ZM682 250Q682 322 649 387T546 497T381 542Q272 542 184 459T95 250Q95 132 178 45T389 -42Q515 -42 598 45T682 250ZM311 250Q311 285 332 304T375 328Q376 328 382 328T392 329Q424 326 445 305T466 250Q466 217 445 195T389 172Q354 172 333 195T311 250'],
  
      // RIGHT TACK
      0x22A2: [695,0,611,55,555,'55 678Q55 679 56 681T58 684T61 688T65 691T70 693T77 694Q88 692 95 679V367H540Q555 359 555 347Q555 334 540 327H95V15Q88 2 77 0Q73 0 70 1T65 3T61 6T59 9T57 13T55 16V678'],
  
      // LEFT TACK
      0x22A3: [695,0,611,54,555,'515 678Q515 679 516 681T518 684T521 688T525 691T530 693T537 694Q548 692 555 679V15Q548 2 537 0Q533 0 530 1T525 3T521 6T519 9T517 13T515 16V327H71Q70 327 67 329T59 336T55 347T59 358T66 365T71 367H515V678'],
  
      // DOWN TACK
      0x22A4: [668,0,778,55,723,'55 642T55 648T59 659T66 666T71 668H708Q723 660 723 648T708 628H409V15Q402 2 391 0Q387 0 384 1T379 3T375 6T373 9T371 13T369 16V628H71Q70 628 67 630T59 637'],
  
      // UP TACK
      0x22A5: [669,0,778,54,723,'369 652Q369 653 370 655T372 658T375 662T379 665T384 667T391 668Q402 666 409 653V40H708Q723 32 723 20T708 0H71Q70 0 67 2T59 9T55 20T59 31T66 38T71 40H369V652'],
  
      // TRUE
      0x22A8: [750,249,867,119,812,'139 -249H137Q125 -249 119 -235V251L120 737Q130 750 139 750Q152 750 159 735V367H796Q811 359 811 347Q811 336 797 328L479 327H161L159 328V172L161 173H797Q798 172 800 171T803 169T805 167T808 164T809 162T810 158T811 153Q811 140 796 133H159V-235Q151 -249 141 -249H139'],
  
      // DIAMOND OPERATOR
      0x22C4: [488,-12,500,12,488,'242 486Q245 488 250 488Q256 488 258 486Q262 484 373 373T486 258T488 250T486 242T373 127T258 14Q256 12 250 12Q245 12 242 14Q237 16 127 126T14 242Q12 245 12 250T14 258Q16 263 126 373T242 486ZM439 250L250 439L61 250L250 61L439 250'],
  
      // DOT OPERATOR
      0x22C5: [310,-190,278,78,199,'78 250Q78 274 95 292T138 310Q162 310 180 294T199 251Q199 226 182 208T139 190T96 207T78 250'],
  
      // STAR OPERATOR
      0x22C6: [486,-16,500,3,497,'210 282Q210 284 225 381T241 480Q241 484 245 484Q249 486 251 486Q258 486 260 477T272 406Q275 390 276 380Q290 286 290 282L388 299Q484 314 487 314H488Q497 314 497 302Q497 297 434 266Q416 257 404 251L315 206L361 118Q372 98 383 75T401 40L407 28Q407 16 395 16Q394 16 392 16L390 17L250 159L110 17L108 16Q106 16 105 16Q93 16 93 28L99 40Q105 52 116 75T139 118L185 206L96 251Q6 296 4 300Q3 301 3 302Q3 314 12 314H13Q16 314 112 299L210 282'],
  
      // BOWTIE
      0x22C8: [505,5,900,26,873,'833 50T833 250T832 450T659 351T487 250T658 150T832 50Q833 50 833 250ZM873 10Q866 -5 854 -5Q851 -5 845 -3L449 226L260 115Q51 -5 43 -5Q39 -5 35 -1T28 7L26 11V489Q33 505 43 505Q51 505 260 385L449 274L845 503Q851 505 853 505Q866 505 873 490V10ZM412 250L67 450Q66 450 66 250T67 50Q69 51 240 150T412 250'],
  
      // VERTICAL ELLIPSIS
      0x22EE: [900,30,278,78,199,'78 30Q78 54 95 72T138 90Q162 90 180 74T199 31Q199 6 182 -12T139 -30T96 -13T78 30ZM78 440Q78 464 95 482T138 500Q162 500 180 484T199 441Q199 416 182 398T139 380T96 397T78 440ZM78 840Q78 864 95 882T138 900Q162 900 180 884T199 841Q199 816 182 798T139 780T96 797T78 840'],
  
      // MIDLINE HORIZONTAL ELLIPSIS
      0x22EF: [310,-190,1172,78,1093,'78 250Q78 274 95 292T138 310Q162 310 180 294T199 251Q199 226 182 208T139 190T96 207T78 250ZM525 250Q525 274 542 292T585 310Q609 310 627 294T646 251Q646 226 629 208T586 190T543 207T525 250ZM972 250Q972 274 989 292T1032 310Q1056 310 1074 294T1093 251Q1093 226 1076 208T1033 190T990 207T972 250'],
  
      // DOWN RIGHT DIAGONAL ELLIPSIS
      0x22F1: [820,-100,1282,133,1148,'133 760Q133 784 150 802T193 820Q217 820 235 804T254 761Q254 736 237 718T194 700T151 717T133 760ZM580 460Q580 484 597 502T640 520Q664 520 682 504T701 461Q701 436 684 418T641 400T598 417T580 460ZM1027 160Q1027 184 1044 202T1087 220Q1111 220 1129 204T1148 161Q1148 136 1131 118T1088 100T1045 117T1027 160'],
  
      // LEFT CEILING
      0x2308: [750,250,444,174,422,'174 734Q178 746 190 750H298H369Q400 750 411 747T422 730T411 713T372 709Q365 709 345 709T310 710H214V-235Q206 -248 196 -250Q192 -250 189 -249T184 -247T180 -244T178 -241T176 -237T174 -234V734'],
  
      // RIGHT CEILING
      0x2309: [750,250,444,21,269,'21 717T21 730T32 746T75 750H147H256Q266 742 269 735V-235Q262 -248 251 -250Q247 -250 244 -249T239 -247T235 -244T233 -241T231 -237T229 -234V710H133Q119 710 99 710T71 709Q43 709 32 713'],
  
      // LEFT FLOOR
      0x230A: [751,251,444,174,423,'174 734Q174 735 175 737T177 740T180 744T184 747T189 749T196 750Q206 748 214 735V-210H310H373Q401 -210 411 -213T422 -230T411 -247T369 -251Q362 -251 338 -251T298 -250H190Q178 -246 174 -234V734'],
  
      // RIGHT FLOOR
      0x230B: [751,250,444,21,269,'229 734Q229 735 230 737T232 740T235 744T239 747T244 749T251 750Q262 748 269 735V-235Q266 -240 256 -249L147 -250H77Q43 -250 32 -247T21 -230T32 -213T72 -209Q79 -209 99 -209T133 -210H229V734'],
  
      // stix-small down curve
      0x2322: [388,-122,1000,55,944,'55 141Q55 149 72 174T125 234T209 303T329 360T478 388H526Q649 383 765 319Q814 291 858 250T923 179T944 141Q944 133 938 128T924 122Q914 124 912 125T902 139Q766 328 500 328Q415 328 342 308T225 258T150 199T102 148T84 124Q81 122 75 122Q55 127 55 141'],
  
      // stix-small up curve
      0x2323: [378,-134,1000,55,944,'923 378Q944 378 944 358Q944 345 912 311T859 259Q710 134 500 134Q288 134 140 259Q55 336 55 358Q55 366 61 372T75 378Q78 378 84 376Q86 376 101 356T147 310T221 257T339 212T500 193Q628 193 734 236Q841 282 903 363Q914 378 923 378'],
  
      // UPPER LEFT OR LOWER RIGHT CURLY BRACKET SECTION
      0x23B0: [744,244,412,56,357,'357 741V726Q357 720 349 715Q261 655 242 539Q240 526 240 454T239 315T239 247Q240 235 240 124V40Q240 -17 233 -53T201 -130Q155 -206 78 -244H69H64Q58 -244 57 -243T56 -234Q56 -232 56 -231V-225Q56 -218 63 -215Q153 -153 170 -39Q172 -25 173 119V219Q173 245 174 249Q173 258 173 376V460Q173 515 178 545T201 611Q244 695 327 741L334 744H354L357 741'],
  
      // UPPER RIGHT OR LOWER LEFT CURLY BRACKET SECTION
      0x23B1: [744,244,412,55,357,'78 744Q153 706 196 640T239 492V376Q239 341 239 314T238 271T238 253Q239 251 239 223V119V49Q239 -39 254 -85Q263 -111 275 -134T301 -172T326 -197T346 -213T356 -221T357 -232V-241L354 -244H334Q264 -209 222 -146T174 -12Q173 -6 173 95Q173 134 173 191T174 250Q173 258 173 382V451Q173 542 159 585Q145 626 120 658T75 706T56 723V731Q56 741 57 742T66 744H78'],
  
      // MATHEMATICAL LEFT ANGLE BRACKET
      0x27E8: [750,250,389,109,333,'333 -232Q332 -239 327 -244T313 -250Q303 -250 296 -240Q293 -233 202 6T110 250T201 494T296 740Q299 745 306 749L309 750Q312 750 313 750Q331 750 333 732Q333 727 243 489Q152 252 152 250T243 11Q333 -227 333 -232'],
  
      // MATHEMATICAL RIGHT ANGLE BRACKET
      0x27E9: [750,250,389,55,279,'55 732Q56 739 61 744T75 750Q85 750 92 740Q95 733 186 494T278 250T187 6T92 -240Q85 -250 75 -250Q67 -250 62 -245T55 -232Q55 -227 145 11Q236 248 236 250T145 489Q55 727 55 732'],
  
      // MATHEMATICAL LEFT FLATTENED PARENTHESIS
      0x27EE: [744,244,412,173,357,'357 741V726Q357 720 349 715Q261 655 242 539Q240 526 240 394V331Q240 259 239 250Q240 242 240 119V49Q240 -42 254 -85Q263 -111 275 -134T301 -172T326 -197T346 -213T356 -221T357 -232V-241L354 -244H334Q264 -209 222 -146T174 -12Q173 -6 173 95Q173 134 173 191T174 250Q173 260 173 376V460Q173 515 178 545T201 611Q244 695 327 741L334 744H354L357 741'],
  
      // MATHEMATICAL RIGHT FLATTENED PARENTHESIS
      0x27EF: [744,244,412,55,240,'78 744Q153 706 196 640T239 492V376Q239 339 239 311T238 269T238 252Q240 236 240 124V40Q240 -18 233 -53T202 -130Q156 -206 79 -244H70H65Q58 -244 57 -242T56 -231T57 -220T64 -215Q153 -154 170 -39Q173 -18 174 119V247Q173 249 173 382V451Q173 542 159 585Q145 626 120 658T75 706T56 723V731Q56 741 57 742T66 744H78'],
  
      // LONG LEFTWARDS ARROW
      0x27F5: [511,11,1609,55,1525,'165 270H1510Q1525 262 1525 250T1510 230H165Q167 228 182 216T211 189T244 152T277 96T303 25Q308 7 308 0Q308 -11 288 -11Q281 -11 278 -11T272 -7T267 2T263 21Q245 94 195 151T73 236Q58 242 55 247Q55 254 59 257T73 264Q121 283 158 314T215 375T247 434T264 480L267 497Q269 503 270 505T275 509T288 511Q308 511 308 500Q308 493 303 475Q293 438 278 406T246 352T215 315T185 287T165 270'],
  
      // LONG RIGHTWARDS ARROW
      0x27F6: [511,11,1638,84,1553,'84 237T84 250T98 270H1444Q1328 357 1301 493Q1301 494 1301 496T1300 499Q1300 511 1317 511H1320Q1329 511 1332 510T1338 506T1341 497T1344 481T1352 456Q1374 389 1425 336T1544 261Q1553 258 1553 250Q1553 244 1548 241T1524 231T1486 212Q1445 186 1415 152T1370 85T1349 35T1341 4Q1339 -6 1336 -8T1320 -11Q1300 -11 1300 0Q1300 7 1305 25Q1337 151 1444 230H98Q84 237 84 250'],
  
      // LONG LEFT RIGHT ARROW
      0x27F7: [511,11,1859,55,1803,'165 270H1694Q1578 357 1551 493Q1551 494 1551 496T1550 499Q1550 511 1567 511H1570Q1579 511 1582 510T1588 506T1591 497T1594 481T1602 456Q1624 389 1675 336T1794 261Q1803 258 1803 250Q1803 244 1798 241T1774 231T1736 212Q1695 186 1665 152T1620 85T1599 35T1591 4Q1589 -6 1586 -8T1570 -11Q1550 -11 1550 0Q1550 7 1555 25Q1587 151 1694 230H165Q167 228 182 216T211 189T244 152T277 96T303 25Q308 7 308 0Q308 -11 288 -11Q281 -11 278 -11T272 -7T267 2T263 21Q245 94 195 151T73 236Q58 242 55 247Q55 254 59 257T73 264Q121 283 158 314T215 375T247 434T264 480L267 497Q269 503 270 505T275 509T288 511Q308 511 308 500Q308 493 303 475Q293 438 278 406T246 352T215 315T185 287T165 270'],
  
      // LONG LEFTWARDS DOUBLE ARROW
      0x27F8: [525,24,1609,56,1554,'274 173H1539Q1540 172 1542 171T1545 169T1547 167T1550 164T1551 162T1552 158T1553 153Q1553 140 1538 133H318L328 123Q379 69 414 0Q419 -13 419 -17Q419 -24 399 -24Q388 -24 385 -23T377 -12Q332 77 253 144T72 237Q62 240 59 242T56 250T59 257T70 262T89 268T119 278T160 296Q303 366 377 512Q382 522 385 523T401 525Q419 524 419 515Q419 510 414 500Q379 431 328 377L318 367H1538Q1553 359 1553 347Q1553 336 1539 328L1221 327H903L900 328L602 327H274L264 319Q225 289 147 250Q148 249 165 241T210 217T264 181L274 173'],
  
      // LONG RIGHTWARDS DOUBLE ARROW
      0x27F9: [525,24,1638,56,1582,'1218 514Q1218 525 1234 525Q1239 525 1242 525T1247 525T1251 524T1253 523T1255 520T1257 517T1260 512Q1297 438 1358 381T1469 300T1565 263Q1582 258 1582 250T1573 239T1536 228T1478 204Q1334 134 1260 -12Q1256 -21 1253 -22T1238 -24Q1218 -24 1218 -17Q1218 -13 1223 0Q1258 69 1309 123L1319 133H70Q56 140 56 153Q56 168 72 173H1363L1373 181Q1412 211 1490 250Q1489 251 1472 259T1427 283T1373 319L1363 327H710L707 328L390 327H72Q56 332 56 347Q56 360 70 367H1319L1309 377Q1276 412 1247 458T1218 514'],
  
      // LONG LEFT RIGHT DOUBLE ARROW
      0x27FA: [525,24,1858,56,1802,'1438 514Q1438 525 1454 525Q1459 525 1462 525T1467 525T1471 524T1473 523T1475 520T1477 517T1480 512Q1517 438 1578 381T1689 300T1785 263Q1802 258 1802 250T1793 239T1756 228T1698 204Q1554 134 1480 -12Q1476 -21 1473 -22T1458 -24Q1438 -24 1438 -17Q1438 -13 1443 0Q1478 69 1529 123L1539 133H318L328 123Q379 69 414 0Q419 -13 419 -17Q419 -24 399 -24Q388 -24 385 -23T377 -12Q332 77 253 144T72 237Q62 240 59 242T56 250T59 257T70 262T89 268T119 278T160 296Q303 366 377 512Q382 522 385 523T401 525Q419 524 419 515Q419 510 414 500Q379 431 328 377L318 367H1539L1529 377Q1496 412 1467 458T1438 514ZM274 173H1583L1593 181Q1632 211 1710 250Q1709 251 1692 259T1647 283T1593 319L1583 327H930L927 328L602 327H274L264 319Q225 289 147 250Q148 249 165 241T210 217T264 181L274 173'],
  
      // LONG RIGHTWARDS ARROW FROM BAR
      0x27FC: [511,11,1638,54,1553,'95 155V109Q95 83 92 73T75 63Q61 63 58 74T54 130Q54 140 54 180T55 250Q55 421 57 425Q61 437 75 437Q88 437 91 428T95 393V345V270H1444Q1328 357 1301 493Q1301 494 1301 496T1300 499Q1300 511 1317 511H1320Q1329 511 1332 510T1338 506T1341 497T1344 481T1352 456Q1374 389 1425 336T1544 261Q1553 258 1553 250Q1553 244 1548 241T1524 231T1486 212Q1445 186 1415 152T1370 85T1349 35T1341 4Q1339 -6 1336 -8T1320 -11Q1300 -11 1300 0Q1300 7 1305 25Q1337 151 1444 230H95V155'],
  
      // PRECEDES ABOVE SINGLE-LINE EQUALS SIGN
      0x2AAF: [636,138,778,84,694,'84 346Q84 359 91 363T117 367Q120 367 126 367T137 366Q388 370 512 430T653 609Q657 636 676 636Q685 635 689 629T694 618V612Q689 566 672 528T626 463T569 417T500 383T435 362T373 346Q379 345 404 339T440 330T477 318T533 296Q592 266 630 223T681 145T694 78Q694 57 674 57Q662 57 657 67T652 92T640 135T606 191Q500 320 137 326H114Q104 326 98 327T88 332T84 346ZM84 -131T84 -118T98 -98H679Q694 -106 694 -118T679 -138H98Q84 -131 84 -118'],
  
      // SUCCEEDS ABOVE SINGLE-LINE EQUALS SIGN
      0x2AB0: [636,138,778,83,694,'84 614Q84 636 102 636Q115 636 119 626T125 600T137 556T171 501Q277 372 640 366H661Q694 366 694 346T661 326H640Q578 325 526 321T415 307T309 280T222 237T156 172T124 83Q122 66 118 62T103 57Q100 57 98 57T95 58T93 59T90 62T85 67Q83 71 83 80Q88 126 105 164T151 229T208 275T277 309T342 330T404 346Q401 347 380 351T345 360T302 373T245 396Q125 455 92 565Q84 599 84 614ZM84 -131T84 -118T98 -98H679Q694 -106 694 -118T679 -138H98Q84 -131 84 -118']
  };

  SVG.FONTDATA.FONTS['MathJax_Math-italic'] = {
    directory: 'Math/Italic',
    family: 'MathJax_Math',
    id: 'MJMATHI',
    style: 'italic',
    skew: {
      0x41: 0.139,
      0x42: 0.0833,
      0x43: 0.0833,
      0x44: 0.0556,
      0x45: 0.0833,
      0x46: 0.0833,
      0x47: 0.0833,
      0x48: 0.0556,
      0x49: 0.111,
      0x4A: 0.167,
      0x4B: 0.0556,
      0x4C: 0.0278,
      0x4D: 0.0833,
      0x4E: 0.0833,
      0x4F: 0.0833,
      0x50: 0.0833,
      0x51: 0.0833,
      0x52: 0.0833,
      0x53: 0.0833,
      0x54: 0.0833,
      0x55: 0.0278,
      0x58: 0.0833,
      0x5A: 0.0833,
      0x63: 0.0556,
      0x64: 0.167,
      0x65: 0.0556,
      0x66: 0.167,
      0x67: 0.0278,
      0x68: -0.0278,
      0x6C: 0.0833,
      0x6F: 0.0556,
      0x70: 0.0833,
      0x71: 0.0833,
      0x72: 0.0556,
      0x73: 0.0556,
      0x74: 0.0833,
      0x75: 0.0278,
      0x76: 0.0278,
      0x77: 0.0833,
      0x78: 0.0278,
      0x79: 0.0556,
      0x7A: 0.0556,
      0x393: 0.0833,
      0x394: 0.167,
      0x398: 0.0833,
      0x39B: 0.167,
      0x39E: 0.0833,
      0x3A0: 0.0556,
      0x3A3: 0.0833,
      0x3A5: 0.0556,
      0x3A6: 0.0833,
      0x3A8: 0.0556,
      0x3A9: 0.0833,
      0x3B1: 0.0278,
      0x3B2: 0.0833,
      0x3B4: 0.0556,
      0x3B5: 0.0833,
      0x3B6: 0.0833,
      0x3B7: 0.0556,
      0x3B8: 0.0833,
      0x3B9: 0.0556,
      0x3BC: 0.0278,
      0x3BD: 0.0278,
      0x3BE: 0.111,
      0x3BF: 0.0556,
      0x3C1: 0.0833,
      0x3C2: 0.0833,
      0x3C4: 0.0278,
      0x3C5: 0.0278,
      0x3C6: 0.0833,
      0x3C7: 0.0556,
      0x3C8: 0.111,
      0x3D1: 0.0833,
      0x3D5: 0.0833,
      0x3F1: 0.0833,
      0x3F5: 0.0556
    },
  
      // SPACE
      0x20: [0,0,250,0,0,''],
  
      // SOLIDUS
      0x2F: [716,215,778,139,638,'166 -215T159 -215T147 -212T141 -204T139 -197Q139 -190 144 -183Q157 -157 378 274T602 707Q605 716 618 716Q625 716 630 712T636 703T638 696Q638 691 406 241T170 -212Q166 -215 159 -215'],
  
      // LATIN CAPITAL LETTER A
      0x41: [716,0,750,35,726,'208 74Q208 50 254 46Q272 46 272 35Q272 34 270 22Q267 8 264 4T251 0Q249 0 239 0T205 1T141 2Q70 2 50 0H42Q35 7 35 11Q37 38 48 46H62Q132 49 164 96Q170 102 345 401T523 704Q530 716 547 716H555H572Q578 707 578 706L606 383Q634 60 636 57Q641 46 701 46Q726 46 726 36Q726 34 723 22Q720 7 718 4T704 0Q701 0 690 0T651 1T578 2Q484 2 455 0H443Q437 6 437 9T439 27Q443 40 445 43L449 46H469Q523 49 533 63L521 213H283L249 155Q208 86 208 74ZM516 260Q516 271 504 416T490 562L463 519Q447 492 400 412L310 260L413 259Q516 259 516 260'],
  
      // LATIN CAPITAL LETTER B
      0x42: [683,0,759,35,756,'231 637Q204 637 199 638T194 649Q194 676 205 682Q206 683 335 683Q594 683 608 681Q671 671 713 636T756 544Q756 480 698 429T565 360L555 357Q619 348 660 311T702 219Q702 146 630 78T453 1Q446 0 242 0Q42 0 39 2Q35 5 35 10Q35 17 37 24Q42 43 47 45Q51 46 62 46H68Q95 46 128 49Q142 52 147 61Q150 65 219 339T288 628Q288 635 231 637ZM649 544Q649 574 634 600T585 634Q578 636 493 637Q473 637 451 637T416 636H403Q388 635 384 626Q382 622 352 506Q352 503 351 500L320 374H401Q482 374 494 376Q554 386 601 434T649 544ZM595 229Q595 273 572 302T512 336Q506 337 429 337Q311 337 310 336Q310 334 293 263T258 122L240 52Q240 48 252 48T333 46Q422 46 429 47Q491 54 543 105T595 229'],
  
      // LATIN CAPITAL LETTER C
      0x43: [705,22,715,50,760,'50 252Q50 367 117 473T286 641T490 704Q580 704 633 653Q642 643 648 636T656 626L657 623Q660 623 684 649Q691 655 699 663T715 679T725 690L740 705H746Q760 705 760 698Q760 694 728 561Q692 422 692 421Q690 416 687 415T669 413H653Q647 419 647 422Q647 423 648 429T650 449T651 481Q651 552 619 605T510 659Q484 659 454 652T382 628T299 572T226 479Q194 422 175 346T156 222Q156 108 232 58Q280 24 350 24Q441 24 512 92T606 240Q610 253 612 255T628 257Q648 257 648 248Q648 243 647 239Q618 132 523 55T319 -22Q206 -22 128 53T50 252'],
  
      // LATIN CAPITAL LETTER D
      0x44: [683,0,828,33,803,'287 628Q287 635 230 637Q207 637 200 638T193 647Q193 655 197 667T204 682Q206 683 403 683Q570 682 590 682T630 676Q702 659 752 597T803 431Q803 275 696 151T444 3L430 1L236 0H125H72Q48 0 41 2T33 11Q33 13 36 25Q40 41 44 43T67 46Q94 46 127 49Q141 52 146 61Q149 65 218 339T287 628ZM703 469Q703 507 692 537T666 584T629 613T590 629T555 636Q553 636 541 636T512 636T479 637H436Q392 637 386 627Q384 623 313 339T242 52Q242 48 253 48T330 47Q335 47 349 47T373 46Q499 46 581 128Q617 164 640 212T683 339T703 469'],
  
      // LATIN CAPITAL LETTER E
      0x45: [680,0,738,31,764,'492 213Q472 213 472 226Q472 230 477 250T482 285Q482 316 461 323T364 330H312Q311 328 277 192T243 52Q243 48 254 48T334 46Q428 46 458 48T518 61Q567 77 599 117T670 248Q680 270 683 272Q690 274 698 274Q718 274 718 261Q613 7 608 2Q605 0 322 0H133Q31 0 31 11Q31 13 34 25Q38 41 42 43T65 46Q92 46 125 49Q139 52 144 61Q146 66 215 342T285 622Q285 629 281 629Q273 632 228 634H197Q191 640 191 642T193 659Q197 676 203 680H757Q764 676 764 669Q764 664 751 557T737 447Q735 440 717 440H705Q698 445 698 453L701 476Q704 500 704 528Q704 558 697 578T678 609T643 625T596 632T532 634H485Q397 633 392 631Q388 629 386 622Q385 619 355 499T324 377Q347 376 372 376H398Q464 376 489 391T534 472Q538 488 540 490T557 493Q562 493 565 493T570 492T572 491T574 487T577 483L544 351Q511 218 508 216Q505 213 492 213'],
  
      // LATIN CAPITAL LETTER F
      0x46: [680,0,643,31,749,'48 1Q31 1 31 11Q31 13 34 25Q38 41 42 43T65 46Q92 46 125 49Q139 52 144 61Q146 66 215 342T285 622Q285 629 281 629Q273 632 228 634H197Q191 640 191 642T193 659Q197 676 203 680H742Q749 676 749 669Q749 664 736 557T722 447Q720 440 702 440H690Q683 445 683 453Q683 454 686 477T689 530Q689 560 682 579T663 610T626 626T575 633T503 634H480Q398 633 393 631Q388 629 386 623Q385 622 352 492L320 363H375Q378 363 398 363T426 364T448 367T472 374T489 386Q502 398 511 419T524 457T529 475Q532 480 548 480H560Q567 475 567 470Q567 467 536 339T502 207Q500 200 482 200H470Q463 206 463 212Q463 215 468 234T473 274Q473 303 453 310T364 317H309L277 190Q245 66 245 60Q245 46 334 46H359Q365 40 365 39T363 19Q359 6 353 0H336Q295 2 185 2Q120 2 86 2T48 1'],
  
      // LATIN CAPITAL LETTER G
      0x47: [705,22,786,50,760,'50 252Q50 367 117 473T286 641T490 704Q580 704 633 653Q642 643 648 636T656 626L657 623Q660 623 684 649Q691 655 699 663T715 679T725 690L740 705H746Q760 705 760 698Q760 694 728 561Q692 422 692 421Q690 416 687 415T669 413H653Q647 419 647 422Q647 423 648 429T650 449T651 481Q651 552 619 605T510 659Q492 659 471 656T418 643T357 615T294 567T236 496T189 394T158 260Q156 242 156 221Q156 173 170 136T206 79T256 45T308 28T353 24Q407 24 452 47T514 106Q517 114 529 161T541 214Q541 222 528 224T468 227H431Q425 233 425 235T427 254Q431 267 437 273H454Q494 271 594 271Q634 271 659 271T695 272T707 272Q721 272 721 263Q721 261 719 249Q714 230 709 228Q706 227 694 227Q674 227 653 224Q646 221 643 215T629 164Q620 131 614 108Q589 6 586 3Q584 1 581 1Q571 1 553 21T530 52Q530 53 528 52T522 47Q448 -22 322 -22Q201 -22 126 55T50 252'],
  
      // LATIN CAPITAL LETTER H
      0x48: [683,0,831,31,888,'228 637Q194 637 192 641Q191 643 191 649Q191 673 202 682Q204 683 219 683Q260 681 355 681Q389 681 418 681T463 682T483 682Q499 682 499 672Q499 670 497 658Q492 641 487 638H485Q483 638 480 638T473 638T464 637T455 637Q416 636 405 634T387 623Q384 619 355 500Q348 474 340 442T328 395L324 380Q324 378 469 378H614L615 381Q615 384 646 504Q674 619 674 627T617 637Q594 637 587 639T580 648Q580 650 582 660Q586 677 588 679T604 682Q609 682 646 681T740 680Q802 680 835 681T871 682Q888 682 888 672Q888 645 876 638H874Q872 638 869 638T862 638T853 637T844 637Q805 636 794 634T776 623Q773 618 704 340T634 58Q634 51 638 51Q646 48 692 46H723Q729 38 729 37T726 19Q722 6 716 0H701Q664 2 567 2Q533 2 504 2T458 2T437 1Q420 1 420 10Q420 15 423 24Q428 43 433 45Q437 46 448 46H454Q481 46 514 49Q520 50 522 50T528 55T534 64T540 82T547 110T558 153Q565 181 569 198Q602 330 602 331T457 332H312L279 197Q245 63 245 58Q245 51 253 49T303 46H334Q340 38 340 37T337 19Q333 6 327 0H312Q275 2 178 2Q144 2 115 2T69 2T48 1Q31 1 31 10Q31 12 34 24Q39 43 44 45Q48 46 59 46H65Q92 46 125 49Q139 52 144 61Q147 65 216 339T285 628Q285 635 228 637'],
  
      // LATIN CAPITAL LETTER I
      0x49: [683,0,440,26,504,'43 1Q26 1 26 10Q26 12 29 24Q34 43 39 45Q42 46 54 46H60Q120 46 136 53Q137 53 138 54Q143 56 149 77T198 273Q210 318 216 344Q286 624 286 626Q284 630 284 631Q274 637 213 637H193Q184 643 189 662Q193 677 195 680T209 683H213Q285 681 359 681Q481 681 487 683H497Q504 676 504 672T501 655T494 639Q491 637 471 637Q440 637 407 634Q393 631 388 623Q381 609 337 432Q326 385 315 341Q245 65 245 59Q245 52 255 50T307 46H339Q345 38 345 37T342 19Q338 6 332 0H316Q279 2 179 2Q143 2 113 2T65 2T43 1'],
  
      // LATIN CAPITAL LETTER J
      0x4A: [683,22,555,57,633,'447 625Q447 637 354 637H329Q323 642 323 645T325 664Q329 677 335 683H352Q393 681 498 681Q541 681 568 681T605 682T619 682Q633 682 633 672Q633 670 630 658Q626 642 623 640T604 637Q552 637 545 623Q541 610 483 376Q420 128 419 127Q397 64 333 21T195 -22Q137 -22 97 8T57 88Q57 130 80 152T132 174Q177 174 182 130Q182 98 164 80T123 56Q115 54 115 53T122 44Q148 15 197 15Q235 15 271 47T324 130Q328 142 387 380T447 625'],
  
      // LATIN CAPITAL LETTER K
      0x4B: [683,0,849,31,889,'285 628Q285 635 228 637Q205 637 198 638T191 647Q191 649 193 661Q199 681 203 682Q205 683 214 683H219Q260 681 355 681Q389 681 418 681T463 682T483 682Q500 682 500 674Q500 669 497 660Q496 658 496 654T495 648T493 644T490 641T486 639T479 638T470 637T456 637Q416 636 405 634T387 623L306 305Q307 305 490 449T678 597Q692 611 692 620Q692 635 667 637Q651 637 651 648Q651 650 654 662T659 677Q662 682 676 682Q680 682 711 681T791 680Q814 680 839 681T869 682Q889 682 889 672Q889 650 881 642Q878 637 862 637Q787 632 726 586Q710 576 656 534T556 455L509 418L518 396Q527 374 546 329T581 244Q656 67 661 61Q663 59 666 57Q680 47 717 46H738Q744 38 744 37T741 19Q737 6 731 0H720Q680 3 625 3Q503 3 488 0H478Q472 6 472 9T474 27Q478 40 480 43T491 46H494Q544 46 544 71Q544 75 517 141T485 216L427 354L359 301L291 248L268 155Q245 63 245 58Q245 51 253 49T303 46H334Q340 37 340 35Q340 19 333 5Q328 0 317 0Q314 0 280 1T180 2Q118 2 85 2T49 1Q31 1 31 11Q31 13 34 25Q38 41 42 43T65 46Q92 46 125 49Q139 52 144 61Q147 65 216 339T285 628'],
  
      // LATIN CAPITAL LETTER L
      0x4C: [683,2,681,32,647,'228 637Q194 637 192 641Q191 643 191 649Q191 673 202 682Q204 683 217 683Q271 680 344 680Q485 680 506 683H518Q524 677 524 674T522 656Q517 641 513 637H475Q406 636 394 628Q387 624 380 600T313 336Q297 271 279 198T252 88L243 52Q243 48 252 48T311 46H328Q360 46 379 47T428 54T478 72T522 106T564 161Q580 191 594 228T611 270Q616 273 628 273H641Q647 264 647 262T627 203T583 83T557 9Q555 4 553 3T537 0T494 -1Q483 -1 418 -1T294 0H116Q32 0 32 10Q32 17 34 24Q39 43 44 45Q48 46 59 46H65Q92 46 125 49Q139 52 144 61Q147 65 216 339T285 628Q285 635 228 637'],
  
      // LATIN CAPITAL LETTER M
      0x4D: [684,0,970,35,1051,'289 629Q289 635 232 637Q208 637 201 638T194 648Q194 649 196 659Q197 662 198 666T199 671T201 676T203 679T207 681T212 683T220 683T232 684Q238 684 262 684T307 683Q386 683 398 683T414 678Q415 674 451 396L487 117L510 154Q534 190 574 254T662 394Q837 673 839 675Q840 676 842 678T846 681L852 683H948Q965 683 988 683T1017 684Q1051 684 1051 673Q1051 668 1048 656T1045 643Q1041 637 1008 637Q968 636 957 634T939 623Q936 618 867 340T797 59Q797 55 798 54T805 50T822 48T855 46H886Q892 37 892 35Q892 19 885 5Q880 0 869 0Q864 0 828 1T736 2Q675 2 644 2T609 1Q592 1 592 11Q592 13 594 25Q598 41 602 43T625 46Q652 46 685 49Q699 52 704 61Q706 65 742 207T813 490T848 631L654 322Q458 10 453 5Q451 4 449 3Q444 0 433 0Q418 0 415 7Q413 11 374 317L335 624L267 354Q200 88 200 79Q206 46 272 46H282Q288 41 289 37T286 19Q282 3 278 1Q274 0 267 0Q265 0 255 0T221 1T157 2Q127 2 95 1T58 0Q43 0 39 2T35 11Q35 13 38 25T43 40Q45 46 65 46Q135 46 154 86Q158 92 223 354T289 629'],
  
      // LATIN CAPITAL LETTER N
      0x4E: [683,0,803,31,888,'234 637Q231 637 226 637Q201 637 196 638T191 649Q191 676 202 682Q204 683 299 683Q376 683 387 683T401 677Q612 181 616 168L670 381Q723 592 723 606Q723 633 659 637Q635 637 635 648Q635 650 637 660Q641 676 643 679T653 683Q656 683 684 682T767 680Q817 680 843 681T873 682Q888 682 888 672Q888 650 880 642Q878 637 858 637Q787 633 769 597L620 7Q618 0 599 0Q585 0 582 2Q579 5 453 305L326 604L261 344Q196 88 196 79Q201 46 268 46H278Q284 41 284 38T282 19Q278 6 272 0H259Q228 2 151 2Q123 2 100 2T63 2T46 1Q31 1 31 10Q31 14 34 26T39 40Q41 46 62 46Q130 49 150 85Q154 91 221 362L289 634Q287 635 234 637'],
  
      // LATIN CAPITAL LETTER O
      0x4F: [704,22,763,50,740,'740 435Q740 320 676 213T511 42T304 -22Q207 -22 138 35T51 201Q50 209 50 244Q50 346 98 438T227 601Q351 704 476 704Q514 704 524 703Q621 689 680 617T740 435ZM637 476Q637 565 591 615T476 665Q396 665 322 605Q242 542 200 428T157 216Q157 126 200 73T314 19Q404 19 485 98T608 313Q637 408 637 476'],
  
      // LATIN CAPITAL LETTER P
      0x50: [683,0,642,33,751,'287 628Q287 635 230 637Q206 637 199 638T192 648Q192 649 194 659Q200 679 203 681T397 683Q587 682 600 680Q664 669 707 631T751 530Q751 453 685 389Q616 321 507 303Q500 302 402 301H307L277 182Q247 66 247 59Q247 55 248 54T255 50T272 48T305 46H336Q342 37 342 35Q342 19 335 5Q330 0 319 0Q316 0 282 1T182 2Q120 2 87 2T51 1Q33 1 33 11Q33 13 36 25Q40 41 44 43T67 46Q94 46 127 49Q141 52 146 61Q149 65 218 339T287 628ZM645 554Q645 567 643 575T634 597T609 619T560 635Q553 636 480 637Q463 637 445 637T416 636T404 636Q391 635 386 627Q384 621 367 550T332 412T314 344Q314 342 395 342H407H430Q542 342 590 392Q617 419 631 471T645 554'],
  
      // LATIN CAPITAL LETTER Q
      0x51: [704,194,791,50,740,'399 -80Q399 -47 400 -30T402 -11V-7L387 -11Q341 -22 303 -22Q208 -22 138 35T51 201Q50 209 50 244Q50 346 98 438T227 601Q351 704 476 704Q514 704 524 703Q621 689 680 617T740 435Q740 255 592 107Q529 47 461 16L444 8V3Q444 2 449 -24T470 -66T516 -82Q551 -82 583 -60T625 -3Q631 11 638 11Q647 11 649 2Q649 -6 639 -34T611 -100T557 -165T481 -194Q399 -194 399 -87V-80ZM636 468Q636 523 621 564T580 625T530 655T477 665Q429 665 379 640Q277 591 215 464T153 216Q153 110 207 59Q231 38 236 38V46Q236 86 269 120T347 155Q372 155 390 144T417 114T429 82T435 55L448 64Q512 108 557 185T619 334T636 468ZM314 18Q362 18 404 39L403 49Q399 104 366 115Q354 117 347 117Q344 117 341 117T337 118Q317 118 296 98T274 52Q274 18 314 18'],
  
      // LATIN CAPITAL LETTER R
      0x52: [683,21,759,33,755,'230 637Q203 637 198 638T193 649Q193 676 204 682Q206 683 378 683Q550 682 564 680Q620 672 658 652T712 606T733 563T739 529Q739 484 710 445T643 385T576 351T538 338L545 333Q612 295 612 223Q612 212 607 162T602 80V71Q602 53 603 43T614 25T640 16Q668 16 686 38T712 85Q717 99 720 102T735 105Q755 105 755 93Q755 75 731 36Q693 -21 641 -21H632Q571 -21 531 4T487 82Q487 109 502 166T517 239Q517 290 474 313Q459 320 449 321T378 323H309L277 193Q244 61 244 59Q244 55 245 54T252 50T269 48T302 46H333Q339 38 339 37T336 19Q332 6 326 0H311Q275 2 180 2Q146 2 117 2T71 2T50 1Q33 1 33 10Q33 12 36 24Q41 43 46 45Q50 46 61 46H67Q94 46 127 49Q141 52 146 61Q149 65 218 339T287 628Q287 635 230 637ZM630 554Q630 586 609 608T523 636Q521 636 500 636T462 637H440Q393 637 386 627Q385 624 352 494T319 361Q319 360 388 360Q466 361 492 367Q556 377 592 426Q608 449 619 486T630 554'],
  
      // LATIN CAPITAL LETTER S
      0x53: [705,22,613,52,645,'308 24Q367 24 416 76T466 197Q466 260 414 284Q308 311 278 321T236 341Q176 383 176 462Q176 523 208 573T273 648Q302 673 343 688T407 704H418H425Q521 704 564 640Q565 640 577 653T603 682T623 704Q624 704 627 704T632 705Q645 705 645 698T617 577T585 459T569 456Q549 456 549 465Q549 471 550 475Q550 478 551 494T553 520Q553 554 544 579T526 616T501 641Q465 662 419 662Q362 662 313 616T263 510Q263 480 278 458T319 427Q323 425 389 408T456 390Q490 379 522 342T554 242Q554 216 546 186Q541 164 528 137T492 78T426 18T332 -20Q320 -22 298 -22Q199 -22 144 33L134 44L106 13Q83 -14 78 -18T65 -22Q52 -22 52 -14Q52 -11 110 221Q112 227 130 227H143Q149 221 149 216Q149 214 148 207T144 186T142 153Q144 114 160 87T203 47T255 29T308 24'],
  
      // LATIN CAPITAL LETTER T
      0x54: [677,0,584,21,704,'40 437Q21 437 21 445Q21 450 37 501T71 602L88 651Q93 669 101 677H569H659Q691 677 697 676T704 667Q704 661 687 553T668 444Q668 437 649 437Q640 437 637 437T631 442L629 445Q629 451 635 490T641 551Q641 586 628 604T573 629Q568 630 515 631Q469 631 457 630T439 622Q438 621 368 343T298 60Q298 48 386 46Q418 46 427 45T436 36Q436 31 433 22Q429 4 424 1L422 0Q419 0 415 0Q410 0 363 1T228 2Q99 2 64 0H49Q43 6 43 9T45 27Q49 40 55 46H83H94Q174 46 189 55Q190 56 191 56Q196 59 201 76T241 233Q258 301 269 344Q339 619 339 625Q339 630 310 630H279Q212 630 191 624Q146 614 121 583T67 467Q60 445 57 441T43 437H40'],
  
      // LATIN CAPITAL LETTER U
      0x55: [683,22,683,60,767,'107 637Q73 637 71 641Q70 643 70 649Q70 673 81 682Q83 683 98 683Q139 681 234 681Q268 681 297 681T342 682T362 682Q378 682 378 672Q378 670 376 658Q371 641 366 638H364Q362 638 359 638T352 638T343 637T334 637Q295 636 284 634T266 623Q265 621 238 518T184 302T154 169Q152 155 152 140Q152 86 183 55T269 24Q336 24 403 69T501 205L552 406Q599 598 599 606Q599 633 535 637Q511 637 511 648Q511 650 513 660Q517 676 519 679T529 683Q532 683 561 682T645 680Q696 680 723 681T752 682Q767 682 767 672Q767 650 759 642Q756 637 737 637Q666 633 648 597Q646 592 598 404Q557 235 548 205Q515 105 433 42T263 -22Q171 -22 116 34T60 167V183Q60 201 115 421Q164 622 164 628Q164 635 107 637'],
  
      // LATIN CAPITAL LETTER V
      0x56: [683,22,583,52,769,'52 648Q52 670 65 683H76Q118 680 181 680Q299 680 320 683H330Q336 677 336 674T334 656Q329 641 325 637H304Q282 635 274 635Q245 630 242 620Q242 618 271 369T301 118L374 235Q447 352 520 471T595 594Q599 601 599 609Q599 633 555 637Q537 637 537 648Q537 649 539 661Q542 675 545 679T558 683Q560 683 570 683T604 682T668 681Q737 681 755 683H762Q769 676 769 672Q769 655 760 640Q757 637 743 637Q730 636 719 635T698 630T682 623T670 615T660 608T652 599T645 592L452 282Q272 -9 266 -16Q263 -18 259 -21L241 -22H234Q216 -22 216 -15Q213 -9 177 305Q139 623 138 626Q133 637 76 637H59Q52 642 52 648'],
  
      // LATIN CAPITAL LETTER W
      0x57: [683,22,944,51,1048,'436 683Q450 683 486 682T553 680Q604 680 638 681T677 682Q695 682 695 674Q695 670 692 659Q687 641 683 639T661 637Q636 636 621 632T600 624T597 615Q597 603 613 377T629 138L631 141Q633 144 637 151T649 170T666 200T690 241T720 295T759 362Q863 546 877 572T892 604Q892 619 873 628T831 637Q817 637 817 647Q817 650 819 660Q823 676 825 679T839 682Q842 682 856 682T895 682T949 681Q1015 681 1034 683Q1048 683 1048 672Q1048 666 1045 655T1038 640T1028 637Q1006 637 988 631T958 617T939 600T927 584L923 578L754 282Q586 -14 585 -15Q579 -22 561 -22Q546 -22 542 -17Q539 -14 523 229T506 480L494 462Q472 425 366 239Q222 -13 220 -15T215 -19Q210 -22 197 -22Q178 -22 176 -15Q176 -12 154 304T131 622Q129 631 121 633T82 637H58Q51 644 51 648Q52 671 64 683H76Q118 680 176 680Q301 680 313 683H323Q329 677 329 674T327 656Q322 641 318 637H297Q236 634 232 620Q262 160 266 136L501 550L499 587Q496 629 489 632Q483 636 447 637Q428 637 422 639T416 648Q416 650 418 660Q419 664 420 669T421 676T424 680T428 682T436 683'],
  
      // LATIN CAPITAL LETTER X
      0x58: [683,0,828,26,852,'42 0H40Q26 0 26 11Q26 15 29 27Q33 41 36 43T55 46Q141 49 190 98Q200 108 306 224T411 342Q302 620 297 625Q288 636 234 637H206Q200 643 200 645T202 664Q206 677 212 683H226Q260 681 347 681Q380 681 408 681T453 682T473 682Q490 682 490 671Q490 670 488 658Q484 643 481 640T465 637Q434 634 411 620L488 426L541 485Q646 598 646 610Q646 628 622 635Q617 635 609 637Q594 637 594 648Q594 650 596 664Q600 677 606 683H618Q619 683 643 683T697 681T738 680Q828 680 837 683H845Q852 676 852 672Q850 647 840 637H824Q790 636 763 628T722 611T698 593L687 584Q687 585 592 480L505 384Q505 383 536 304T601 142T638 56Q648 47 699 46Q734 46 734 37Q734 35 732 23Q728 7 725 4T711 1Q708 1 678 1T589 2Q528 2 496 2T461 1Q444 1 444 10Q444 11 446 25Q448 35 450 39T455 44T464 46T480 47T506 54Q523 62 523 64Q522 64 476 181L429 299Q241 95 236 84Q232 76 232 72Q232 53 261 47Q262 47 267 47T273 46Q276 46 277 46T280 45T283 42T284 35Q284 26 282 19Q279 6 276 4T261 1Q258 1 243 1T201 2T142 2Q64 2 42 0'],
  
      // LATIN CAPITAL LETTER Y
      0x59: [683,-1,581,30,763,'66 637Q54 637 49 637T39 638T32 641T30 647T33 664T42 682Q44 683 56 683Q104 680 165 680Q288 680 306 683H316Q322 677 322 674T320 656Q316 643 310 637H298Q242 637 242 624Q242 619 292 477T343 333L346 336Q350 340 358 349T379 373T411 410T454 461Q546 568 561 587T577 618Q577 634 545 637Q528 637 528 647Q528 649 530 661Q533 676 535 679T549 683Q551 683 578 682T657 680Q684 680 713 681T746 682Q763 682 763 673Q763 669 760 657T755 643Q753 637 734 637Q662 632 617 587Q608 578 477 424L348 273L322 169Q295 62 295 57Q295 46 363 46Q379 46 384 45T390 35Q390 33 388 23Q384 6 382 4T366 1Q361 1 324 1T232 2Q170 2 138 2T102 1Q84 1 84 9Q84 14 87 24Q88 27 89 30T90 35T91 39T93 42T96 44T101 45T107 45T116 46T129 46Q168 47 180 50T198 63Q201 68 227 171L252 274L129 623Q128 624 127 625T125 627T122 629T118 631T113 633T105 634T96 635T83 636T66 637'],
  
      // LATIN CAPITAL LETTER Z
      0x5A: [683,0,683,58,723,'58 8Q58 23 64 35Q64 36 329 334T596 635L586 637Q575 637 512 637H500H476Q442 637 420 635T365 624T311 598T266 548T228 469Q227 466 226 463T224 458T223 453T222 450L221 448Q218 443 202 443Q185 443 182 453L214 561Q228 606 241 651Q249 679 253 681Q256 683 487 683H718Q723 678 723 675Q723 673 717 649Q189 54 188 52L185 49H274Q369 50 377 51Q452 60 500 100T579 247Q587 272 590 277T603 282H607Q628 282 628 271Q547 5 541 2Q538 0 300 0H124Q58 0 58 8'],
  
      // LATIN SMALL LETTER A
      0x61: [441,10,529,33,506,'33 157Q33 258 109 349T280 441Q331 441 370 392Q386 422 416 422Q429 422 439 414T449 394Q449 381 412 234T374 68Q374 43 381 35T402 26Q411 27 422 35Q443 55 463 131Q469 151 473 152Q475 153 483 153H487Q506 153 506 144Q506 138 501 117T481 63T449 13Q436 0 417 -8Q409 -10 393 -10Q359 -10 336 5T306 36L300 51Q299 52 296 50Q294 48 292 46Q233 -10 172 -10Q117 -10 75 30T33 157ZM351 328Q351 334 346 350T323 385T277 405Q242 405 210 374T160 293Q131 214 119 129Q119 126 119 118T118 106Q118 61 136 44T179 26Q217 26 254 59T298 110Q300 114 325 217T351 328'],
  
      // LATIN SMALL LETTER B
      0x62: [694,11,429,40,422,'73 647Q73 657 77 670T89 683Q90 683 161 688T234 694Q246 694 246 685T212 542Q204 508 195 472T180 418L176 399Q176 396 182 402Q231 442 283 442Q345 442 383 396T422 280Q422 169 343 79T173 -11Q123 -11 82 27T40 150V159Q40 180 48 217T97 414Q147 611 147 623T109 637Q104 637 101 637H96Q86 637 83 637T76 640T73 647ZM336 325V331Q336 405 275 405Q258 405 240 397T207 376T181 352T163 330L157 322L136 236Q114 150 114 114Q114 66 138 42Q154 26 178 26Q211 26 245 58Q270 81 285 114T318 219Q336 291 336 325'],
  
      // LATIN SMALL LETTER C
      0x63: [442,12,433,34,430,'34 159Q34 268 120 355T306 442Q362 442 394 418T427 355Q427 326 408 306T360 285Q341 285 330 295T319 325T330 359T352 380T366 386H367Q367 388 361 392T340 400T306 404Q276 404 249 390Q228 381 206 359Q162 315 142 235T121 119Q121 73 147 50Q169 26 205 26H209Q321 26 394 111Q403 121 406 121Q410 121 419 112T429 98T420 83T391 55T346 25T282 0T202 -11Q127 -11 81 37T34 159'],
  
      // LATIN SMALL LETTER D
      0x64: [694,10,520,33,523,'366 683Q367 683 438 688T511 694Q523 694 523 686Q523 679 450 384T375 83T374 68Q374 26 402 26Q411 27 422 35Q443 55 463 131Q469 151 473 152Q475 153 483 153H487H491Q506 153 506 145Q506 140 503 129Q490 79 473 48T445 8T417 -8Q409 -10 393 -10Q359 -10 336 5T306 36L300 51Q299 52 296 50Q294 48 292 46Q233 -10 172 -10Q117 -10 75 30T33 157Q33 205 53 255T101 341Q148 398 195 420T280 442Q336 442 364 400Q369 394 369 396Q370 400 396 505T424 616Q424 629 417 632T378 637H357Q351 643 351 645T353 664Q358 683 366 683ZM352 326Q329 405 277 405Q242 405 210 374T160 293Q131 214 119 129Q119 126 119 118T118 106Q118 61 136 44T179 26Q233 26 290 98L298 109L352 326'],
  
      // LATIN SMALL LETTER E
      0x65: [443,11,466,39,430,'39 168Q39 225 58 272T107 350T174 402T244 433T307 442H310Q355 442 388 420T421 355Q421 265 310 237Q261 224 176 223Q139 223 138 221Q138 219 132 186T125 128Q125 81 146 54T209 26T302 45T394 111Q403 121 406 121Q410 121 419 112T429 98T420 82T390 55T344 24T281 -1T205 -11Q126 -11 83 42T39 168ZM373 353Q367 405 305 405Q272 405 244 391T199 357T170 316T154 280T149 261Q149 260 169 260Q282 260 327 284T373 353'],
  
      // LATIN SMALL LETTER F
      0x66: [705,205,490,55,550,'118 -162Q120 -162 124 -164T135 -167T147 -168Q160 -168 171 -155T187 -126Q197 -99 221 27T267 267T289 382V385H242Q195 385 192 387Q188 390 188 397L195 425Q197 430 203 430T250 431Q298 431 298 432Q298 434 307 482T319 540Q356 705 465 705Q502 703 526 683T550 630Q550 594 529 578T487 561Q443 561 443 603Q443 622 454 636T478 657L487 662Q471 668 457 668Q445 668 434 658T419 630Q412 601 403 552T387 469T380 433Q380 431 435 431Q480 431 487 430T498 424Q499 420 496 407T491 391Q489 386 482 386T428 385H372L349 263Q301 15 282 -47Q255 -132 212 -173Q175 -205 139 -205Q107 -205 81 -186T55 -132Q55 -95 76 -78T118 -61Q162 -61 162 -103Q162 -122 151 -136T127 -157L118 -162'],
  
      // LATIN SMALL LETTER G
      0x67: [442,205,477,10,480,'311 43Q296 30 267 15T206 0Q143 0 105 45T66 160Q66 265 143 353T314 442Q361 442 401 394L404 398Q406 401 409 404T418 412T431 419T447 422Q461 422 470 413T480 394Q480 379 423 152T363 -80Q345 -134 286 -169T151 -205Q10 -205 10 -137Q10 -111 28 -91T74 -71Q89 -71 102 -80T116 -111Q116 -121 114 -130T107 -144T99 -154T92 -162L90 -164H91Q101 -167 151 -167Q189 -167 211 -155Q234 -144 254 -122T282 -75Q288 -56 298 -13Q311 35 311 43ZM384 328L380 339Q377 350 375 354T369 368T359 382T346 393T328 402T306 405Q262 405 221 352Q191 313 171 233T151 117Q151 38 213 38Q269 38 323 108L331 118L384 328'],
  
      // LATIN SMALL LETTER H
      0x68: [694,11,576,48,555,'137 683Q138 683 209 688T282 694Q294 694 294 685Q294 674 258 534Q220 386 220 383Q220 381 227 388Q288 442 357 442Q411 442 444 415T478 336Q478 285 440 178T402 50Q403 36 407 31T422 26Q450 26 474 56T513 138Q516 149 519 151T535 153Q555 153 555 145Q555 144 551 130Q535 71 500 33Q466 -10 419 -10H414Q367 -10 346 17T325 74Q325 90 361 192T398 345Q398 404 354 404H349Q266 404 205 306L198 293L164 158Q132 28 127 16Q114 -11 83 -11Q69 -11 59 -2T48 16Q48 30 121 320L195 616Q195 629 188 632T149 637H128Q122 643 122 645T124 664Q129 683 137 683'],
  
      // LATIN SMALL LETTER I
      0x69: [661,11,345,21,302,'184 600Q184 624 203 642T247 661Q265 661 277 649T290 619Q290 596 270 577T226 557Q211 557 198 567T184 600ZM21 287Q21 295 30 318T54 369T98 420T158 442Q197 442 223 419T250 357Q250 340 236 301T196 196T154 83Q149 61 149 51Q149 26 166 26Q175 26 185 29T208 43T235 78T260 137Q263 149 265 151T282 153Q302 153 302 143Q302 135 293 112T268 61T223 11T161 -11Q129 -11 102 10T74 74Q74 91 79 106T122 220Q160 321 166 341T173 380Q173 404 156 404H154Q124 404 99 371T61 287Q60 286 59 284T58 281T56 279T53 278T49 278T41 278H27Q21 284 21 287'],
  
      // LATIN SMALL LETTER J
      0x6A: [661,204,412,-12,403,'297 596Q297 627 318 644T361 661Q378 661 389 651T403 623Q403 595 384 576T340 557Q322 557 310 567T297 596ZM288 376Q288 405 262 405Q240 405 220 393T185 362T161 325T144 293L137 279Q135 278 121 278H107Q101 284 101 286T105 299Q126 348 164 391T252 441Q253 441 260 441T272 442Q296 441 316 432Q341 418 354 401T367 348V332L318 133Q267 -67 264 -75Q246 -125 194 -164T75 -204Q25 -204 7 -183T-12 -137Q-12 -110 7 -91T53 -71Q70 -71 82 -81T95 -112Q95 -148 63 -167Q69 -168 77 -168Q111 -168 139 -140T182 -74L193 -32Q204 11 219 72T251 197T278 308T289 365Q289 372 288 376'],
  
      // LATIN SMALL LETTER K
      0x6B: [694,11,521,48,503,'121 647Q121 657 125 670T137 683Q138 683 209 688T282 694Q294 694 294 686Q294 679 244 477Q194 279 194 272Q213 282 223 291Q247 309 292 354T362 415Q402 442 438 442Q468 442 485 423T503 369Q503 344 496 327T477 302T456 291T438 288Q418 288 406 299T394 328Q394 353 410 369T442 390L458 393Q446 405 434 405H430Q398 402 367 380T294 316T228 255Q230 254 243 252T267 246T293 238T320 224T342 206T359 180T365 147Q365 130 360 106T354 66Q354 26 381 26Q429 26 459 145Q461 153 479 153H483Q499 153 499 144Q499 139 496 130Q455 -11 378 -11Q333 -11 305 15T277 90Q277 108 280 121T283 145Q283 167 269 183T234 206T200 217T182 220H180Q168 178 159 139T145 81T136 44T129 20T122 7T111 -2Q98 -11 83 -11Q66 -11 57 -1T48 16Q48 26 85 176T158 471L195 616Q196 629 188 632T149 637H144Q134 637 131 637T124 640T121 647'],
  
      // LATIN SMALL LETTER L
      0x6C: [695,12,298,38,266,'117 59Q117 26 142 26Q179 26 205 131Q211 151 215 152Q217 153 225 153H229Q238 153 241 153T246 151T248 144Q247 138 245 128T234 90T214 43T183 6T137 -11Q101 -11 70 11T38 85Q38 97 39 102L104 360Q167 615 167 623Q167 626 166 628T162 632T157 634T149 635T141 636T132 637T122 637Q112 637 109 637T101 638T95 641T94 647Q94 649 96 661Q101 680 107 682T179 688Q194 689 213 690T243 693T254 694Q266 694 266 686Q266 675 193 386T118 83Q118 81 118 75T117 65V59'],
  
      // LATIN SMALL LETTER M
      0x6D: [443,11,878,21,857,'21 287Q22 293 24 303T36 341T56 388T88 425T132 442T175 435T205 417T221 395T229 376L231 369Q231 367 232 367L243 378Q303 442 384 442Q401 442 415 440T441 433T460 423T475 411T485 398T493 385T497 373T500 364T502 357L510 367Q573 442 659 442Q713 442 746 415T780 336Q780 285 742 178T704 50Q705 36 709 31T724 26Q752 26 776 56T815 138Q818 149 821 151T837 153Q857 153 857 145Q857 144 853 130Q845 101 831 73T785 17T716 -10Q669 -10 648 17T627 73Q627 92 663 193T700 345Q700 404 656 404H651Q565 404 506 303L499 291L466 157Q433 26 428 16Q415 -11 385 -11Q372 -11 364 -4T353 8T350 18Q350 29 384 161L420 307Q423 322 423 345Q423 404 379 404H374Q288 404 229 303L222 291L189 157Q156 26 151 16Q138 -11 108 -11Q95 -11 87 -5T76 7T74 17Q74 30 112 181Q151 335 151 342Q154 357 154 369Q154 405 129 405Q107 405 92 377T69 316T57 280Q55 278 41 278H27Q21 284 21 287'],
  
      // LATIN SMALL LETTER N
      0x6E: [443,11,600,21,580,'21 287Q22 293 24 303T36 341T56 388T89 425T135 442Q171 442 195 424T225 390T231 369Q231 367 232 367L243 378Q304 442 382 442Q436 442 469 415T503 336T465 179T427 52Q427 26 444 26Q450 26 453 27Q482 32 505 65T540 145Q542 153 560 153Q580 153 580 145Q580 144 576 130Q568 101 554 73T508 17T439 -10Q392 -10 371 17T350 73Q350 92 386 193T423 345Q423 404 379 404H374Q288 404 229 303L222 291L189 157Q156 26 151 16Q138 -11 108 -11Q95 -11 87 -5T76 7T74 17Q74 30 112 180T152 343Q153 348 153 366Q153 405 129 405Q91 405 66 305Q60 285 60 284Q58 278 41 278H27Q21 284 21 287'],
  
      // LATIN SMALL LETTER O
      0x6F: [441,11,485,34,476,'201 -11Q126 -11 80 38T34 156Q34 221 64 279T146 380Q222 441 301 441Q333 441 341 440Q354 437 367 433T402 417T438 387T464 338T476 268Q476 161 390 75T201 -11ZM121 120Q121 70 147 48T206 26Q250 26 289 58T351 142Q360 163 374 216T388 308Q388 352 370 375Q346 405 306 405Q243 405 195 347Q158 303 140 230T121 120'],
  
      // LATIN SMALL LETTER P
      0x70: [443,194,503,-39,497,'23 287Q24 290 25 295T30 317T40 348T55 381T75 411T101 433T134 442Q209 442 230 378L240 387Q302 442 358 442Q423 442 460 395T497 281Q497 173 421 82T249 -10Q227 -10 210 -4Q199 1 187 11T168 28L161 36Q160 35 139 -51T118 -138Q118 -144 126 -145T163 -148H188Q194 -155 194 -157T191 -175Q188 -187 185 -190T172 -194Q170 -194 161 -194T127 -193T65 -192Q-5 -192 -24 -194H-32Q-39 -187 -39 -183Q-37 -156 -26 -148H-6Q28 -147 33 -136Q36 -130 94 103T155 350Q156 355 156 364Q156 405 131 405Q109 405 94 377T71 316T59 280Q57 278 43 278H29Q23 284 23 287ZM178 102Q200 26 252 26Q282 26 310 49T356 107Q374 141 392 215T411 325V331Q411 405 350 405Q339 405 328 402T306 393T286 380T269 365T254 350T243 336T235 326L232 322Q232 321 229 308T218 264T204 212Q178 106 178 102'],
  
      // LATIN SMALL LETTER Q
      0x71: [442,194,446,33,460,'33 157Q33 258 109 349T280 441Q340 441 372 389Q373 390 377 395T388 406T404 418Q438 442 450 442Q454 442 457 439T460 434Q460 425 391 149Q320 -135 320 -139Q320 -147 365 -148H390Q396 -156 396 -157T393 -175Q389 -188 383 -194H370Q339 -192 262 -192Q234 -192 211 -192T174 -192T157 -193Q143 -193 143 -185Q143 -182 145 -170Q149 -154 152 -151T172 -148Q220 -148 230 -141Q238 -136 258 -53T279 32Q279 33 272 29Q224 -10 172 -10Q117 -10 75 30T33 157ZM352 326Q329 405 277 405Q242 405 210 374T160 293Q131 214 119 129Q119 126 119 118T118 106Q118 61 136 44T179 26Q233 26 290 98L298 109L352 326'],
  
      // LATIN SMALL LETTER R
      0x72: [443,11,451,21,430,'21 287Q22 290 23 295T28 317T38 348T53 381T73 411T99 433T132 442Q161 442 183 430T214 408T225 388Q227 382 228 382T236 389Q284 441 347 441H350Q398 441 422 400Q430 381 430 363Q430 333 417 315T391 292T366 288Q346 288 334 299T322 328Q322 376 378 392Q356 405 342 405Q286 405 239 331Q229 315 224 298T190 165Q156 25 151 16Q138 -11 108 -11Q95 -11 87 -5T76 7T74 17Q74 30 114 189T154 366Q154 405 128 405Q107 405 92 377T68 316T57 280Q55 278 41 278H27Q21 284 21 287'],
  
      // LATIN SMALL LETTER S
      0x73: [443,10,469,53,419,'131 289Q131 321 147 354T203 415T300 442Q362 442 390 415T419 355Q419 323 402 308T364 292Q351 292 340 300T328 326Q328 342 337 354T354 372T367 378Q368 378 368 379Q368 382 361 388T336 399T297 405Q249 405 227 379T204 326Q204 301 223 291T278 274T330 259Q396 230 396 163Q396 135 385 107T352 51T289 7T195 -10Q118 -10 86 19T53 87Q53 126 74 143T118 160Q133 160 146 151T160 120Q160 94 142 76T111 58Q109 57 108 57T107 55Q108 52 115 47T146 34T201 27Q237 27 263 38T301 66T318 97T323 122Q323 150 302 164T254 181T195 196T148 231Q131 256 131 289'],
  
      // LATIN SMALL LETTER T
      0x74: [626,11,361,19,330,'26 385Q19 392 19 395Q19 399 22 411T27 425Q29 430 36 430T87 431H140L159 511Q162 522 166 540T173 566T179 586T187 603T197 615T211 624T229 626Q247 625 254 615T261 596Q261 589 252 549T232 470L222 433Q222 431 272 431H323Q330 424 330 420Q330 398 317 385H210L174 240Q135 80 135 68Q135 26 162 26Q197 26 230 60T283 144Q285 150 288 151T303 153H307Q322 153 322 145Q322 142 319 133Q314 117 301 95T267 48T216 6T155 -11Q125 -11 98 4T59 56Q57 64 57 83V101L92 241Q127 382 128 383Q128 385 77 385H26'],
  
      // LATIN SMALL LETTER U
      0x75: [442,11,572,21,551,'21 287Q21 295 30 318T55 370T99 420T158 442Q204 442 227 417T250 358Q250 340 216 246T182 105Q182 62 196 45T238 27T291 44T328 78L339 95Q341 99 377 247Q407 367 413 387T427 416Q444 431 463 431Q480 431 488 421T496 402L420 84Q419 79 419 68Q419 43 426 35T447 26Q469 29 482 57T512 145Q514 153 532 153Q551 153 551 144Q550 139 549 130T540 98T523 55T498 17T462 -8Q454 -10 438 -10Q372 -10 347 46Q345 45 336 36T318 21T296 6T267 -6T233 -11Q189 -11 155 7Q103 38 103 113Q103 170 138 262T173 379Q173 380 173 381Q173 390 173 393T169 400T158 404H154Q131 404 112 385T82 344T65 302T57 280Q55 278 41 278H27Q21 284 21 287'],
  
      // LATIN SMALL LETTER V
      0x76: [443,11,485,21,467,'173 380Q173 405 154 405Q130 405 104 376T61 287Q60 286 59 284T58 281T56 279T53 278T49 278T41 278H27Q21 284 21 287Q21 294 29 316T53 368T97 419T160 441Q202 441 225 417T249 361Q249 344 246 335Q246 329 231 291T200 202T182 113Q182 86 187 69Q200 26 250 26Q287 26 319 60T369 139T398 222T409 277Q409 300 401 317T383 343T365 361T357 383Q357 405 376 424T417 443Q436 443 451 425T467 367Q467 340 455 284T418 159T347 40T241 -11Q177 -11 139 22Q102 54 102 117Q102 148 110 181T151 298Q173 362 173 380'],
  
      // LATIN SMALL LETTER W
      0x77: [443,11,716,21,690,'580 385Q580 406 599 424T641 443Q659 443 674 425T690 368Q690 339 671 253Q656 197 644 161T609 80T554 12T482 -11Q438 -11 404 5T355 48Q354 47 352 44Q311 -11 252 -11Q226 -11 202 -5T155 14T118 53T104 116Q104 170 138 262T173 379Q173 380 173 381Q173 390 173 393T169 400T158 404H154Q131 404 112 385T82 344T65 302T57 280Q55 278 41 278H27Q21 284 21 287Q21 293 29 315T52 366T96 418T161 441Q204 441 227 416T250 358Q250 340 217 250T184 111Q184 65 205 46T258 26Q301 26 334 87L339 96V119Q339 122 339 128T340 136T341 143T342 152T345 165T348 182T354 206T362 238T373 281Q402 395 406 404Q419 431 449 431Q468 431 475 421T483 402Q483 389 454 274T422 142Q420 131 420 107V100Q420 85 423 71T442 42T487 26Q558 26 600 148Q609 171 620 213T632 273Q632 306 619 325T593 357T580 385'],
  
      // LATIN SMALL LETTER X
      0x78: [442,11,572,35,522,'52 289Q59 331 106 386T222 442Q257 442 286 424T329 379Q371 442 430 442Q467 442 494 420T522 361Q522 332 508 314T481 292T458 288Q439 288 427 299T415 328Q415 374 465 391Q454 404 425 404Q412 404 406 402Q368 386 350 336Q290 115 290 78Q290 50 306 38T341 26Q378 26 414 59T463 140Q466 150 469 151T485 153H489Q504 153 504 145Q504 144 502 134Q486 77 440 33T333 -11Q263 -11 227 52Q186 -10 133 -10H127Q78 -10 57 16T35 71Q35 103 54 123T99 143Q142 143 142 101Q142 81 130 66T107 46T94 41L91 40Q91 39 97 36T113 29T132 26Q168 26 194 71Q203 87 217 139T245 247T261 313Q266 340 266 352Q266 380 251 392T217 404Q177 404 142 372T93 290Q91 281 88 280T72 278H58Q52 284 52 289'],
  
      // LATIN SMALL LETTER Y
      0x79: [443,205,490,21,497,'21 287Q21 301 36 335T84 406T158 442Q199 442 224 419T250 355Q248 336 247 334Q247 331 231 288T198 191T182 105Q182 62 196 45T238 27Q261 27 281 38T312 61T339 94Q339 95 344 114T358 173T377 247Q415 397 419 404Q432 431 462 431Q475 431 483 424T494 412T496 403Q496 390 447 193T391 -23Q363 -106 294 -155T156 -205Q111 -205 77 -183T43 -117Q43 -95 50 -80T69 -58T89 -48T106 -45Q150 -45 150 -87Q150 -107 138 -122T115 -142T102 -147L99 -148Q101 -153 118 -160T152 -167H160Q177 -167 186 -165Q219 -156 247 -127T290 -65T313 -9T321 21L315 17Q309 13 296 6T270 -6Q250 -11 231 -11Q185 -11 150 11T104 82Q103 89 103 113Q103 170 138 262T173 379Q173 380 173 381Q173 390 173 393T169 400T158 404H154Q131 404 112 385T82 344T65 302T57 280Q55 278 41 278H27Q21 284 21 287'],
  
      // LATIN SMALL LETTER Z
      0x7A: [442,11,465,35,468,'347 338Q337 338 294 349T231 360Q211 360 197 356T174 346T162 335T155 324L153 320Q150 317 138 317Q117 317 117 325Q117 330 120 339Q133 378 163 406T229 440Q241 442 246 442Q271 442 291 425T329 392T367 375Q389 375 411 408T434 441Q435 442 449 442H462Q468 436 468 434Q468 430 463 420T449 399T432 377T418 358L411 349Q368 298 275 214T160 106L148 94L163 93Q185 93 227 82T290 71Q328 71 360 90T402 140Q406 149 409 151T424 153Q443 153 443 143Q443 138 442 134Q425 72 376 31T278 -11Q252 -11 232 6T193 40T155 57Q111 57 76 -3Q70 -11 59 -11H54H41Q35 -5 35 -2Q35 13 93 84Q132 129 225 214T340 322Q352 338 347 338'],
  
      // GREEK CAPITAL LETTER GAMMA
      0x393: [680,-1,615,31,721,'49 1Q31 1 31 10Q31 12 34 24Q39 43 44 45Q48 46 59 46H65Q92 46 125 49Q139 52 144 61Q146 66 215 342T285 622Q285 629 281 629Q273 632 228 634H197Q191 640 191 642T193 661Q197 674 203 680H714Q721 676 721 669Q721 664 708 557T694 447Q692 440 674 440H662Q655 445 655 454Q655 455 658 480T661 534Q661 572 652 592Q638 619 603 626T501 634H471Q398 633 393 630Q389 628 386 622Q385 619 315 341T245 60Q245 46 333 46H345Q366 46 366 35Q366 33 363 21T358 6Q356 1 339 1Q334 1 292 1T187 2Q122 2 88 2T49 1'],
  
      // GREEK CAPITAL LETTER THETA
      0x398: [704,22,763,50,740,'740 435Q740 320 676 213T511 42T304 -22Q207 -22 138 35T51 201Q50 209 50 244Q50 346 98 438T227 601Q351 704 476 704Q514 704 524 703Q621 689 680 617T740 435ZM640 466Q640 523 625 565T583 628T532 658T479 668Q370 668 273 559T151 255Q150 245 150 213Q150 156 165 116T207 55T259 26T313 17Q385 17 451 63T561 184Q590 234 615 312T640 466ZM510 276Q510 278 512 288L515 298Q515 299 384 299H253L250 285Q246 271 244 268T231 265H227Q216 265 214 266T207 274Q207 278 223 345T244 416Q247 419 260 419H263Q280 419 280 408Q280 406 278 396L275 386Q275 385 406 385H537L540 399Q544 413 546 416T559 419H563Q574 419 576 418T583 410Q583 403 566 339Q549 271 544 267Q542 265 538 265H530H527Q510 265 510 276'],
  
      // GREEK CAPITAL LETTER LAMDA
      0x39B: [716,0,694,35,670,'135 2Q114 2 90 2T60 1Q35 1 35 11Q35 28 42 40Q45 46 55 46Q119 46 151 94Q153 97 325 402T498 709Q505 716 526 716Q543 716 549 710Q550 709 560 548T580 224T591 57Q594 52 595 52Q603 47 638 46H663Q670 39 670 35Q669 12 657 0H644Q613 2 530 2Q497 2 469 2T424 2T405 1Q388 1 388 10Q388 15 391 24Q392 27 393 32T395 38T397 41T401 44T406 45T415 46Q473 46 487 64L472 306Q468 365 465 426T459 518L457 550Q456 550 328 322T198 88Q196 80 196 77Q196 49 243 46Q261 46 261 35Q261 34 259 22Q256 7 254 4T240 0Q237 0 211 1T135 2'],
  
      // GREEK CAPITAL LETTER XI
      0x39E: [678,0,742,53,777,'222 668Q222 670 229 677H654Q677 677 705 677T740 678Q764 678 770 676T777 667Q777 662 764 594Q761 579 757 559T751 528L749 519Q747 512 729 512H717Q710 519 710 525Q712 532 715 559T719 591Q718 595 711 595Q682 598 486 598Q252 598 246 592Q239 587 228 552L216 517Q214 512 197 512H185Q178 517 178 522Q178 524 198 591T222 668ZM227 262Q218 262 215 262T209 266L207 270L227 356Q247 435 250 439Q253 443 260 443H267H280Q287 438 287 433Q287 430 285 420T280 402L278 393Q278 392 431 392H585L590 415Q595 436 598 439T612 443H628Q635 438 635 433Q635 431 615 351T594 268Q592 262 575 262H572Q556 262 556 272Q556 280 560 293L565 313H258L252 292Q248 271 245 267T230 262H227ZM60 0Q53 4 53 11Q53 14 68 89T84 169Q88 176 98 176H104H116Q123 169 123 163Q122 160 117 127T112 88Q112 80 243 80H351H454Q554 80 574 81T597 88V89Q603 100 610 121T622 157T630 174Q633 176 646 176H658Q665 171 665 166Q665 164 643 89T618 7Q616 2 607 1T548 0H335H60'],
  
      // GREEK CAPITAL LETTER PI
      0x3A0: [681,0,831,31,887,'48 1Q31 1 31 10Q31 12 34 24Q39 43 44 45Q48 46 59 46H65Q92 46 125 49Q139 52 144 61Q146 66 215 342T285 622Q285 629 281 629Q273 632 228 634H197Q191 640 191 642T193 661Q197 674 203 680H541Q621 680 709 680T812 681Q841 681 855 681T877 679T886 676T887 670Q887 663 885 656Q880 637 875 635Q871 634 860 634H854Q827 634 794 631Q780 628 775 619Q773 614 704 338T634 58Q634 51 638 51Q646 48 692 46H723Q729 38 729 37T726 19Q722 6 716 0H701Q664 2 567 2Q533 2 504 2T458 2T437 1Q420 1 420 10Q420 15 423 24Q428 43 433 45Q437 46 448 46H454Q481 46 514 49Q528 52 533 61Q536 67 572 209T642 491T678 632Q678 634 533 634H388Q387 631 316 347T245 59Q245 55 246 54T253 50T270 48T303 46H334Q340 38 340 37T337 19Q333 6 327 0H312Q275 2 178 2Q144 2 115 2T69 2T48 1'],
  
      // GREEK CAPITAL LETTER SIGMA
      0x3A3: [683,0,780,58,806,'65 0Q58 4 58 11Q58 16 114 67Q173 119 222 164L377 304Q378 305 340 386T261 552T218 644Q217 648 219 660Q224 678 228 681Q231 683 515 683H799Q804 678 806 674Q806 667 793 559T778 448Q774 443 759 443Q747 443 743 445T739 456Q739 458 741 477T743 516Q743 552 734 574T710 609T663 627T596 635T502 637Q480 637 469 637H339Q344 627 411 486T478 341V339Q477 337 477 336L457 318Q437 300 398 265T322 196L168 57Q167 56 188 56T258 56H359Q426 56 463 58T537 69T596 97T639 146T680 225Q686 243 689 246T702 250H705Q726 250 726 239Q726 238 683 123T639 5Q637 1 610 1Q577 0 348 0H65'],
  
      // GREEK CAPITAL LETTER UPSILON
      0x3A5: [706,0,583,28,700,'45 535Q34 535 31 536T28 544Q28 554 39 578T70 631T126 683T206 705Q230 705 251 698T295 671T330 612T344 514Q344 477 342 473V472Q343 472 347 480T361 509T380 547Q471 704 596 704Q615 704 625 702Q659 692 679 663T700 595Q700 565 696 552T687 537T670 535Q656 535 653 536T649 543Q649 544 649 550T650 562Q650 589 629 605T575 621Q502 621 448 547T365 361Q290 70 290 60Q290 46 379 46H404Q410 40 410 39T408 19Q404 6 398 0H381Q340 2 225 2Q184 2 149 2T94 2T69 1Q61 1 58 1T53 4T51 10Q51 11 53 23Q54 25 55 30T56 36T58 40T60 43T62 44T67 46T73 46T82 46H89Q144 46 163 49T190 62L198 93Q206 124 217 169T241 262T262 350T274 404Q281 445 281 486V494Q281 621 185 621Q147 621 116 601T74 550Q71 539 66 537T45 535'],
  
      // GREEK CAPITAL LETTER PHI
      0x3A6: [683,0,667,24,642,'356 624Q356 637 267 637H243Q237 642 237 645T239 664Q243 677 249 683H264Q342 681 429 681Q565 681 571 683H583Q589 677 589 674T587 656Q582 641 578 637H540Q516 637 504 637T479 633T463 630T454 623T448 613T443 597T438 576Q436 566 434 556T430 539L428 533Q442 533 472 526T543 502T613 451T642 373Q642 301 567 241T386 158L336 150Q332 150 331 146Q310 66 310 60Q310 46 399 46H424Q430 40 430 39T428 19Q424 6 418 0H401Q360 2 247 2Q207 2 173 2T119 2T95 1Q87 1 84 1T79 4T77 10Q77 11 79 23Q80 25 81 30T82 36T84 40T86 43T88 44T93 46T99 46T108 46H115Q170 46 189 49T216 62Q220 74 228 107L239 150L223 152Q139 164 82 205T24 311Q24 396 125 462Q207 517 335 533L346 578Q356 619 356 624ZM130 291Q130 203 241 188H249Q249 190 287 342L325 495H324Q313 495 291 491T229 466T168 414Q130 357 130 291ZM536 393Q536 440 507 463T418 496L341 187L351 189Q443 201 487 255Q536 314 536 393'],
  
      // GREEK CAPITAL LETTER PSI
      0x3A8: [683,0,612,21,692,'216 151Q48 174 48 329Q48 361 56 403T65 458Q65 482 58 494T43 507T28 510T21 520Q21 528 23 534T29 544L32 546H72H94Q110 546 119 544T139 536T154 514T159 476V465Q159 445 149 399T138 314Q142 229 197 201Q223 187 226 190L233 218Q240 246 253 300T280 407Q333 619 333 625Q333 637 244 637H220Q214 642 214 645T216 664Q220 677 226 683H241Q321 681 405 681Q543 681 549 683H560Q566 677 566 674T564 656Q559 641 555 637H517Q448 636 436 628Q429 623 423 600T373 404L320 192Q370 201 419 248Q451 281 469 317T500 400T518 457Q529 486 542 505T569 532T594 543T621 546H644H669Q692 546 692 536Q691 509 676 509Q623 509 593 399Q587 377 579 355T552 301T509 244T446 195T359 159Q324 151 314 151Q311 151 310 150T298 106T287 60Q287 46 376 46H401Q407 40 407 39T405 19Q401 6 395 0H378Q337 2 224 2Q184 2 150 2T96 2T72 1Q64 1 61 1T56 4T54 10Q54 11 56 23Q57 25 58 30T59 36T61 40T63 43T65 44T70 46T76 46T85 46H92Q147 46 166 49T193 62L204 106Q216 149 216 151'],
  
      // GREEK CAPITAL LETTER OMEGA
      0x3A9: [704,0,772,80,786,'125 84Q127 78 194 76H243V78Q243 122 208 215T165 350Q164 359 162 389Q162 522 272 610Q328 656 396 680T525 704Q628 704 698 661Q734 637 755 601T781 544T786 504Q786 439 747 374T635 226T537 109Q518 81 518 77Q537 76 557 76Q608 76 620 78T640 92Q646 100 656 119T673 155T683 172Q690 173 698 173Q718 173 718 162Q718 161 681 82T642 2Q639 0 550 0H461Q455 5 455 9T458 28Q472 78 510 149T584 276T648 402T677 525Q677 594 636 631T530 668Q476 668 423 641T335 568Q284 499 271 400Q270 388 270 348Q270 298 277 228T285 115Q285 82 280 49T271 6Q269 1 258 1T175 0H87Q83 3 80 7V18Q80 22 82 98Q84 156 85 163T91 172Q94 173 104 173T119 172Q124 169 124 126Q125 104 125 84'],
  
      // GREEK SMALL LETTER ALPHA
      0x3B1: [442,11,640,34,603,'34 156Q34 270 120 356T309 442Q379 442 421 402T478 304Q484 275 485 237V208Q534 282 560 374Q564 388 566 390T582 393Q603 393 603 385Q603 376 594 346T558 261T497 161L486 147L487 123Q489 67 495 47T514 26Q528 28 540 37T557 60Q559 67 562 68T577 70Q597 70 597 62Q597 56 591 43Q579 19 556 5T512 -10H505Q438 -10 414 62L411 69L400 61Q390 53 370 41T325 18T267 -2T203 -11Q124 -11 79 39T34 156ZM208 26Q257 26 306 47T379 90L403 112Q401 255 396 290Q382 405 304 405Q235 405 183 332Q156 292 139 224T121 120Q121 71 146 49T208 26'],
  
      // GREEK SMALL LETTER BETA
      0x3B2: [705,194,566,23,573,'29 -194Q23 -188 23 -186Q23 -183 102 134T186 465Q208 533 243 584T309 658Q365 705 429 705H431Q493 705 533 667T573 570Q573 465 469 396L482 383Q533 332 533 252Q533 139 448 65T257 -10Q227 -10 203 -2T165 17T143 40T131 59T126 65L62 -188Q60 -194 42 -194H29ZM353 431Q392 431 427 419L432 422Q436 426 439 429T449 439T461 453T472 471T484 495T493 524T501 560Q503 569 503 593Q503 611 502 616Q487 667 426 667Q384 667 347 643T286 582T247 514T224 455Q219 439 186 308T152 168Q151 163 151 147Q151 99 173 68Q204 26 260 26Q302 26 349 51T425 137Q441 171 449 214T457 279Q457 337 422 372Q380 358 347 358H337Q258 358 258 389Q258 396 261 403Q275 431 353 431'],
  
      // GREEK SMALL LETTER GAMMA
      0x3B3: [441,216,518,11,543,'31 249Q11 249 11 258Q11 275 26 304T66 365T129 418T206 441Q233 441 239 440Q287 429 318 386T371 255Q385 195 385 170Q385 166 386 166L398 193Q418 244 443 300T486 391T508 430Q510 431 524 431H537Q543 425 543 422Q543 418 522 378T463 251T391 71Q385 55 378 6T357 -100Q341 -165 330 -190T303 -216Q286 -216 286 -188Q286 -138 340 32L346 51L347 69Q348 79 348 100Q348 257 291 317Q251 355 196 355Q148 355 108 329T51 260Q49 251 47 251Q45 249 31 249'],
  
      // GREEK SMALL LETTER DELTA
      0x3B4: [717,10,444,36,451,'195 609Q195 656 227 686T302 717Q319 716 351 709T407 697T433 690Q451 682 451 662Q451 644 438 628T403 612Q382 612 348 641T288 671T249 657T235 628Q235 584 334 463Q401 379 401 292Q401 169 340 80T205 -10H198Q127 -10 83 36T36 153Q36 286 151 382Q191 413 252 434Q252 435 245 449T230 481T214 521T201 566T195 609ZM112 130Q112 83 136 55T204 27Q233 27 256 51T291 111T309 178T316 232Q316 267 309 298T295 344T269 400L259 396Q215 381 183 342T137 256T118 179T112 130'],
  
      // GREEK SMALL LETTER EPSILON
      0x3B5: [452,23,466,27,428,'190 -22Q124 -22 76 11T27 107Q27 174 97 232L107 239L99 248Q76 273 76 304Q76 364 144 408T290 452H302Q360 452 405 421Q428 405 428 392Q428 381 417 369T391 356Q382 356 371 365T338 383T283 392Q217 392 167 368T116 308Q116 289 133 272Q142 263 145 262T157 264Q188 278 238 278H243Q308 278 308 247Q308 206 223 206Q177 206 142 219L132 212Q68 169 68 112Q68 39 201 39Q253 39 286 49T328 72T345 94T362 105Q376 103 376 88Q376 79 365 62T334 26T275 -8T190 -22'],
  
      // GREEK SMALL LETTER ZETA
      0x3B6: [704,204,438,44,471,'296 643Q298 704 324 704Q342 704 342 687Q342 682 339 664T336 633Q336 623 337 618T338 611Q339 612 341 612Q343 614 354 616T374 618L384 619H394Q471 619 471 586Q467 548 386 546H372Q338 546 320 564L311 558Q235 506 175 398T114 190Q114 171 116 155T125 127T137 104T153 86T171 72T192 61T213 53T235 46T256 39L322 16Q389 -10 389 -80Q389 -119 364 -154T300 -202Q292 -204 274 -204Q247 -204 225 -196Q210 -192 193 -182T172 -167Q167 -159 173 -148Q180 -139 191 -139Q195 -139 221 -153T283 -168Q298 -166 310 -152T322 -117Q322 -91 302 -75T250 -51T183 -29T116 4T65 62T44 160Q44 287 121 410T293 590L302 595Q296 613 296 643'],
  
      // GREEK SMALL LETTER ETA
      0x3B7: [443,216,497,21,503,'21 287Q22 290 23 295T28 317T38 348T53 381T73 411T99 433T132 442Q156 442 175 435T205 417T221 395T229 376L231 369Q231 367 232 367L243 378Q304 442 382 442Q436 442 469 415T503 336V326Q503 302 439 53Q381 -182 377 -189Q364 -216 332 -216Q319 -216 310 -208T299 -186Q299 -177 358 57L420 307Q423 322 423 345Q423 404 379 404H374Q288 404 229 303L222 291L189 157Q156 26 151 16Q138 -11 108 -11Q95 -11 87 -5T76 7T74 17Q74 30 114 189T154 366Q154 405 128 405Q107 405 92 377T68 316T57 280Q55 278 41 278H27Q21 284 21 287'],
  
      // GREEK SMALL LETTER THETA
      0x3B8: [705,10,469,35,462,'35 200Q35 302 74 415T180 610T319 704Q320 704 327 704T339 705Q393 701 423 656Q462 596 462 495Q462 380 417 261T302 66T168 -10H161Q125 -10 99 10T60 63T41 130T35 200ZM383 566Q383 668 330 668Q294 668 260 623T204 521T170 421T157 371Q206 370 254 370L351 371Q352 372 359 404T375 484T383 566ZM113 132Q113 26 166 26Q181 26 198 36T239 74T287 161T335 307L340 324H145Q145 321 136 286T120 208T113 132'],
  
      // GREEK SMALL LETTER IOTA
      0x3B9: [442,10,354,48,333,'139 -10Q111 -10 92 0T64 25T52 52T48 74Q48 89 55 109T85 199T135 375L137 384Q139 394 140 397T145 409T151 422T160 431T173 439T190 442Q202 442 213 435T225 410Q225 404 214 358T181 238T137 107Q126 74 126 54Q126 43 126 39T130 31T142 27H147Q206 27 255 78Q272 98 281 114T290 138T295 149T313 153Q321 153 324 153T329 152T332 149T332 143Q332 106 276 48T145 -10H139'],
  
      // GREEK SMALL LETTER KAPPA
      0x3BA: [442,11,576,48,554,'83 -11Q70 -11 62 -4T51 8T49 17Q49 30 96 217T147 414Q160 442 193 442Q205 441 213 435T223 422T225 412Q225 401 208 337L192 270Q193 269 208 277T235 292Q252 304 306 349T396 412T467 431Q489 431 500 420T512 391Q512 366 494 347T449 327Q430 327 418 338T405 368Q405 370 407 380L397 375Q368 360 315 315L253 266L240 257H245Q262 257 300 251T366 230Q422 203 422 150Q422 140 417 114T411 67Q411 26 437 26Q484 26 513 137Q516 149 519 151T535 153Q554 153 554 144Q554 121 527 64T457 -7Q447 -10 431 -10Q386 -10 360 17T333 90Q333 108 336 122T339 146Q339 170 320 186T271 209T222 218T185 221H180L155 122Q129 22 126 16Q113 -11 83 -11'],
  
      // GREEK SMALL LETTER LAMDA
      0x3BB: [694,12,583,47,557,'166 673Q166 685 183 694H202Q292 691 316 644Q322 629 373 486T474 207T524 67Q531 47 537 34T546 15T551 6T555 2T556 -2T550 -11H482Q457 3 450 18T399 152L354 277L340 262Q327 246 293 207T236 141Q211 112 174 69Q123 9 111 -1T83 -12Q47 -12 47 20Q47 37 61 52T199 187Q229 216 266 252T321 306L338 322Q338 323 288 462T234 612Q214 657 183 657Q166 657 166 673'],
  
      // GREEK SMALL LETTER MU
      0x3BC: [442,216,603,23,580,'58 -216Q44 -216 34 -208T23 -186Q23 -176 96 116T173 414Q186 442 219 442Q231 441 239 435T249 423T251 413Q251 401 220 279T187 142Q185 131 185 107V99Q185 26 252 26Q261 26 270 27T287 31T302 38T315 45T327 55T338 65T348 77T356 88T365 100L372 110L408 253Q444 395 448 404Q461 431 491 431Q504 431 512 424T523 412T525 402L449 84Q448 79 448 68Q448 43 455 35T476 26Q485 27 496 35Q517 55 537 131Q543 151 547 152Q549 153 557 153H561Q580 153 580 144Q580 138 575 117T555 63T523 13Q510 0 491 -8Q483 -10 467 -10Q446 -10 429 -4T402 11T385 29T376 44T374 51L368 45Q362 39 350 30T324 12T288 -4T246 -11Q199 -11 153 12L129 -85Q108 -167 104 -180T92 -202Q76 -216 58 -216'],
  
      // GREEK SMALL LETTER NU
      0x3BD: [442,2,494,45,530,'74 431Q75 431 146 436T219 442Q231 442 231 434Q231 428 185 241L137 51H140L150 55Q161 59 177 67T214 86T261 119T312 165Q410 264 445 394Q458 442 496 442Q509 442 519 434T530 411Q530 390 516 352T469 262T388 162T267 70T106 5Q81 -2 71 -2Q66 -2 59 -1T51 1Q45 5 45 11Q45 13 88 188L132 364Q133 377 125 380T86 385H65Q59 391 59 393T61 412Q65 431 74 431'],
  
      // GREEK SMALL LETTER XI
      0x3BE: [704,205,438,21,443,'268 632Q268 704 296 704Q314 704 314 687Q314 682 311 664T308 635T309 620V616H315Q342 619 360 619Q443 619 443 586Q439 548 358 546H344Q326 546 317 549T290 566Q257 550 226 505T195 405Q195 381 201 364T211 342T218 337Q266 347 298 347Q375 347 375 314Q374 297 359 288T327 277T280 275Q234 275 208 283L195 286Q149 260 119 214T88 130Q88 116 90 108Q101 79 129 63T229 20Q238 17 243 15Q337 -21 354 -33Q383 -53 383 -94Q383 -137 351 -171T273 -205Q240 -205 202 -190T158 -167Q156 -163 156 -159Q156 -151 161 -146T176 -140Q182 -140 189 -143Q232 -168 274 -168Q286 -168 292 -165Q313 -151 313 -129Q313 -112 301 -104T232 -75Q214 -68 204 -64Q198 -62 171 -52T136 -38T107 -24T78 -8T56 12T36 37T26 66T21 103Q21 149 55 206T145 301L154 307L148 313Q141 319 136 323T124 338T111 358T103 382T99 413Q99 471 143 524T259 602L271 607Q268 618 268 632'],
  
      // GREEK SMALL LETTER OMICRON
      0x3BF: [441,11,485,34,476,'201 -11Q126 -11 80 38T34 156Q34 221 64 279T146 380Q222 441 301 441Q333 441 341 440Q354 437 367 433T402 417T438 387T464 338T476 268Q476 161 390 75T201 -11ZM121 120Q121 70 147 48T206 26Q250 26 289 58T351 142Q360 163 374 216T388 308Q388 352 370 375Q346 405 306 405Q243 405 195 347Q158 303 140 230T121 120'],
  
      // GREEK SMALL LETTER PI
      0x3C0: [431,11,570,19,573,'132 -11Q98 -11 98 22V33L111 61Q186 219 220 334L228 358H196Q158 358 142 355T103 336Q92 329 81 318T62 297T53 285Q51 284 38 284Q19 284 19 294Q19 300 38 329T93 391T164 429Q171 431 389 431Q549 431 553 430Q573 423 573 402Q573 371 541 360Q535 358 472 358H408L405 341Q393 269 393 222Q393 170 402 129T421 65T431 37Q431 20 417 5T381 -10Q370 -10 363 -7T347 17T331 77Q330 86 330 121Q330 170 339 226T357 318T367 358H269L268 354Q268 351 249 275T206 114T175 17Q164 -11 132 -11'],
  
      // GREEK SMALL LETTER RHO
      0x3C1: [442,216,517,23,510,'58 -216Q25 -216 23 -186Q23 -176 73 26T127 234Q143 289 182 341Q252 427 341 441Q343 441 349 441T359 442Q432 442 471 394T510 276Q510 219 486 165T425 74T345 13T266 -10H255H248Q197 -10 165 35L160 41L133 -71Q108 -168 104 -181T92 -202Q76 -216 58 -216ZM424 322Q424 359 407 382T357 405Q322 405 287 376T231 300Q217 269 193 170L176 102Q193 26 260 26Q298 26 334 62Q367 92 389 158T418 266T424 322'],
  
      // GREEK SMALL LETTER FINAL SIGMA
      0x3C2: [442,107,363,30,405,'31 207Q31 306 115 374T302 442Q341 442 373 430T405 400Q405 392 399 383T379 374Q373 375 348 390T296 405Q222 405 160 357T98 249Q98 232 103 218T112 195T132 175T154 159T186 141T219 122Q234 114 255 102T286 85T299 78L302 74Q306 71 308 69T315 61T322 51T328 40T332 25T334 8Q334 -31 305 -69T224 -107Q194 -107 163 -92Q156 -88 156 -80Q156 -73 162 -67T178 -61Q186 -61 190 -63Q209 -71 224 -71Q244 -71 253 -59T263 -30Q263 -25 263 -21T260 -12T255 -4T248 3T239 9T227 17T213 25T195 34T174 46Q170 48 150 58T122 74T97 90T70 112T51 137T36 169T31 207'],
  
      // GREEK SMALL LETTER SIGMA
      0x3C3: [431,11,571,31,572,'184 -11Q116 -11 74 34T31 147Q31 247 104 333T274 430Q275 431 414 431H552Q553 430 555 429T559 427T562 425T565 422T567 420T569 416T570 412T571 407T572 401Q572 357 507 357Q500 357 490 357T476 358H416L421 348Q439 310 439 263Q439 153 359 71T184 -11ZM361 278Q361 358 276 358Q152 358 115 184Q114 180 114 178Q106 141 106 117Q106 67 131 47T188 26Q242 26 287 73Q316 103 334 153T356 233T361 278'],
  
      // GREEK SMALL LETTER TAU
      0x3C4: [431,13,437,18,517,'39 284Q18 284 18 294Q18 301 45 338T99 398Q134 425 164 429Q170 431 332 431Q492 431 497 429Q517 424 517 402Q517 388 508 376T485 360Q479 358 389 358T299 356Q298 355 283 274T251 109T233 20Q228 5 215 -4T186 -13Q153 -13 153 20V30L203 192Q214 228 227 272T248 336L254 357Q254 358 208 358Q206 358 197 358T183 359Q105 359 61 295Q56 287 53 286T39 284'],
  
      // GREEK SMALL LETTER UPSILON
      0x3C5: [443,10,540,21,523,'413 384Q413 406 432 424T473 443Q492 443 507 425T523 367Q523 334 508 270T468 153Q424 63 373 27T282 -10H268Q220 -10 186 2T135 36T111 78T104 121Q104 170 138 262T173 379Q173 380 173 381Q173 390 173 393T169 400T158 404H154Q131 404 112 385T82 344T65 302T57 280Q55 278 41 278H27Q21 284 21 287Q21 299 34 333T82 404T161 441Q200 441 225 419T250 355Q248 336 247 334Q247 331 232 291T201 199T185 118Q185 68 211 47T275 26Q317 26 355 57T416 132T452 216T465 277Q465 301 457 318T439 343T421 361T413 384'],
  
      // GREEK SMALL LETTER PHI
      0x3C6: [442,218,654,50,618,'92 210Q92 176 106 149T142 108T185 85T220 72L235 70L237 71L250 112Q268 170 283 211T322 299T370 375T429 423T502 442Q547 442 582 410T618 302Q618 224 575 152T457 35T299 -10Q273 -10 273 -12L266 -48Q260 -83 252 -125T241 -179Q236 -203 215 -212Q204 -218 190 -218Q159 -215 159 -185Q159 -175 214 -2L209 0Q204 2 195 5T173 14T147 28T120 46T94 71T71 103T56 142T50 190Q50 238 76 311T149 431H162Q183 431 183 423Q183 417 175 409Q134 361 114 300T92 210ZM574 278Q574 320 550 344T486 369Q437 369 394 329T323 218Q309 184 295 109L286 64Q304 62 306 62Q423 62 498 131T574 278'],
  
      // GREEK SMALL LETTER CHI
      0x3C7: [443,204,626,24,600,'576 -125Q576 -147 547 -175T487 -204H476Q394 -204 363 -157Q334 -114 293 26L284 59Q283 58 248 19T170 -66T92 -151T53 -191Q49 -194 43 -194Q36 -194 31 -189T25 -177T38 -154T151 -30L272 102L265 131Q189 405 135 405Q104 405 87 358Q86 351 68 351Q48 351 48 361Q48 369 56 386T89 423T148 442Q224 442 258 400Q276 375 297 320T330 222L341 180Q344 180 455 303T573 429Q579 431 582 431Q600 431 600 414Q600 407 587 392T477 270Q356 138 353 134L362 102Q392 -10 428 -89T490 -168Q504 -168 517 -156T536 -126Q539 -116 543 -115T557 -114T571 -115Q576 -118 576 -125'],
  
      // GREEK SMALL LETTER PSI
      0x3C8: [694,205,651,21,634,'161 441Q202 441 226 417T250 358Q250 338 218 252T187 127Q190 85 214 61Q235 43 257 37Q275 29 288 29H289L371 360Q455 691 456 692Q459 694 472 694Q492 694 492 687Q492 678 411 356Q329 28 329 27T335 26Q421 26 498 114T576 278Q576 302 568 319T550 343T532 361T524 384Q524 405 541 424T583 443Q602 443 618 425T634 366Q634 337 623 288T605 220Q573 125 492 57T329 -11H319L296 -104Q272 -198 272 -199Q270 -205 252 -205H239Q233 -199 233 -197Q233 -192 256 -102T279 -9Q272 -8 265 -8Q106 14 106 139Q106 174 139 264T173 379Q173 380 173 381Q173 390 173 393T169 400T158 404H154Q131 404 112 385T82 344T65 302T57 280Q55 278 41 278H27Q21 284 21 287Q21 299 34 333T82 404T161 441'],
  
      // GREEK SMALL LETTER OMEGA
      0x3C9: [443,12,622,15,604,'495 384Q495 406 514 424T555 443Q574 443 589 425T604 364Q604 334 592 278T555 155T483 38T377 -11Q297 -11 267 66Q266 68 260 61Q201 -11 125 -11Q15 -11 15 139Q15 230 56 325T123 434Q135 441 147 436Q160 429 160 418Q160 406 140 379T94 306T62 208Q61 202 61 187Q61 124 85 100T143 76Q201 76 245 129L253 137V156Q258 297 317 297Q348 297 348 261Q348 243 338 213T318 158L308 135Q309 133 310 129T318 115T334 97T358 83T393 76Q456 76 501 148T546 274Q546 305 533 325T508 357T495 384'],
  
      // GREEK THETA SYMBOL
      0x3D1: [705,11,591,21,563,'537 500Q537 474 533 439T524 383L521 362Q558 355 561 351Q563 349 563 345Q563 321 552 318Q542 318 521 323L510 326Q496 261 459 187T362 51T241 -11Q100 -11 100 105Q100 139 127 242T154 366Q154 405 128 405Q107 405 92 377T68 316T57 280Q55 278 41 278H27Q21 284 21 287Q21 291 27 313T47 368T79 418Q103 442 134 442Q169 442 201 419T233 344Q232 330 206 228T180 98Q180 26 247 26Q292 26 332 90T404 260L427 349Q422 349 398 359T339 392T289 440Q265 476 265 520Q265 590 312 647T417 705Q463 705 491 670T528 592T537 500ZM464 564Q464 668 413 668Q373 668 339 622T304 522Q304 494 317 470T349 431T388 406T421 391T435 387H436L443 415Q450 443 457 485T464 564'],
  
      // GREEK PHI SYMBOL
      0x3D5: [694,205,596,42,579,'409 688Q413 694 421 694H429H442Q448 688 448 686Q448 679 418 563Q411 535 404 504T392 458L388 442Q388 441 397 441T429 435T477 418Q521 397 550 357T579 260T548 151T471 65T374 11T279 -10H275L251 -105Q245 -128 238 -160Q230 -192 227 -198T215 -205H209Q189 -205 189 -198Q189 -193 211 -103L234 -11Q234 -10 226 -10Q221 -10 206 -8T161 6T107 36T62 89T43 171Q43 231 76 284T157 370T254 422T342 441Q347 441 348 445L378 567Q409 686 409 688ZM122 150Q122 116 134 91T167 53T203 35T237 27H244L337 404Q333 404 326 403T297 395T255 379T211 350T170 304Q152 276 137 237Q122 191 122 150ZM500 282Q500 320 484 347T444 385T405 400T381 404H378L332 217L284 29Q284 27 285 27Q293 27 317 33T357 47Q400 66 431 100T475 170T494 234T500 282'],
  
      // GREEK PI SYMBOL
      0x3D6: [431,10,828,19,823,'206 -10Q158 -10 136 24T114 110Q114 233 199 349L205 358H184Q144 358 121 347Q108 340 95 330T75 312T61 295T53 285Q51 284 38 284Q19 284 19 294Q19 300 38 329T93 391T164 429Q171 431 532 431Q799 431 803 430Q823 423 823 402Q823 377 801 364Q790 358 766 358Q748 358 748 357Q748 355 749 348T752 327T754 297Q754 258 738 207T693 107T618 24T520 -10Q488 -10 466 2T432 36T416 77T411 120Q411 128 410 128T404 122Q373 71 323 31T206 -10ZM714 296Q714 316 707 358H251Q250 357 244 348T230 328T212 301T193 267T176 229T164 187T159 144Q159 62 222 62Q290 62 349 127T432 285Q433 286 434 288T435 291T437 293T440 294T444 294T452 294H466Q472 288 472 286Q472 285 464 244T456 170Q456 62 534 62Q604 62 659 139T714 296'],
  
      // GREEK RHO SYMBOL
      0x3F1: [442,194,517,67,510,'205 -174Q136 -174 102 -153T67 -76Q67 -25 91 85T127 234Q143 289 182 341Q252 427 341 441Q343 441 349 441T359 442Q432 442 471 394T510 276Q510 169 431 80T253 -10Q226 -10 204 -2T169 19T146 44T132 64L128 73Q128 72 124 53T116 5T112 -44Q112 -68 117 -78T150 -95T236 -102Q327 -102 356 -111T386 -154Q386 -166 384 -178Q381 -190 378 -192T361 -194H348Q342 -188 342 -179Q342 -169 315 -169Q294 -169 264 -171T205 -174ZM424 322Q424 359 407 382T357 405Q322 405 287 376T231 300Q221 276 204 217Q188 152 188 116Q188 68 210 47T259 26Q297 26 334 62Q367 92 389 158T418 266T424 322'],
  
      // GREEK LUNATE EPSILON SYMBOL
      0x3F5: [431,11,406,40,382,'227 -11Q149 -11 95 41T40 174Q40 262 87 322Q121 367 173 396T287 430Q289 431 329 431H367Q382 426 382 411Q382 385 341 385H325H312Q191 385 154 277L150 265H327Q340 256 340 246Q340 228 320 219H138V217Q128 187 128 143Q128 77 160 52T231 26Q258 26 284 36T326 57T343 68Q350 68 354 58T358 39Q358 36 357 35Q354 31 337 21T289 0T227 -11'],
  
      // INCREMENT
      0x2206: [716,0,833,48,788,'']
  };

  SVG.FONTDATA.FONTS['MathJax_Main'][0x22EE][0]  += 400;  // adjust height for \vdots
  SVG.FONTDATA.FONTS['MathJax_Main'][0x22F1][0]  += 700;  // adjust height for \ddots

  //
  //  Add some spacing characters (more will come later)
  //
  MathJax.Hub.Insert(SVG.FONTDATA.FONTS['MathJax_Main'],{
    0x2000: [0,0,500,0,0,{space:1}],     // en quad
    0x2001: [0,0,1000,0,0,{space:1}],    // em quad
    0x2002: [0,0,500,0,0,{space:1}],     // en space
    0x2003: [0,0,1000,0,0,{space:1}],    // em space
    0x2004: [0,0,333,0,0,{space:1}],     // 3-per-em space
    0x2005: [0,0,250,0,0,{space:1}],     // 4-per-em space
    0x2006: [0,0,167,0,0,{space:1}],     // 6-per-em space
    0x2009: [0,0,167,0,0,{space:1}],     // thin space
    0x200A: [0,0,83,0,0,{space:1}],      // hair space
    0x200B: [0,0,0,0,0,{space:1}],       // zero-width space
    0xEEE0: [0,0,-575,0,0,{space:1}],
    0xEEE1: [0,0,-300,0,0,{space:1}],
    0xEEE8: [0,0,25,0,0,{space:1}]
  });

  HUB.Register.StartupHook("SVG Jax Require",function () {
    HUB.Register.LoadHook(SVG.fontDir+"/Size4/Regular/Main.js",function () {
      SVG.FONTDATA.FONTS['MathJax_Size4'][0xE154][0] += 200;  // adjust height for brace extender
      SVG.FONTDATA.FONTS['MathJax_Size4'][0xE154][1] += 200;  // adjust depth for brace extender
    });
    
    SVG.FONTDATA.FONTS['MathJax_Main'][0x2245][2] -= 222; // fix incorrect right bearing in font
    HUB.Register.LoadHook(SVG.fontDir+"/Main/Bold/MathOperators.js",function () {
      SVG.FONTDATA.FONTS['MathJax_Main-bold'][0x2245][2] -= 106; // fix incorrect right bearing in font
    });

    HUB.Register.LoadHook(SVG.fontDir+"/Typewriter/Regular/BasicLatin.js",function () {
      SVG.FONTDATA.FONTS['MathJax_Typewriter'][0x20][2] += 275; // fix incorrect width
    });

    AJAX.loadComplete(SVG.fontDir + "/fontdata.js");
  });
  
})(MathJax.OutputJax.SVG,MathJax.ElementJax.mml,MathJax.Ajax,MathJax.Hub);

