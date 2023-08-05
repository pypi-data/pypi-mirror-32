/*************************************************************
 *
 *  MathJax/jax/output/HTML-CSS/fonts/STIX/General/Italic/LatinExtendedAdditional.js
 *
 *  Copyright (c) 2009-2015 The MathJax Consortium
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
 *
 */

MathJax.Hub.Insert(
  MathJax.OutputJax['HTML-CSS'].FONTDATA.FONTS['STIXGeneral-italic'],
  {
    0x1E80: [880,18,833,71,906],       // LATIN CAPITAL LETTER W WITH GRAVE
    0x1E81: [664,18,667,15,648],       // LATIN SMALL LETTER W WITH GRAVE
    0x1E82: [876,18,833,71,906],       // LATIN CAPITAL LETTER W WITH ACUTE
    0x1E83: [664,18,667,15,648],       // LATIN SMALL LETTER W WITH ACUTE
    0x1E84: [818,18,833,71,906],       // LATIN CAPITAL LETTER W WITH DIAERESIS
    0x1E85: [606,18,667,15,648],       // LATIN SMALL LETTER W WITH DIAERESIS
    0x1EF2: [880,0,556,78,633],        // LATIN CAPITAL LETTER Y WITH GRAVE
    0x1EF3: [664,206,444,-24,426]      // LATIN SMALL LETTER Y WITH GRAVE
  }
);

MathJax.Ajax.loadComplete(MathJax.OutputJax["HTML-CSS"].fontDir + "/General/Italic/LatinExtendedAdditional.js");
