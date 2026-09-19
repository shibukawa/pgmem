package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ENRMetadataGetTupDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = F_table_open(m, v6, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
			F_relation_close(m, v8, int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = v12
				return v17
			}
		}
	} else {
		v17 = v3
		return v17
	}
}
func F_EnterParallelMode(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_EnterParallelMode[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = v4 + int32(1)
	return
}
func F_EstimateSnapshotSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = F_mul_size(m, v5, int32(4))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_add_size(m, int32(24), v7)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v13 <= int32(0) {
				v27 = v11
				return v27
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
				if v16 == int32(1) {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
					if v19 != int32(1) {
						v27 = v11
						return v27
					} else {
						v23 = F_mul_size(m, v13, int32(4))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = F_add_size(m, v11, v23)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								v27 = v25
								return v27
							}
						}
					}
				} else {
					v23 = F_mul_size(m, v13, int32(4))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = F_add_size(m, v11, v23)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = v25
							return v27
						}
					}
				}
			}
		}
	}
}
func F_ExecARInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if l4 == int32(0) {
		if v7 != 0 {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)))
			if v16 != 0 {
				v22 = int32(0)
				F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					return
				}
			} else {
				if l4 == int32(0) {
					return
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
					if v19 != int32(1) {
						return
					} else {
						v22 = int32(0)
						F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			if l4 == int32(0) {
				return
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
				if v19 != int32(1) {
					return
				} else {
					v22 = int32(0)
					F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
		if v10 == int32(0) {
			if v7 != 0 {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)))
				if v16 != 0 {
					v22 = int32(0)
					F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						return
					}
				} else {
					if l4 == int32(0) {
						return
					} else {
						v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
						if v19 != int32(1) {
							return
						} else {
							v22 = int32(0)
							F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				if l4 == int32(0) {
					return
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
					if v19 != int32(1) {
						return
					} else {
						v22 = int32(0)
						F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
			if v13 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_ExecARInsertTriggers_0), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecARInsertTriggers_1), int32(2556), int32(_a_F_ExecARInsertTriggers_2))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if v7 != 0 {
					v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)))
					if v16 != 0 {
						v22 = int32(0)
						F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							return
						}
					} else {
						if l4 == int32(0) {
							return
						} else {
							v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
							if v19 != int32(1) {
								return
							} else {
								v22 = int32(0)
								F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					if l4 == int32(0) {
						return
					} else {
						v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
						if v19 != int32(1) {
							return
						} else {
							v22 = int32(0)
							F_AfterTriggerSaveEvent(m, l0, l1, v22, v22, v22, int32(1), v22, l2, l3, v22, l4, v22)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecBSInsertTriggers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v19 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L22
	}
L2:
	;
	m.G0 = v9 + int32(48)
	return
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	if v22 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+56))
	v28 = F_before_stmt_triggers_fired(m, v26, int32(3))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v28 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(34359738810)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v34 <= int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v42 = v3
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v46 = v43 + v42*int32(60)
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+12)))
	if v47&int32(71) != int32(6) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L2
L11:
	;
	v74 = v42 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v74 < v75 {
		v42 = v74
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v53 = int32(0)
	v56 = F_TriggerEnabled(m, l0, l1, v46, v52, v53, v53, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if v56 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v46
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v65 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v68 = v65
	goto L17
L16:
	;
	v66 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v70 = F_ExecCallTriggerFunc(m, v9+int32(4), v42, v63, v64, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L19
	}
L18:
	;
	v68 = v66
	goto L17
L19:
	;
	if v70 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L11
L21:
	;
	goto L10
L22:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_ExecBSInsertTriggers_0), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_ExecBSInsertTriggers_1), int32(2448), int32(_a_F_ExecBSInsertTriggers_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecBSUpdateTriggers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v3
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v14
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v20 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L23
	}
L2:
	;
	m.G0 = v10 + int32(48)
	return
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)))
	if v23 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	v29 = F_before_stmt_triggers_fired(m, v27, int32(2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v29 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v31 = F_ExecGetAllUpdatedCols(m, l1, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(42949673402)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v38 <= int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v45 = int32(0)
	goto L10
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v52 = v49 + v45*int32(60)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+12)))
	if v53&int32(83) != int32(18) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L2
L12:
	;
	v79 = v45 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v79 < v80 {
		v45 = v79
		goto L10
	} else {
		goto L22
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v59 = int32(0)
	v61 = F_TriggerEnabled(m, l0, l1, v52, v58, v31, v59, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v61 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v52
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v70 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v73 = v70
	goto L18
L17:
	;
	v71 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L19
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v75 = F_ExecCallTriggerFunc(m, v10+int32(4), v45, v68, v69, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v73 = v71
	goto L18
L20:
	;
	if v75 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	goto L11
L23:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(_a_F_ExecBSUpdateTriggers_0), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_ExecBSUpdateTriggers_1), int32(2949), int32(_a_F_ExecBSUpdateTriggers_2))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecBatchInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int64
	_ = v61
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l4
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	v21 = m.T0[v20].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l5, l1, l2, l3, v14+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v23 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l4 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v34 = int32(0)
	goto L5
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21+v34<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v42
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	F_ExecARInsertTriggers(m, l5, l1, v40, int32(0), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v56 = int32(0)
	if base.B2i32(l6 == v56)|base.B2i32(v54 <= v56) != 0 {
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v48 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_ExecWithCheckOptions(m, int32(0), l1, v40, l5)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v53 = v34 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v53 < v54 {
		v34 = v53
		goto L5
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	goto L6
L13:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l5)+112))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+112)) = v61 + base.I64_extend_i32_u(v54)
	goto L3
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(0)
	m.G0 = v14 + int32(16)
	return
L15:
	;
	v78 = int32(0)
	if l4 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v91 = int32(0)
	v94 = v78
	goto L19
L17:
	;
	v140 = v78
	goto L18
L18:
	;
	v144 = v140 << (uint(int32(2)) % 32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2+v144)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	m.T0[v148].(func(*base.Module, int32))(m, v146)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v98 = v94 << (uint(int32(2)) % 32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2+v98)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	m.T0[v102].(func(*base.Module, int32))(m, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	if l4&int32(1) == int32(0) {
		goto L14
	} else {
		goto L26
	}
L21:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98+l3)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	m.T0[v108].(func(*base.Module, int32))(m, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v112 = v98 | int32(4)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2+v112)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	m.T0[v116].(func(*base.Module, int32))(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v112+l3)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	m.T0[v122].(func(*base.Module, int32))(m, v120)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v125 = int32(2)
	v126 = v94 + v125
	v128 = v91 + v125
	if v128 != l4&int32(2147483646) {
		v91 = v128
		v94 = v126
		goto L19
	} else {
		goto L25
	}
L25:
	;
	goto L20
L26:
	;
	v140 = v126
	goto L18
L27:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v144+l3)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	m.T0[v154].(func(*base.Module, int32))(m, v152)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L14
}
func F_ExecBuildAuxRowMark(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v12 = F_palloc0(m, int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v18 != int32(5) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L78
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L75
	}
L5:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v174 != v175 {
		goto L52
	} else {
		goto L53
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v17
	v23 = v9 + int32(96)
	v28 = F_pg_snprintf(m, v23, int32(32), int32(_a_F_ExecBuildAuxRowMark_0), v9+int32(80))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v17
	v105 = v9 + int32(96)
	v110 = F_pg_snprintf(m, v105, int32(32), int32(_a_F_ExecBuildAuxRowMark_1), v9+int32(48))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L32
	}
L9:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v86)
	if v86 != 0 {
		goto L5
	} else {
		goto L28
	}
L11:
	;
	v86 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v37 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = int32(0)
	if v40 < v37 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v78 = int32(0)
	goto L16
L16:
	;
	v86 = base.I32_extend16_s(v78)
	goto L10
L17:
	;
	v43 = v37
	goto L19
L18:
	;
	v43 = v40
	goto L19
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v46 = int32(0)
	goto L21
L20:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+8)))
	v78 = v70
	goto L16
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v44+v46<<(uint(int32(2))%32))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+26)))
	if v56 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v86 = int32(0)
	goto L10
L23:
	;
	v67 = v46 + int32(1)
	if v67 != v43 {
		v46 = v67
		goto L21
	} else {
		goto L27
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v59 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v62 = F_strcmp(m, v59, v23)
	mBase = m.M
	if v62 == int32(0) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	goto L22
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v23
	F_errmsg_internal(m, int32(_a_F_ExecBuildAuxRowMark_2), v9-int32(-64))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2598), int32(_a_F_ExecBuildAuxRowMark_4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	if l1 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)) = uint16(v168)
	if v168 == int32(0) {
		goto L4
	} else {
		goto L51
	}
L34:
	;
	v168 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v119 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v122 = int32(0)
	if v122 < v119 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v160 = int32(0)
	goto L39
L39:
	;
	v168 = base.I32_extend16_s(v160)
	goto L33
L40:
	;
	v125 = v119
	goto L42
L41:
	;
	v125 = v122
	goto L42
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v128 = int32(0)
	goto L44
L43:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+8)))
	v160 = v152
	goto L39
L44:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v126+v128<<(uint(int32(2))%32))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+26)))
	if v138 != int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v168 = int32(0)
	goto L33
L46:
	;
	v149 = v128 + int32(1)
	if v149 != v125 {
		v128 = v149
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	if v141 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v144 = F_strcmp(m, v141, v105)
	mBase = m.M
	if v144 == int32(0) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	goto L45
L51:
	;
	goto L5
L52:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v177
	v180 = v9 + int32(96)
	v181 = int32(32)
	v185 = F_pg_snprintf(m, v180, v181, int32(_a_F_ExecBuildAuxRowMark_5), v9+v181)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	m.G0 = v9 + int32(128)
	return v12
L55:
	;
	if l1 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+6)) = uint16(v243)
	if v243 == int32(0) {
		goto L3
	} else {
		goto L74
	}
L57:
	;
	v243 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v194 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v197 = int32(0)
	if v197 < v194 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v235 = int32(0)
	goto L62
L62:
	;
	v243 = base.I32_extend16_s(v235)
	goto L56
L63:
	;
	v200 = v194
	goto L65
L64:
	;
	v200 = v197
	goto L65
L65:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v203 = int32(0)
	goto L67
L66:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212)+8)))
	v235 = v227
	goto L62
L67:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v201+v203<<(uint(int32(2))%32))))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+26)))
	if v213 != int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v243 = int32(0)
	goto L56
L69:
	;
	v224 = v203 + int32(1)
	if v224 != v200 {
		v203 = v224
		goto L67
	} else {
		goto L73
	}
L70:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)+12))
	if v216 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v219 = F_strcmp(m, v216, v180)
	mBase = m.M
	if v219 == int32(0) {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	goto L68
L74:
	;
	goto L54
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(96)
	F_errmsg_internal(m, int32(_a_F_ExecBuildAuxRowMark_2), v9)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2607), int32(_a_F_ExecBuildAuxRowMark_4))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(96)
	F_errmsg_internal(m, int32(_a_F_ExecBuildAuxRowMark_2), v9+int32(16))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2617), int32(_a_F_ExecBuildAuxRowMark_4))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecCleanTypeFromTL(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_ExecTypeFromTLInternal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ExecDropSingleTupleTableSlot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	m.T0[v4].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		m.T0[v8].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v11 == int32(0) {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
				if v19&int32(16) != 0 {
					F_pfree(m, l0)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						return
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v22 != 0 {
						F_pfree(m, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v25 == int32(0) {
								F_pfree(m, l0)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							} else {
								F_pfree(m, v25)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v25 == int32(0) {
							F_pfree(m, l0)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						} else {
							F_pfree(m, v25)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				if v14 < int32(0) {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
					if v19&int32(16) != 0 {
						F_pfree(m, l0)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							return
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v22 != 0 {
							F_pfree(m, v22)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v25 == int32(0) {
									F_pfree(m, l0)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										return
									}
								} else {
									F_pfree(m, v25)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v32 = m.ExcPending
										if v32 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v25 == int32(0) {
								F_pfree(m, l0)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							} else {
								F_pfree(m, v25)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				} else {
					F_DecrTupleDescRefCount(m, v11)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
						if v19&int32(16) != 0 {
							F_pfree(m, l0)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								return
							}
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v22 != 0 {
								F_pfree(m, v22)
								mBase = m.M
								v24 = m.ExcPending
								if v24 != 0 {
									return
								} else {
									v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v25 == int32(0) {
										F_pfree(m, l0)
										mBase = m.M
										v32 = m.ExcPending
										if v32 != 0 {
											return
										} else {
											return
										}
									} else {
										F_pfree(m, v25)
										mBase = m.M
										v29 = m.ExcPending
										if v29 != 0 {
											return
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v32 = m.ExcPending
											if v32 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v25 == int32(0) {
									F_pfree(m, l0)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										return
									}
								} else {
									F_pfree(m, v25)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v32 = m.ExcPending
										if v32 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecForceStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v11 == int32(_a_F_ExecForceStoreMinimalTuple_0) {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v14&int32(4) != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			F_pfree(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v23 = v20 & int32(-5)
				v24 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v24)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v24
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v24)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
				v34 = v23 & int32(_a_F_ExecForceStoreMinimalTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v34)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v37 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = l0 - v37
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v36 + v37
				if l2 == v24 {
				} else {
					v46 = v34 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v46)
				}
				m.G0 = v9 + int32(32)
				return
			}
		} else {
			v23 = v14
			v24 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v24)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v24
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v24)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
			v34 = v23 & int32(_a_F_ExecForceStoreMinimalTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v34)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v37 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = l0 - v37
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v36 + v37
			if l2 == v24 {
			} else {
				v46 = v34 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v46)
			}
			m.G0 = v9 + int32(32)
			return
		}
	} else {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		m.T0[v48].(func(*base.Module, int32))(m, l1)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v52 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l0 - v52
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v51 + v52
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			F_heap_deform_tuple(m, v9+int32(12), v60, v61, v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v67 = v65 & int32(_a_F_ExecForceStoreMinimalTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v67)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v70)
				if l2 == int32(0) {
					m.G0 = v9 + int32(32)
					return
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
					m.T0[v75].(func(*base.Module, int32))(m, l1)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
							return
						}
					}
				}
			}
		}
	}
}
func F_ExecGather(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[0]))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v20 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	if v25 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+20))
	F_MemoryContextReset(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L29
	}
L9:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v91 != 0 {
		goto L26
	} else {
		goto L27
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+160)))
	if v29 != int32(1) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v34 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	F_LaunchParallelWorkers(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L18
	}
L13:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v38 = F_ExecInitParallelPlan(m, v33, v28, v32, v25, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_ExecParallelReinitialize(m, v33, v34, v32)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v38
	v44 = v38
	goto L12
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v44 = v43
	goto L12
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28)+164))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+164)) = v50 + v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+168)) = v54 + v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	if int32(0) < v58 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(0)
	goto L9
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_ExecParallelCreateReaders(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v79
	goto L19
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v64
	v68 = F_palloc(m, v64<<(uint(int32(2))%32))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v73 = v71 << (uint(int32(2)) % 32)
	if v73 == int32(0) {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+40))
	base.MemoryCopy(m, v68, v77, v73)
	goto L19
L26:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGather[1])))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+80)))
	v98 = v93 & (v94 ^ int32(1))
	goto L28
L27:
	;
	v98 = int32(1)
	goto L28
L28:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v99)
	v102 = v98 & v99
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v102)
	goto L8
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L31
L30:
	;
	m.G0 = v12 + int32(16)
	return v312
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v126 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+4)))
	if v276&int32(2) != 0 {
		v312 = int32(0)
		goto L30
	} else {
		goto L95
	}
L33:
	;
	goto L32
L34:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v242 != int32(1) {
		goto L31
	} else {
		goto L82
	}
L35:
	;
	v227 = int32(0)
	v229 = F_ExecStoreMinimalTuple(m, v163, v115, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L80
	}
L36:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	m.T0[v224].(func(*base.Module, int32))(m, v115)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L79
	}
L37:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v129 != int32(1) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[0]))
	if v133 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	v137 = v126
	goto L43
L43:
	;
	v138 = int32(0)
	if v137 <= v138 {
		goto L34
	} else {
		goto L45
	}
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v137 = v136
	goto L43
L45:
	;
	v143 = v138
	goto L46
L46:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[0]))
	if v151 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154+v155<<(uint(int32(2))%32))))
	v163 = F_TupleQueueReaderNext(m, v159, int32(1), v12+int32(15))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v165 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v170 = v168 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v170
	if v170 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if v163 != 0 {
		goto L35
	} else {
		goto L71
	}
L56:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v174 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v185 = (v170 - v182) << (uint(int32(2)) % 32)
	if v185 != 0 {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	F_ExecParallelFinish(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v177 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	goto L34
L66:
	;
	goto L65
L67:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v189 = v186 + v182<<(uint(int32(2))%32)
	base.MemoryCopy(m, v189, v189+int32(4), v185)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v194 < v195 {
		goto L46
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(0)
	goto L46
L71:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v201 = v199 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v201 < v203 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v205 = v201
	goto L74
L73:
	;
	v205 = int32(0)
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v205
	v208 = v143 + int32(1)
	if v208 < v203 {
		v143 = v208
		goto L46
	} else {
		goto L75
	}
L75:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v210 != 0 {
		goto L34
	} else {
		goto L76
	}
L76:
	;
	v211 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[2]))
	v217 = F_WaitLatch(m, v213, int32(33), v211, int32(134217741))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = int32(0)
	goto L78
L78:
	;
	v143 = v211
	goto L46
L79:
	;
	v267 = v115
	goto L33
L80:
	;
	if v115 == int32(0) {
		v312 = v227
		goto L30
	} else {
		goto L81
	}
L81:
	;
	v267 = v115
	goto L33
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v246 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+24))
	v249 = v247
	goto L85
L84:
	;
	v249 = int32(0)
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+172)) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v116)+52))
	if v251 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_ExecReScan(m, v116)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v255 = m.T0[v254].(func(*base.Module, int32) int32)(m, v116)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+172)) = int32(0)
	if v255 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+4)))
	if v259&int32(2) == int32(0) {
		v267 = v255
		goto L33
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v264 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v264)
	goto L31
L94:
	;
	goto L93
L95:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v279 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v312 = v267
	goto L30
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v267
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+72))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v283)+16))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	m.T0[v287].(func(*base.Module, int32))(m, v285)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v290 = int32(_a_F_ExecGather_0)
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3]))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v284)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3])) = v293
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v283)+24))
	v299 = m.T0[v298].(func(*base.Module, int32, int32, int32) int32)(m, v283+int32(4), v284, int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3])) = v291
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285)+4)))
	v305 = v303 & int32(_a_F_ExecGather_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v285)+4)) = uint16(v305)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v285)+12))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	*(*uint16)(unsafe.Add(mBase, uint32(v285)+6)) = uint16(v308)
	v312 = v285
	goto L30
}
func F_ExecGatherMerge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v14 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	if v18 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	F_MemoryContextReset(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L30
	}
L9:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGatherMerge[1])))
	if v81 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+160)))
	if v22 != int32(1) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v27 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	F_LaunchParallelWorkers(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L18
	}
L13:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v31 = F_ExecInitParallelPlan(m, v26, v21, v25, v18, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_ExecParallelReinitialize(m, v26, v27, v25)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v31
	v37 = v31
	goto L12
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v37 = v36
	goto L12
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+164))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v43 + v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+168)) = v47 + v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if int32(0) < v51 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_ExecParallelCreateReaders(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v72 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v72
	goto L9
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v57
	v61 = F_palloc(m, v57<<(uint(int32(2))%32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v66 = v64 << (uint(int32(2)) % 32)
	if v66 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+40))
	base.MemoryCopy(m, v61, v70, v66)
	goto L9
L25:
	;
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v87)
	goto L8
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v84 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)) = uint8(v85)
	goto L25
L29:
	;
	goto L28
L30:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v98 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if v273 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v103
	if v103 < v101 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	v258 = F_gather_merge_readnext(m, l0, v256, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L80
	}
L35:
	;
	v110 = int32(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+8)) = uint8(v150)
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(0)
	goto L42
L38:
	;
	v117 = v110 << (uint(int32(4)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v117+v118)+4)) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v122+v117)+8)) = v120
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*uint8)(unsafe.Add(mBase, uint32(v126+v117)+12)) = uint8(v120)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v132 = v110 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+v132<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	m.T0[v138].(func(*base.Module, int32))(m, v136)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	if v132 != v101 {
		v110 = v132
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v101 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_binaryheap_build(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L79
	}
L44:
	;
	v158 = int32(1)
	goto L45
L45:
	;
	v165 = int32(0)
	goto L47
L46:
	;
	goto L43
L47:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[0]))
	if v172 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v101 <= int32(0) {
		goto L43
	} else {
		goto L70
	}
L49:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v165 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L51
L53:
	;
	v206 = v165 + int32(1)
	if v206 <= v101 {
		v165 = v206
		goto L47
	} else {
		goto L69
	}
L54:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v185+v165<<(uint(int32(2))%32))))
	if v189 != 0 {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)))
	if v177 != 0 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v179 = int32(4)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v165<<(uint(v179)%32)-v179))))
	if v184 != 0 {
		goto L53
	} else {
		goto L59
	}
L58:
	;
	goto L53
L59:
	;
	goto L54
L60:
	;
	F_load_tuple_array(m, l0, v165)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L68
	}
L61:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+4)))
	if v190&int32(2) == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v195 = F_gather_merge_readnext(m, l0, v165, v158)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v195 == int32(0) {
		goto L53
	} else {
		goto L66
	}
L66:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_binaryheap_add_unordered(m, v199, v165)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	goto L53
L68:
	;
	goto L53
L69:
	;
	goto L48
L70:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v213 = int32(1)
	goto L71
L71:
	;
	v219 = int32(4)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210+v213<<(uint(v219)%32)-v219))))
	if v224 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L46
L73:
	;
	v227 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228+v213<<(uint(int32(2))%32))))
	if v232 == v227 {
		v158 = v227
		goto L45
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v241 = v213 + int32(1)
	if v241 <= v101 {
		v213 = v241
		goto L71
	} else {
		goto L78
	}
L76:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+4)))
	if v235&int32(2) != 0 {
		v158 = v227
		goto L45
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	goto L72
L79:
	;
	v253 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v253)
	goto L31
L80:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v258 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_binaryheap_replace_first(m, v260, v256)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v263 = F_binaryheap_remove_first(m, v260)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L85
	}
L84:
	;
	goto L31
L85:
	;
	goto L31
L86:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v276 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v338 = int32(0)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v339+v340<<(uint(int32(2))%32))))
	if v344 == v338 {
		v383 = v338
		goto L103
	} else {
		goto L104
	}
L89:
	;
	return int32(0)
L90:
	;
	goto L91
L91:
	;
	v285 = int32(0)
	goto L92
L92:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v292 = v289 + v285<<(uint(int32(4))%32)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+8))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v293 < v294 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	return int32(0)
L94:
	;
	v298 = v293
	goto L97
L95:
	;
	goto L96
L96:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v325 = v285 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v323+v325<<(uint(int32(2))%32))))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	m.T0[v331].(func(*base.Module, int32))(m, v329)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L101
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+8)) = v298 + int32(1)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306+v298<<(uint(int32(2))%32))))
	F_pfree(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L99
	}
L98:
	;
	goto L96
L99:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v292)+8))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v313 < v314 {
		v298 = v313
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v325 < v334 {
		v285 = v325
		goto L92
	} else {
		goto L102
	}
L102:
	;
	goto L93
L103:
	;
	return v383
L104:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+4)))
	if v347&int32(2) != 0 {
		v383 = v338
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v350 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	return v344
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v344
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+72))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v355)+16))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+8))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	m.T0[v359].(func(*base.Module, int32))(m, v357)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v362 = int32(_a_F_ExecGatherMerge_0)
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2]))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v356)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2])) = v365
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v355)+24))
	v371 = m.T0[v370].(func(*base.Module, int32, int32, int32) int32)(m, v355+int32(4), v356, int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2])) = v363
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357)+4)))
	v377 = v375 & int32(_a_F_ExecGatherMerge_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v357)+4)) = uint16(v377)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	*(*uint16)(unsafe.Add(mBase, uint32(v357)+6)) = uint16(v380)
	v383 = v357
	goto L103
}
func F_ExecInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 float64
	_ = v572
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v783 int64
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v932 float64
	_ = v932
	var v943 int32
	_ = v943
	v21 = m.G0
	v23 = v21 + int32(-64)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+132))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+200))
	if v30 == int32(0) {
		v53 = l1
		v54 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	m.T0[v57].(func(*base.Module, int32))(m, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L17
	}
L2:
	;
	v33 = F_ExecFindPartition(m, v27, l1, v30, l2, v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+204))
	if v37 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v38 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v45 = F_ExecGetRootToChildMap(m, v33, v26)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L14
	}
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v42 = l2
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v42
	goto L7
L11:
	;
	v41 = int32(0)
	goto L13
L12:
	;
	v41 = l2
	goto L13
L13:
	;
	v42 = v41
	goto L10
L14:
	;
	if v45 == int32(0) {
		v53 = v33
		v54 = l2
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v33)+204))
	v51 = F_execute_attr_map_slot(m, v49, l2, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v53 = v33
	v54 = v51
	goto L1
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+48))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+116)))
	if v62 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	if v70 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v65 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_ExecOpenIndices(m, v53, base.B2i32(v29 != int32(0)))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	m.G0 = v23 - int32(-64)
	return v943
L23:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v929 == int32(0) {
		v943 = v922
		goto L22
	} else {
		goto L293
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L3
	} else {
		goto L290
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L3
	} else {
		goto L287
	}
L26:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L3
	} else {
		goto L283
	}
L27:
	;
	if l3 != 0 {
		goto L239
	} else {
		goto L240
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v53)+84))
	if v95 != 0 {
		goto L45
	} else {
		goto L46
	}
L29:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)))
	if v73 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
	if v76 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v87 = v70
	goto L32
L32:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+10)))
	if v88 != int32(1) {
		goto L28
	} else {
		goto L42
	}
L33:
	;
	F_ExecPendingInserts(m, v26)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v79 = F_ExecBRInsertTriggers(m, v26, v53, v54)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	if v79 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v943 = int32(0)
	goto L22
L39:
	;
	goto L40
L40:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	if v84 == int32(0) {
		goto L28
	} else {
		goto L41
	}
L41:
	;
	v87 = v84
	goto L32
L42:
	;
	v91 = int32(0)
	v92 = F_ExecIRInsertTriggers(m, v26, v53, v54)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	if v92 != 0 {
		v765 = v54
		v770 = v91
		goto L27
	} else {
		goto L44
	}
L44:
	;
	v943 = v91
	goto L22
L45:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+36)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	if v100 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+36)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	if v224 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L48:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v53)+104))
	if int32(2) <= v109 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+17)))
	if v103 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	F_ExecComputeStoredGenerated(m, v53, v26, v54, int32(3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	if v109 == v112 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v53)+84))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+52))
	v212 = m.T0[v211].(func(*base.Module, int32, int32, int32, int32) int32)(m, v26, v53, v54, v25)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L78
	}
L55:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v53)+108))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v53)+112))
	F_ExecBatchInsert(m, v27, v53, v114, v115, v109, v26, l3)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v118 = int32(_a_F_ExecInsert_0)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v26)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v53)+108))
	if v123 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v53)+104))
	v129 = F_palloc(m, v126<<(uint(int32(2))%32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v53)+100))
	if v139 <= v138 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+108)) = v129
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v53)+104))
	v135 = F_palloc(m, v132<<(uint(int32(2))%32))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+112)) = v135
	goto L61
L64:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v142 = F_CreateTupleDescCopy(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L67
	}
L65:
	;
	v170 = v138
	goto L66
L66:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v53)+108))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+v170<<(uint(int32(2))%32))))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+32))
	m.T0[v178].(func(*base.Module, int32, int32))(m, v176, v54)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L3
	} else {
		goto L71
	}
L67:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v145 = F_CreateTupleDescCopy(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v148 = F_MakeTupleTableSlot(m, v142, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v53)+108))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v150+v151<<(uint(int32(2))%32)))) = v148
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v157 = F_MakeTupleTableSlot(m, v145, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v53)+112))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v159+v160<<(uint(int32(2))%32)))) = v157
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v53)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+100)) = v165 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	v170 = v169
	goto L66
L71:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v53)+112))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181+v182<<(uint(int32(2))%32))))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+32))
	m.T0[v188].(func(*base.Module, int32, int32))(m, v186, v25)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	if v191|base.B2i32(v112 == v109) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v203 = v191
	goto L75
L74:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
	v195 = F_lappend(m, v194, v53)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L3
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+96)) = v203 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v119
	v943 = int32(0)
	goto L22
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+188)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v26)+192))
	v199 = F_lappend(m, v198, v27)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+192)) = v199
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	v203 = v202
	goto L75
L78:
	;
	if v212 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v943 = int32(0)
	goto L22
L80:
	;
	goto L81
L81:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v212)+36)) = v218
	v765 = v212
	v770 = int32(0)
	goto L27
L82:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	switch v233 - int32(2) {
	case 0:
		v245 = v233
		goto L86
	default:
		goto L87
	case 3:
		goto L88
	}
L83:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+17)))
	if v227 != int32(1) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	F_ExecComputeStoredGenerated(m, v53, v26, v54, int32(3))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v246 != 0 {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	v245 = int32(1)
	goto L86
L88:
	;
	v236 = int32(2)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v27)+216))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	if v240 == v236 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v243 = v236
	goto L91
L90:
	;
	v243 = int32(1)
	goto L91
L91:
	;
	v245 = v243
	goto L86
L92:
	;
	F_ExecWithCheckOptions(m, v245, v53, v54, v26)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	if v250 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	F_ExecConstraints(m, v53, v54, v26)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L3
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v60)+48))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+131)))
	if v254 != int32(1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	if v29 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L101:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v53)+200))
	if v257 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	if v258 == int32(0) {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v266 = F_ExecPartitionCheck(m, v53, v54, v26, int32(1))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L3
	} else {
		goto L107
	}
L105:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+8)))
	if v261 != int32(1) {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	goto L100
L108:
	;
	v745 = int32(0)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v60)+188))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)+80))
	m.T0[v750].(func(*base.Module, int32, int32, int32, int32, int32))(m, v60, v54, v746, v745, v745)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L3
	} else {
		goto L236
	}
L109:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v271 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v274 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+30)) = uint16(v274)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+26)) = int32(-1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v53)+156))
	goto L111
L111:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[1]))
	if v302 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L3
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)) = uint8(v305)
	v308 = v21 + int32(-32)
	v311 = F_ExecCheckIndexConstraints(m, v53, v54, v26, v308, v21+int32(-38), v278)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L3
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	if v311 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if base.B2i32(v29 != int32(2)) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v668 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L3
	} else {
		goto L224
	}
L121:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v53)+160))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+16))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+64))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v324 = F_ExecUpdateLockMode(m, v323, v53)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L3
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v637 = int32(0)
	v638 = F_ExecGetReturningSlot(m, v26, v53)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L3
	} else {
		goto L214
	}
L124:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v326)+64))
	v329 = int32(0)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v322)+188))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+104))
	v335 = m.T0[v334].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v322, v308, v327, v321, v328, v324, v329, v329, v21+int32(-24))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L3
	} else {
		goto L132
	}
L125:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)+12))
	m.T0[v634].(func(*base.Module, int32))(m, v321)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L3
	} else {
		goto L213
	}
L126:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecCheckTupleVisible(m, v543, v322, v321)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L3
	} else {
		goto L194
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L3
	} else {
		goto L191
	}
L128:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[2]))
	if v511 < int32(2) {
		goto L125
	} else {
		goto L186
	}
L129:
	;
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[2]))
	if v491 < int32(2) {
		goto L125
	} else {
		goto L181
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L3
	} else {
		goto L178
	}
L131:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+20))
	v342 = m.T0[v341].(func(*base.Module, int32, int32, int32) int32)(m, v321, int32(-2), v21+int32(-25))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L133
	}
L132:
	;
	switch v335 {
	case 0:
		goto L126
	case 1:
		goto L131
	case 2:
		goto L130
	case 3:
		goto L129
	case 4:
		goto L128
	default:
		goto L127
	}
L133:
	;
	if base.Ui32(v342) < base.Ui32(int32(3)) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L3
	} else {
		goto L174
	}
L135:
	;
	v463 = int32(0)
	goto L134
L136:
	;
	goto L137
L137:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[3]))
	if v354 == v342 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v463 = int32(1)
	goto L134
L139:
	;
	goto L140
L140:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[4]))
	if v358 <= int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v463 = v455
	goto L134
L142:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[5]))
	if v362 == int32(0) {
		v455 = int32(0)
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[6]))
	v426 = int32(0)
	v428 = v358 - int32(1)
	goto L164
L145:
	;
	v367 = v362
	goto L146
L146:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v367)+20))
	if v372 == int32(4) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v455 = int32(0)
	goto L141
L148:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v367)+80))
	if v419 != 0 {
		v367 = v419
		goto L146
	} else {
		goto L163
	}
L149:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	if v375 == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v378 = int32(1)
	if v342 == v375 {
		v455 = v378
		goto L141
	} else {
		goto L151
	}
L151:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v367)+52))
	v382 = v380 - int32(1)
	if v382 < int32(0) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v387 = int32(0)
	v389 = v382
	goto L153
L153:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v367)+48))
	v395 = int32(2)
	v396 = base.I32_div_s(v389-v387, v395)
	v397 = v396 + v387
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v393+v397<<(uint(v395)%32))))
	if v401 == v342 {
		v455 = v378
		goto L141
	} else {
		goto L155
	}
L154:
	;
	goto L148
L155:
	;
	v405 = F_TransactionIdPrecedes(m, v401, v342)
	mBase = m.M
	if v405 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v406 = v397 + int32(1)
	goto L158
L157:
	;
	v406 = v387
	goto L158
L158:
	;
	if v405 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v409 = v389
	goto L161
L160:
	;
	v409 = v397 - int32(1)
	goto L161
L161:
	;
	if v406 <= v409 {
		v387 = v406
		v389 = v409
		goto L153
	} else {
		goto L162
	}
L162:
	;
	goto L154
L163:
	;
	goto L147
L164:
	;
	v433 = int32(2)
	v434 = base.I32_div_s(v428-v426, v433)
	v435 = v434 + v426
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v424+v435<<(uint(v433)%32))))
	v440 = base.B2i32(v439 == v342)
	if v439 == v342 {
		v455 = v440
		goto L141
	} else {
		goto L166
	}
L165:
	;
	v455 = v440
	goto L141
L166:
	;
	v443 = base.B2i32(base.Ui32(v439) < base.Ui32(v342))
	if base.Ui32(v439) < base.Ui32(v342) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v444 = v435 + int32(1)
	goto L169
L168:
	;
	v444 = v426
	goto L169
L169:
	;
	if base.Ui32(v439) < base.Ui32(v342) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v447 = v428
	goto L172
L171:
	;
	v447 = v435 - int32(1)
	goto L172
L172:
	;
	if v444 <= v447 {
		v426 = v444
		v428 = v447
		goto L164
	} else {
		goto L173
	}
L173:
	;
	goto L165
L174:
	;
	if v463 != 0 {
		goto L26
	} else {
		goto L175
	}
L175:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInsert_1), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L3
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2793), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L3
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInsert_4), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L3
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2803), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L3
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L3
	} else {
		goto L182
	}
L182:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	F_errmsg(m, int32(_a_F_ExecInsert_5), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L3
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2810), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L3
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L3
	} else {
		goto L187
	}
L187:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L3
	} else {
		goto L188
	}
L188:
	;
	F_errmsg(m, int32(_a_F_ExecInsert_6), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2826), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L3
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v335
	F_errmsg_internal(m, int32(_a_F_ExecInsert_7), v23)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2833), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L3
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	v546 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v320)+12)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v320)+8)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v320)+4)) = v321
	if v318 == v546 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v577 != 0 {
		goto L201
	} else {
		goto L202
	}
L196:
	;
	v552 = int32(_a_F_ExecInsert_0)
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v320)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v555
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v318)+20))
	v560 = m.T0[v559].(func(*base.Module, int32, int32, int32) int32)(m, v318, v320, v21+int32(-1))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L3
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v553
	if v560 != 0 {
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+12))
	m.T0[v565].(func(*base.Module, int32))(m, v321)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	v568 = int32(0)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	if v569 == v568 {
		v922 = v568
		goto L23
	} else {
		goto L200
	}
L200:
	;
	v572 = *(*float64)(unsafe.Add(mBase, uint32(v569)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v569)+240)) = base.F64_add(v572, float64(1))
	v922 = v568
	goto L23
L201:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	F_ExecWithCheckOptions(m, int32(3), v53, v321, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L3
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v53)+160))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)+12))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)+72))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v583)+16))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+8))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	m.T0[v587].(func(*base.Module, int32))(m, v585)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L3
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	v590 = int32(_a_F_ExecInsert_0)
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v584)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v593
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v583)+24))
	v599 = m.T0[v598].(func(*base.Module, int32, int32, int32) int32)(m, v583+int32(4), v584, int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L3
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v591
	v603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v585)+4)))
	v605 = v603 & int32(_a_F_ExecInsert_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v585)+4)) = uint16(v605)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v585)+12))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)))
	*(*uint16)(unsafe.Add(mBase, uint32(v585)+6)) = uint16(v608)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v53)+160))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v613)+8))
	v615 = F_ExecUpdate(m, l0, v53, v21+int32(-32), int32(0), v321, v614, l3)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L3
	} else {
		goto L208
	}
L207:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+12))
	m.T0[v630].(func(*base.Module, int32))(m, v321)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L3
	} else {
		goto L212
	}
L208:
	;
	if v615 == int32(0) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v53)+152))
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+8)))
	if v620&int32(2) == int32(0) {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v615)+8))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)+28))
	m.T0[v626].(func(*base.Module, int32))(m, v615)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L211
	}
L211:
	;
	goto L207
L212:
	;
	v922 = v615
	goto L23
L213:
	;
	goto L111
L214:
	;
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[2]))
	if v641 < int32(2) {
		v922 = v637
		goto L23
	} else {
		goto L215
	}
L215:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[7]))
	if v646 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecInsert[8])))
	if v648&int32(1) == int32(0) {
		goto L25
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v644)+188))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)+60))
	v658 = m.T0[v657].(func(*base.Module, int32, int32, int32, int32) int32)(m, v644, v21+int32(-32), int32(_a_F_ExecInsert_9), v638)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L3
	} else {
		goto L220
	}
L219:
	;
	goto L218
L220:
	;
	if v658 == int32(0) {
		goto L24
	} else {
		goto L221
	}
L221:
	;
	F_ExecCheckTupleVisible(m, v26, v644, v638)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L3
	} else {
		goto L222
	}
L222:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v638)+8))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v664)+12))
	m.T0[v665].(func(*base.Module, int32))(m, v638)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L3
	} else {
		goto L223
	}
L223:
	;
	v922 = v637
	goto L23
L224:
	;
	v670 = m.G0
	v672 = v670 - int32(16)
	m.G0 = v672
	v674 = int32(_a_F_ExecInsert_10)
	v675 = int32(1)
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[9]))
	v679 = v677 + v675
	if base.Ui32(v679) <= base.Ui32(v675) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v682 = v675
	goto L227
L226:
	;
	v682 = v679
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[9])) = v682
	*(*int64)(unsafe.Add(mBase, uint32(v672)+8)) = int64(74027918874902528)
	*(*int32)(unsafe.Add(mBase, uint32(v672))) = v668
	*(*int32)(unsafe.Add(mBase, uint32(v672)+4)) = v682
	v689 = int32(0)
	v691 = F_LockAcquire(m, v672, int32(7), v689, v689)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L3
	} else {
		goto L228
	}
L228:
	;
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[9]))
	m.G0 = v672 + int32(16)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	v699 = int32(0)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v60)+188))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+84))
	m.T0[v702].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v60, v54, v698, v699, v699, v694)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L3
	} else {
		goto L229
	}
L229:
	;
	v705 = int32(0)
	v710 = F_ExecInsertIndexTuples(m, v53, v54, v26, v705, int32(1), v21+int32(-39), v278, v705)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L3
	} else {
		goto L230
	}
L230:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v60)+188))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+88))
	m.T0[v718].(func(*base.Module, int32, int32, int32, int32))(m, v60, v54, v694, (v712^int32(-1))&int32(1))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L3
	} else {
		goto L231
	}
L231:
	;
	v721 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	v723 = m.G0
	v725 = v723 - int32(16)
	m.G0 = v725
	*(*int32)(unsafe.Add(mBase, uint32(v725))) = v721
	*(*int64)(unsafe.Add(mBase, uint32(v725)+8)) = int64(74027918874902528)
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v725)+4)) = v731
	v735 = F_LockRelease(m, v725, int32(7), int32(0))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L3
	} else {
		goto L233
	}
L233:
	;
	m.G0 = v725 + int32(16)
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)))
	if v740 != int32(1) {
		v765 = v54
		v770 = v710
		goto L27
	} else {
		goto L234
	}
L234:
	;
	F_list_free(m, v710)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L3
	} else {
		goto L235
	}
L235:
	;
	goto L111
L236:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v753 <= int32(0) {
		v765 = v54
		v770 = v745
		goto L27
	} else {
		goto L237
	}
L237:
	;
	v756 = int32(0)
	v761 = F_ExecInsertIndexTuples(m, v53, v54, v26, v756, v756, v756, v756, v756)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L3
	} else {
		goto L238
	}
L238:
	;
	v765 = v54
	v770 = v761
	goto L27
L239:
	;
	v783 = *(*int64)(unsafe.Add(mBase, uint32(v26)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+112)) = v783 + int64(1)
	goto L241
L240:
	;
	goto L241
L241:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v27)+204))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	if v788 != int32(2) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	F_ExecARInsertTriggers(m, v26, v53, v765, v770, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L3
	} else {
		goto L253
	}
L243:
	;
	v806 = v787
	goto L242
L244:
	;
	goto L245
L245:
	;
	if v787 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v806 = int32(0)
	goto L242
L247:
	;
	goto L248
L248:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787)+2)))
	if v794 != int32(1) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v806 = v787
	goto L242
L250:
	;
	goto L251
L251:
	;
	v797 = int32(0)
	F_ExecARUpdateTriggers(m, v26, v53, v797, v797, v797, v797, v765, v797, v787, v797)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L3
	} else {
		goto L252
	}
L252:
	;
	v806 = v797
	goto L242
L253:
	;
	F_list_free(m, v770)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L3
	} else {
		goto L254
	}
L254:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v53)+116))
	if v811 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	F_ExecWithCheckOptions(m, int32(0), v53, v765, v26)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L3
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v53)+152))
	if v815 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L257
L259:
	;
	if l4 != 0 {
		goto L279
	} else {
		goto L280
	}
L260:
	;
	v866 = int32(0)
	goto L259
L261:
	;
	goto L262
L262:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v819 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v843 = F_ExecProcessReturning(m, l0, v53, int32(3), v840, v765, v25)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L3
	} else {
		goto L271
	}
L264:
	;
	v840 = int32(0)
	goto L263
L265:
	;
	goto L266
L266:
	;
	v823 = F_ExecGetRootToChildMap(m, v53, v26)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L3
	} else {
		goto L267
	}
L267:
	;
	if v823 == int32(0) {
		v840 = v819
		goto L263
	} else {
		goto L268
	}
L268:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v823)+8))
	v828 = F_ExecGetReturningSlot(m, v26, v53)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L3
	} else {
		goto L269
	}
L269:
	;
	v830 = F_execute_attr_map_slot(m, v827, v819, v828)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L3
	} else {
		goto L270
	}
L270:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v832)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v830)+36)) = v833
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v836 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v835)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v830)+32)) = uint16(v836)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v835)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v830)+28)) = v838
	v840 = v830
	goto L263
L271:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v845 == int32(0) {
		v866 = v843
		goto L259
	} else {
		goto L272
	}
L272:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v843)+8))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v848)+28))
	m.T0[v849].(func(*base.Module, int32))(m, v843)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L3
	} else {
		goto L273
	}
L273:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v840)+8))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v852)+12))
	m.T0[v853].(func(*base.Module, int32))(m, v840)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L3
	} else {
		goto L274
	}
L274:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v840 != v856 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v856)+8))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)+12))
	m.T0[v859].(func(*base.Module, int32))(m, v856)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L3
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	v866 = v843
	goto L259
L278:
	;
	goto L277
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v765
	goto L281
L280:
	;
	goto L281
L281:
	;
	if l5 == int32(0) {
		v943 = v866
		goto L22
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v53
	v943 = v866
	goto L22
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(_a_F_ExecInsert_11)
	F_errmsg(m, int32(_a_F_ExecInsert_12), v21+int32(-48))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L3
	} else {
		goto L284
	}
L284:
	;
	F_errhint(m, int32(_a_F_ExecInsert_13), int32(0))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L3
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2790), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L3
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInsert_14), int32(0))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L3
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_15), int32(1264), int32(_a_F_ExecInsert_16))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L3
	} else {
		goto L289
	}
L289:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L290:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInsert_17), int32(0))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L3
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(409), int32(_a_F_ExecInsert_18))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L3
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L293:
	;
	v932 = *(*float64)(unsafe.Add(mBase, uint32(v929)+224))
	*(*float64)(unsafe.Add(mBase, uint32(v929)+224)) = base.F64_add(v932, float64(1))
	v943 = v922
	goto L22
}
func F_ExecNamedTuplestoreScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(739), int32(740))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ExecSampleScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(746), int32(747))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ExecScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	F_MemoryContextReset(m, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v19|v20 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v15 + int32(16)
	return v285
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[0]))
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	goto L44
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v18 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v113 = m.T0[l1].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L43
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+72))
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+64))
	v42 = F_bms_is_member(m, v40, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v53 = int32(1)
	v54 = v37 - v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v56 = v54 + v55
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v57 == v53 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	if v42 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v47 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v47 != 0 {
		v285 = v46
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	m.T0[v50].(func(*base.Module, int32))(m, v46)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v285 = v46
	goto L3
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	m.T0[v62].(func(*base.Module, int32))(m, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v66 = v54 << (uint(int32(2)) % 32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v67)))
	if v69 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v285 = v60
	goto L3
L25:
	;
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v70)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	if v72&int32(2) != 0 {
		v285 = v4
		goto L3
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83+v66)))
	if v85 == int32(0) {
		goto L11
	} else {
		goto L34
	}
L28:
	;
	v75 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v69)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v75 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	m.T0[v80].(func(*base.Module, int32))(m, v69)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v285 = v69
	goto L3
L33:
	;
	goto L32
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v89)
	v93 = F_EvalPlanQualFetchRowMark(m, v18, v37, v88)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if base.B2i32(v88 == int32(0))|base.B2i32(v93 == int32(0)) != 0 {
		v285 = v4
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
	if v98&int32(2) != 0 {
		v285 = v4
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v101 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v88)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v101 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	m.T0[v106].(func(*base.Module, int32))(m, v88)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v285 = v88
	goto L3
L42:
	;
	goto L41
L43:
	;
	v285 = v113
	goto L3
L44:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[0]))
	if v128 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v18 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v211
	if v20 != 0 {
		goto L87
	} else {
		goto L88
	}
L51:
	;
	if v19 == int32(0) {
		v285 = v219
		goto L3
	} else {
		goto L84
	}
L52:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
	if v214&int32(2) == int32(0) {
		goto L50
	} else {
		goto L83
	}
L53:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	m.T0[v207].(func(*base.Module, int32))(m, v204)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L82
	}
L54:
	;
	v201 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v158)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L80
	}
L55:
	;
	if v197 != 0 {
		v211 = v197
		goto L52
	} else {
		goto L79
	}
L56:
	;
	v194 = m.T0[l1].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L78
	}
L57:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+72))
	if v134 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)+64))
	v139 = F_bms_is_member(m, v137, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v146 = int32(1)
	v147 = v134 - v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v149 = v147 + v148
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v150 == v146 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	if v139 == int32(0) {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v144 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v144 != 0 {
		v197 = v143
		goto L55
	} else {
		goto L64
	}
L64:
	;
	v204 = v143
	goto L53
L65:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v204 = v153
	goto L53
L66:
	;
	goto L67
L67:
	;
	v155 = v147 << (uint(int32(2)) % 32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155+v156)))
	if v158 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v159)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
	if v161&int32(2) == int32(0) {
		goto L54
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167+v155)))
	if v169 == int32(0) {
		goto L56
	} else {
		goto L72
	}
L71:
	;
	v219 = int32(0)
	goto L51
L72:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v173)
	v175 = int32(0)
	v178 = F_EvalPlanQualFetchRowMark(m, v18, v134, v172)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if base.B2i32(v172 == v175)|base.B2i32(v178 == int32(0)) != 0 {
		v219 = v175
		goto L51
	} else {
		goto L74
	}
L74:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+4)))
	if v183&int32(2) != 0 {
		v219 = v175
		goto L51
	} else {
		goto L75
	}
L75:
	;
	v186 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v172)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v186 == int32(0) {
		v204 = v172
		goto L53
	} else {
		goto L77
	}
L77:
	;
	v211 = v172
	goto L52
L78:
	;
	v197 = v194
	goto L55
L79:
	;
	v219 = int32(0)
	goto L51
L80:
	;
	if v201 != 0 {
		v211 = v158
		goto L52
	} else {
		goto L81
	}
L81:
	;
	v204 = v158
	goto L53
L82:
	;
	v211 = v204
	goto L52
L83:
	;
	v219 = v211
	goto L51
L84:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	m.T0[v227].(func(*base.Module, int32))(m, v225)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v285 = v225
	goto L3
L86:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v274 != 0 {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v231 = int32(_a_F_ExecScan_0)
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v234
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v239 = m.T0[v238].(func(*base.Module, int32, int32, int32) int32)(m, v20, v21, v15+int32(15))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v19 == int32(0) {
		v285 = v211
		goto L3
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v232
	if v239 == int32(0) {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	m.T0[v251].(func(*base.Module, int32))(m, v249)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v254 = int32(_a_F_ExecScan_0)
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1]))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v248)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v257
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v263 = m.T0[v262].(func(*base.Module, int32, int32, int32) int32)(m, v19+int32(4), v248, int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v255
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249)+4)))
	v269 = v267 & int32(_a_F_ExecScan_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+4)) = uint16(v269)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+6)) = uint16(v272)
	v285 = v249
	goto L3
L95:
	;
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v274)+240)) = base.F64_add(v275, float64(1))
	goto L97
L96:
	;
	goto L97
L97:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	F_MemoryContextReset(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L44
}
func F_ExecShutdownNode_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v12 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = F_planstate_tree_walker_impl(m, l0, int32(634), l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)))
	if v15 != int32(1) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	F_InstrStartNode(m, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v23 - int32(418) {
	case 0:
		goto L15
	case 1:
		goto L14
	default:
		goto L10
	case 5:
		goto L11
	case 14:
		goto L16
	case 15:
		goto L13
	case 16:
		goto L12
	}
L10:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v112 == int32(0) {
		goto L1
	} else {
		goto L78
	}
L11:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v102 != 0 {
		goto L73
	} else {
		goto L74
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v61 != 0 {
		goto L51
	} else {
		goto L52
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v47 != 0 {
		goto L37
	} else {
		goto L38
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+44))
	if v44 != 0 {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+160))
	if v40 != 0 {
		goto L29
	} else {
		goto L30
	}
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v26 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ExecParallelFinish(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v29 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pfree(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v34 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	F_ExecParallelCleanup(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L10
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L27
L29:
	;
	m.T0[v40].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L10
L32:
	;
	goto L31
L33:
	;
	m.T0[v44].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L10
L36:
	;
	goto L35
L37:
	;
	F_ExecParallelFinish(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v50 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	F_pfree(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v55 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_ExecParallelCleanup(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L10
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	goto L47
L49:
	;
	goto L10
L50:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v70 == int32(0) {
		goto L49
	} else {
		goto L57
	}
L51:
	;
	if v60 != 0 {
		v69 = v60
		goto L50
	} else {
		goto L54
	}
L52:
	;
	v66 = v60
	goto L53
L53:
	;
	if v66 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	v63 = F_palloc0(m, int32(20))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v63
	v66 = v63
	goto L53
L56:
	;
	v69 = v66
	goto L50
L57:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v74 < v73 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v76 = v73
	goto L60
L59:
	;
	v76 = v74
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	if v79 < v78 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v81 = v78
	goto L63
L62:
	;
	v81 = v79
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
	if v84 < v83 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v86 = v83
	goto L66
L65:
	;
	v86 = v84
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v70)+52))
	if v89 < v88 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v91 = v88
	goto L69
L68:
	;
	v91 = v89
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70)+104))
	if base.Ui32(v94) < base.Ui32(v93) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v96 = v93
	goto L72
L71:
	;
	v96 = v94
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v96
	goto L49
L73:
	;
	F_ExecHashTableDetachBatch(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L10
L76:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_ExecHashTableDetach(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+4)))
	if v115 != int32(1) {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_InstrStopNode(m, v112, float64(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L80
	}
L80:
	;
	goto L1
}
func F_ExecStorePinnedBufferHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 == int32(_a_F_ExecStorePinnedBufferHeapTuple_0) {
		v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v8&int32(4) != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			F_pfree(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v17 = v14 & int32(-5)
				v18 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
				v24 = v17 & int32(_a_F_ExecStorePinnedBufferHeapTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v24)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v26
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v28)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
				if l2 != v30 {
					if v30 != 0 {
						F_ReleaseBuffer(m, v30)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l2
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l2
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
						return
					}
				} else {
					if l2 == int32(0) {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
						return
					} else {
						F_ReleaseBuffer(m, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
							return
						}
					}
				}
			}
		} else {
			v17 = v8
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
			v24 = v17 & int32(_a_F_ExecStorePinnedBufferHeapTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v26
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v28)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
			if l2 != v30 {
				if v30 != 0 {
					F_ReleaseBuffer(m, v30)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l2
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l2
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
					return
				}
			} else {
				if l2 == int32(0) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
					return
				} else {
					F_ReleaseBuffer(m, l2)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ExecStorePinnedBufferHeapTuple_2), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ExecStorePinnedBufferHeapTuple_3), int32(1620), int32(_a_F_ExecStorePinnedBufferHeapTuple_4))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_ExecUpdateEpilogue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v7 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v10 <= v7 {
		v27 = v7
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = int32(0)
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
		if v33 == int32(3) {
			v36 = int32(208)
		} else {
			v36 = int32(204)
		}
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v8+v36)))
		F_ExecARUpdateTriggers(m, v28, l2, v29, v29, l3, l4, l5, v27, v38, int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			F_list_free(m, v27)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
				if v44 != 0 {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_ExecWithCheckOptions(m, int32(0), l2, l5, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			}
		}
	} else {
		v13 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v14 == v13 {
			v27 = v13
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v29 = int32(0)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
			if v33 == int32(3) {
				v36 = int32(208)
			} else {
				v36 = int32(204)
			}
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v8+v36)))
			F_ExecARUpdateTriggers(m, v28, l2, v29, v29, l3, l4, l5, v27, v38, int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_list_free(m, v27)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
					if v44 != 0 {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_ExecWithCheckOptions(m, int32(0), l2, l5, v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = int32(0)
			v24 = F_ExecInsertIndexTuples(m, l2, l5, v17, int32(1), v19, v19, v19, base.B2i32(v14 == int32(2)))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v27 = v24
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v29 = int32(0)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+104))
				if v33 == int32(3) {
					v36 = int32(208)
				} else {
					v36 = int32(204)
				}
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v8+v36)))
				F_ExecARUpdateTriggers(m, v28, l2, v29, v29, l3, l4, l5, v27, v38, int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_list_free(m, v27)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
						if v44 != 0 {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_ExecWithCheckOptions(m, int32(0), l2, l5, v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								return
							}
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_ExtendSUBTRANS(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	if l0&int32(2047) != 0 {
		v8 = base.B2i32(l0 != int32(3))
	} else {
		v8 = int32(0)
	}
	if v8 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendSUBTRANS[0]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
		v15 = int32(base.Ui32(l0) >> (uint(int32(11)) % 32))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_ExtendSUBTRANS[1])))
		v18 = base.I32_rem_u_s(v15, v17)
		v21 = v13 + v18<<(uint(int32(7))%32)
		v23 = F_LWLockAcquire(m, v21, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v27 = F_SimpleLruZeroPage(m, int32(_a_F_ExtendSUBTRANS_0), base.I64_extend_i32_u(v15))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_LWLockRelease(m, v21)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		return
	}
}
func F___expo2(m *base.Module, l0 float64, l1 float64) float64 {
	var v3 float64
	_ = v3
	var v7 float64
	_ = v7
	v3 = float64(2.247116418577895e+307)
	v7 = F_exp(m, base.F64_add(l0, float64(-1416.0996898839683)))
	return base.F64_mul(base.F64_mul(base.F64_mul(l1, v3), v7), v3)
}
func F__equalAccessPriv(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v45
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = F_equal(m, v39, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v6 == int32(0) {
		v45 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 != v7 {
		v45 = v3
		goto L1
	} else {
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v45 = v3
	goto L1
L15:
	;
	goto L2
L16:
	;
	return int32(0)
L17:
	;
	v45 = v41
	goto L1
}
func F__equalCreateUserMappingStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = F_equal(m, v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v56
L2:
	;
	return int32(0)
L3:
	;
	if v8 == int32(0) {
		v56 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v47 != v48 {
		v56 = v3
		goto L1
	} else {
		goto L19
	}
L6:
	;
	if v14 == int32(0) {
		v56 = v3
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v14 != v15 {
		v56 = v3
		goto L1
	} else {
		goto L18
	}
L9:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if base.B2i32(v20 == int32(0))|base.B2i32(v20 != v23) != 0 {
		v41 = v20
		v42 = v23
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v41-v42 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v26 = v15
	v27 = v14
	goto L13
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+1)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v31 == int32(0) {
		v41 = v31
		v42 = v30
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v41 = v31
	v42 = v30
	goto L11
L15:
	;
	v34 = int32(1)
	if v31 == v30 {
		v26 = v26 + v34
		v27 = v27 + v34
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v56 = v3
	goto L1
L18:
	;
	goto L5
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v52 = F_equal(m, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v56 = v52
	goto L1
}
func F_each_worker(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	v2 = l1
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v17 = F_palloc0(m, int32(28))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v20 = F_palloc0(m, int32(40))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_InitMaterializedSRF(m, l0, int32(2))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(1347)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(1348)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(1349)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(1350)
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v17)+21)) = uint8(v39)
					*(*uint8)(unsafe.Add(mBase, uint32(v17)+20)) = uint8(v2)
					v44 = F_pg_detoast_datum_packed(m, v14)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v46 = int32(1)
						v47 = v44 + v46
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
						v52 = v50 & v46
						if v52 != 0 {
							v53 = v47
						} else {
							v53 = v44 + int32(4)
						}
						if v50 == int32(1) {
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
							if v59 == int32(18) {
								v62 = int32(16)
							} else {
								v62 = int32(0)
							}
							if base.Ui32((v59-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v69 = int32(4)
							} else {
								v69 = v62
							}
							v80 = v69
						} else {
							v70 = int32(1)
							if v52 != 0 {
								v80 = int32(base.Ui32(v50)>>(uint(v70)%32)) - v70
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
								v80 = int32(base.Ui32(v74)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker[0]))
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
						v85 = F_makeJsonLexContextCstringLen(m, v11+int32(12), v53, v80, v83, int32(1))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v17))) = v85
							v89 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker[1]))
							v94 = F_AllocSetContextCreateInternal(m, v89, int32(_a_F_each_worker_0), int32(0), int32(_a_F_each_worker_1), int32(_a_F_each_worker_2))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v94
								v98 = v11 + int32(12)
								v99 = F_pg_parse_json(m, v98, v20)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									if v99 != 0 {
										F_json_errsave_error(m, v99, v98, int32(0))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
											F_MemoryContextDelete(m, v104)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return
											} else {
												F_freeJsonLexContext(m, v11+int32(12))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													v111 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v111)
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
										F_MemoryContextDelete(m, v104)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return
										} else {
											F_freeJsonLexContext(m, v11+int32(12))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return
											} else {
												v111 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v111)
												m.G0 = v11 + int32(80)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ean13_in(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13871(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ec_member_matches_ctid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v6 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v7 == v6 {
		v26 = v6
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v10 != int32(6) {
			v26 = v6
		} else {
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+8)))
			if v13 != int32(_a_F_ec_member_matches_ctid_0) {
				v26 = v6
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				if v16 != int32(27) {
					v26 = v6
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
					if v19 != v20 {
						v26 = v6
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
						if v22 != 0 {
							v26 = v6
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
							v26 = base.B2i32(v23 == int32(0))
						}
					}
				}
			}
		}
	}
	return v26
}
func F_element_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_element_match[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_eq_v_b(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1-int32(4))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9-v10 < v8 {
		v81 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v81
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = v13 + v9 - v8
	if base.Ui32(int32(4)) <= base.Ui32(v8) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v77 != 0 {
		v81 = v3
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v77 = int32(0)
	goto L3
L5:
	;
	v51 = v46
	v52 = v47
	v53 = v48
	goto L15
L6:
	;
	if (v15|l1)&int32(3) != 0 {
		v46 = v15
		v47 = l1
		v48 = v8
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v39 = v15
	v40 = l1
	v41 = v8
	goto L8
L8:
	;
	if v41 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v23 = v15
	v24 = l1
	v25 = v8
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != v29 {
		v46 = v23
		v47 = v24
		v48 = v25
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v39 = v34
	v40 = v32
	v41 = v36
	goto L8
L12:
	;
	v31 = int32(4)
	v32 = v24 + v31
	v34 = v23 + v31
	v36 = v25 - v31
	if base.Ui32(int32(3)) < base.Ui32(v36) {
		v23 = v34
		v24 = v32
		v25 = v36
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v46 = v39
	v47 = v40
	v48 = v41
	goto L5
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 == v57 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v77 = v56 - v57
	goto L3
L17:
	;
	v59 = int32(1)
	v64 = v53 - v59
	if v64 != 0 {
		v51 = v51 + v59
		v52 = v52 + v59
		v53 = v64
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9 - v8
	v81 = int32(1)
	goto L1
}
func F_eqsel(m *base.Module, l0 int32) int32 {
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = F_eqsel_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_Float8GetDatum(m, v3)
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_erfc2(m *base.Module, l0 int32, l1 float64) float64 {
	var v12 float64
	_ = v12
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v142 float64
	_ = v142
	var v143 float64
	_ = v143
	var v148 float64
	_ = v148
	var v151 float64
	_ = v151
	var v160 float64
	_ = v160
	if base.Ui32(l0) <= base.Ui32(int32(1072955391)) {
		v12 = base.F64_add(base.F64_abs(l1), float64(-1))
		return base.F64_sub(float64(0.15493708848953247), base.F64_div(base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, float64(-0.002166375594868791)), float64(0.035478304325618236))), float64(-0.11089469428239668))), float64(0.31834661990116175))), float64(-0.3722078760357013))), float64(0.41485611868374833))), float64(-0.0023621185607526594)), base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, base.F64_add(base.F64_mul(v12, float64(0.011984499846799107)), float64(0.01363708391202905))), float64(0.12617121980876164))), float64(0.07182865441419627))), float64(0.540397917702171))), float64(0.10642088040084423))), float64(1))))
	} else {
		v55 = base.F64_abs(l1)
		v57 = base.F64_div(float64(1), base.F64_mul(v55, v55))
		if base.Ui32(l0) <= base.Ui32(int32(1074191212)) {
			v142 = base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, float64(-9.814329344169145)), float64(-81.2874355063066))), float64(-184.60509290671104))), float64(-162.39666946257347))), float64(-62.375332450326006))), float64(-10.558626225323291))), float64(-0.6938585727071818))), float64(-0.009864944034847148))
			v143 = base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, float64(-0.0604244152148581)), float64(6.570249770319282))), float64(108.63500554177944))), float64(429.00814002756783))), float64(645.3872717332679))), float64(434.56587747522923))), float64(137.65775414351904))), float64(19.651271667439257))
		} else {
			v142 = base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, float64(-483.5191916086514)), float64(-1025.0951316110772))), float64(-637.5664433683896))), float64(-160.63638485582192))), float64(-17.757954917754752))), float64(-0.799283237680523))), float64(-0.0098649429247001))
			v143 = base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, base.F64_add(base.F64_mul(v57, float64(-22.44095244658582)), float64(474.52854120695537))), float64(2553.0504064331644))), float64(3199.8582195085955))), float64(1536.729586084437))), float64(325.7925129965739))), float64(30.33806074348246))
		}
		v148 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v55) & int64(-4294967296))
		v151 = F_exp(m, base.F64_sub(float64(-0.5625), base.F64_mul(v148, v148)))
		v160 = F_exp(m, base.F64_add(base.F64_mul(base.F64_sub(v148, v55), base.F64_add(v55, v148)), base.F64_div(v142, base.F64_add(base.F64_mul(v57, v143), float64(1)))))
		return base.F64_div(base.F64_mul(v151, v160), v55)
	}
}
func F_errcode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_errcode[0]))
	if v4 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_errcode[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_errcode_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_errcode_1), int32(859), int32(_a_F_errcode_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4*int32(100))+uint32(_c_F_errcode[1]))) = l0
		return
	}
}
func F_errcontext_msg(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(_a_F_errcontext_msg_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[0])) = v14 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[1]))
	if int32(0) <= v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(_a_F_errcontext_msg_1)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[2]))
	v26 = v19 * int32(100)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errcontext_msg[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[2])) = v29
	v32 = v10 + int32(16)
	F_initStringInfo(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L26
	}
L4:
	;
	return
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errcontext_msg[4])))
	if v35 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_appendStringInfoString(m, v32, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errcontext_msg[5])))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[6])) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v47 = F_appendStringInfoVA(m, v10+int32(16), l0, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	F_appendStringInfoChar(m, v32, int32(10))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = v47
	goto L15
L13:
	;
	goto L14
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errcontext_msg[4])))
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v57 = v10 + int32(16)
	F_enlargeStringInfo(m, v57, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errcontext_msg[5])))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[6])) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
	v64 = F_appendStringInfoVA(m, v57, l0, l1)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v64 != 0 {
		v54 = v64
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	F_pfree(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v77 = F_pstrdup(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_errcontext_msg[4]))) = v77
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[2])) = v23
	v85 = int32(_a_F_errcontext_msg_0)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[0])) = v87 - int32(1)
	m.G0 = v10 + int32(32)
	return
L26:
	;
	F_errmsg_internal(m, int32(_a_F_errcontext_msg_2), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_errcontext_msg_3), int32(1393), int32(_a_F_errcontext_msg_4))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errfinish(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	v4 = int32(0)
	v6 = int32(_a_F_errfinish_0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[0])) = v8 + int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[1]))
	if v4 <= v13 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v258 = F_fflush(m, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L46
	} else {
		goto L79
	}
L2:
	;
	v240 = int32(_a_F_errfinish_0)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[0])) = v242 - int32(1)
	v247 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[2])) = v247
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[3])) = v247
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[4])) = v247
	F_pg_re_throw(m)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L46
	} else {
		goto L78
	}
L3:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L46
	} else {
		goto L75
	}
L6:
	;
	v19 = F_strlen(m, l0)
	mBase = m.M
	v26 = v19 + int32(1)
	goto L11
L7:
	;
	v69 = v4
	goto L8
L8:
	;
	v71 = v13 * int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_c_F_errfinish[5]))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_c_F_errfinish[6]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_c_F_errfinish[7]))) = v69
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_c_F_errfinish[8])))
	v78 = int32(_a_F_errfinish_1)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[9]))
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[9])) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_c_F_errfinish[11])))
	if v84|base.B2i32(l2 == int32(0)) != 0 {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	goto L9
L11:
	;
	v28 = int32(0)
	if v26 == v28 {
		v38 = v28
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v38 = v33
	goto L10
L13:
	;
	v32 = v26 - int32(1)
	v33 = l0 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34 != int32(47) {
		v26 = v32
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v41 = v38 + int32(1)
	goto L17
L16:
	;
	v41 = l0
	goto L17
L17:
	;
	v45 = F_strlen(m, v41)
	mBase = m.M
	v52 = v45 + int32(1)
	goto L20
L18:
	;
	if v64 != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	goto L18
L20:
	;
	v54 = int32(0)
	if v52 == v54 {
		v64 = v54
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v64 = v59
	goto L19
L22:
	;
	v58 = v52 - int32(1)
	v59 = v41 + v58
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v60 != int32(92) {
		v52 = v58
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v67 = v64 + int32(1)
	goto L26
L25:
	;
	v67 = v41
	goto L26
L26:
	;
	v69 = v67
	goto L8
L27:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[12]))
	if v158 != 0 {
		goto L49
	} else {
		goto L50
	}
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[13]))
	if v89 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[14]))
	if v93 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v96 == int32(0) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v99 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v102 = v93
	goto L33
L33:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if base.B2i32(v109 == int32(0))|base.B2i32(v109 != v112) != 0 {
		v130 = v109
		v131 = v112
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v138 = m.G0
	v140 = v138 - int32(16)
	m.G0 = v140
	F_initStringInfo(m, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L46
	} else {
		goto L47
	}
L35:
	;
	if v130-v131 != 0 {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	goto L35
L37:
	;
	v115 = l2
	v116 = v102
	goto L38
L38:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v120 == int32(0) {
		v130 = v120
		v131 = v119
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v130 = v120
	v131 = v119
	goto L36
L40:
	;
	v123 = int32(1)
	if v120 == v119 {
		v115 = v115 + v123
		v116 = v116 + v123
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v133 = F_strlen(m, v102)
	mBase = m.M
	v136 = v133 + v102 + int32(1)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v137 != 0 {
		v102 = v136
		goto L33
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	goto L34
L45:
	;
	goto L27
L46:
	;
	return
L47:
	;
	F_appendStringInfoString(m, v140, int32(_a_F_errfinish_2))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+uint32(_c_F_errfinish[11]))) = v147
	m.G0 = v140 + int32(16)
	goto L27
L49:
	;
	v159 = v158
	goto L52
L50:
	;
	goto L51
L51:
	;
	if v77 == int32(21) {
		goto L2
	} else {
		goto L56
	}
L52:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	m.T0[v165].(func(*base.Module, int32))(m, v164)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L46
	} else {
		goto L54
	}
L53:
	;
	goto L51
L54:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v168 != 0 {
		v159 = v168
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	F_EmitErrorReport(m)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L46
	} else {
		goto L57
	}
L57:
	;
	F_FreeErrorDataContents(m, v71+int32(_a_F_errfinish_3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L46
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[9])) = v79
	v182 = int32(_a_F_errfinish_4)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[1]))
	v185 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[1])) = v184 - v185
	v188 = int32(_a_F_errfinish_0)
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[0])) = v190 - v185
	if v77 == int32(22) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[15]))
	if v197 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	if int32(23) <= v77 {
		goto L1
	} else {
		goto L70
	}
L62:
	;
	v206 = F_fflush(m, int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L46
	} else {
		goto L65
	}
L63:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[16]))
	if v199 != int32(2) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[16])) = int32(0)
	goto L62
L65:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[17]))
	if v209 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[17])) = int32(3)
	goto L68
L67:
	;
	goto L68
L68:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L46
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[18]))
	if v221 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L46
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return
L74:
	;
	goto L73
L75:
	;
	F_errmsg_internal(m, int32(_a_F_errfinish_5), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L46
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_errfinish_6), int32(482), int32(_a_F_errfinish_7))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L46
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errhint(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn13875(m, l0, l1, int32(_a_F_errhint_0), int32(1324))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_errstart_cold(m *base.Module, l0 int32, l1 int32) {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = F_errstart(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_errtableconstraint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
	v6 = F_get_namespace_name(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_err_generic_string(m, int32(115), v6)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			F_err_generic_string(m, int32(116), v11+int32(4))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_err_generic_string(m, int32(110), l1)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_estonian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v735 int32
	_ = v735
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1257 int32
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1387 int32
	_ = v1387
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1493 int32
	_ = v1493
	var v1509 int32
	_ = v1509
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1677 int32
	_ = v1677
	var v1693 int32
	_ = v1693
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1752 int32
	_ = v1752
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1831 int32
	_ = v1831
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1843 int32
	_ = v1843
	var v1859 int32
	_ = v1859
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1911 int32
	_ = v1911
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2006 int32
	_ = v2006
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2018 int32
	_ = v2018
	var v2034 int32
	_ = v2034
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2097 int32
	_ = v2097
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	v10 = F_find_among(m, l0, int32(_a_F_estonian_UTF_8_stem_0), int32(290))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v2097
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v134
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v160 = v150
	goto L67
L3:
	;
	return int32(0)
L4:
	;
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v134 = v16
	goto L2
L6:
	;
	goto L7
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 < v19 {
		v134 = v19
		goto L2
	} else {
		goto L8
	}
L8:
	;
	switch v10 - int32(1) {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	case 5:
		goto L22
	case 6:
		goto L21
	case 7:
		goto L20
	case 8:
		goto L19
	case 9:
		goto L18
	case 10:
		goto L17
	case 11:
		goto L16
	case 12:
		goto L15
	case 13:
		goto L14
	case 14:
		goto L13
	case 15:
		goto L12
	case 16:
		goto L11
	case 17:
		goto L10
	default:
		goto L9
	}
L9:
	;
	return int32(0)
L10:
	;
	v127 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L62
	}
L11:
	;
	v121 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L60
	}
L12:
	;
	v115 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L58
	}
L13:
	;
	v109 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L56
	}
L14:
	;
	v103 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_5))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L54
	}
L15:
	;
	v97 = F_slice_from_s(m, l0, int32(7), int32(_a_F_estonian_UTF_8_stem_6))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L52
	}
L16:
	;
	v91 = F_slice_from_s(m, l0, int32(7), int32(_a_F_estonian_UTF_8_stem_7))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L50
	}
L17:
	;
	v85 = F_slice_from_s(m, l0, int32(6), int32(_a_F_estonian_UTF_8_stem_8))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L48
	}
L18:
	;
	v79 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_9))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L46
	}
L19:
	;
	v73 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_10))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L3
	} else {
		goto L44
	}
L20:
	;
	v67 = F_slice_from_s(m, l0, int32(6), int32(_a_F_estonian_UTF_8_stem_11))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L42
	}
L21:
	;
	v61 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L40
	}
L22:
	;
	v55 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_13))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L38
	}
L23:
	;
	v49 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_14))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L36
	}
L24:
	;
	v43 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_15))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L34
	}
L25:
	;
	v37 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L32
	}
L26:
	;
	v31 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_17))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L30
	}
L27:
	;
	v25 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_18))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	if int32(0) <= v25 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v2097 = v25
	goto L1
L30:
	;
	if int32(0) <= v31 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	v2097 = v31
	goto L1
L32:
	;
	if int32(0) <= v37 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	v2097 = v37
	goto L1
L34:
	;
	if int32(0) <= v43 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v2097 = v43
	goto L1
L36:
	;
	if int32(0) <= v49 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v2097 = v49
	goto L1
L38:
	;
	if int32(0) <= v55 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v2097 = v55
	goto L1
L40:
	;
	if int32(0) <= v61 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v2097 = v61
	goto L1
L42:
	;
	if int32(0) <= v67 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v2097 = v67
	goto L1
L44:
	;
	if int32(0) <= v73 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	v2097 = v73
	goto L1
L46:
	;
	if int32(0) <= v79 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	v2097 = v79
	goto L1
L48:
	;
	if int32(0) <= v85 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	v2097 = v85
	goto L1
L50:
	;
	if int32(0) <= v91 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	v2097 = v91
	goto L1
L52:
	;
	if int32(0) <= v97 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	v2097 = v97
	goto L1
L54:
	;
	if int32(0) <= v103 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v2097 = v103
	goto L1
L56:
	;
	if int32(0) <= v109 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	v2097 = v109
	goto L1
L58:
	;
	if int32(0) <= v115 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v2097 = v115
	goto L1
L60:
	;
	if int32(0) <= v121 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v2097 = v121
	goto L1
L62:
	;
	if v127 < int32(0) {
		v2097 = v127
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L9
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	if v386 < v388 {
		goto L118
	} else {
		goto L119
	}
L65:
	;
	if v255 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L66:
	;
	v255 = v227
	goto L65
L67:
	;
	if v151 <= v160 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v255 = int32(-1)
	goto L65
L70:
	;
	goto L71
L71:
	;
	v167 = int32(1)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v152))))
	if base.Ui32(v169) < base.Ui32(int32(192)) {
		v226 = v169
		v227 = v167
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if int32(252) < v226 {
		goto L85
	} else {
		goto L86
	}
L73:
	;
	v173 = v160 + int32(1)
	if v173 == v151 {
		v226 = v169
		v227 = v167
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v152))))
	v178 = v176 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v169) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v152))))
	v194 = v192 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v169) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v182 = v160 + int32(2)
	if v182 != v151 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v226 = v169<<(uint(int32(6))%32)&int32(1984) | v178
	v227 = int32(2)
	goto L72
L79:
	;
	goto L78
L80:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152+v198))))
	v226 = v211&int32(63) | (v169<<(uint(int32(18))%32)&int32(_a_F_estonian_UTF_8_stem_19) | v178<<(uint(int32(12))%32) | v194<<(uint(int32(6))%32))
	v227 = int32(4)
	goto L72
L81:
	;
	v198 = v160 + int32(3)
	if v198 != v151 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v226 = v169<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v178<<(uint(int32(6))%32) | v194
	v227 = int32(3)
	goto L72
L84:
	;
	goto L83
L85:
	;
	v244 = v227 + v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v244
	v160 = v244
	goto L67
L86:
	;
	v231 = v226 - int32(97)
	if v231 < int32(0) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v231)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v237)>>(uint(v231&int32(7))%32))&int32(1) != 0 {
		goto L66
	} else {
		goto L88
	}
L88:
	;
	goto L85
L90:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v383 = v258
	goto L64
L91:
	;
	goto L92
L92:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v280 = v270
	goto L95
L93:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v376 < int32(0) {
		v383 = v377
		goto L64
	} else {
		goto L117
	}
L94:
	;
	v376 = v347
	goto L93
L95:
	;
	if v271 <= v280 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v376 = int32(-1)
	goto L93
L98:
	;
	goto L99
L99:
	;
	v287 = int32(1)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v272))))
	if base.Ui32(v289) < base.Ui32(int32(192)) {
		v346 = v289
		v347 = v287
		goto L100
	} else {
		goto L101
	}
L100:
	;
	if int32(252) < v346 {
		goto L94
	} else {
		goto L113
	}
L101:
	;
	v293 = v280 + int32(1)
	if v293 == v271 {
		v346 = v289
		v347 = v287
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293+v272))))
	v298 = v296 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v289) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302+v272))))
	v314 = v312 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v289) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v302 = v280 + int32(2)
	if v302 != v271 {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v346 = v289<<(uint(int32(6))%32)&int32(1984) | v298
	v347 = int32(2)
	goto L100
L107:
	;
	goto L106
L108:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+v318))))
	v346 = v331&int32(63) | (v289<<(uint(int32(18))%32)&int32(_a_F_estonian_UTF_8_stem_19) | v298<<(uint(int32(12))%32) | v314<<(uint(int32(6))%32))
	v347 = int32(4)
	goto L100
L109:
	;
	v318 = v280 + int32(3)
	if v318 != v271 {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v346 = v289<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v298<<(uint(int32(6))%32) | v314
	v347 = int32(3)
	goto L100
L112:
	;
	goto L111
L113:
	;
	v351 = v346 - int32(97)
	if v351 < int32(0) {
		goto L94
	} else {
		goto L114
	}
L114:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v351)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v357)>>(uint(v351&int32(7))%32))&int32(1) == int32(0) {
		goto L94
	} else {
		goto L115
	}
L115:
	;
	v365 = v347 + v280
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365
	v280 = v365
	goto L95
L117:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = v380 + v376
	v383 = v377
	goto L64
L118:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v753
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	if v753 < v756 {
		goto L202
	} else {
		goto L203
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v388
	v393 = v386 - int32(1)
	if v393 <= v388 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	goto L118
L121:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v393))))
	if v397 != int32(105) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v402 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_21), int32(2))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	if v402 == int32(0) {
		goto L120
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v407
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L127
L125:
	;
	if v462 < int32(0) {
		goto L118
	} else {
		goto L144
	}
L127:
	;
	goto L128
L128:
	;
	goto L129
L129:
	;
	v417 = v407
	v419 = int32(4)
	goto L132
L131:
	;
	v462 = v444
	goto L125
L132:
	;
	if v417 <= v6 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L131
L134:
	;
	v462 = int32(-1)
	goto L125
L135:
	;
	goto L136
L136:
	;
	v424 = v417 - int32(1)
	v426 = int32(*(*int8)(unsafe.Add(mBase, uint32(v410+v424))))
	if base.B2i32(int32(0) <= v426)|base.B2i32(v424 <= v6) != 0 {
		v444 = v424
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v448 = int32(1)
	if v448 < v419 {
		v417 = v444
		v419 = v419 - v448
		goto L132
	} else {
		goto L143
	}
L138:
	;
	v432 = v424
	goto L139
L139:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v432))))
	if base.Ui32(int32(191)) < base.Ui32(v437) {
		v444 = v432
		goto L137
	} else {
		goto L141
	}
L140:
	;
	v444 = v6
	goto L137
L141:
	;
	v441 = v432 - int32(1)
	if v6 < v441 {
		v432 = v441
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	goto L133
L144:
	;
	v465 = v407 - v409
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v465 + v466
	switch v402 - int32(1) {
	case 0:
		goto L146
	case 1:
		goto L145
	default:
		goto L118
	}
L145:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L177
L146:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L149
L147:
	;
	if v599 != 0 {
		goto L118
	} else {
		goto L170
	}
L148:
	;
	v599 = v592
	goto L147
L149:
	;
	if v483 <= v484 {
		v592 = int32(-1)
		goto L148
	} else {
		goto L151
	}
L150:
	;
	v592 = int32(0)
	goto L148
L151:
	;
	v501 = int32(1)
	v502 = v483 - v501
	v504 = int32(*(*int8)(unsafe.Add(mBase, uint32(v485+v502))))
	v506 = v504 & int32(255)
	if base.B2i32(v502 == v484)|base.B2i32(int32(0) <= v504) != 0 {
		v564 = v506
		v568 = v501
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if int32(252) < v564 {
		goto L160
	} else {
		goto L161
	}
L153:
	;
	v513 = v506 & int32(63)
	v515 = v483 - int32(2)
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v515))))
	v519 = v517 << (uint(int32(6)) % 32)
	if base.B2i32(v515 != v484)&base.B2i32(base.Ui32(v517) < base.Ui32(int32(192))) == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v564 = v519&int32(1984) | v513
	v568 = int32(2)
	goto L152
L155:
	;
	goto L156
L156:
	;
	v532 = v519&int32(4032) | v513
	v534 = v483 - int32(3)
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v534))))
	if base.B2i32(v534 != v484)&base.B2i32(base.Ui32(v536) < base.Ui32(int32(224))) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v564 = v536<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v532
	v568 = int32(3)
	goto L152
L158:
	;
	goto L159
L159:
	;
	v554 = int32(4)
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483+v485-v554))))
	v564 = v536<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v556&int32(7)<<(uint(int32(18))%32) | v532
	v568 = v554
	goto L152
L160:
	;
	v599 = v568
	goto L147
L161:
	;
	goto L162
L162:
	;
	v570 = v564 - int32(97)
	if v570 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v599 = v568
	goto L147
L164:
	;
	goto L165
L165:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v570)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[1]))))
	if int32(base.Ui32(v576)>>(uint(v570&int32(7))%32))&int32(1) == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v599 = v568
	goto L147
L167:
	;
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v483 - v568
	goto L169
L169:
	;
	goto L150
L170:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v600 + v465
	v605 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L3
	} else {
		goto L171
	}
L171:
	;
	if v605 != 0 {
		goto L118
	} else {
		goto L172
	}
L172:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v607 + v465
	v610 = F_slice_del(m, l0)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	if int32(0) <= v610 {
		goto L118
	} else {
		goto L174
	}
L174:
	;
	v2097 = v610
	goto L1
L175:
	;
	if v742 != 0 {
		goto L118
	} else {
		goto L198
	}
L176:
	;
	v742 = v735
	goto L175
L177:
	;
	if v626 <= v627 {
		v735 = int32(-1)
		goto L176
	} else {
		goto L179
	}
L178:
	;
	v735 = int32(0)
	goto L176
L179:
	;
	v644 = int32(1)
	v645 = v626 - v644
	v647 = int32(*(*int8)(unsafe.Add(mBase, uint32(v628+v645))))
	v649 = v647 & int32(255)
	if base.B2i32(v645 == v627)|base.B2i32(int32(0) <= v647) != 0 {
		v707 = v649
		v711 = v644
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if int32(382) < v707 {
		goto L188
	} else {
		goto L189
	}
L181:
	;
	v656 = v649 & int32(63)
	v658 = v626 - int32(2)
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v658))))
	v662 = v660 << (uint(int32(6)) % 32)
	if base.B2i32(v658 != v627)&base.B2i32(base.Ui32(v660) < base.Ui32(int32(192))) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v707 = v662&int32(1984) | v656
	v711 = int32(2)
	goto L180
L183:
	;
	goto L184
L184:
	;
	v675 = v662&int32(4032) | v656
	v677 = v626 - int32(3)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628+v677))))
	if base.B2i32(v677 != v627)&base.B2i32(base.Ui32(v679) < base.Ui32(int32(224))) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v707 = v679<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v675
	v711 = int32(3)
	goto L180
L186:
	;
	goto L187
L187:
	;
	v697 = int32(4)
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626+v628-v697))))
	v707 = v679<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v699&int32(7)<<(uint(int32(18))%32) | v675
	v711 = v697
	goto L180
L188:
	;
	v742 = v711
	goto L175
L189:
	;
	goto L190
L190:
	;
	v713 = v707 - int32(98)
	if v713 < int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v742 = v711
	goto L175
L192:
	;
	goto L193
L193:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v713)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[2]))))
	if int32(base.Ui32(v719)>>(uint(v713&int32(7))%32))&int32(1) == int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v742 = v711
	goto L175
L195:
	;
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v626 - v711
	goto L197
L197:
	;
	goto L178
L198:
	;
	v743 = F_slice_del(m, l0)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	if int32(0) <= v743 {
		goto L118
	} else {
		goto L200
	}
L200:
	;
	v2097 = v743
	goto L1
L201:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1911
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L477
L202:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v936
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v938)))
	if v936 < v939 {
		goto L242
	} else {
		goto L243
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v753
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v756
	if v753 <= v756 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v759
	goto L202
L205:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v764 = int32(1)
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762+v753-v764))))
	if base.B2i32(v766&int32(224) != int32(96))|base.B2i32(v764<<(uint(v766)%32)&int32(_a_F_estonian_UTF_8_stem_24) == int32(0)) != 0 {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v780 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_25), int32(21))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L3
	} else {
		goto L207
	}
L207:
	;
	if v780 == int32(0) {
		goto L204
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v759
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v785
	switch v780 - int32(1) {
	case 0:
		goto L211
	case 1:
		goto L210
	case 2:
		goto L209
	default:
		goto L201
	}
L209:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L218
L210:
	;
	v795 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_26))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L3
	} else {
		goto L214
	}
L211:
	;
	v789 = F_slice_del(m, l0)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L3
	} else {
		goto L212
	}
L212:
	;
	if int32(0) <= v789 {
		goto L201
	} else {
		goto L213
	}
L213:
	;
	v2097 = v789
	goto L1
L214:
	;
	if int32(0) <= v795 {
		goto L201
	} else {
		goto L215
	}
L215:
	;
	v2097 = v795
	goto L1
L216:
	;
	if v927 != 0 {
		goto L202
	} else {
		goto L239
	}
L217:
	;
	v927 = v920
	goto L216
L218:
	;
	if v811 <= v812 {
		v920 = int32(-1)
		goto L217
	} else {
		goto L220
	}
L219:
	;
	v920 = int32(0)
	goto L217
L220:
	;
	v829 = int32(1)
	v830 = v811 - v829
	v832 = int32(*(*int8)(unsafe.Add(mBase, uint32(v813+v830))))
	v834 = v832 & int32(255)
	if base.B2i32(v830 == v812)|base.B2i32(int32(0) <= v832) != 0 {
		v892 = v834
		v896 = v829
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if int32(252) < v892 {
		goto L229
	} else {
		goto L230
	}
L222:
	;
	v841 = v834 & int32(63)
	v843 = v811 - int32(2)
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813+v843))))
	v847 = v845 << (uint(int32(6)) % 32)
	if base.B2i32(v843 != v812)&base.B2i32(base.Ui32(v845) < base.Ui32(int32(192))) == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v892 = v847&int32(1984) | v841
	v896 = int32(2)
	goto L221
L224:
	;
	goto L225
L225:
	;
	v860 = v847&int32(4032) | v841
	v862 = v811 - int32(3)
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813+v862))))
	if base.B2i32(v862 != v812)&base.B2i32(base.Ui32(v864) < base.Ui32(int32(224))) == int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v892 = v864<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v860
	v896 = int32(3)
	goto L221
L227:
	;
	goto L228
L228:
	;
	v882 = int32(4)
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811+v813-v882))))
	v892 = v864<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v884&int32(7)<<(uint(int32(18))%32) | v860
	v896 = v882
	goto L221
L229:
	;
	v927 = v896
	goto L216
L230:
	;
	goto L231
L231:
	;
	v898 = v892 - int32(97)
	if v898 < int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v927 = v896
	goto L216
L233:
	;
	goto L234
L234:
	;
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v898)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v904)>>(uint(v898&int32(7))%32))&int32(1) == int32(0) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v927 = v896
	goto L216
L236:
	;
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v811 - v896
	goto L238
L238:
	;
	goto L219
L239:
	;
	v928 = F_slice_del(m, l0)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L3
	} else {
		goto L240
	}
L240:
	;
	if int32(0) <= v928 {
		goto L201
	} else {
		goto L241
	}
L241:
	;
	v2097 = v928
	goto L1
L242:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v996
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v998)))
	if v996 < v999 {
		goto L258
	} else {
		goto L259
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v936
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v939
	if v936-int32(3) <= v939 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v942
	goto L242
L245:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v949 = int32(1)
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947+v936-v949))))
	if base.B2i32(v951&int32(224) != int32(96))|base.B2i32(v949<<(uint(v951)%32)&int32(_a_F_estonian_UTF_8_stem_27) == int32(0)) != 0 {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v965 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_28), int32(12))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L3
	} else {
		goto L247
	}
L247:
	;
	if v965 == int32(0) {
		goto L244
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v942
	v970 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v970
	switch v965 - int32(1) {
	case 0:
		goto L251
	case 1:
		goto L250
	case 2:
		goto L249
	default:
		goto L242
	}
L249:
	;
	v988 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_29))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L3
	} else {
		goto L256
	}
L250:
	;
	v982 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_30))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L3
	} else {
		goto L254
	}
L251:
	;
	v976 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_31))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L3
	} else {
		goto L252
	}
L252:
	;
	if int32(0) <= v976 {
		goto L242
	} else {
		goto L253
	}
L253:
	;
	v2097 = v976
	goto L1
L254:
	;
	if int32(0) <= v982 {
		goto L242
	} else {
		goto L255
	}
L255:
	;
	v2097 = v982
	goto L1
L256:
	;
	if int32(0) <= v988 {
		goto L242
	} else {
		goto L257
	}
L257:
	;
	v2097 = v988
	goto L1
L258:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1244
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1246)))
	if v1244 < v1247 {
		goto L316
	} else {
		goto L317
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v996
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v999
	if v996 <= v999 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1002
	goto L258
L261:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1007 = int32(1)
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005+v996-v1007))))
	if base.B2i32(v1009&int32(224) != int32(96))|base.B2i32(v1007<<(uint(v1009)%32)&int32(_a_F_estonian_UTF_8_stem_32) == int32(0)) != 0 {
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v1023 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_33), int32(10))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L3
	} else {
		goto L263
	}
L263:
	;
	if v1023 == int32(0) {
		goto L260
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1002
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1028
	switch v1023 - int32(1) {
	case 0:
		goto L267
	case 1:
		goto L266
	default:
		goto L265
	}
L265:
	;
	v1234 = F_slice_del(m, l0)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L3
	} else {
		goto L314
	}
L266:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L296
L267:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L270
L268:
	;
	if v1161 == int32(0) {
		goto L265
	} else {
		goto L291
	}
L269:
	;
	v1161 = v1154
	goto L268
L270:
	;
	if v1045 <= v1046 {
		v1154 = int32(-1)
		goto L269
	} else {
		goto L272
	}
L271:
	;
	v1154 = int32(0)
	goto L269
L272:
	;
	v1063 = int32(1)
	v1064 = v1045 - v1063
	v1066 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1047+v1064))))
	v1068 = v1066 & int32(255)
	if base.B2i32(v1064 == v1046)|base.B2i32(int32(0) <= v1066) != 0 {
		v1126 = v1068
		v1130 = v1063
		goto L273
	} else {
		goto L274
	}
L273:
	;
	if int32(117) < v1126 {
		goto L281
	} else {
		goto L282
	}
L274:
	;
	v1075 = v1068 & int32(63)
	v1077 = v1045 - int32(2)
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047+v1077))))
	v1081 = v1079 << (uint(int32(6)) % 32)
	if base.B2i32(v1077 != v1046)&base.B2i32(base.Ui32(v1079) < base.Ui32(int32(192))) == int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1126 = v1081&int32(1984) | v1075
	v1130 = int32(2)
	goto L273
L276:
	;
	goto L277
L277:
	;
	v1094 = v1081&int32(4032) | v1075
	v1096 = v1045 - int32(3)
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047+v1096))))
	if base.B2i32(v1096 != v1046)&base.B2i32(base.Ui32(v1098) < base.Ui32(int32(224))) == int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1126 = v1098<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1094
	v1130 = int32(3)
	goto L273
L279:
	;
	goto L280
L280:
	;
	v1116 = int32(4)
	v1118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045+v1047-v1116))))
	v1126 = v1098<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1118&int32(7)<<(uint(int32(18))%32) | v1094
	v1130 = v1116
	goto L273
L281:
	;
	v1161 = v1130
	goto L268
L282:
	;
	goto L283
L283:
	;
	v1132 = v1126 - int32(97)
	if v1132 < int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1161 = v1130
	goto L268
L285:
	;
	goto L286
L286:
	;
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1132)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1138)>>(uint(v1132&int32(7))%32))&int32(1) == int32(0) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1161 = v1130
	goto L268
L288:
	;
	goto L289
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1045 - v1130
	goto L290
L290:
	;
	goto L271
L291:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1164 + (v1028 - v1032)
	v1170 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L3
	} else {
		goto L292
	}
L292:
	;
	if v1170 != 0 {
		goto L265
	} else {
		goto L293
	}
L293:
	;
	goto L258
L294:
	;
	if v1225 < int32(0) {
		goto L258
	} else {
		goto L313
	}
L296:
	;
	goto L297
L297:
	;
	goto L298
L298:
	;
	v1180 = v1028
	v1182 = int32(4)
	goto L301
L300:
	;
	v1225 = v1207
	goto L294
L301:
	;
	if v1180 <= v1002 {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	goto L300
L303:
	;
	v1225 = int32(-1)
	goto L294
L304:
	;
	goto L305
L305:
	;
	v1187 = v1180 - int32(1)
	v1189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1173+v1187))))
	if base.B2i32(int32(0) <= v1189)|base.B2i32(v1187 <= v1002) != 0 {
		v1207 = v1187
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1211 = int32(1)
	if v1211 < v1182 {
		v1180 = v1207
		v1182 = v1182 - v1211
		goto L301
	} else {
		goto L312
	}
L307:
	;
	v1195 = v1187
	goto L308
L308:
	;
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1173+v1195))))
	if base.Ui32(int32(191)) < base.Ui32(v1200) {
		v1207 = v1195
		goto L306
	} else {
		goto L310
	}
L309:
	;
	v1207 = v1002
	goto L306
L310:
	;
	v1204 = v1195 - int32(1)
	if v1002 < v1204 {
		v1195 = v1204
		goto L308
	} else {
		goto L311
	}
L311:
	;
	goto L309
L312:
	;
	goto L302
L313:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1228 + (v1028 - v1172)
	goto L265
L314:
	;
	if int32(0) <= v1234 {
		goto L258
	} else {
		goto L315
	}
L315:
	;
	v2097 = v1234
	goto L1
L316:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1536
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1538)))
	if v1536 < v1539 {
		goto L396
	} else {
		goto L397
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1244
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1247
	if v1244 <= v1247 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1250
	goto L316
L319:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253+v1244-int32(1)))))
	if v1257&int32(254) != int32(100) {
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1264 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_34), int32(7))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L3
	} else {
		goto L321
	}
L321:
	;
	if v1264 == int32(0) {
		goto L318
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1250
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1269
	switch v1264 - int32(1) {
	case 0:
		goto L326
	case 1:
		goto L325
	case 2:
		goto L324
	case 3:
		goto L323
	default:
		goto L316
	}
L323:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L368
L324:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L335
L325:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1282 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L3
	} else {
		goto L329
	}
L326:
	;
	v1275 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_35))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L3
	} else {
		goto L327
	}
L327:
	;
	if int32(0) <= v1275 {
		goto L316
	} else {
		goto L328
	}
L328:
	;
	v2097 = v1275
	goto L1
L329:
	;
	if v1282 != 0 {
		goto L316
	} else {
		goto L330
	}
L330:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1284 + (v1269 - v1279)
	v1288 = F_slice_del(m, l0)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L3
	} else {
		goto L331
	}
L331:
	;
	if int32(0) <= v1288 {
		goto L316
	} else {
		goto L332
	}
L332:
	;
	v2097 = v1288
	goto L1
L333:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1348 = v1346 + (v1269 - v1292)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1348
	if int32(0) <= v1345 {
		goto L352
	} else {
		goto L353
	}
L335:
	;
	goto L336
L336:
	;
	goto L337
L337:
	;
	v1300 = v1269
	v1302 = int32(4)
	goto L340
L339:
	;
	v1345 = v1327
	goto L333
L340:
	;
	if v1300 <= v1250 {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	goto L339
L342:
	;
	v1345 = int32(-1)
	goto L333
L343:
	;
	goto L344
L344:
	;
	v1307 = v1300 - int32(1)
	v1309 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1293+v1307))))
	if base.B2i32(int32(0) <= v1309)|base.B2i32(v1307 <= v1250) != 0 {
		v1327 = v1307
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1331 = int32(1)
	if v1331 < v1302 {
		v1300 = v1327
		v1302 = v1302 - v1331
		goto L340
	} else {
		goto L351
	}
L346:
	;
	v1315 = v1307
	goto L347
L347:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1293+v1315))))
	if base.Ui32(int32(191)) < base.Ui32(v1320) {
		v1327 = v1315
		goto L345
	} else {
		goto L349
	}
L348:
	;
	v1327 = v1250
	goto L345
L349:
	;
	v1324 = v1315 - int32(1)
	if v1250 < v1324 {
		v1315 = v1324
		goto L347
	} else {
		goto L350
	}
L350:
	;
	goto L348
L351:
	;
	goto L341
L352:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1348 <= v1352 {
		goto L355
	} else {
		goto L356
	}
L353:
	;
	goto L354
L354:
	;
	v1383 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_36))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L3
	} else {
		goto L364
	}
L355:
	;
	v1377 = F_slice_del(m, l0)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L3
	} else {
		goto L362
	}
L356:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1356 = int32(1)
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1354+v1348-v1356))))
	if base.Ui32(v1356) < base.Ui32((v1358-int32(115))&int32(255)) {
		goto L355
	} else {
		goto L357
	}
L357:
	;
	v1367 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_37), int32(5))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L3
	} else {
		goto L359
	}
L358:
	;
	v1373 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_38))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L3
	} else {
		goto L360
	}
L359:
	;
	switch v1367 - int32(1) {
	case 0:
		goto L358
	case 1:
		goto L355
	default:
		goto L316
	}
L360:
	;
	if int32(0) <= v1373 {
		goto L316
	} else {
		goto L361
	}
L361:
	;
	v2097 = v1373
	goto L1
L362:
	;
	if int32(0) <= v1377 {
		goto L316
	} else {
		goto L363
	}
L363:
	;
	v2097 = v1377
	goto L1
L364:
	;
	if int32(0) <= v1383 {
		goto L316
	} else {
		goto L365
	}
L365:
	;
	v2097 = v1383
	goto L1
L366:
	;
	if v1516 != 0 {
		goto L389
	} else {
		goto L390
	}
L367:
	;
	v1516 = v1509
	goto L366
L368:
	;
	if v1400 <= v1401 {
		v1509 = int32(-1)
		goto L367
	} else {
		goto L370
	}
L369:
	;
	v1509 = int32(0)
	goto L367
L370:
	;
	v1418 = int32(1)
	v1419 = v1400 - v1418
	v1421 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1402+v1419))))
	v1423 = v1421 & int32(255)
	if base.B2i32(v1419 == v1401)|base.B2i32(int32(0) <= v1421) != 0 {
		v1481 = v1423
		v1485 = v1418
		goto L371
	} else {
		goto L372
	}
L371:
	;
	if int32(117) < v1481 {
		goto L379
	} else {
		goto L380
	}
L372:
	;
	v1430 = v1423 & int32(63)
	v1432 = v1400 - int32(2)
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1402+v1432))))
	v1436 = v1434 << (uint(int32(6)) % 32)
	if base.B2i32(v1432 != v1401)&base.B2i32(base.Ui32(v1434) < base.Ui32(int32(192))) == int32(0) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1481 = v1436&int32(1984) | v1430
	v1485 = int32(2)
	goto L371
L374:
	;
	goto L375
L375:
	;
	v1449 = v1436&int32(4032) | v1430
	v1451 = v1400 - int32(3)
	v1453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1402+v1451))))
	if base.B2i32(v1451 != v1401)&base.B2i32(base.Ui32(v1453) < base.Ui32(int32(224))) == int32(0) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1481 = v1453<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1449
	v1485 = int32(3)
	goto L371
L377:
	;
	goto L378
L378:
	;
	v1471 = int32(4)
	v1473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400+v1402-v1471))))
	v1481 = v1453<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1473&int32(7)<<(uint(int32(18))%32) | v1449
	v1485 = v1471
	goto L371
L379:
	;
	v1516 = v1485
	goto L366
L380:
	;
	goto L381
L381:
	;
	v1487 = v1481 - int32(97)
	if v1487 < int32(0) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1516 = v1485
	goto L366
L383:
	;
	goto L384
L384:
	;
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1487)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1493)>>(uint(v1487&int32(7))%32))&int32(1) == int32(0) {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1516 = v1485
	goto L366
L386:
	;
	goto L387
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1400 - v1485
	goto L388
L388:
	;
	goto L369
L389:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1517 + (v1269 - v1387)
	v1523 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9))
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L3
	} else {
		goto L392
	}
L390:
	;
	goto L391
L391:
	;
	v1527 = F_slice_del(m, l0)
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L3
	} else {
		goto L394
	}
L392:
	;
	if v1523 == int32(0) {
		goto L316
	} else {
		goto L393
	}
L393:
	;
	goto L391
L394:
	;
	if int32(0) <= v1527 {
		goto L316
	} else {
		goto L395
	}
L395:
	;
	v2097 = v1527
	goto L1
L396:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1713
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1715)))
	if v1713 < v1716 {
		goto L433
	} else {
		goto L434
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1536
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1539
	if v1536 <= v1539 {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1542
	goto L396
L399:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1547 = int32(1)
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1545+v1536-v1547))))
	if base.B2i32(v1549&int32(224) != int32(96))|base.B2i32(v1547<<(uint(v1549)%32)&int32(_a_F_estonian_UTF_8_stem_39) == int32(0)) != 0 {
		goto L398
	} else {
		goto L400
	}
L400:
	;
	v1563 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_40), int32(3))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L3
	} else {
		goto L401
	}
L401:
	;
	if v1563 == int32(0) {
		goto L398
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1542
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1568
	switch v1563 - int32(1) {
	case 0:
		goto L404
	case 1:
		goto L403
	default:
		goto L396
	}
L403:
	;
	v1705 = F_slice_del(m, l0)
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L3
	} else {
		goto L431
	}
L404:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L407
L405:
	;
	if v1700 != 0 {
		goto L396
	} else {
		goto L428
	}
L406:
	;
	v1700 = v1693
	goto L405
L407:
	;
	if v1584 <= v1585 {
		v1693 = int32(-1)
		goto L406
	} else {
		goto L409
	}
L408:
	;
	v1693 = int32(0)
	goto L406
L409:
	;
	v1602 = int32(1)
	v1603 = v1584 - v1602
	v1605 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1586+v1603))))
	v1607 = v1605 & int32(255)
	if base.B2i32(v1603 == v1585)|base.B2i32(int32(0) <= v1605) != 0 {
		v1665 = v1607
		v1669 = v1602
		goto L410
	} else {
		goto L411
	}
L410:
	;
	if int32(117) < v1665 {
		goto L418
	} else {
		goto L419
	}
L411:
	;
	v1614 = v1607 & int32(63)
	v1616 = v1584 - int32(2)
	v1618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586+v1616))))
	v1620 = v1618 << (uint(int32(6)) % 32)
	if base.B2i32(v1616 != v1585)&base.B2i32(base.Ui32(v1618) < base.Ui32(int32(192))) == int32(0) {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1665 = v1620&int32(1984) | v1614
	v1669 = int32(2)
	goto L410
L413:
	;
	goto L414
L414:
	;
	v1633 = v1620&int32(4032) | v1614
	v1635 = v1584 - int32(3)
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586+v1635))))
	if base.B2i32(v1635 != v1585)&base.B2i32(base.Ui32(v1637) < base.Ui32(int32(224))) == int32(0) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1665 = v1637<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1633
	v1669 = int32(3)
	goto L410
L416:
	;
	goto L417
L417:
	;
	v1655 = int32(4)
	v1657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1584+v1586-v1655))))
	v1665 = v1637<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1657&int32(7)<<(uint(int32(18))%32) | v1633
	v1669 = v1655
	goto L410
L418:
	;
	v1700 = v1669
	goto L405
L419:
	;
	goto L420
L420:
	;
	v1671 = v1665 - int32(97)
	if v1671 < int32(0) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1700 = v1669
	goto L405
L422:
	;
	goto L423
L423:
	;
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1671)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1677)>>(uint(v1671&int32(7))%32))&int32(1) == int32(0) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1700 = v1669
	goto L405
L425:
	;
	goto L426
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1584 - v1669
	goto L427
L427:
	;
	goto L408
L428:
	;
	v1701 = F_slice_del(m, l0)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L3
	} else {
		goto L429
	}
L429:
	;
	if int32(0) <= v1701 {
		goto L396
	} else {
		goto L430
	}
L430:
	;
	v2097 = v1701
	goto L1
L431:
	;
	if int32(0) <= v1705 {
		goto L396
	} else {
		goto L432
	}
L432:
	;
	v2097 = v1705
	goto L1
L433:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1874
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1876)))
	if v1874 < v1877 {
		goto L201
	} else {
		goto L466
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1713
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1716
	if v1713 <= v1716 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1719
	goto L433
L436:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1722+v1713-int32(1)))))
	if v1726 != int32(105) {
		goto L435
	} else {
		goto L437
	}
L437:
	;
	v1731 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_41), int32(1))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L3
	} else {
		goto L438
	}
L438:
	;
	if v1731 == int32(0) {
		goto L435
	} else {
		goto L439
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1719
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1736
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L442
L440:
	;
	if v1866 != 0 {
		goto L433
	} else {
		goto L463
	}
L441:
	;
	v1866 = v1859
	goto L440
L442:
	;
	if v1736 <= v1719 {
		v1859 = int32(-1)
		goto L441
	} else {
		goto L444
	}
L443:
	;
	v1859 = int32(0)
	goto L441
L444:
	;
	v1768 = int32(1)
	v1769 = v1736 - v1768
	v1771 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1752+v1769))))
	v1773 = v1771 & int32(255)
	if base.B2i32(v1769 == v1719)|base.B2i32(int32(0) <= v1771) != 0 {
		v1831 = v1773
		v1835 = v1768
		goto L445
	} else {
		goto L446
	}
L445:
	;
	if int32(117) < v1831 {
		goto L453
	} else {
		goto L454
	}
L446:
	;
	v1780 = v1773 & int32(63)
	v1782 = v1736 - int32(2)
	v1784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+v1782))))
	v1786 = v1784 << (uint(int32(6)) % 32)
	if base.B2i32(v1782 != v1719)&base.B2i32(base.Ui32(v1784) < base.Ui32(int32(192))) == int32(0) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v1831 = v1786&int32(1984) | v1780
	v1835 = int32(2)
	goto L445
L448:
	;
	goto L449
L449:
	;
	v1799 = v1786&int32(4032) | v1780
	v1801 = v1736 - int32(3)
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1752+v1801))))
	if base.B2i32(v1801 != v1719)&base.B2i32(base.Ui32(v1803) < base.Ui32(int32(224))) == int32(0) {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v1831 = v1803<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1799
	v1835 = int32(3)
	goto L445
L451:
	;
	goto L452
L452:
	;
	v1821 = int32(4)
	v1823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1736+v1752-v1821))))
	v1831 = v1803<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1823&int32(7)<<(uint(int32(18))%32) | v1799
	v1835 = v1821
	goto L445
L453:
	;
	v1866 = v1835
	goto L440
L454:
	;
	goto L455
L455:
	;
	v1837 = v1831 - int32(97)
	if v1837 < int32(0) {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v1866 = v1835
	goto L440
L457:
	;
	goto L458
L458:
	;
	v1843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1837)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1843)>>(uint(v1837&int32(7))%32))&int32(1) == int32(0) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v1866 = v1835
	goto L440
L460:
	;
	goto L461
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1736 - v1835
	goto L462
L462:
	;
	goto L443
L463:
	;
	v1867 = F_slice_del(m, l0)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L3
	} else {
		goto L464
	}
L464:
	;
	if int32(0) <= v1867 {
		goto L433
	} else {
		goto L465
	}
L465:
	;
	v2097 = v1867
	goto L1
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1874
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1877
	v1883 = v1874 - int32(1)
	if v1883 <= v1877 {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1880
	goto L201
L468:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1885+v1883))))
	if base.B2i32(v1887 != int32(117))&base.B2i32(v1887 != int32(97)) != 0 {
		goto L467
	} else {
		goto L469
	}
L469:
	;
	v1895 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_42), int32(4))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L3
	} else {
		goto L470
	}
L470:
	;
	if v1895 == int32(0) {
		goto L467
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1880
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1900
	v1902 = F_slice_del(m, l0)
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L3
	} else {
		goto L472
	}
L472:
	;
	if int32(0) <= v1902 {
		goto L201
	} else {
		goto L473
	}
L473:
	;
	v2097 = v1902
	goto L1
L474:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2094
	v2097 = int32(1)
	goto L1
L475:
	;
	if v2041 != 0 {
		goto L474
	} else {
		goto L498
	}
L476:
	;
	v2041 = v2034
	goto L475
L477:
	;
	if v1911 <= v1926 {
		v2034 = int32(-1)
		goto L476
	} else {
		goto L479
	}
L478:
	;
	v2034 = int32(0)
	goto L476
L479:
	;
	v1943 = int32(1)
	v1944 = v1911 - v1943
	v1946 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1927+v1944))))
	v1948 = v1946 & int32(255)
	if base.B2i32(v1944 == v1926)|base.B2i32(int32(0) <= v1946) != 0 {
		v2006 = v1948
		v2010 = v1943
		goto L480
	} else {
		goto L481
	}
L480:
	;
	if int32(252) < v2006 {
		goto L488
	} else {
		goto L489
	}
L481:
	;
	v1955 = v1948 & int32(63)
	v1957 = v1911 - int32(2)
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1927+v1957))))
	v1961 = v1959 << (uint(int32(6)) % 32)
	if base.B2i32(v1957 != v1926)&base.B2i32(base.Ui32(v1959) < base.Ui32(int32(192))) == int32(0) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v2006 = v1961&int32(1984) | v1955
	v2010 = int32(2)
	goto L480
L483:
	;
	goto L484
L484:
	;
	v1974 = v1961&int32(4032) | v1955
	v1976 = v1911 - int32(3)
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1927+v1976))))
	if base.B2i32(v1976 != v1926)&base.B2i32(base.Ui32(v1978) < base.Ui32(int32(224))) == int32(0) {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	v2006 = v1978<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1974
	v2010 = int32(3)
	goto L480
L486:
	;
	goto L487
L487:
	;
	v1996 = int32(4)
	v1998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1911+v1927-v1996))))
	v2006 = v1978<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1998&int32(7)<<(uint(int32(18))%32) | v1974
	v2010 = v1996
	goto L480
L488:
	;
	v2041 = v2010
	goto L475
L489:
	;
	goto L490
L490:
	;
	v2012 = v2006 - int32(97)
	if v2012 < int32(0) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	v2041 = v2010
	goto L475
L492:
	;
	goto L493
L493:
	;
	v2018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2012)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2018)>>(uint(v2012&int32(7))%32))&int32(1) == int32(0) {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v2041 = v2010
	goto L475
L495:
	;
	goto L496
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1911 - v2010
	goto L497
L497:
	;
	goto L478
L498:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v2043)))
	if v2042 < v2044 {
		goto L474
	} else {
		goto L499
	}
L499:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2042
	v2048 = v2042 - int32(1)
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2048 <= v2049 {
		goto L474
	} else {
		goto L500
	}
L500:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2051+v2048))))
	if base.B2i32(v2053&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2053)%32)&int32(_a_F_estonian_UTF_8_stem_43) == int32(0)) != 0 {
		goto L474
	} else {
		goto L501
	}
L501:
	;
	v2067 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_44), int32(3))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L3
	} else {
		goto L502
	}
L502:
	;
	if v2067 == int32(0) {
		goto L474
	} else {
		goto L503
	}
L503:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2071
	switch v2067 - int32(1) {
	case 0:
		goto L506
	case 1:
		goto L505
	case 2:
		goto L504
	default:
		goto L474
	}
L504:
	;
	v2089 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_45))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L3
	} else {
		goto L511
	}
L505:
	;
	v2083 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_46))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L3
	} else {
		goto L509
	}
L506:
	;
	v2077 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_47))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L3
	} else {
		goto L507
	}
L507:
	;
	if int32(0) <= v2077 {
		goto L474
	} else {
		goto L508
	}
L508:
	;
	v2097 = v2077
	goto L1
L509:
	;
	if int32(0) <= v2083 {
		goto L474
	} else {
		goto L510
	}
L510:
	;
	v2097 = v2083
	goto L1
L511:
	;
	if v2089 < int32(0) {
		v2097 = v2089
		goto L1
	} else {
		goto L512
	}
L512:
	;
	goto L474
}
func F_examine_expression(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_palloc0(m, int32(248))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
		v17 = F_exprType(m, l0)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v17
			v20 = F_exprTypmod(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v20
				v23 = F_exprCollation(m, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v23
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					v29 = F_SearchSysCacheCopy(m, int32(82), v27, int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 != 0 {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
							v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
							v33 = v31 + v32
							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v33
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_examine_expression[0]))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v36
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+184)) = v40
							v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+76)))
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+204)) = uint16(v42)
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)))
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+214)) = uint8(v44)
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+128)))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+188)) = v40
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+219)) = uint8(v46)
							v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+76)))
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+206)) = uint16(v49)
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)))
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+215)) = uint8(v51)
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+128)))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v40
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+220)) = uint8(v53)
							v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+76)))
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+208)) = uint16(v56)
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)))
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+216)) = uint8(v58)
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+128)))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = v40
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+221)) = uint8(v60)
							v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+76)))
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+210)) = uint16(v63)
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)))
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+217)) = uint8(v65)
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+128)))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+200)) = v40
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+222)) = uint8(v67)
							v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+76)))
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+212)) = uint16(v70)
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)))
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+218)) = uint8(v72)
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+128)))
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+223)) = uint8(v74)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
							if v76 != 0 {
								v78 = F_OidFunctionCall1Coll(m, v76, int32(0), v12)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									if v78 != 0 {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
										if v84 == int32(0) {
											F_pfree(m, v29)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v12)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													v95 = int32(0)
													m.G0 = v9 + int32(16)
													return v95
												}
											}
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if int32(0) < v87 {
												v95 = v12
												m.G0 = v9 + int32(16)
												return v95
											} else {
												F_pfree(m, v29)
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v12)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														v95 = int32(0)
														m.G0 = v9 + int32(16)
														return v95
													}
												}
											}
										}
									} else {
										F_pfree(m, v29)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v12)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return int32(0)
											} else {
												v95 = int32(0)
												m.G0 = v9 + int32(16)
												return v95
											}
										}
									}
								}
							} else {
								v80 = F_std_typanalyze(m, v12)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									if v80 == int32(0) {
										F_pfree(m, v29)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v12)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return int32(0)
											} else {
												v95 = int32(0)
												m.G0 = v9 + int32(16)
												return v95
											}
										}
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
										if v84 == int32(0) {
											F_pfree(m, v29)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v12)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													v95 = int32(0)
													m.G0 = v9 + int32(16)
													return v95
												}
											}
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if int32(0) < v87 {
												v95 = v12
												m.G0 = v9 + int32(16)
												return v95
											} else {
												F_pfree(m, v29)
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v12)
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														v95 = int32(0)
														m.G0 = v9 + int32(16)
														return v95
													}
												}
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v104
								F_errmsg_internal(m, int32(_a_F_examine_expression_0), v9)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_examine_expression_1), int32(642), int32(_a_F_examine_expression_2))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_examine_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v610 int32
	_ = v610
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v958 int32
	_ = v958
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1194 int32
	_ = v1194
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v18
	v26 = F_exprType(m, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v26
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	if v30 == int32(0) {
		v71 = l1
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v84 = v71
	goto L24
L4:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v71 = int32(0)
	goto L3
L6:
	;
	goto L7
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v36 != int32(319) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = F_expression_tree_walker_impl(m, l1, int32(1488), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v49 = l1
	goto L14
L11:
	;
	if v41 == int32(0) {
		v71 = l1
		goto L3
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v65 = F_expression_tree_mutator_impl(m, v49, int32(1489), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v58 != int32(319) {
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v71 = int32(0)
	goto L3
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v61 != 0 {
		v49 = v61
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v71 = v65
	goto L3
L19:
	;
	m.G0 = v16 + int32(80)
	return
L20:
	;
	F_bms_free(m, v577)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L211
	}
L21:
	;
	v711 = l1
	v712 = v5
	goto L20
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L208
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L205
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v93 != int32(27) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L202
	}
L26:
	;
	if v93 != int32(6) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v84 = v658
	goto L24
L28:
	;
	goto L25
L29:
	;
	goto L28
L30:
	;
	v574 = F_pull_varnos(m, l0, v84)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L168
	}
L31:
	;
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if l2 != v98 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v84
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v102 = F_find_base_rel(m, l0, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v107
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v84)+8)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v102)+108))
	if v111 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v172)
	v184 = l0
	v188 = v84
	goto L54
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v112 <= int32(0) {
		v172 = int32(0)
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v172 = int32(0)
	goto L37
L41:
	;
	v115 = int32(0)
	if v115 < v112 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v118 = v112
	goto L44
L43:
	;
	v118 = v115
	goto L44
L44:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v122 = int32(0)
	goto L45
L45:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v119+v122<<(uint(int32(2))%32))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+101)))
	if v138 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L40
L47:
	;
	v154 = v122 + int32(1)
	if v154 != v118 {
		v122 = v154
		goto L45
	} else {
		goto L53
	}
L48:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+40))
	if v141 != int32(1) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if v145 != v109 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v147 = int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v137)+88))
	if v148 == int32(0) {
		v172 = v147
		goto L37
	} else {
		goto L51
	}
L51:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+100)))
	if v151 != 0 {
		v172 = v147
		goto L37
	} else {
		goto L52
	}
L52:
	;
	goto L47
L53:
	;
	goto L46
L54:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v184)+36))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197+v198<<(uint(int32(2))%32))))
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_examine_variable[0]))
	if v204 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L19
L56:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	switch v229 {
	case 0:
		goto L68
	case 1:
		goto L67
	default:
		goto L19
	case 6:
		goto L66
	}
L57:
	;
	v207 = int32(*(*int16)(unsafe.Add(mBase, uint32(v188)+8)))
	v208 = m.T0[v204].(func(*base.Module, int32, int32, int32, int32) int32)(m, v184, v202, v207, l3)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v208 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v212 == int32(0) {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v215 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_0), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_2), int32(_a_F_examine_variable_3))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	if v422 == int32(0) {
		goto L19
	} else {
		goto L109
	}
L66:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+92)))
	if v261 != 0 {
		goto L19
	} else {
		goto L78
	}
L67:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+20)))
	if v250 != 0 {
		goto L19
	} else {
		goto L75
	}
L68:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v202)+16))
	v232 = int32(*(*int16)(unsafe.Add(mBase, uint32(v188)+8)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+20)))
	v234 = F_SearchSysCache3(m, int32(65), v231, v232, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1490)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v234
	if v234 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v240 = int32(*(*int16)(unsafe.Add(mBase, uint32(v188)+8)))
	v243 = F_bms_make_singleton(m, v240+int32(7))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v248 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v248)
	goto L19
L73:
	;
	v245 = F_all_rows_selectable(m, v184, v239, v243)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v245)
	goto L19
L75:
	;
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+8)))
	if v251 == int32(0) {
		goto L19
	} else {
		goto L76
	}
L76:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v257 = F_find_base_rel(m, v184, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v416 = v188 + int32(8)
	v421 = v257 + int32(140)
	goto L65
L78:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+8)))
	if v262 == int32(0) {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v202)+88))
	v269 = v184
	v272 = v267
	goto L81
L80:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+48))
	if v303 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L81:
	;
	if v272 == int32(0) {
		goto L80
	} else {
		goto L83
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v269)+16))
	if v285 != 0 {
		v269 = v285
		v272 = v272 - int32(1)
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v202)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v290
	F_errmsg_internal(m, int32(_a_F_examine_variable_4), v16-int32(-64))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_5), int32(_a_F_examine_variable_3))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v269)+76))
	if v390 == int32(0) {
		goto L29
	} else {
		goto L106
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L103
	}
L90:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v306 <= int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v202)+84))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v303)+12))
	v316 = int32(0)
	goto L92
L92:
	;
	v326 = v316 << (uint(int32(2)) % 32)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v310+v326)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if base.B2i32(v332 == int32(0))|base.B2i32(v332 != v335) != 0 {
		v353 = v332
		v354 = v335
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L89
L94:
	;
	if v353-v354 == int32(0) {
		goto L88
	} else {
		goto L101
	}
L95:
	;
	goto L94
L96:
	;
	v338 = v329
	v339 = v309
	goto L97
L97:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+1)))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+1)))
	if v343 == int32(0) {
		v353 = v343
		v354 = v342
		goto L95
	} else {
		goto L99
	}
L98:
	;
	v353 = v343
	v354 = v342
	goto L95
L99:
	;
	v346 = int32(1)
	if v343 == v342 {
		v338 = v338 + v346
		v339 = v339 + v346
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v359 = v316 + int32(1)
	if v306 != v359 {
		v316 = v359
		goto L92
	} else {
		goto L102
	}
L102:
	;
	goto L93
L103:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v202)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v378
	F_errmsg_internal(m, int32(_a_F_examine_variable_6), v16+int32(16))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_7), int32(_a_F_examine_variable_3))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	if v393 <= v316 {
		goto L29
	} else {
		goto L107
	}
L107:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v395+v326)))
	if v397 <= int32(0) {
		goto L23
	} else {
		goto L108
	}
L108:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v184)+8))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+16))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v416 = v188 + int32(8)
	v421 = v402 + v397<<(uint(int32(2))%32) - int32(4)
	goto L65
L109:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)+144))
	if v426 != 0 {
		goto L19
	} else {
		goto L110
	}
L110:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v425)+108))
	if v427 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v425)+96))
	if v428 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v430 = v428
	goto L114
L113:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425)+76))
	v430 = v429
	goto L114
L114:
	;
	v431 = int32(*(*int16)(unsafe.Add(mBase, uint32(v416))))
	if v430 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	if v469 == int32(0) {
		goto L22
	} else {
		goto L128
	}
L116:
	;
	goto L115
L117:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	if v435 <= int32(0) {
		v469 = int32(0)
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v469 = int32(0)
	goto L116
L120:
	;
	v438 = int32(0)
	if v438 < v435 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v441 = v435
	goto L123
L122:
	;
	v441 = v438
	goto L123
L123:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v430)+12))
	v446 = int32(0)
	goto L124
L124:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v442+v446<<(uint(int32(2))%32))))
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v454)+8)))
	if v455 == v431&int32(_a_F_examine_variable_8) {
		v469 = v454
		goto L116
	} else {
		goto L126
	}
L125:
	;
	goto L119
L126:
	;
	v458 = v446 + int32(1)
	if v458 != v441 {
		v446 = v458
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+26)))
	if v473 == int32(1) {
		goto L22
	} else {
		goto L129
	}
L129:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v425)+120))
	if v476 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v477 != int32(1) {
		goto L19
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v425)+100))
	if v520 != 0 {
		goto L147
	} else {
		goto L148
	}
L133:
	;
	v480 = int32(0)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v469)+16))
	if base.B2i32(v482 == v480)|base.B2i32(v476 == v480) != 0 {
		v515 = v480
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v515 == int32(0) {
		goto L19
	} else {
		goto L146
	}
L135:
	;
	goto L134
L136:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if int32(0) < v488 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v492 = int32(0)
	goto L140
L138:
	;
	goto L139
L139:
	;
	v515 = v480
	goto L135
L140:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v476)+12))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v496+v492<<(uint(int32(2))%32))))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+4))
	if v482 == v501 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	v515 = int32(1)
	goto L135
L143:
	;
	goto L144
L144:
	;
	v505 = v492 + int32(1)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v505 < v506 {
		v492 = v505
		goto L140
	} else {
		goto L145
	}
L145:
	;
	goto L141
L146:
	;
	v518 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v518)
	goto L19
L147:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v521 != int32(1) {
		goto L19
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+40)))
	if v564 != 0 {
		goto L19
	} else {
		goto L164
	}
L150:
	;
	v524 = int32(0)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v469)+16))
	if base.B2i32(v526 == v524)|base.B2i32(v520 == v524) != 0 {
		v559 = v524
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v559 == int32(0) {
		goto L19
	} else {
		goto L163
	}
L152:
	;
	goto L151
L153:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if int32(0) < v532 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v536 = int32(0)
	goto L157
L155:
	;
	goto L156
L156:
	;
	v559 = v524
	goto L152
L157:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v520)+12))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v540+v536<<(uint(int32(2))%32))))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	if v526 == v545 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L156
L159:
	;
	v559 = int32(1)
	goto L152
L160:
	;
	goto L161
L161:
	;
	v549 = v536 + int32(1)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v549 < v550 {
		v536 = v549
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	v562 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v562)
	goto L19
L164:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	if v565 == int32(0) {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	if v568 != int32(6) {
		goto L19
	} else {
		goto L166
	}
L166:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v565)+28))
	if v571 == int32(0) {
		v184 = v422
		v188 = v565
		goto L54
	} else {
		goto L167
	}
L167:
	;
	goto L55
L168:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v577 = F_bms_difference(m, v574, v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	if v577 == int32(0) {
		goto L21
	} else {
		goto L170
	}
L170:
	;
	v583 = int32(0)
	if v577 == v583 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v654
	v711 = v84
	v712 = v656
	goto L20
L172:
	;
	if v637 != 0 {
		goto L187
	} else {
		goto L188
	}
L173:
	;
	v637 = int32(0)
	goto L172
L174:
	;
	goto L175
L175:
	;
	v591 = int32(1)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v577)+4))
	if v592 <= v591 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v595 = v591
	goto L178
L177:
	;
	v595 = v592
	goto L178
L178:
	;
	v600 = int32(0)
	v602 = int32(-1)
	goto L180
L179:
	;
	v637 = v629
	goto L172
L180:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v577+int32(8)+v600<<(uint(int32(2))%32))))
	if v610 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(76)))) = v621
	v629 = int32(1)
	goto L179
L182:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v610)))|base.B2i32(int32(0) <= v602) != 0 {
		v629 = v583
		goto L179
	} else {
		goto L185
	}
L183:
	;
	v621 = v602
	goto L184
L184:
	;
	v623 = v600 + int32(1)
	if v623 != v595 {
		v600 = v623
		v602 = v621
		goto L180
	} else {
		goto L186
	}
L185:
	;
	v621 = base.I32_ctz(v610) | v600<<(uint(int32(5))%32)
	goto L184
L186:
	;
	goto L181
L187:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if l2 != v639 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	if l2 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L190:
	;
	v641 = l2
	goto L192
L191:
	;
	v641 = int32(0)
	goto L192
L192:
	;
	if v641 != 0 {
		goto L21
	} else {
		goto L193
	}
L193:
	;
	v642 = F_find_base_rel(m, l0, v639)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v654 = v642
	v656 = v642
	goto L171
L195:
	;
	v646 = F_find_join_rel(m, l0, v574)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v648 = F_bms_is_member(m, l2, v574)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L199
	}
L198:
	;
	v654 = v646
	v656 = v5
	goto L171
L199:
	;
	if v648 == int32(0) {
		goto L21
	} else {
		goto L200
	}
L200:
	;
	v652 = F_find_base_rel(m, l0, l2)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v654 = v652
	v656 = v5
	goto L171
L202:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v202)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v663
	F_errmsg_internal(m, int32(_a_F_examine_variable_9), v16+int32(32))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_10), int32(_a_F_examine_variable_3))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v202)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v679
	F_errmsg_internal(m, int32(_a_F_examine_variable_11), v16+int32(48))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_12), int32(_a_F_examine_variable_3))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	v697 = int32(*(*int16)(unsafe.Add(mBase, uint32(v416))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v697
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v696
	F_errmsg_internal(m, int32(_a_F_examine_variable_13), v16)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_14), int32(_a_F_examine_variable_3))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v711
	v716 = F_exprType(m, v711)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v716
	v719 = F_exprTypmod(m, v711)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v719
	if v712 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	F_bms_free(m, v574)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L1
	} else {
		goto L338
	}
L215:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v725 = int32(0)
	if base.B2i32(v574 == v725)|base.B2i32(v724 == v725) != 0 {
		v770 = v725
		goto L217
	} else {
		goto L218
	}
L216:
	;
	if v770 != 0 {
		goto L229
	} else {
		goto L230
	}
L217:
	;
	goto L216
L218:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v724)+4))
	if v735 < v736 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v738 = v735
	goto L221
L220:
	;
	v738 = v736
	goto L221
L221:
	;
	if v738 <= int32(1) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v741 = int32(1)
	goto L224
L223:
	;
	v741 = v738
	goto L224
L224:
	;
	v742 = int32(8)
	v747 = int32(0)
	goto L225
L225:
	;
	v754 = v747 << (uint(int32(2)) % 32)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v724+v742+v754)))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v574+v742+v754)))
	v759 = v756 & v758
	v761 = base.B2i32(v759 != int32(0))
	if v759 != 0 {
		v770 = v761
		goto L217
	} else {
		goto L227
	}
L226:
	;
	v770 = v761
	goto L217
L227:
	;
	v763 = v747 + int32(1)
	if v763 != v741 {
		v747 = v763
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v773 = F_remove_nulling_relids(m, v711, v771, int32(0))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	v775 = v711
	goto L231
L231:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v712)+108))
	if v776 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v775 = v773
	goto L231
L233:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v712)+112))
	if v989 == int32(0) {
		goto L214
	} else {
		goto L293
	}
L234:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	if v779 <= int32(0) {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v792 = v5
	goto L236
L236:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v776)+12))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v795+v792<<(uint(int32(2))%32))))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+84))
	if v800 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	goto L233
L238:
	;
	v973 = v792 + int32(1)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	if v973 < v974 {
		v792 = v973
		goto L236
	} else {
		goto L292
	}
L239:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v800)+12))
	if v803 == int32(0) {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v799)+36))
	if int32(0) < v806 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v811 = int32(0)
	v812 = v803
	v816 = v806
	goto L244
L242:
	;
	goto L243
L243:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v958 != 0 {
		goto L233
	} else {
		goto L291
	}
L244:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v799)+44))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v823+v811<<(uint(int32(2))%32))))
	if v827 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	goto L243
L246:
	;
	if v812 != 0 {
		goto L251
	} else {
		goto L252
	}
L247:
	;
	v940 = v812
	v941 = v816
	goto L248
L248:
	;
	v943 = v811 + int32(1)
	if v943 < v941 {
		v811 = v943
		v812 = v940
		v816 = v941
		goto L244
	} else {
		goto L290
	}
L249:
	;
	v929 = v812 + int32(4)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v799)+84))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v931)+12))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v931)+4))
	if base.Ui32(v929) < base.Ui32(v932+v933<<(uint(int32(2))%32)) {
		goto L287
	} else {
		goto L288
	}
L250:
	;
	v925 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v925)
	goto L249
L251:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v812)))
	if v830 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L252:
	;
	goto L253
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L284
	}
L254:
	;
	v839 = F_equal(m, v775, v838)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L259
	}
L255:
	;
	v838 = int32(0)
	goto L254
L256:
	;
	goto L257
L257:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v830)))
	if v834 != int32(27) {
		v838 = v830
		goto L254
	} else {
		goto L258
	}
L258:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v830)+4))
	v838 = v837
	goto L254
L259:
	;
	if v839 == int32(0) {
		goto L249
	} else {
		goto L260
	}
L260:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799)+101)))
	if base.B2i32(v843 != int32(1))|v811 != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_examine_variable[1]))
	if v857 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L262:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v799)+40))
	if v847 != int32(1) {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v799)+88))
	if v850 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v799)+100)))
	if v851 != int32(1) {
		goto L261
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v854 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v854)
	goto L261
L267:
	;
	goto L266
L268:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v799)+88))
	if v885 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L269:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v799)+4))
	v864 = m.T0[v857].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v860, base.I32_extend16_s(v811+int32(1)), l3)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	if v864 == int32(0) {
		goto L268
	} else {
		goto L271
	}
L271:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v868 == int32(0) {
		goto L249
	} else {
		goto L272
	}
L272:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v871 != 0 {
		goto L233
	} else {
		goto L273
	}
L273:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_0), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_15), int32(_a_F_examine_variable_16))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v799)+4))
	v890 = int32(16)
	v897 = F_SearchSysCache3(m, int32(65), v889, (v811<<(uint(v890)%32)+int32(_a_F_examine_variable_17))>>(uint(v890)%32), int32(0))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v911 != 0 {
		goto L233
	} else {
		goto L283
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1490)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v897
	if v897 == int32(0) {
		goto L250
	} else {
		goto L281
	}
L281:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v799)+12))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v904)+68))
	v907 = F_all_rows_selectable(m, l0, v905, int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v907)
	goto L279
L283:
	;
	goto L249
L284:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_18), int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_19), int32(_a_F_examine_variable_16))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	v938 = v929
	goto L289
L288:
	;
	v938 = int32(0)
	goto L289
L289:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v799)+36))
	v940 = v938
	v941 = v939
	goto L248
L290:
	;
	goto L245
L291:
	;
	goto L238
L292:
	;
	goto L237
L293:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v989)+4))
	if v992 <= int32(0) {
		goto L214
	} else {
		goto L294
	}
L294:
	;
	v1006 = int32(0)
	goto L295
L295:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1009 != 0 {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	goto L214
L297:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v1024 != 0 {
		goto L214
	} else {
		goto L301
	}
L298:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v712)+68))
	v1023 = v1009 + v1010<<(uint(int32(2))%32)
	goto L297
L299:
	;
	goto L300
L300:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1014)+52))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1015)+12))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v712)+68))
	v1023 = v1016 + v1017<<(uint(int32(2))%32) - int32(4)
	goto L297
L301:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v989)+12))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1025+v1006<<(uint(int32(2))%32))))
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+16)))
	if v1030 != int32(101) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1177 = v1006 + int32(1)
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v989)+4))
	if v1177 < v1178 {
		v1006 = v1177
		goto L295
	} else {
		goto L337
	}
L303:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+8)))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+20)))
	if v1033 != v1035 {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+24))
	if v1037 == int32(0) {
		goto L302
	} else {
		goto L305
	}
L305:
	;
	v1040 = int32(0)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+4))
	if v1041 <= v1040 {
		goto L302
	} else {
		goto L306
	}
L306:
	;
	v1045 = v1040
	goto L307
L307:
	;
	v1057 = int32(0)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+12))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1058+v1045<<(uint(int32(2))%32))))
	if v1062 == v1057 {
		v1069 = v1057
		goto L309
	} else {
		goto L310
	}
L308:
	;
	goto L302
L309:
	;
	v1070 = F_equal(m, v775, v1069)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1062)))
	if v1065 != int32(27) {
		v1069 = v1062
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+4))
	v1069 = v1068
	goto L309
L312:
	;
	if v1070 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+4))
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+20)))
	v1074 = m.G0
	v1076 = v1074 - int32(48)
	m.G0 = v1076
	v1079 = F_SearchSysCache2(m, int32(62), v1072, v1073)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L318
	}
L314:
	;
	goto L315
L315:
	;
	v1160 = v1045 + int32(1)
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+4))
	if v1160 < v1161 {
		v1045 = v1160
		goto L307
	} else {
		goto L336
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1491)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1114
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v712)+68))
	v1156 = F_all_rows_selectable(m, l0, v1154, int32(0))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L335
	}
L317:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L1
	} else {
		goto L332
	}
L318:
	;
	if v1079 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1085 = F_SysCacheGetAttr(m, int32(62), v1079, int32(6), v1076+int32(47))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L329
	}
L322:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076)+47)))
	if v1087 == int32(1) {
		goto L317
	} else {
		goto L323
	}
L323:
	;
	v1090 = F_DatumGetExpandedArray(m, v1085)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_deconstruct_expanded_array(m, v1090)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+48))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1094+v1045<<(uint(int32(2))%32))))
	v1099 = F_pg_detoast_datum(m, v1098)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1099)))
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+40)) = v1099
	v1103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+36)) = v1103
	*(*uint16)(unsafe.Add(mBase, uint32(v1076)+32)) = uint16(v1103)
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+28)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+24)) = int32(base.Ui32(v1101) >> (uint(int32(2)) % 32))
	v1114 = F_heap_copytuple(m, v1076+int32(24))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	F_ReleaseCatCache(m, v1079)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	m.G0 = v1076 + int32(48)
	goto L316
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = v1072
	F_errmsg_internal(m, int32(_a_F_examine_variable_20), v1076)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(_a_F_examine_variable_21), int32(2414), int32(_a_F_examine_variable_22))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+20)) = v1072
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+16)) = int32(101)
	F_errmsg_internal(m, int32(_a_F_examine_variable_23), v1076+int32(16))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F_examine_variable_21), int32(2421), int32(_a_F_examine_variable_22))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v1156)
	goto L302
L336:
	;
	goto L308
L337:
	;
	goto L296
L338:
	;
	goto L19
}
func F_exec_run_select(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v12 == v4 {
		if l2 != 0 {
			v17 = int32(4)
		} else {
			v17 = int32(2052)
		}
		F_exec_prepare_plan(m, l0, l1, v17)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l1
				v25 = v23
			} else {
				v25 = v4
			}
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
			if l2 != 0 {
				v31 = F_SPI_cursor_open_internal(m, int32(0), v27, v25, v26&int32(1))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31
					if v31 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_run_select_0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v102 = *(*int32)(unsafe.Add(mBase, _c_F_exec_run_select[0]))
							v103 = F_SPI_result_code_string(m, v102)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v103
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v100
								F_errmsg_internal(m, int32(_a_F_exec_run_select_1), v10+int32(32))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_3), int32(_a_F_exec_run_select_4))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						if v36 != 0 {
							F_SPI_freetuptable(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v39
								v41 = int32(10)
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								if v42 == v39 {
									v90 = v41
									m.G0 = v10 + int32(48)
									return v90
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
									F_MemoryContextReset(m, v45)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v90 = v41
										m.G0 = v10 + int32(48)
										return v90
									}
								}
							}
						} else {
							v39 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v39
							v41 = int32(10)
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							if v42 == v39 {
								v90 = v41
								m.G0 = v10 + int32(48)
								return v90
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
								F_MemoryContextReset(m, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v90 = v41
									m.G0 = v10 + int32(48)
									return v90
								}
							}
						}
					}
				}
			} else {
				v52 = F_SPI_execute_plan_with_paramlist(m, v27, v25, v26&int32(1), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 != int32(5) {
						if v52 == int32(6) {
							F_errstart_cold(m, int32(21), int32(_a_F_exec_run_select_0))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16801924))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_exec_run_select_5), int32(0))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int32(0)
									} else {
										F_set_errcontext_domain(m, int32(_a_F_exec_run_select_0))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v131
											F_errcontext_msg(m, int32(_a_F_exec_run_select_6), v10+int32(16))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_7), int32(_a_F_exec_run_select_4))
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(_a_F_exec_run_select_0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16801924))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_exec_run_select_8), int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										F_set_errcontext_domain(m, int32(_a_F_exec_run_select_0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
											F_errcontext_msg(m, int32(_a_F_exec_run_select_6), v10)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_9), int32(_a_F_exec_run_select_4))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, _c_F_exec_run_select[1]))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v83
						v86 = *(*int64)(unsafe.Add(mBase, _c_F_exec_run_select[2]))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v86
						v90 = int32(5)
						m.G0 = v10 + int32(48)
						return v90
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		if v22 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l1
			v25 = v23
		} else {
			v25 = v4
		}
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if l2 != 0 {
			v31 = F_SPI_cursor_open_internal(m, int32(0), v27, v25, v26&int32(1))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31
				if v31 == int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_run_select_0))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v102 = *(*int32)(unsafe.Add(mBase, _c_F_exec_run_select[0]))
						v103 = F_SPI_result_code_string(m, v102)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v103
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v100
							F_errmsg_internal(m, int32(_a_F_exec_run_select_1), v10+int32(32))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_3), int32(_a_F_exec_run_select_4))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v36 != 0 {
						F_SPI_freetuptable(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v39
							v41 = int32(10)
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							if v42 == v39 {
								v90 = v41
								m.G0 = v10 + int32(48)
								return v90
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
								F_MemoryContextReset(m, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									v90 = v41
									m.G0 = v10 + int32(48)
									return v90
								}
							}
						}
					} else {
						v39 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v39
						v41 = int32(10)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						if v42 == v39 {
							v90 = v41
							m.G0 = v10 + int32(48)
							return v90
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
							F_MemoryContextReset(m, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v90 = v41
								m.G0 = v10 + int32(48)
								return v90
							}
						}
					}
				}
			}
		} else {
			v52 = F_SPI_execute_plan_with_paramlist(m, v27, v25, v26&int32(1), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if v52 != int32(5) {
					if v52 == int32(6) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_run_select_0))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16801924))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_exec_run_select_5), int32(0))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									F_set_errcontext_domain(m, int32(_a_F_exec_run_select_0))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
									} else {
										v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v131
										F_errcontext_msg(m, int32(_a_F_exec_run_select_6), v10+int32(16))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_7), int32(_a_F_exec_run_select_4))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_run_select_0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16801924))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_exec_run_select_8), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									F_set_errcontext_domain(m, int32(_a_F_exec_run_select_0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
										F_errcontext_msg(m, int32(_a_F_exec_run_select_6), v10)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_9), int32(_a_F_exec_run_select_4))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, _c_F_exec_run_select[1]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v83
					v86 = *(*int64)(unsafe.Add(mBase, _c_F_exec_run_select[2]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v86
					v90 = int32(5)
					m.G0 = v10 + int32(48)
					return v90
				}
			}
		}
	}
}
func F_executeStartsWith(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v5 = int32(2)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6 != int32(1) {
		v82 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v82
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v9 != int32(1) {
		v82 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v13 < v12 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v82 = int32(0)
	goto L1
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(int32(4)) <= base.Ui32(v12) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if v78 != 0 {
		goto L4
	} else {
		goto L24
	}
L7:
	;
	v78 = int32(0)
	goto L6
L8:
	;
	v52 = v47
	v53 = v48
	v54 = v49
	goto L18
L9:
	;
	if (v15|v16)&int32(3) != 0 {
		v47 = v15
		v48 = v16
		v49 = v12
		goto L8
	} else {
		goto L12
	}
L10:
	;
	v40 = v15
	v41 = v16
	v42 = v12
	goto L11
L11:
	;
	if v42 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L12:
	;
	v24 = v15
	v25 = v16
	v26 = v12
	goto L13
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v29 != v30 {
		v47 = v24
		v48 = v25
		v49 = v26
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v40 = v35
	v41 = v33
	v42 = v37
	goto L11
L15:
	;
	v32 = int32(4)
	v33 = v25 + v32
	v35 = v24 + v32
	v37 = v26 - v32
	if base.Ui32(int32(3)) < base.Ui32(v37) {
		v24 = v35
		v25 = v33
		v26 = v37
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v47 = v40
	v48 = v41
	v49 = v42
	goto L8
L18:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 == v58 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v78 = v57 - v58
	goto L6
L20:
	;
	v60 = int32(1)
	v65 = v54 - v60
	if v65 != 0 {
		v52 = v52 + v60
		v53 = v53 + v60
		v54 = v65
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	goto L7
L24:
	;
	return int32(1)
}
func F_expand_dynamic_library_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = Fn13880(m, l0, int32(47))
	mBase = m.M
	if v10 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_expand_dynamic_library_name[0]))
		v15 = F_find_in_path(m, l0, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v63 = v15
				m.G0 = v7 + int32(32)
				return v63
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_expand_dynamic_library_name_0)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				v23 = F_psprintf(m, int32(_a_F_expand_dynamic_library_name_1), v7)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_expand_dynamic_library_name[0]))
					v27 = F_find_in_path(m, v23, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v23)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v27 == int32(0) {
								v61 = F_pstrdup(m, l0)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v63 = v61
									m.G0 = v7 + int32(32)
									return v63
								}
							} else {
								v63 = v27
								m.G0 = v7 + int32(32)
								return v63
							}
						}
					}
				}
			}
		}
	} else {
		v35 = F_substitute_path_macro(m, l0, int32(_a_F_expand_dynamic_library_name_2), int32(_a_F_expand_dynamic_library_name_3))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = F_pg_file_exists(m, v35)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v37 != 0 {
					v63 = v35
					m.G0 = v7 + int32(32)
					return v63
				} else {
					F_pfree(m, v35)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a_F_expand_dynamic_library_name_0)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
						v47 = F_psprintf(m, int32(_a_F_expand_dynamic_library_name_1), v7+int32(16))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v51 = F_substitute_path_macro(m, v47, int32(_a_F_expand_dynamic_library_name_2), int32(_a_F_expand_dynamic_library_name_3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v47)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = F_pg_file_exists(m, v51)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										if v55 != 0 {
											v63 = v51
											m.G0 = v7 + int32(32)
											return v63
										} else {
											F_pfree(m, v51)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												v61 = F_pstrdup(m, l0)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													v63 = v61
													m.G0 = v7 + int32(32)
													return v63
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_expand_insert_targetlist(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = v17
	goto L3
L2:
	;
	v18 = v4
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+120)))
	if int32(0) < v21 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = int32(1)
	v32 = v18
	v33 = v24
	v36 = v4
	goto L7
L5:
	;
	v131 = v18
	v133 = int32(1)
	v135 = v4
	goto L6
L6:
	;
	if v131 != 0 {
		goto L36
	} else {
		goto L37
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v131 = v118
	v133 = v21 + v24
	v135 = v120
	goto L6
L9:
	;
	v120 = F_lappend(m, v36, v116)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L21
	} else {
		goto L33
	}
L10:
	;
	v65 = v39 + v40<<(uint(int32(4))%32) + v33*int32(100) - int32(80)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+91)))
	if v66 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+26)))
	if v44 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+8)))
	if v33 != v45 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v48 = v32 + int32(4)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v48) < base.Ui32(v50+v51<<(uint(int32(2))%32)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v56 = v48
	goto L16
L15:
	;
	v56 = int32(0)
	goto L16
L16:
	;
	v116 = v43
	v118 = v56
	goto L9
L17:
	;
	v111 = F_pstrdup(m, v65+int32(4))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L21
	} else {
		goto L31
	}
L18:
	;
	v71 = int32(0)
	v74 = int32(1)
	v76 = F_makeConst(m, int32(23), int32(-1), v71, int32(4), v71, v74, v74)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)+68))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+90)))
	if v81 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	return int32(0)
L22:
	;
	v107 = v76
	goto L17
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v65)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v82
	v86 = F_getBaseTypeAndTypmod(m, v80, v15+int32(12))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v65)+76))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v65)+96))
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+72)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+82)))
	v100 = F_coerce_null_to_domain(m, v80, v96, v97, v98, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v65)+96))
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v65)+72)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+82)))
	v94 = F_makeConst(m, v86, v88, v89, v90, int32(0), int32(1), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v107 = v94
	goto L17
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v102 == int32(7) {
		v107 = v100
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v105 = F_eval_const_expressions(m, l0, v100)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L21
	} else {
		goto L30
	}
L30:
	;
	v107 = v105
	goto L17
L31:
	;
	v114 = F_makeTargetEntry(m, v107, base.I32_extend16_s(v33), v111, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v116 = v114
	v118 = v32
	goto L9
L33:
	;
	if base.B2i32(v33 == v21) == int32(0) {
		v32 = v118
		v33 = v33 + int32(1)
		v36 = v120
		goto L7
	} else {
		goto L34
	}
L34:
	;
	goto L8
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L21
	} else {
		goto L48
	}
L36:
	;
	v143 = v131
	v145 = v133
	v147 = v135
	goto L39
L37:
	;
	v181 = v135
	goto L38
L38:
	;
	m.G0 = v15 + int32(16)
	return v181
L39:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+26)))
	if v151 == int32(0) {
		goto L35
	} else {
		goto L41
	}
L40:
	;
	v181 = v160
	goto L38
L41:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+8)))
	if v154 != v145 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v156 = F_flatCopyTargetEntry(m, v150)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L21
	} else {
		goto L45
	}
L43:
	;
	v159 = v150
	goto L44
L44:
	;
	v160 = F_lappend(m, v147, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L21
	} else {
		goto L46
	}
L45:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+8)) = uint16(v145)
	v159 = v156
	goto L44
L46:
	;
	v165 = v143 + int32(4)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v165) < base.Ui32(v166+v167<<(uint(int32(2))%32)) {
		v143 = v165
		v145 = v145 + int32(1)
		v147 = v160
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	F_errmsg_internal(m, int32(_a_F_expand_insert_targetlist_0), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_expand_insert_targetlist_1), int32(504), int32(_a_F_expand_insert_targetlist_2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_expand_planner_arrays(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = int32(2)
	v9 = l1 + v6
	v11 = v9 << (uint(v7) % 32)
	v12 = F_repalloc0(m, v5, v6<<(uint(v7)%32), v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v12
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v19 = F_repalloc0(m, v15, v16<<(uint(int32(2))%32), v11)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v19
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v26 = F_repalloc0(m, v22, v23<<(uint(int32(2))%32), v11)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v30 = v26
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v30
					return
				}
			} else {
				v28 = F_palloc0(m, v11)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = v28
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v30
					return
				}
			}
		}
	}
}
func F_extractRemainingColumns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	v8 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v15 == v8 {
		v61 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		v61 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = v8
	v34 = v8
	goto L4
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v29<<(uint(int32(2))%32))))
	v40 = F_bms_add_member(m, v34, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v61 = v40
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	v45 = v29 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v45 < v46 {
		v29 = v45
		v34 = v40
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v63 <= v62 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v178 = v8
	goto L11
L11:
	;
	return v178
L12:
	;
	return int32(0)
L13:
	;
	goto L14
L14:
	;
	v76 = v62
	v80 = v8
	goto L15
L15:
	;
	v83 = v76 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v87 = v84 + v76<<(uint(int32(2))%32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v90 == int32(0) {
		v163 = v80
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v178 = v163
	goto L11
L17:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v83 < v164 {
		v76 = v83
		v80 = v163
		goto L15
	} else {
		goto L26
	}
L18:
	;
	v93 = F_bms_is_member(m, v83, v61)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	if v93 != 0 {
		v163 = v80
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v96 = F_lappend_int(m, v95, v83)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v101 = F_lappend(m, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v101
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v107 = l1 + v83<<(uint(int32(5))%32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(32))))
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v107-int32(28)))))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(24))))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(20))))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(16))))
	v124 = F_makeVar(m, v110, v113, v116, v119, v122, int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+32)) = v128
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v107-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+36)) = v132
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v124)+40)) = uint16(v136)
	F_markNullableIfNeeded(m, l0, v124)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v140 = F_lappend(m, v104, v124)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v140
	v143 = int32(5)
	v145 = l6 + v80<<(uint(v143)%32)
	v148 = l1 + v76<<(uint(v143)%32)
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v148)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v145)+24)) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v148)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v145)+16)) = v151
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v145)+8)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v155
	v163 = v80 + int32(1)
	goto L17
L26:
	;
	goto L16
}
func F_extract_actual_join_clauses(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	if l0 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v5
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v23<<(uint(int32(2))%32))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L1
L6:
	;
	v111 = v23 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v111 < v112 {
		v23 = v111
		goto L4
	} else {
		goto L44
	}
L7:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v105 = F_lappend(m, v104, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L42
	} else {
		goto L43
	}
L8:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v97 != int32(7) {
		goto L35
	} else {
		goto L36
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v33 = int32(0)
	if v32 == v33 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+10)))
	if v87 != 0 {
		goto L6
	} else {
		goto L27
	}
L12:
	;
	if v86 != 0 {
		goto L8
	} else {
		goto L26
	}
L13:
	;
	v86 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	if l1 == int32(0) {
		v79 = v33
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v86 = v79
	goto L12
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v43 < v42 {
		v79 = v33
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v45 = int32(1)
	if v42 <= v45 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v48 = v45
	goto L21
L20:
	;
	v48 = v42
	goto L21
L21:
	;
	v49 = int32(8)
	v54 = int32(0)
	goto L22
L22:
	;
	v61 = v54 << (uint(int32(2)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v32+v49+v61)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1+v49+v61)))
	v68 = v63 & (v65 ^ int32(-1))
	v70 = base.B2i32(v68 == int32(0))
	if v68 != 0 {
		v79 = v70
		goto L16
	} else {
		goto L24
	}
L23:
	;
	v79 = v70
	goto L16
L24:
	;
	v72 = v54 + int32(1)
	if v72 != v48 {
		v54 = v72
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L11
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v89 != int32(7) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v102 = l3
	v103 = v88
	goto L7
L29:
	;
	goto L30
L30:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+24)))
	if v92 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v102 = l3
	v103 = v88
	goto L7
L32:
	;
	goto L33
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	if v93 == int32(0) {
		v102 = l3
		v103 = v88
		goto L7
	} else {
		goto L34
	}
L34:
	;
	goto L6
L35:
	;
	v102 = l2
	v103 = v96
	goto L7
L36:
	;
	goto L37
L37:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+24)))
	if v100 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v102 = l2
	v103 = v96
	goto L7
L39:
	;
	goto L40
L40:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	if v101 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v102 = l2
	v103 = v96
	goto L7
L42:
	;
	return
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v105
	goto L6
L44:
	;
	goto L5
}
