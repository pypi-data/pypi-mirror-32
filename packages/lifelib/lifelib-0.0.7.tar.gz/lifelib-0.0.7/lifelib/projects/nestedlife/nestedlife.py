"""The main module to build a simplelife model.

This module contains only one function :py:func:`build`,
which creates a model from source modules and return it.

If this module is run as a script, the :py:func:`build` function is called
and the created model is available as ``model`` global variable.
"""

import os
import modelx as mx

def build(load_saved=False):
    """Build a model and return it.

    Read input data from `input.xlsm`, create `Input` space and its
    subspace and cells and populate them with the data.

    Args:
        load_saved: If ``True``, input data is read from `lifelib.mx` file
            instead of `input.xlsm`, which is saved when
            :py:func:`build_input <simplelife.build_input.build_input>`
            is executed last time. Defaults to ``False``
    """

    # Make sure the current directory is this folder
    os.chdir(os.path.abspath(os.path.dirname(__file__)))

    # ------------------------------------------------------------------------
    # Build Input space

    from build_input import build_input

    if load_saved:
        model = mx.open_model('nestedlife.mx')
        input = model.Input
    else:
        model = mx.new_model(name='nestedlife')
        input = build_input(model, 'input.xlsm')
        model.save('nestedlife.mx')

    # ------------------------------------------------------------------------
    # Build CommFunc space

    lifetable_refs = {'Input': input}

    def lifetable_params(Sex, IntRate, TableID):
        refs={'MortalityTable': Input.MortalityTables(TableID).MortalityTable}
        return {'bases': _self,
                'refs': refs}

    lifetable = model.import_module(
        module_='lifetable',
        name='LifeTable',
        formula=lifetable_params,
        refs=lifetable_refs)

    # ------------------------------------------------------------------------
    # Build Policy space

    from policy import policy_attrs

    policy_refs = {'PolicyData': input.PolicyData,
                   'ProductSpec': input.ProductSpec,
                   'LifeTable': lifetable,
                   'PolicyAttrs': policy_attrs}

    def policy_params(PolicyID):
        refs = {attr: PolicyData[PolicyID].cells[attr] for attr in PolicyAttrs}
        alias = {'PremTerm': refs['PolicyTerm'],
                 'x': refs['IssueAge'],
                 'm': refs['PolicyTerm'],
                 'n': refs['PolicyTerm']}

        refs.update(alias)
        return {'bases': _self,
                'refs': refs}

    policy = model.import_module(
        module_='policy',
        name='Policy',
        formula=policy_params,
        refs=policy_refs)

    # ------------------------------------------------------------------------
    # Build Assumptions space

    asmp_refs = {'Policy': policy,
                 'ProductSpec': input.ProductSpec,
                 'MortalityTables': input.MortalityTables,
                 'asmp': input.Assumptions,
                 'asmp_tbl': input.AssumptionTables}

    def asmp_params(PolicyID):
        refs = {'pol': Policy[PolicyID]}
        alias = {'prd': refs['pol'].Product,
                 'polt': refs['pol'].PolicyType,
                 'gen': refs['pol'].Gen}
        refs.update(alias)
        return {'bases': _self,
                'refs': refs}

    asmp = model.import_module(
        module_='assumptions',
        name='Assumptions',
        formula=asmp_params,
        refs=asmp_refs)

    asmp.allow_none = True

    # ------------------------------------------------------------------------
    # Build Assumptions space

    def econ_params(ScenID):
        refs = {'Scenario': Input.Scenarios[ScenID]}
        return {'bases': _self,
                'refs': refs}

    economic = model.import_module(
        module_='economic',
        name='Economic',
        formula=econ_params,
        refs={'asmp': asmp,
              'Input': input})

    # ------------------------------------------------------------------------
    # Build Projection space
    
    # Model tree structure
    # 
    # lifelib --+
    #           +--BaseProjection
    #           +--OuterProjection[PolicyID] <--- BaseProjection
    #                    +--InnerProjection[t] <-- BaseProjection

    proj_refs = {'Pol': policy,
                 'Asmp': asmp,
                 'Scen': economic}

    def proj_params(PolicyID, ScenID=1):
        refs = {'pol': Pol[PolicyID],
                'asmp': Asmp[PolicyID],
                'scen': Scen[ScenID]}
        return {'bases': _self,
                'refs': refs}

    pvmixin = model.import_module(
        module_='present_values',
        name='PV_Mixin')

    baseproj = model.import_module(
        module_='projection',
        name='BaseProjection',
        bases=pvmixin)


    outerproj = model.new_space(
        bases=baseproj,
        name='OuterProjection',
        formula=proj_params,
        refs=proj_refs)

    def innerproj_params(t0):
        refs = {'pol': _self.parent.pol,
                'asmp': _self.parent.asmp,
                'scen': _self.parent.scen,
                'outer': _self.parent,
                'DiscRate': _self.parent.scen.DiscRate}
        
        return {'bases': _self,
                'refs': refs}

    innerproj = outerproj.new_space(
        bases=baseproj,
        name='InnerProjection',
        formula=innerproj_params)

    return model


if __name__ == '__main__':
    import pandas as pd
    import numpy as np
    model = build(load_saved=True)
    outer = model.OuterProjection

    vars = ['prj_incm_Premium',
            'prj_bnft_Surrender',
            'prj_bnft_Death',
            'prj_exps_Maint',
            'prj_exps_CommTotal',
            'prj_exps_Acq']

    polid = 171

    for cells in vars:
        list(getattr(outer[polid], cells)(t) for t in range(50))

    cfs = outer[polid].frame[vars].sort_index().dropna()
    cfs[vars[1:]] = cfs[vars[1:]].mul(-1)

    [outer[polid].prj_NetLiabilityCashflow[t] for t in range(50)]


    ncf = outer[polid].prj_NetLiabilityCashflow.frame.sort_index()

    inner = outer[polid].InnerProjection
    ncf = [[inner[t0].prj_NetLiabilityCashflow[t] for t in range(t0, 10)]
           for t0 in range(0, 6)]

    data = []
    for t0 in range(0, 6):
        for t in range(0, 11):
            data.append(
                {'t0': t0,
                 't': t,
                 'NCF': inner[t0].prj_NetLiabilityCashflow[t]
                    if t >= t0 else np.nan})

    df = pd.DataFrame(data)

    import seaborn as sns
    import matplotlib.pyplot as plt
