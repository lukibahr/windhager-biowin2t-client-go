# windhager-biowint2-client-go

!!! REFACTOR PROJECT LIKE THIS: https://github.com/google/go-github

Windhager BioWinTouch 2 goes cloud native

Hint: python implementation is located in separate branch

## Accessing the UI

The UI is served by the Windhager pellet appliance built-in MES inifinity controller. Behind the UI, a REST API serves all values that are avaialable through the ui or the touch panel.
The Username, named `Service` has access to the UI via basic authentication and the password can be fetched using the Windhager Connect platform [https://connect.windhager.com](https://connect.windhager.com)

More information about the mes infinity controller can be found here: [https://www.windhager.com/int_en/products/control/mes-infinity/](https://www.windhager.com/int_en/products/control/mes-infinity/)

## Curling the API

The api endpoint is available under `http://192.168.2.121/api/1.0/lookup/<OID>`. To add the digest authentication, use the `--digest` parameter like the following: `curl http://192.168.2.140/api/1.0/lookup/1/60/0/100/9 --digest -u "$USERNAME:$PASSWORD"`

Sample response looks like the following:

```bash
{
    "OID": "/1/60/0/12/101/0",
    "groupNr": 12,
    "maxValue": "14.0",
    "memberNr": 101,
    "minValue": "6.0",
    "name": "12-101",
    "step": "0.1",
    "stepId": 0,
    "subtypeId": -1,
    "timestamp": "2021-09-07 11:23:03",
    "typeId": 15,
    "unit": "kg",
    "unitId": 45,
    "value": "6.0",
    "writeProt": false
}
```
### BioWIn2

- Laufzeit bis Hauptreinigung: http://192.168.2.121/api/1.0/lookup/1/60/0/98/9 (done)
- Laufzeit bis Reinigung: http://192.168.2.121/api/1.0/lookup/1/60/0/98/8
- Betriebsstunden: http://192.168.2.121/api/1.0/lookup/1/60/0/98/4
- Anzahl der Brennerstarts: http://192.168.2.121/api/1.0/lookup/1/60/0/98/3
- Temperatur Abgas: http://192.168.2.121/api/1.0/lookup/1/60/0/98/1
- Aktuelle Kesselleistung: http://192.168.2.121/api/1.0/lookup/1/60/0/98/0

- Kesseltemperatur Istwert: http://192.168.2.121/api/1.0/lookup/1/60/0/100/1
- Brennkammertemperatur: http://192.168.2.121/api/1.0/lookup/1/60/0/100/2
- Aktuelle Betriebsphase: http://192.168.2.121/api/1.0/lookup/1/60/0/100/3
- Brennstoffmenge Förderschnecke Istwert: http://192.168.2.121/api/1.0/lookup/1/60/0/100/9

#### Betriebsphasen

- (3) Standby: In dieser Betriebsphase wird von der vorhandenen Regelung keine Wärme-anforderung übertragen. Der Brenner ist ausgeschaltet und der Kesseltempe-ratur-Sollwert ist 0 °C
- () Vorspülen: Das Saugzuggebläse läuft, der Brennraum des FireWIN wird mit Frischluftdurchspült. Diese Phase kann einige Minuten dauern bevor der Brenner inBetrieb geht
- (6) Zündphase: Das Saugzuggebläse läuft, Pellets werden in den Brennertopf gefördert undentzündet. Wird eine Flammenbildung erkannt, wird in die Flammenstabili-sierung übergegangen
- () Flammenstabilisierung: Nach dem Zündvorgang wird eine gleichmäßige Verbrennung aufgebaut undanschließend in den Modulationsbetrieb geschaltet
- () Modulationsbetrieb: Der Brenner ist im Modulationsbetrieb. Die Leistung wird stufenlos zwischen30 % und 100 % geregelt
- Ausbrand: Die Verbrennung wird eingestellt. Der Pelletstransport in den Brennertopf wirdgestoppt, das Saugzuggebläse läuft nach, bis die restlichen Pellets verbranntsind und der Brennertopf abgekühlt ist
- Brenner AUS: Die Wärmeanforderung von der Regelung ist vorhandenen, aber die Kessel-temperatur (Istwert) ist höher als der Kesseltemperatur-Sollwert. Daher ist die Verbrennung eingestellt und der Brenner ausgeschaltet


### Huber/WW

- Warmwasser temp ist-wert (oben) soll-wert unten: http://192.168.2.121/api/1.0/lookup/1/15/0/114
- Aussentemperatur: http://192.168.2.121/api/1.0/lookup/1/15/0/115
- Vorlauftemperatur (ist-Soll): http://192.168.2.121/api/1.0/lookup/1/15/0/116
- Kesseltemperatur (ist-Soll): http://192.168.2.121/api/1.0/lookup/1/16/0/117

### Urban

- Aussentemperatur: http://192.168.2.121/api/1.0/lookup/1/15/1/115
- Vorlauftemperatur (ist-Soll): http://192.168.2.121/api/1.0/lookup/1/15/1/116
- Kesseltemperatur (ist-Soll): http://192.168.2.121/api/1.0/lookup/1/16/1/117


### Bahr

- Aussentemperatur: http://192.168.2.121/api/1.0/lookup/1/16/1/115
- Vorlauftemperatur (ist-Soll): http://192.168.2.121/api/1.0/lookup/1/16/1/116
- Kesseltemperatur (ist-Soll): http://192.168.2.121/api/1.0/lookup/1/16/1/117


### Basler

- Aussentemperatur: http://192.168.2.121/api/1.0/lookup/1/16/0/115
- Vorlauftemperatur (ist-Soll): http://192.168.2.121/api/1.0/lookup/1/16/0/116
- Kesseltemperatur (ist-Soll): http://192.168.2.121/api/1.0/lookup/1/16/0/117



        'Alarmcode',                  '1/60/0/155/4',
        'Betriebsart',                '1/60/0/155/1',
        'Betriebsphase',              '1/60/0/155/2',
        'Betriebsstunden',            '1/60/0/156/2',
        'Laufzeit_bis_Reinigung',     '1/60/0/156/3',
        'WarmWasser_IST',             '1/15/0/114/0',
        'WarmWasser_SOLL',            '1/15/0/114/1',
        'Aussentemperatur',           '1/15/0/115/0',
        'VorlaufTemp_IST',            '1/15/0/116/0',
        'VorlaufTemp_SOLL',           '1/15/0/116/1',
        'KesselTemp_IST',             '1/60/0/155/0',
        'KesselTemp_SOLL',            '1/60/0/156/0',
        'PufferTemp_Unten',           '1/15/0/118/0',
        'KesselReinigung',            '1/60/0/156/3',
        'Kesselleistung',             '1/60/0/156/7',
        'AbgasTemperatur',            '1/60/0/156/9',
        'Brennerstarts',              '1/60/0/156/8',
        'Pelletsverbrauch',           '1/60/0/156/10'
