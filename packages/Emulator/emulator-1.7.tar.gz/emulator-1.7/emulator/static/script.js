var __global_static_device_id = "";
var __global_static_div_type = 0;
var __global_static_form_name = "";

var devices = []
var file_path = "../static/images/";

var events = {
    "switch": {
        "color": ["red", "blue", "green"],
        "meter": ["kWh", "W", "A", "V"]
    },
    "light": {
        "color": ["red", "blue", "green", "cold_w", "warm_w"]
    },
    "multi": {
        "alarm_burglar": ["temper_removed_cover"],
        "sensor_temp": ["C", "F"],
        "sensor_lumin": ["Lux"],
        "sensor_humid": ["%"],
        "sensor_uv": ["index"],
        "battery": ["full", "half", "low", "dead"]
    },
    "dimmer": {
        "alarm_heat": ["overheat"],
        "alarm_power": ["surge", "voltage_drop", "over_current", "load_error"],
        "alarm_system": ["hw_failure"],
        "sensor_power": ["W"]
    },
    "flood": {
        "alarm_water": ["leak"],
        "alarm_gp": ["gp"],
        "sensor_temp": ["C"],
        "battery": ["full", "half", "low", "dead"]
    },
    "door": {
        "alarm_lock": ["rf_not_locked"],
        "battery": ["full", "half", "low", "dead"]
    },
    "alarm": {
        "alarm_burglar": ["temper_removed_cover"],
        "alarm_heat": ["overheat"],
        "alarm_power": ["ac_off", "ac_on"],
        "alarm_gas": ["CO"],
        "alarm_fire": ["smoke", "smoke_test"],
        "battery": ["full", "half", "low", "dead"]
    }
}


function getCurrentCell(myTable) {
    var currentRowIndex = myTable.rows.length;
    var currentColIndex;
    var currentCell;

    if (currentRowIndex > 0) // If at least one row exists, get the number of columns in that row
    {
        currentColIndex = myTable.rows[currentRowIndex - 1].cells.length;
    }

    if (currentRowIndex == 0) // If no rows exist, create a new row
    {
        currentRow = myTable.insertRow(currentRowIndex);
        currentCell = myTable.rows[currentRowIndex].insertCell(-1);
    } else if (currentColIndex > 3) // If a row has 4 cells, then create a new row
    {
        currentRow = myTable.insertRow(currentRowIndex);
        currentCell = currentRow.insertCell(-1);
    } else // Insert data into the current cell 
    {
        currentCell = myTable.rows[currentRowIndex - 1].insertCell(-1);
    }

    return currentCell;
}

function addRow(type, status) {
    var color = "";
    var div_tag = document.createElement("div");

    var img_div_tag = document.createElement("div");
    img_div_tag.setAttribute("style", "width:100%;height:80%");
    var img = document.createElement("img");
    img.setAttribute("type", "button");
    img.setAttribute("class", "w3-circle w3-margin-top");
    img.setAttribute("style", "width:100%;object-fit:contain");
    img.setAttribute("id", __global_static_device_id);
    img.onclick = function() { device_on_click(this.id, __global_static_form_name); };
    img.src = file_path + devices[type] + ".png";
    img_div_tag.appendChild(img);

    var toggle_div_tag = document.createElement("div");
    toggle_div_tag.setAttribute("style", "width:100%;height:20%");
    var toggle = document.createElement("button");
    toggle.setAttribute("id", "toggle_" + __global_static_device_id);
    toggle.setAttribute("class", "w3-button w3-block w3-section");
    if (status == "false") {
        status = "Off";
        toggle.setAttribute("style", "color:white;background-color:#4f78e6;");
    } else {
        status = "On";
        toggle.setAttribute("style", "color:white;background-color:#0fa22e;");
    }
    //toggle.setAttribute("innerHTML", status);
    toggle.setAttribute("value", status);
    toggle.setAttribute("type", "submit");
    toggle.setAttribute("onclick", "javascript:toggle_switch(this.id); return false;");
    toggle_div_tag.appendChild(toggle);

    div_tag.appendChild(img_div_tag);
    div_tag.appendChild(toggle_div_tag);

    return div_tag;
}

function init() {
    var dummy_msg = "Init!";
    $.ajax({
        type: "POST",
        url: "/init_onload",
        data: dummy_msg,
        contentType: "application/json;charset=UTF-8",
        success: function(result) {
            init2(result);
            //populate_devices_table(result);
        }
    });
}

function init2(result1) {
    var dummy_msg = "Init2";
    $.ajax({
        type: "POST",
        url: "/init2_onload",
        data: dummy_msg,
        contentType: "application/json;charset=UTF-8",
        success: function(result) {
            populate_devices_to_add_table(result);
            populate_devices_table(result1);
        }
    });
}

/*
<img id="switch" name="switch" src="../static/images/switch.png" alt="switch" 
style="width:100%" class="w3-circle w3-margin-top" 
onclick="javascript:return device_on_click(this.id, 'device_prop');">
*/

function populate_devices_to_add_table(result) {
    var old_tbody = document.getElementById("tbody_add");
    var new_tbody = document.createElement("tbody");
    new_tbody.setAttribute("id", "tbody2");

    files = JSON.parse(result);
    console.log(files);

    for (var i = 0; i < files.length; i++) {
        var tr_tag;
        if (files[i] == "plus") {
            if ((i - 1) < files.length) {
                var tmp = files[i];
                files[i] = files[files.length - 1];
                files[files.length - 1] = tmp;
            }
        }
        devices.push(files[i]);
        console.log(devices);

        if (i == 0)
            tr_tag = document.createElement("tr");

        var td_tag = document.createElement("td");

        var img = document.createElement("img");
        img.setAttribute("class", "w3-circle w3-margin-top");
        img.setAttribute("style", "width:100%;");
        img.setAttribute("id", files[i]);
        img.setAttribute("name", files[i]);
        img.setAttribute("alt", files[i]);
        img.setAttribute("onclick", "javascript:return device_on_click(this.id, 'device_prop');");
        img.src = file_path + files[i] + ".png";

        td_tag.appendChild(img);
        tr_tag.appendChild(td_tag);

        if ((i + 1) % 4 == 0) {
            new_tbody.appendChild(tr_tag);
            tr_tag = document.createElement("tr");
            console.log("row done");
        }
    }

    old_tbody.parentNode.replaceChild(new_tbody, old_tbody);

    var new_tbody2 = document.getElementById("tbody2");
    new_tbody2.id = "tbody_add";

    console.log(devices);
}

function modal_hide_ok(name) {
    document.getElementById(name).style.display = "none";

    var form = "";
    for (var j = 0; j < devices.length; j++) {
        if (__global_static_device_id.indexOf(devices[j]) > -1) {
            form = devices[j] + "_form";
            break;
        }
    }

    var form_elements = document.getElementById(form);
    var jsondata = "{ ";
    var i;

    for (i = 0; i < form_elements.length; i++) {
        var name = form_elements.elements[i].getAttribute("name");
        if (name.indexOf('services') > -1) {
            if (form_elements.elements[i].checked == true)
                jsondata += "\"" + name + "\" : \"" + form_elements.elements[i].checked + "\", ";
            else
                jsondata += "\"" + name + "\" : \"\", ";
        } else
            jsondata += "\"" + name + "\" : \"" + form_elements.elements[i].value + "\", ";
    }
    var status = document.getElementById("toggle_" + __global_static_device_id).value;
    if (status == "On")
        jsondata += "\"" + "device_status" + "\" : \"" + String(true) + "\", ";
    else
        jsondata += "\"" + "device_status" + "\" : \"" + String(false) + "\", ";
    jsondata += "\"id\" : \"" + String(__global_static_device_id) + "\"} ";

    $.ajax({
        type: "POST",
        contentType: "application/json;charset=UTF-8",
        data: JSON.stringify(jsondata),
        dataType: "json",
        url: "/properties",
        success: function(result) {
            console.log(result);
        },
        error: function(error) {
            console.log(error);
        }
    });
}

function toggle_switch(id) {
    var value = {};
    var button = document.getElementById(id);
    if (button.value == "On") {
        button.value = "Off";
        //button.innerHTML = "Off";
        button.style = "background-color:#4f78e6";
        value[id] = false;
    } else {
        button.value = "On";
        //button.innerHTML = "On";
        button.style = "background-color:#0fa22e";
        value[id] = true;
    }
    // Based on the result from the ajax call, switch on/off the device
    $.ajax({
        type: "POST",
        url: "/toggle_device_status",
        data: JSON.stringify(value),
        contentType: "application/json;charset=UTF-8",
        success: function(result) {
            console.log(result);
        }
    });
}

function device_on_click(id, form) {
    __global_static_device_id = id;
    __global_static_form_name = form;

    if (__global_static_div_type == 1) {
        var modal = document.getElementById("modal");
        modal.style.display = "block";

        var remove_button = document.getElementById("remove_device");
        remove_button.setAttribute("style", "visibility:visible");

        $.ajax({
            type: "POST",
            url: "/get_properties",
            data: id,
            contentType: "application/json;charset=UTF-8",
            success: function(result) {
                populate_device_properties(result);
            }
        });
    } else {
        var myTable = document.getElementById("tbody1");
        var type;
        for (var j = 0; j < devices.length; j++) {
            if (__global_static_device_id == devices[j])
                type = j;
        }

        __global_static_device_id += Math.floor(Math.random() * 10000);

        $.ajax({
            type: "POST",
            url: "/add_device",
            data: __global_static_device_id,
            contentType: "application/json;charset=UTF-8",
            success: function(result) {
                console.log(result);
                var data = JSON.parse(result);
                var currentCell = getCurrentCell(myTable);
                var img = addRow(type, data["val"]["device_status"]);
                currentCell.appendChild(img);
            }
        });
    }
}

function populate_device_properties(result) {
    console.log("populate_device_properties: received: " + result);
    if (result == "Error")
        return

    var prop = JSON.parse(result);
    var odiv_prop = document.getElementById("device_prop");

    var ndiv_prop = document.createElement('div');
    ndiv_prop.setAttribute("id", "device_prop2");
    ndiv_prop.setAttribute("class", "modal-content w3-container");
    ndiv_prop.setAttribute("style", "max-width:600px;max-height:600px;overflow:auto");

    var type;
    for (var j = 0; j < devices.length; j++) {
        if (__global_static_device_id.indexOf(devices[j]) > -1) {
            type = j;
            break;
        }
    }

    // Create heading
    var hp = document.createElement("p");
    var hp_text = document.createTextNode("Device Properties");
    hp.appendChild(hp_text);
    ndiv_prop.appendChild(hp);

    // Create form elements
    var form_elements = document.createElement("form");
    form_elements.setAttribute("id", devices[type] + "_form");
    form_elements.setAttribute("role", "form");
    form_elements.setAttribute("method", "POST");

    // Text input 1
    var text1 = document.createElement("input");
    var label1 = document.createElement("label");
    var label_name1 = document.createTextNode("\tProduct Name");
    var paragraph1 = document.createElement("p");

    text1.setAttribute("name", "product_name");
    text1.setAttribute("class", "w3-input");
    text1.setAttribute("type", "text");
    text1.setAttribute("value", prop["product_name"]);

    label1.setAttribute("for", "Product Name");
    label1.appendChild(text1);
    label1.appendChild(label_name1);

    paragraph1.appendChild(label1);

    form_elements.appendChild(paragraph1);

    // Text input 2
    var text2 = document.createElement("input");
    var label2 = document.createElement("label");
    var label_name2 = document.createTextNode("\tDevice Id");
    var paragraph2 = document.createElement("p");

    text2.setAttribute("name", "device_id");
    text2.setAttribute("class", "w3-input");
    text2.setAttribute("type", "text");
    text2.setAttribute("value", prop["device_id"]);

    label2.setAttribute("for", "Device Id");
    label2.appendChild(text2);
    label2.appendChild(label_name2);

    paragraph2.appendChild(label2);
    form_elements.appendChild(paragraph2);

    // Create form elements for device interfaces
    var events_keys = Object.keys(events[devices[type]]);
    var events_length = events_keys.length;

    if ("services" in prop) {
        for (var evt = 0; evt < events_length; evt++) {
            var evt_name = events_keys[evt].toString();
            var paragraph = document.createElement("p");
            paragraph.innerHTML = events_keys[evt];
            paragraph.appendChild(document.createElement("br"));

            var evt_value = events[devices[type]][events_keys[evt]];
            for (var arr = 0; arr < evt_value.length; arr++) {
                var checkBox = document.createElement("input");
                var label = document.createElement("label");
                var label_name = document.createTextNode("\t\t" + evt_value[arr]);

                checkBox.setAttribute("name", "services_" + evt_value[arr]);
                checkBox.setAttribute("type", "checkbox");
                checkBox.setAttribute("id", evt_value[arr]);

                label.setAttribute("for", "\t\t" + evt_value[arr]);
                label.appendChild(checkBox);
                label.appendChild(label_name);

                paragraph.appendChild(label);

                form_elements.appendChild(paragraph);
            }
        }
    } else {
        for (var evt = 0; evt < events_length; evt++) {
            var evt_name = events_keys[evt].toString();
            var paragraph = document.createElement("p");
            paragraph.innerHTML = events_keys[evt];
            paragraph.appendChild(document.createElement("br"));

            var evt_value = events[devices[type]][events_keys[evt]];
            for (var arr = 0; arr < evt_value.length; arr++) {
                var checkBox = document.createElement("input");
                var label = document.createElement("label");
                var label_name = document.createTextNode("\t\t" + evt_value[arr]);

                checkBox.setAttribute("name", "services_" + evt_value[arr]);
                checkBox.setAttribute("type", "checkbox");
                checkBox.setAttribute("id", evt_value[arr]);

                if (prop["services_" + evt_value[arr]] == "true") {
                    checkBox.setAttribute("checked", "true");
                    if (evt_name == "battery") {
                        var button = document.getElementById("toggle_" + __global_static_device_id);
                        if (evt_value[arr] == "dead") {
                            button.value = "Off";
                            //button.innerHTML = "Off";
                            button.style = "background-color:#4f78e6";
                        } else {
                            button.value = "On";
                            //button.innerHTML = "On";
                            button.style = "background-color:#0fa22e";
                        }
                    }

                }

                label.setAttribute("for", "\t\t" + evt_value[arr]);
                label.appendChild(checkBox);
                label.appendChild(label_name);

                paragraph.appendChild(label);

                form_elements.appendChild(paragraph);
            }
        }
    }

    ndiv_prop.appendChild(form_elements);
    odiv_prop.parentNode.replaceChild(ndiv_prop, odiv_prop);

    var new_div2 = document.getElementById("device_prop2");
    new_div2.id = "device_prop";
}

function populate_devices_table(result) {
    var id_array = [];
    var st_array = [];

    var s = JSON.parse(result);
    JSON.parse(result).forEach(o => (id_array.push(o.id), st_array.push(o.device_status)));
    console.log(st_array)

    var new_tbody = document.createElement("tbody");
    new_tbody.setAttribute("id", "tbody2");
    var old_tbody = document.getElementById("tbody1");

    for (var i = 0; i < id_array.length; i++) {
        for (var j = 0; j < devices.length; j++) {
            if (id_array[i].indexOf(devices[j]) > -1) {
                // Device found
                __global_static_device_id = id_array[i];
                var currentCell = getCurrentCell(new_tbody);
                var img = addRow(j, st_array[i]);
                currentCell.appendChild(img);
            }
        }
    }

    old_tbody.parentNode.replaceChild(new_tbody, old_tbody);

    var new_tbody2 = document.getElementById("tbody2");
    new_tbody2.id = "tbody1";
    console.log("populate_devices_table: Done!");
}

function modal_hide_cancel(name) {
    document.getElementById(name).style.display = "none";
}

function remove_device(name) {
    document.getElementById(name).style.display = "none";

    id = String(__global_static_device_id);

    $.ajax({
        type: "POST",
        url: "/remove_device",
        data: id,
        contentType: "application/json;charset=UTF-8",
        success: function(result) {
            populate_devices_table(result);
        }
    });
}

function set_div_type(type) {
    __global_static_div_type = type;
}