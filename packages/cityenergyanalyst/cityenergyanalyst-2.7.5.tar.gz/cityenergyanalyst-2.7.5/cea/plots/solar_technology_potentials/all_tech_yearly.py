from __future__ import division
from __future__ import print_function

import plotly.graph_objs as go
from plotly.offline import plot

from cea.plots.variable_naming import LOGO, COLOR, NAMING

def all_tech_district_yearly(data_frame, pv_analysis_fields, pvt_analysis_fields, sc_fp_analysis_fields, sc_et_analysis_fields, title,
                             output_path):
    E_analysis_fields = []
    Q_analysis_fields = []

    pv_E_analysis_fields_used = data_frame.columns[data_frame.columns.isin(pv_analysis_fields)].tolist()
    E_analysis_fields.extend(pv_E_analysis_fields_used)
    sc_fp_Q_analysis_fields_used = data_frame.columns[data_frame.columns.isin(sc_fp_analysis_fields)].tolist()
    Q_analysis_fields.extend(sc_fp_Q_analysis_fields_used)
    sc_et_Q_analysis_fields_used = data_frame.columns[data_frame.columns.isin(sc_et_analysis_fields)].tolist()
    Q_analysis_fields.extend(sc_et_Q_analysis_fields_used)
    pvt_E_analysis_fields_used = data_frame.columns[data_frame.columns.isin(pvt_analysis_fields[0:5])].tolist()
    E_analysis_fields.extend(pvt_E_analysis_fields_used)
    pvt_Q_analysis_fields_used = data_frame.columns[data_frame.columns.isin(pvt_analysis_fields[5:10])].tolist()
    Q_analysis_fields.extend(pvt_Q_analysis_fields_used)
    data_frame_MWh = data_frame / 1000  # to MWh

    # CALCULATE GRAPH
    traces_graph, x_axis = calc_graph(E_analysis_fields, Q_analysis_fields, data_frame_MWh)

    # CALCULATE TABLE
    traces_table = calc_table(E_analysis_fields, Q_analysis_fields, data_frame_MWh)

    # PLOT GRAPH
    traces_graph.append(traces_table)

    # CREATE BUTTON
    annotations = list(
        [dict(
            text='<b>In this plot, users can explore the combined potentials of all solar technologies.</b><br>'
                 'Instruction:'
                 'Click on the technologies to install on each building surface.<br>'
                 'Example: PV_walls_east_E + PVT_walls_south_E/Q + SC_FP_roofs_top_Q <br><br>'
            , x=0.8, y=1.1,
            xanchor='left', xref='paper', yref='paper', align='left', showarrow=False, bgcolor="rgb(254,220,198)")])

    layout = go.Layout(images=LOGO, title=title, barmode='stack', annotations=annotations,
                       yaxis=dict(title='Electricity/Thermal Potential [MWh/yr]', domain=[0.35, 1]),
                       xaxis=dict(title='Building'), legend=dict(x=1, y=0.1, xanchor='left'))
    fig = go.Figure(data=traces_graph, layout=layout)
    plot(fig, auto_open=False, filename=output_path)

    return {'data': traces_graph, 'layout': layout}


def calc_graph(E_analysis_fields, Q_analysis_fields, data_frame):
    # calculate graph
    graph = []

    data_frame['total_E'] = total_E = data_frame[E_analysis_fields].sum(axis=1)
    data_frame['total_Q'] = total_Q = data_frame[Q_analysis_fields].sum(axis=1)
    # data_frame = data_frame.sort_values(by='total', ascending=False) # this will get the maximum value to the left
    for field in E_analysis_fields:
        y = data_frame[field]
        total_perc = (y / total_E.sum() * 100).round(2)
        total_perc_txt = ["(" + str(x) + " %)" for x in total_perc]
        if field.split('_')[0] == 'PVT':
            trace1 = go.Bar(x=data_frame.index, y=y, name=field.split('_kWh', 1)[0], text=total_perc_txt,
                            marker=dict(color=COLOR[field]), visible='legendonly',
                            width=0.3, offset=-0.35, legendgroup=field.split('_')[1])
        else:
            trace1 = go.Bar(x=data_frame.index, y=y, name=field.split('_kWh', 1)[0], text=total_perc_txt,
                            marker=dict(color=COLOR[field]), visible='legendonly',
                            width=0.3, offset=-0.35)
        graph.append(trace1)

    for field in Q_analysis_fields:
        y = data_frame[field]
        total_perc = (y / total_Q * 100).round(2).values
        total_perc_txt = ["(" + str(x) + " %)" for x in total_perc]
        if field.split('_')[0] == 'PVT':
            trace2 = go.Bar(x=data_frame.index, y=y, name=field.split('_kWh', 1)[0], text=total_perc_txt,
                            marker=dict(color=COLOR[field], line=dict(
                                color="rgb(105,105,105)", width=1)), opacity=0.7, visible='legendonly', base=0,
                            width=0.3, offset=0, legendgroup=field.split('_')[1])
        elif field.split('_')[1]== 'FP':
            trace2 = go.Bar(x=data_frame.index, y=y, name=field.split('_kWh', 1)[0], text=total_perc_txt,
                            marker=dict(color=COLOR[field], line=dict(
                                color="rgb(105,105,105)", width=1)), opacity=0.7, visible='legendonly', base=0,
                            width=0.3, offset=0)
        elif field.split('_')[1]=='ET':
            trace2 = go.Bar(x=data_frame.index, y=y, name=field.split('_kWh', 1)[0], text=total_perc_txt,
                            marker=dict(color=COLOR[field], line=dict(
                                color="rgb(105,105,105)", width=1)), opacity=0.7, visible='legendonly', base=0,
                            width=0.3, offset=0)
        else: raise ValueError('the specified analysis field is not in the right form: ',field)

        graph.append(trace2)

    return graph, data_frame.index,


def calc_table(E_analysis_fields, Q_analysis_fields, data_frame):
    analysis_fields = []
    total_perc = []
    median = []
    # find the three highest
    anchors = []
    load_names = []

    E_median = data_frame[E_analysis_fields].median().round(2).tolist()
    E_total = data_frame[E_analysis_fields].sum().round(2).tolist()
    if sum(E_total) > 0:
        E_total_perc = [str(x) + " (" + str(round(x / sum(E_total) * 100, 1)) + " %)" for x in E_total]
        for field in E_analysis_fields:
            anchors.append(calc_top_three_anchor_loads(data_frame, field))
            load_names.append(NAMING[field] + ' (' + field.split('_kWh', 1)[0] + ')')
    else:
        E_total_perc = ['0 (0%)']*len(E_total)
        for field in E_analysis_fields:
            anchors.append('-')
            load_names.append(NAMING[field] + ' (' + field.split('_kWh', 1)[0] + ')')
    analysis_fields.extend(E_analysis_fields)
    total_perc.extend(E_total_perc)
    median.extend(E_median)

    Q_median = (data_frame[Q_analysis_fields].median() / 1000).round(2).tolist()  # to MWh
    Q_total = data_frame[Q_analysis_fields].sum().round(2).tolist()
    if sum(Q_total) > 0:
        Q_total_perc = [str(x) + " (" + str(round(x / sum(Q_total) * 100, 1)) + " %)" for x in Q_total]
        for field in Q_analysis_fields:
            anchors.append(calc_top_three_anchor_loads(data_frame, field))
            load_names.append(NAMING[field] + ' (' + field.split('_kWh', 1)[0] + ')')
    else:
        Q_total_perc = ['0 (0%)']*len(Q_total)
        for field in Q_analysis_fields:
            anchors.append('-')
            load_names.append(NAMING[field] + ' (' + field.split('_kWh', 1)[0] + ')')

    analysis_fields.extend(Q_analysis_fields)
    total_perc.extend(Q_total_perc)
    median.extend(Q_median)

    analysis_fields = filter(None, [x for field in analysis_fields for x in field.split('_kWh', 1)])




    table = go.Table(domain=dict(x=[0, 1], y=[0.0, 0.2]),
                     header=dict(values=['Surface', 'Total [MWh/yr]', 'Median [MWh/yr]', 'Top 3 most irradiated']),
                     cells=dict(values=[load_names, total_perc, median, anchors]))

    return table


def calc_top_three_anchor_loads(data_frame, field):
    data_frame = data_frame.sort_values(by=field, ascending=False)
    anchor_list = data_frame[:3].index.values
    return anchor_list
