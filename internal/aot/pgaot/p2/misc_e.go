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
			F_sequence_close(m, v8, int32(0))
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
	var v41 int32
	_ = v41
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
	v41 = v3
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v46 = v43 + v41*int32(60)
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
	v74 = v41 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v74 < v75 {
		v41 = v74
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
	v70 = F_ExecCallTriggerFunc(m, v9+int32(4), v41, v63, v64, v69)
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
	var v60 int64
	_ = v60
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
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
		goto L15
	} else {
		goto L16
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
	if l6 == int32(0) {
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
	if v54 <= int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l5)+112))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+112)) = v60 + base.I64_extend_i32_u(v54)
	goto L3
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = int32(0)
	m.G0 = v14 + int32(16)
	return
L16:
	;
	v77 = int32(1)
	v79 = int32(0)
	if l4 != v77 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v90 = int32(0)
	v93 = v79
	goto L20
L18:
	;
	v137 = v79
	goto L19
L19:
	;
	if l4&v77 == int32(0) {
		goto L15
	} else {
		goto L27
	}
L20:
	;
	v97 = v93 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2+v97)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	m.T0[v101].(func(*base.Module, int32))(m, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v137 = v125
	goto L19
L22:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l3+v97)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	m.T0[v107].(func(*base.Module, int32))(m, v105)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v111 = v97 | int32(4)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l2+v111)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	m.T0[v115].(func(*base.Module, int32))(m, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l3+v111)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	m.T0[v121].(func(*base.Module, int32))(m, v119)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v124 = int32(2)
	v125 = v93 + v124
	v127 = v90 + v124
	if v127 != l4&int32(2147483646) {
		v90 = v127
		v93 = v125
		goto L20
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v143 = v137 << (uint(int32(2)) % 32)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l2+v143)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	m.T0[v147].(func(*base.Module, int32))(m, v145)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l3+v143)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	m.T0[v153].(func(*base.Module, int32))(m, v151)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L15
}
func F_ExecBuildAuxRowMark(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	v11 = F_palloc0(m, int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v17 != int32(5) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L78
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L75
	}
L5:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v176 != v177 {
		goto L52
	} else {
		goto L53
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v16
	v27 = F_pg_snprintf(m, v8+int32(96), int32(32), int32(_a_F_ExecBuildAuxRowMark_0), v8+int32(80))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v16
	v112 = F_pg_snprintf(m, v8+int32(96), int32(32), int32(_a_F_ExecBuildAuxRowMark_1), v8+int32(48))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L32
	}
L9:
	;
	v31 = int32(0)
	if l1 == v31 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)) = uint16(v86)
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v38 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v41 = int32(0)
	if v41 < v38 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v78 = v31
	goto L16
L16:
	;
	v86 = base.I32_extend16_s(v78)
	goto L10
L17:
	;
	v44 = v38
	goto L19
L18:
	;
	v44 = v41
	goto L19
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v47 = int32(0)
	goto L21
L20:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+8)))
	v78 = v71
	goto L16
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45+v47<<(uint(int32(2))%32))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+26)))
	if v57 != int32(1) {
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
	v68 = v47 + int32(1)
	if v68 != v44 {
		v47 = v68
		goto L21
	} else {
		goto L27
	}
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v60 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v63 = F_strcmp(m, v60, v8+int32(96))
	mBase = m.M
	if v63 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v8 + int32(96)
	F_errmsg_internal(m, int32(_a_F_ExecBuildAuxRowMark_2), v8-int32(-64))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2598), int32(_a_F_ExecBuildAuxRowMark_4))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
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
	v116 = int32(0)
	if l1 == v116 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+8)) = uint16(v171)
	if v171 == int32(0) {
		goto L4
	} else {
		goto L51
	}
L34:
	;
	v171 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v123 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v126 = int32(0)
	if v126 < v123 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v163 = v116
	goto L39
L39:
	;
	v171 = base.I32_extend16_s(v163)
	goto L33
L40:
	;
	v129 = v123
	goto L42
L41:
	;
	v129 = v126
	goto L42
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v132 = int32(0)
	goto L44
L43:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+8)))
	v163 = v156
	goto L39
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v130+v132<<(uint(int32(2))%32))))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+26)))
	if v142 != int32(1) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v171 = int32(0)
	goto L33
L46:
	;
	v153 = v132 + int32(1)
	if v153 != v129 {
		v132 = v153
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	if v145 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v148 = F_strcmp(m, v145, v8+int32(96))
	mBase = m.M
	if v148 == int32(0) {
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
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v179
	v183 = int32(32)
	v187 = F_pg_snprintf(m, v8+int32(96), v183, int32(_a_F_ExecBuildAuxRowMark_5), v8+v183)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	m.G0 = v8 + int32(128)
	return v11
L55:
	;
	v191 = int32(0)
	if l1 == v191 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+6)) = uint16(v246)
	if v246 == int32(0) {
		goto L3
	} else {
		goto L74
	}
L57:
	;
	v246 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v198 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v201 = int32(0)
	if v201 < v198 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v238 = v191
	goto L62
L62:
	;
	v246 = base.I32_extend16_s(v238)
	goto L56
L63:
	;
	v204 = v198
	goto L65
L64:
	;
	v204 = v201
	goto L65
L65:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v207 = int32(0)
	goto L67
L66:
	;
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
	v238 = v231
	goto L62
L67:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v205+v207<<(uint(int32(2))%32))))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+26)))
	if v217 != int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v246 = int32(0)
	goto L56
L69:
	;
	v228 = v207 + int32(1)
	if v228 != v204 {
		v207 = v228
		goto L67
	} else {
		goto L73
	}
L70:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	if v220 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v223 = F_strcmp(m, v220, v8+int32(96))
	mBase = m.M
	if v223 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(96)
	F_errmsg_internal(m, int32(_a_F_ExecBuildAuxRowMark_2), v8)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2607), int32(_a_F_ExecBuildAuxRowMark_4))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(96)
	F_errmsg_internal(m, int32(_a_F_ExecBuildAuxRowMark_2), v8+int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2617), int32(_a_F_ExecBuildAuxRowMark_4))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
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
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	F_MemoryContextReset(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L32
	}
L9:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v90 != 0 {
		goto L29
	} else {
		goto L30
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
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v78
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
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v75 = v73 << (uint(int32(2)) % 32)
	if v75 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L19
L26:
	;
	v76 = F__emscripten_memcpy_bulkmem(m, v68, v72, v75)
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGather[1])))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+80)))
	v97 = v92 & (v93 ^ int32(1))
	goto L31
L30:
	;
	v97 = int32(1)
	goto L31
L31:
	;
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v98)
	v101 = v97 & v98
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v101)
	goto L8
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L34
L33:
	;
	m.G0 = v12 + int32(16)
	return v453
L34:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v125 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+4)))
	if v417&int32(2) != 0 {
		v453 = int32(0)
		goto L33
	} else {
		goto L141
	}
L36:
	;
	goto L35
L37:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v383 != int32(1) {
		goto L34
	} else {
		goto L128
	}
L38:
	;
	v368 = int32(0)
	v370 = F_ExecStoreMinimalTuple(m, v162, v114, v368)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L126
	}
L39:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
	m.T0[v365].(func(*base.Module, int32))(m, v114)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L125
	}
L40:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v128 != int32(1) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[0]))
	if v132 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	v136 = v125
	goto L46
L46:
	;
	v137 = int32(0)
	if v136 <= v137 {
		goto L37
	} else {
		goto L48
	}
L47:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v136 = v135
	goto L46
L48:
	;
	v142 = v137
	goto L49
L49:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[0]))
	if v150 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153+v154<<(uint(int32(2))%32))))
	v162 = F_TupleQueueReaderNext(m, v158, int32(1), v12+int32(15))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v164 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v169 = v167 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v169
	if v169 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v162 != 0 {
		goto L38
	} else {
		goto L117
	}
L59:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v173 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v183 = int32(2)
	v185 = v181 + v182<<(uint(v183)%32)
	v187 = v185 + int32(4)
	v190 = (v169 - v182) << (uint(v183) % 32)
	if v185 == v187 {
		goto L71
	} else {
		goto L72
	}
L62:
	;
	F_ExecParallelFinish(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v176 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	F_pfree(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	goto L37
L69:
	;
	goto L68
L70:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v335 < v336 {
		goto L49
	} else {
		goto L116
	}
L71:
	;
	goto L70
L72:
	;
	v194 = v185 + v190
	if base.Ui32(v187-v194) <= base.Ui32(int32(0)-v190<<(uint(int32(1))%32)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v201 = F___memcpy(m, v185, v187, v190)
	mBase = m.M
	goto L70
L74:
	;
	goto L75
L75:
	;
	v204 = (v185 ^ v187) & int32(3)
	if base.Ui32(v185) < base.Ui32(v187) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	if v306 == int32(0) {
		goto L71
	} else {
		goto L112
	}
L77:
	;
	if base.Ui32(v284) <= base.Ui32(int32(3)) {
		v305 = v283
		v306 = v284
		v307 = v285
		goto L76
	} else {
		goto L108
	}
L78:
	;
	if v204 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	if v204 != 0 {
		v266 = v190
		goto L91
	} else {
		goto L92
	}
L81:
	;
	v305 = v187
	v306 = v190
	v307 = v185
	goto L76
L82:
	;
	goto L83
L83:
	;
	if v185&int32(3) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v283 = v187
	v284 = v190
	v285 = v185
	goto L77
L85:
	;
	goto L86
L86:
	;
	v211 = v187
	v212 = v190
	v213 = v185
	goto L87
L87:
	;
	if v212 == int32(0) {
		goto L71
	} else {
		goto L89
	}
L88:
	;
	v283 = v220
	v284 = v222
	v285 = v224
	goto L77
L89:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v217)
	v219 = int32(1)
	v220 = v211 + v219
	v222 = v212 - v219
	v224 = v213 + v219
	if v224&int32(3) != 0 {
		v211 = v220
		v212 = v222
		v213 = v224
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	if v266 == int32(0) {
		goto L71
	} else {
		goto L104
	}
L92:
	;
	if v194&int32(3) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v231 = v190
	goto L96
L94:
	;
	v246 = v190
	goto L95
L95:
	;
	if base.Ui32(v246) <= base.Ui32(int32(3)) {
		v266 = v246
		goto L91
	} else {
		goto L100
	}
L96:
	;
	if v231 == int32(0) {
		goto L71
	} else {
		goto L98
	}
L97:
	;
	v246 = v237
	goto L95
L98:
	;
	v237 = v231 - int32(1)
	v238 = v185 + v237
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v237))))
	*(*uint8)(unsafe.Add(mBase, uint32(v238))) = uint8(v240)
	if v238&int32(3) != 0 {
		v231 = v237
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v253 = v246
	goto L101
L101:
	;
	v257 = v253 - int32(4)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v187+v257)))
	*(*int32)(unsafe.Add(mBase, uint32(v185+v257))) = v260
	if base.Ui32(int32(3)) < base.Ui32(v257) {
		v253 = v257
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v266 = v257
	goto L91
L103:
	;
	goto L102
L104:
	;
	v273 = v266
	goto L105
L105:
	;
	v277 = v273 - int32(1)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v277))))
	*(*uint8)(unsafe.Add(mBase, uint32(v185+v277))) = uint8(v280)
	if v277 != 0 {
		v273 = v277
		goto L105
	} else {
		goto L107
	}
L106:
	;
	goto L71
L107:
	;
	goto L106
L108:
	;
	v290 = v283
	v291 = v284
	v292 = v285
	goto L109
L109:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = v294
	v296 = int32(4)
	v297 = v290 + v296
	v299 = v292 + v296
	v301 = v291 - v296
	if base.Ui32(int32(3)) < base.Ui32(v301) {
		v290 = v297
		v291 = v301
		v292 = v299
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v305 = v297
	v306 = v301
	v307 = v299
	goto L76
L111:
	;
	goto L110
L112:
	;
	v312 = v305
	v313 = v306
	v314 = v307
	goto L113
L113:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312))))
	*(*uint8)(unsafe.Add(mBase, uint32(v314))) = uint8(v316)
	v318 = int32(1)
	v323 = v313 - v318
	if v323 != 0 {
		v312 = v312 + v318
		v313 = v323
		v314 = v314 + v318
		goto L113
	} else {
		goto L115
	}
L114:
	;
	goto L71
L115:
	;
	goto L114
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(0)
	goto L49
L117:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v342 = v340 + int32(1)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v342 < v344 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v346 = v342
	goto L120
L119:
	;
	v346 = int32(0)
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v346
	v349 = v142 + int32(1)
	if v349 < v344 {
		v142 = v349
		goto L49
	} else {
		goto L121
	}
L121:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v351 != 0 {
		goto L37
	} else {
		goto L122
	}
L122:
	;
	v352 = int32(0)
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[2]))
	v358 = F_WaitLatch(m, v354, int32(33), v352, int32(134217741))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = int32(0)
	goto L124
L124:
	;
	v142 = v352
	goto L49
L125:
	;
	v408 = v114
	goto L36
L126:
	;
	if v114 == int32(0) {
		v453 = v368
		goto L33
	} else {
		goto L127
	}
L127:
	;
	v408 = v114
	goto L36
L128:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v387 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+24))
	v390 = v388
	goto L131
L130:
	;
	v390 = int32(0)
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+172)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v115)+52))
	if v392 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	F_ExecReScan(m, v115)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v396 = m.T0[v395].(func(*base.Module, int32) int32)(m, v115)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+172)) = int32(0)
	if v396 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+4)))
	if v400&int32(2) == int32(0) {
		v408 = v396
		goto L36
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v405 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v405)
	goto L34
L140:
	;
	goto L139
L141:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v420 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v453 = v408
	goto L33
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+12)) = v408
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)+72))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v424)+16))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	m.T0[v428].(func(*base.Module, int32))(m, v426)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v431 = int32(_a_F_ExecGather_0)
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v425)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3])) = v434
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v424)+24))
	v440 = m.T0[v439].(func(*base.Module, int32, int32, int32) int32)(m, v424+int32(4), v425, int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3])) = v432
	v444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v426)+4)))
	v446 = v444 & int32(_a_F_ExecGather_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v426)+4)) = uint16(v446)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v426)+12))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	*(*uint16)(unsafe.Add(mBase, uint32(v426)+6)) = uint16(v449)
	v453 = v426
	goto L33
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
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
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	F_MemoryContextReset(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L33
	}
L9:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGatherMerge[1])))
	if v80 == int32(0) {
		goto L29
	} else {
		goto L30
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
	v71 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v71
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
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v68 = v66 << (uint(int32(2)) % 32)
	if v68 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L9
L25:
	;
	v69 = F__emscripten_memcpy_bulkmem(m, v61, v65, v68)
	mBase = m.M
	goto L27
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v86)
	goto L8
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v83 != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)) = uint8(v84)
	goto L28
L32:
	;
	goto L31
L33:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v97 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if v274 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v102 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v102
	if v102 < v100 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+20))
	v259 = F_gather_merge_readnext(m, l0, v257, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L83
	}
L38:
	;
	v109 = int32(0)
	goto L41
L39:
	;
	goto L40
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+8)) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(0)
	goto L45
L41:
	;
	v116 = v109 << (uint(int32(4)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116+v117)+4)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v121+v116)+8)) = v119
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v116)+12)) = uint8(v119)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v131 = v109 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+v131<<(uint(int32(2))%32))))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	m.T0[v137].(func(*base.Module, int32))(m, v135)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	if v131 != v100 {
		v109 = v131
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v100 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_binaryheap_build(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L82
	}
L47:
	;
	v160 = int32(1)
	goto L48
L48:
	;
	v166 = int32(0)
	goto L50
L49:
	;
	goto L46
L50:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[0]))
	if v173 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v100 <= int32(0) {
		goto L46
	} else {
		goto L73
	}
L52:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v166 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L54
L56:
	;
	v207 = v166 + int32(1)
	if v207 <= v100 {
		v166 = v207
		goto L50
	} else {
		goto L72
	}
L57:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+v166<<(uint(int32(2))%32))))
	if v190 != 0 {
		goto L64
	} else {
		goto L65
	}
L58:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)))
	if v178 != 0 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v180 = int32(4)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v166<<(uint(v180)%32)-v180))))
	if v185 != 0 {
		goto L56
	} else {
		goto L62
	}
L61:
	;
	goto L56
L62:
	;
	goto L57
L63:
	;
	F_load_tuple_array(m, l0, v166)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L71
	}
L64:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+4)))
	if v191&int32(2) == int32(0) {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v196 = F_gather_merge_readnext(m, l0, v166, v160&int32(1))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	if v196 == int32(0) {
		goto L56
	} else {
		goto L69
	}
L69:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_binaryheap_add_unordered(m, v200, v166)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	goto L56
L71:
	;
	goto L56
L72:
	;
	goto L51
L73:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v216 = int32(1)
	goto L74
L74:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211-int32(4)+v216<<(uint(int32(4))%32)))))
	if v225 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L49
L76:
	;
	v228 = int32(0)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v229+v216<<(uint(int32(2))%32))))
	if v233 == v228 {
		v160 = v228
		goto L48
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v242 = v216 + int32(1)
	if v242 <= v100 {
		v216 = v242
		goto L74
	} else {
		goto L81
	}
L79:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+4)))
	if v236&int32(2) != 0 {
		v160 = v228
		goto L48
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	goto L75
L82:
	;
	v254 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v254)
	goto L34
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v259 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_binaryheap_replace_first(m, v261, v257)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v264 = F_binaryheap_remove_first(m, v261)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	goto L34
L88:
	;
	goto L34
L89:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v277 <= int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	v340 = int32(0)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341+v339<<(uint(int32(2))%32))))
	if v345 == v340 {
		v384 = v340
		goto L106
	} else {
		goto L107
	}
L92:
	;
	return int32(0)
L93:
	;
	goto L94
L94:
	;
	v286 = int32(0)
	goto L95
L95:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v293 = v290 + v286<<(uint(int32(4))%32)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v294 < v295 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	return int32(0)
L97:
	;
	v299 = v294
	goto L100
L98:
	;
	goto L99
L99:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v326 = v286 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v324+v326<<(uint(int32(2))%32))))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	m.T0[v332].(func(*base.Module, int32))(m, v330)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L104
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v299 + int32(1)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307+v299<<(uint(int32(2))%32))))
	F_pfree(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	goto L99
L102:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	if v314 < v315 {
		v299 = v314
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v326 < v335 {
		v286 = v326
		goto L95
	} else {
		goto L105
	}
L105:
	;
	goto L96
L106:
	;
	return v384
L107:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+4)))
	if v348&int32(2) != 0 {
		v384 = v340
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v351 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	return v345
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v345
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+72))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v356)+16))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	m.T0[v360].(func(*base.Module, int32))(m, v358)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v363 = int32(_a_F_ExecGatherMerge_0)
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2]))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v357)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2])) = v366
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v356)+24))
	v372 = m.T0[v371].(func(*base.Module, int32, int32, int32) int32)(m, v356+int32(4), v357, int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2])) = v364
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v358)+4)))
	v378 = v376 & int32(_a_F_ExecGatherMerge_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v358)+4)) = uint16(v378)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	*(*uint16)(unsafe.Add(mBase, uint32(v358)+6)) = uint16(v381)
	v384 = v358
	goto L106
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
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
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 float64
	_ = v576
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v788 int64
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v933 int32
	_ = v933
	var v936 float64
	_ = v936
	var v946 int32
	_ = v946
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
		v54 = l1
		v55 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
	m.T0[v60].(func(*base.Module, int32))(m, v55)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
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
	v46 = F_ExecGetRootToChildMap(m, v33, v26)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
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
	if v46 == int32(0) {
		v54 = v33
		v55 = l2
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v33)+204))
	v52 = F_execute_attr_map_slot(m, v50, l2, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v54 = v33
	v55 = v52
	goto L1
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+48))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+116)))
	if v65 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	if v73 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v68 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_ExecOpenIndices(m, v54, base.B2i32(v29 != int32(0)))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
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
	return v946
L23:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v933 == int32(0) {
		v946 = v924
		goto L22
	} else {
		goto L293
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L3
	} else {
		goto L290
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L3
	} else {
		goto L287
	}
L26:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
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
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v54)+84))
	if v98 != 0 {
		goto L45
	} else {
		goto L46
	}
L29:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+8)))
	if v76 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
	if v79 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v90 = v73
	goto L32
L32:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+10)))
	if v91 != int32(1) {
		goto L28
	} else {
		goto L42
	}
L33:
	;
	F_ExecPendingInserts(m, v26)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v82 = F_ExecBRInsertTriggers(m, v26, v54, v55)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	if v82 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v946 = int32(0)
	goto L22
L39:
	;
	goto L40
L40:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	if v87 == int32(0) {
		goto L28
	} else {
		goto L41
	}
L41:
	;
	v90 = v87
	goto L32
L42:
	;
	v94 = int32(0)
	v95 = F_ExecIRInsertTriggers(m, v26, v54, v55)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	if v95 != 0 {
		v770 = v55
		v774 = v94
		goto L27
	} else {
		goto L44
	}
L44:
	;
	v946 = v94
	goto L22
L45:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+36)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v103 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v63)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+36)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	if v226 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L48:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v54)+104))
	if int32(2) <= v112 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+17)))
	if v106 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	F_ExecComputeStoredGenerated(m, v54, v26, v55, int32(3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	if v112 == v115 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v54)+84))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+52))
	v214 = m.T0[v213].(func(*base.Module, int32, int32, int32, int32) int32)(m, v26, v54, v55, v25)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L3
	} else {
		goto L78
	}
L55:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v54)+108))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v54)+112))
	F_ExecBatchInsert(m, v27, v54, v117, v118, v112, v26, l3)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v121 = int32(_a_F_ExecInsert_0)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v26)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v54)+108))
	if v126 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v54)+104))
	v132 = F_palloc(m, v129<<(uint(int32(2))%32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v54)+100))
	if v142 <= v141 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+108)) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v54)+104))
	v138 = F_palloc(m, v135<<(uint(int32(2))%32))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+112)) = v138
	goto L61
L64:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v145 = F_CreateTupleDescCopy(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L67
	}
L65:
	;
	v173 = v141
	goto L66
L66:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v54)+108))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v173<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+32))
	m.T0[v181].(func(*base.Module, int32, int32))(m, v179, v55)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L3
	} else {
		goto L71
	}
L67:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v148 = F_CreateTupleDescCopy(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v151 = F_MakeSingleTupleTableSlot(m, v145, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v54)+108))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v153+v154<<(uint(int32(2))%32)))) = v151
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v160 = F_MakeSingleTupleTableSlot(m, v148, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v54)+112))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v162+v163<<(uint(int32(2))%32)))) = v160
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v54)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+100)) = v168 + int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v173 = v172
	goto L66
L71:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v54)+112))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184+v185<<(uint(int32(2))%32))))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+32))
	m.T0[v191].(func(*base.Module, int32, int32))(m, v189, v25)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	if v115 == v112 {
		v205 = v194
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+96)) = v205 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v122
	v946 = int32(0)
	goto L22
L74:
	;
	if v194 != 0 {
		v205 = v194
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
	v197 = F_lappend(m, v196, v54)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+188)) = v197
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v26)+192))
	v201 = F_lappend(m, v200, v27)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+192)) = v201
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v205 = v204
	goto L73
L78:
	;
	if v214 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v946 = int32(0)
	goto L22
L80:
	;
	goto L81
L81:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+36)) = v220
	v770 = v214
	v774 = int32(0)
	goto L27
L82:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	switch v235 - int32(2) {
	case 0:
		v247 = v235
		goto L86
	default:
		goto L87
	case 3:
		goto L88
	}
L83:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+17)))
	if v229 != int32(1) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	F_ExecComputeStoredGenerated(m, v54, v26, v55, int32(3))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v54)+116))
	if v248 != 0 {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	v247 = int32(1)
	goto L86
L88:
	;
	v238 = int32(2)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v27)+216))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	if v242 == v238 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v245 = v238
	goto L91
L90:
	;
	v245 = int32(1)
	goto L91
L91:
	;
	v247 = v245
	goto L86
L92:
	;
	F_ExecWithCheckOptions(m, v247, v54, v55, v26)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L3
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	if v252 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	F_ExecConstraints(m, v54, v55, v26)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L3
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v63)+48))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+131)))
	if v256 != int32(1) {
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
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v54)+200))
	if v259 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	if v260 == int32(0) {
		goto L100
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v268 = F_ExecPartitionCheck(m, v54, v55, v26, int32(1))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L107
	}
L105:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+8)))
	if v263 != int32(1) {
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
	v750 = int32(0)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v63)+188))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v754)+80))
	m.T0[v755].(func(*base.Module, int32, int32, int32, int32, int32))(m, v63, v55, v751, v750, v750)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L3
	} else {
		goto L236
	}
L109:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v273 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v276 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+30)) = uint16(v276)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+26)) = int32(-1)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
	goto L111
L111:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[1]))
	if v304 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L3
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)) = uint8(v307)
	v313 = F_ExecCheckIndexConstraints(m, v54, v55, v26, v21+int32(-32), v21+int32(-38), v280)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	if v313 == int32(0) {
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
	v673 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L3
	} else {
		goto L224
	}
L121:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v54)+160))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)+16))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+64))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v326 = F_ExecUpdateLockMode(m, v325, v54)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L3
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v642 = int32(0)
	v643 = F_ExecGetReturningSlot(m, v26, v54)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L3
	} else {
		goto L214
	}
L124:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+8))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v330)+64))
	v333 = int32(0)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v324)+188))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+104))
	v339 = m.T0[v338].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v324, v21+int32(-32), v331, v323, v332, v326, v333, v333, v21+int32(-24))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L3
	} else {
		goto L132
	}
L125:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+12))
	m.T0[v639].(func(*base.Module, int32))(m, v323)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L3
	} else {
		goto L213
	}
L126:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecCheckTupleVisible(m, v547, v324, v323)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L3
	} else {
		goto L194
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L3
	} else {
		goto L191
	}
L128:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[2]))
	if v515 < int32(2) {
		goto L125
	} else {
		goto L186
	}
L129:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[2]))
	if v495 < int32(2) {
		goto L125
	} else {
		goto L181
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L3
	} else {
		goto L178
	}
L131:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+20))
	v346 = m.T0[v345].(func(*base.Module, int32, int32, int32) int32)(m, v323, int32(-2), v21+int32(-25))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L3
	} else {
		goto L133
	}
L132:
	;
	switch v339 {
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
	if base.Ui32(v346) < base.Ui32(int32(3)) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L3
	} else {
		goto L174
	}
L135:
	;
	v467 = int32(0)
	goto L134
L136:
	;
	goto L137
L137:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[3]))
	if v358 == v346 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v467 = int32(1)
	goto L134
L139:
	;
	goto L140
L140:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[4]))
	if v362 <= int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v467 = v459
	goto L134
L142:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[5]))
	if v366 == int32(0) {
		v459 = int32(0)
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[6]))
	v430 = int32(0)
	v432 = v362 - int32(1)
	goto L164
L145:
	;
	v371 = v366
	goto L146
L146:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v371)+20))
	if v376 == int32(4) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v459 = int32(0)
	goto L141
L148:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v371)+80))
	if v423 != 0 {
		v371 = v423
		goto L146
	} else {
		goto L163
	}
L149:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	if v379 == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v382 = int32(1)
	if v346 == v379 {
		v459 = v382
		goto L141
	} else {
		goto L151
	}
L151:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v371)+52))
	v386 = v384 - int32(1)
	if v386 < int32(0) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v391 = int32(0)
	v393 = v386
	goto L153
L153:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v371)+48))
	v399 = int32(2)
	v400 = base.I32_div_s(v393-v391, v399)
	v401 = v400 + v391
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v397+v401<<(uint(v399)%32))))
	if v405 == v346 {
		v459 = v382
		goto L141
	} else {
		goto L155
	}
L154:
	;
	goto L148
L155:
	;
	v409 = F_TransactionIdPrecedes(m, v405, v346)
	mBase = m.M
	if v409 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v410 = v401 + int32(1)
	goto L158
L157:
	;
	v410 = v391
	goto L158
L158:
	;
	if v409 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v413 = v393
	goto L161
L160:
	;
	v413 = v401 - int32(1)
	goto L161
L161:
	;
	if v410 <= v413 {
		v391 = v410
		v393 = v413
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
	v437 = int32(2)
	v438 = base.I32_div_s(v432-v430, v437)
	v439 = v438 + v430
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v428+v439<<(uint(v437)%32))))
	v444 = base.B2i32(v443 == v346)
	if v443 == v346 {
		v459 = v444
		goto L141
	} else {
		goto L166
	}
L165:
	;
	v459 = v444
	goto L141
L166:
	;
	v447 = base.B2i32(base.Ui32(v443) < base.Ui32(v346))
	if base.Ui32(v443) < base.Ui32(v346) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v448 = v439 + int32(1)
	goto L169
L168:
	;
	v448 = v430
	goto L169
L169:
	;
	if base.Ui32(v443) < base.Ui32(v346) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v451 = v432
	goto L172
L171:
	;
	v451 = v439 - int32(1)
	goto L172
L172:
	;
	if v448 <= v451 {
		v430 = v448
		v432 = v451
		goto L164
	} else {
		goto L173
	}
L173:
	;
	goto L165
L174:
	;
	if v467 != 0 {
		goto L26
	} else {
		goto L175
	}
L175:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInsert_1), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L3
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2793), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
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
	v488 = m.ExcPending
	if v488 != 0 {
		goto L3
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2803), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
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
	v501 = m.ExcPending
	if v501 != 0 {
		goto L3
	} else {
		goto L182
	}
L182:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	F_errmsg(m, int32(_a_F_ExecInsert_5), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L3
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2810), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
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
	v521 = m.ExcPending
	if v521 != 0 {
		goto L3
	} else {
		goto L187
	}
L187:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L3
	} else {
		goto L188
	}
L188:
	;
	F_errmsg(m, int32(_a_F_ExecInsert_6), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2826), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v339
	F_errmsg_internal(m, int32(_a_F_ExecInsert_7), v23)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2833), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
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
	v550 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v322)+12)) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v322)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v322)+4)) = v323
	if v320 == v550 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v54)+116))
	if v582 != 0 {
		goto L201
	} else {
		goto L202
	}
L196:
	;
	v556 = int32(_a_F_ExecInsert_0)
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v322)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v559
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v320)+20))
	v564 = m.T0[v563].(func(*base.Module, int32, int32, int32) int32)(m, v320, v322, v21+int32(-1))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L3
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v557
	if v564 != 0 {
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v568)+12))
	m.T0[v569].(func(*base.Module, int32))(m, v323)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L3
	} else {
		goto L199
	}
L199:
	;
	v572 = int32(0)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v321)+20))
	if v573 == v572 {
		v924 = v572
		goto L23
	} else {
		goto L200
	}
L200:
	;
	v576 = *(*float64)(unsafe.Add(mBase, uint32(v573)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v573)+240)) = base.F64_add(v576, float64(1))
	v924 = v572
	goto L23
L201:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	F_ExecWithCheckOptions(m, int32(3), v54, v323, v584)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L3
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v54)+160))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+72))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v588)+16))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+8))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v591)+12))
	m.T0[v592].(func(*base.Module, int32))(m, v590)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L3
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	v595 = int32(_a_F_ExecInsert_0)
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v589)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v598
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v588)+24))
	v604 = m.T0[v603].(func(*base.Module, int32, int32, int32) int32)(m, v588+int32(4), v589, int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L3
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v596
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v590)+4)))
	v610 = v608 & int32(_a_F_ExecInsert_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v590)+4)) = uint16(v610)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v590)+12))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	*(*uint16)(unsafe.Add(mBase, uint32(v590)+6)) = uint16(v613)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v54)+160))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)+8))
	v620 = F_ExecUpdate(m, l0, v54, v21+int32(-32), int32(0), v323, v619, l3)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L3
	} else {
		goto L208
	}
L207:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+12))
	m.T0[v635].(func(*base.Module, int32))(m, v323)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L3
	} else {
		goto L212
	}
L208:
	;
	if v620 == int32(0) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v54)+152))
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+8)))
	if v625&int32(2) == int32(0) {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v620)+8))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+28))
	m.T0[v631].(func(*base.Module, int32))(m, v620)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L3
	} else {
		goto L211
	}
L211:
	;
	goto L207
L212:
	;
	v924 = v620
	goto L23
L213:
	;
	goto L111
L214:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[2]))
	if v646 < int32(2) {
		v924 = v642
		goto L23
	} else {
		goto L215
	}
L215:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[7]))
	if v651 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecInsert[8])))
	if v653&int32(1) == int32(0) {
		goto L25
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v649)+188))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)+60))
	v663 = m.T0[v662].(func(*base.Module, int32, int32, int32, int32) int32)(m, v649, v21+int32(-32), int32(_a_F_ExecInsert_9), v643)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L3
	} else {
		goto L220
	}
L219:
	;
	goto L218
L220:
	;
	if v663 == int32(0) {
		goto L24
	} else {
		goto L221
	}
L221:
	;
	F_ExecCheckTupleVisible(m, v26, v649, v643)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L3
	} else {
		goto L222
	}
L222:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v643)+8))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)+12))
	m.T0[v670].(func(*base.Module, int32))(m, v643)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L3
	} else {
		goto L223
	}
L223:
	;
	v924 = v642
	goto L23
L224:
	;
	v675 = m.G0
	v677 = v675 - int32(16)
	m.G0 = v677
	v679 = int32(_a_F_ExecInsert_10)
	v680 = int32(1)
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[9]))
	v684 = v682 + v680
	if base.Ui32(v684) <= base.Ui32(v680) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v687 = v680
	goto L227
L226:
	;
	v687 = v684
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[9])) = v687
	*(*int64)(unsafe.Add(mBase, uint32(v677)+8)) = int64(74027918874902528)
	*(*int32)(unsafe.Add(mBase, uint32(v677))) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v677)+4)) = v687
	v694 = int32(0)
	v696 = F_LockAcquire(m, v677, int32(7), v694, v694)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L3
	} else {
		goto L228
	}
L228:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[9]))
	m.G0 = v677 + int32(16)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v26)+64))
	v704 = int32(0)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v63)+188))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+84))
	m.T0[v707].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v63, v55, v703, v704, v704, v699)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L3
	} else {
		goto L229
	}
L229:
	;
	v710 = int32(0)
	v715 = F_ExecInsertIndexTuples(m, v54, v55, v26, v710, int32(1), v21+int32(-39), v280, v710)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L3
	} else {
		goto L230
	}
L230:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v63)+188))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+88))
	m.T0[v723].(func(*base.Module, int32, int32, int32, int32))(m, v63, v55, v699, (v717^int32(-1))&int32(1))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L3
	} else {
		goto L231
	}
L231:
	;
	v726 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	v728 = m.G0
	v730 = v728 - int32(16)
	m.G0 = v730
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v726
	*(*int64)(unsafe.Add(mBase, uint32(v730)+8)) = int64(74027918874902528)
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v730)+4)) = v736
	v740 = F_LockRelease(m, v730, int32(7), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L3
	} else {
		goto L233
	}
L233:
	;
	m.G0 = v730 + int32(16)
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+25)))
	if v745 != int32(1) {
		v770 = v55
		v774 = v715
		goto L27
	} else {
		goto L234
	}
L234:
	;
	F_list_free(m, v715)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L3
	} else {
		goto L235
	}
L235:
	;
	goto L111
L236:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v758 <= int32(0) {
		v770 = v55
		v774 = v750
		goto L27
	} else {
		goto L237
	}
L237:
	;
	v761 = int32(0)
	v766 = F_ExecInsertIndexTuples(m, v54, v55, v26, v761, v761, v761, v761, v761)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L3
	} else {
		goto L238
	}
L238:
	;
	v770 = v55
	v774 = v766
	goto L27
L239:
	;
	v788 = *(*int64)(unsafe.Add(mBase, uint32(v26)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+112)) = v788 + int64(1)
	goto L241
L240:
	;
	goto L241
L241:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v27)+204))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	if v793 != int32(2) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	F_ExecARInsertTriggers(m, v26, v54, v770, v774, v811)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L3
	} else {
		goto L253
	}
L243:
	;
	v811 = v792
	goto L242
L244:
	;
	goto L245
L245:
	;
	if v792 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v811 = int32(0)
	goto L242
L247:
	;
	goto L248
L248:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v792)+2)))
	if v799 != int32(1) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v811 = v792
	goto L242
L250:
	;
	goto L251
L251:
	;
	v802 = int32(0)
	F_ExecARUpdateTriggers(m, v26, v54, v802, v802, v802, v802, v770, v802, v792, v802)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L3
	} else {
		goto L252
	}
L252:
	;
	v811 = v802
	goto L242
L253:
	;
	F_list_free(m, v774)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L3
	} else {
		goto L254
	}
L254:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v54)+116))
	if v816 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	F_ExecWithCheckOptions(m, int32(0), v54, v770, v26)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L3
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v54)+152))
	if v820 == int32(0) {
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
	v870 = int32(0)
	goto L259
L261:
	;
	goto L262
L262:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v824 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	v849 = F_ExecProcessReturning(m, l0, v54, int32(3), v847, v770, v25)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L3
	} else {
		goto L271
	}
L264:
	;
	v847 = int32(0)
	goto L263
L265:
	;
	goto L266
L266:
	;
	v828 = F_ExecGetRootToChildMap(m, v54, v26)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L3
	} else {
		goto L267
	}
L267:
	;
	if v828 == int32(0) {
		v847 = v824
		goto L263
	} else {
		goto L268
	}
L268:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v828)+8))
	v833 = F_ExecGetReturningSlot(m, v26, v54)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L3
	} else {
		goto L269
	}
L269:
	;
	v835 = F_execute_attr_map_slot(m, v832, v824, v833)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L3
	} else {
		goto L270
	}
L270:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v837)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v835)+36)) = v838
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v835)+28)) = v841
	v843 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v840)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v835)+32)) = uint16(v843)
	v847 = v835
	goto L263
L271:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v851 == int32(0) {
		v870 = v849
		goto L259
	} else {
		goto L272
	}
L272:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v849)+8))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+28))
	m.T0[v855].(func(*base.Module, int32))(m, v849)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L3
	} else {
		goto L273
	}
L273:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v847)+8))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)+12))
	m.T0[v859].(func(*base.Module, int32))(m, v847)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L3
	} else {
		goto L274
	}
L274:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v847 != v862 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v862)+8))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v864)+12))
	m.T0[v865].(func(*base.Module, int32))(m, v862)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
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
	v870 = v849
	goto L259
L278:
	;
	goto L277
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v770
	goto L281
L280:
	;
	goto L281
L281:
	;
	if l5 == int32(0) {
		v946 = v870
		goto L22
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v54
	v946 = v870
	goto L22
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(_a_F_ExecInsert_11)
	F_errmsg(m, int32(_a_F_ExecInsert_12), v21+int32(-48))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L3
	} else {
		goto L284
	}
L284:
	;
	F_errhint(m, int32(_a_F_ExecInsert_13), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L3
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(2790), int32(_a_F_ExecInsert_3))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
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
	v903 = m.ExcPending
	if v903 != 0 {
		goto L3
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_15), int32(1264), int32(_a_F_ExecInsert_16))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
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
	v916 = m.ExcPending
	if v916 != 0 {
		goto L3
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_2), int32(409), int32(_a_F_ExecInsert_18))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
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
	v936 = *(*float64)(unsafe.Add(mBase, uint32(v933)+224))
	*(*float64)(unsafe.Add(mBase, uint32(v933)+224)) = base.F64_add(v936, float64(1))
	v946 = v924
	goto L22
}
func F_ExecNamedTuplestoreScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(738), int32(739))
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
	v4 = F_ExecScan(m, l0, int32(745), int32(746))
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
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
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
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
	var v286 int32
	_ = v286
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
	return v286
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
	goto L45
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
	v112 = m.T0[l1].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L44
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
		v286 = v46
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
	v286 = v46
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
	v286 = v60
	goto L3
L25:
	;
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v70)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	if v72&int32(2) != 0 {
		v286 = v4
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
	v286 = v69
	goto L3
L33:
	;
	goto L32
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v89)
	v91 = F_EvalPlanQualFetchRowMark(m, v18, v37, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v88 == int32(0) {
		v286 = v4
		goto L3
	} else {
		goto L36
	}
L36:
	;
	if v91 == int32(0) {
		v286 = v4
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
	if v97&int32(2) != 0 {
		v286 = v4
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v100 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v88)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v100 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	m.T0[v105].(func(*base.Module, int32))(m, v88)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v286 = v88
	goto L3
L43:
	;
	goto L42
L44:
	;
	v286 = v112
	goto L3
L45:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[0]))
	if v127 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v18 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v209
	if v20 != 0 {
		goto L89
	} else {
		goto L90
	}
L52:
	;
	if v19 == int32(0) {
		v286 = v219
		goto L3
	} else {
		goto L86
	}
L53:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
	if v213&int32(2) == int32(0) {
		goto L51
	} else {
		goto L85
	}
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	m.T0[v206].(func(*base.Module, int32))(m, v201)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L84
	}
L55:
	;
	v199 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v157)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L82
	}
L56:
	;
	if v194 != 0 {
		v209 = v194
		goto L53
	} else {
		goto L81
	}
L57:
	;
	v192 = m.T0[l1].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L80
	}
L58:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+72))
	if v133 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+64))
	v138 = F_bms_is_member(m, v136, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v145 = int32(1)
	v146 = v133 - v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v148 = v146 + v147
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v149 == v145 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	if v138 == int32(0) {
		goto L57
	} else {
		goto L63
	}
L63:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v143 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v143 != 0 {
		v194 = v142
		goto L56
	} else {
		goto L65
	}
L65:
	;
	v201 = v142
	goto L54
L66:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v201 = v152
	goto L54
L67:
	;
	goto L68
L68:
	;
	v154 = v146 << (uint(int32(2)) % 32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v154+v155)))
	if v157 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v158)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
	if v160&int32(2) == int32(0) {
		goto L55
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166+v154)))
	if v168 == int32(0) {
		goto L57
	} else {
		goto L73
	}
L72:
	;
	v219 = int32(0)
	goto L52
L73:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v172 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v172)
	v174 = int32(0)
	v175 = F_EvalPlanQualFetchRowMark(m, v18, v133, v171)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v171 == int32(0) {
		v219 = v174
		goto L52
	} else {
		goto L75
	}
L75:
	;
	if v175 == int32(0) {
		v219 = v174
		goto L52
	} else {
		goto L76
	}
L76:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+4)))
	if v181&int32(2) != 0 {
		v219 = v174
		goto L52
	} else {
		goto L77
	}
L77:
	;
	v184 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v171)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v184 == int32(0) {
		v201 = v171
		goto L54
	} else {
		goto L79
	}
L79:
	;
	v209 = v171
	goto L53
L80:
	;
	v194 = v192
	goto L56
L81:
	;
	v219 = int32(0)
	goto L52
L82:
	;
	if v199 != 0 {
		v209 = v157
		goto L53
	} else {
		goto L83
	}
L83:
	;
	v201 = v157
	goto L54
L84:
	;
	v209 = v201
	goto L53
L85:
	;
	v219 = v209
	goto L52
L86:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	m.T0[v226].(func(*base.Module, int32))(m, v224)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v286 = v224
	goto L3
L88:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v274 != 0 {
		goto L97
	} else {
		goto L98
	}
L89:
	;
	v230 = int32(_a_F_ExecScan_0)
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1]))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v233
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v238 = m.T0[v237].(func(*base.Module, int32, int32, int32) int32)(m, v20, v21, v15+int32(15))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if v19 == int32(0) {
		v286 = v209
		goto L3
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v231
	if v238 == int32(0) {
		goto L88
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
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
		goto L95
	}
L95:
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
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v255
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249)+4)))
	v269 = v267 & int32(_a_F_ExecScan_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+4)) = uint16(v269)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v249)+12))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+6)) = uint16(v272)
	v286 = v249
	goto L3
L97:
	;
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v274)+240)) = base.F64_add(v275, float64(1))
	goto L99
L98:
	;
	goto L99
L99:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	F_MemoryContextReset(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	goto L45
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
	v21 = F_planstate_tree_walker_impl(m, l0, int32(633), l1)
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
	var v13 int32
	_ = v13
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	return v44
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = F_equal(m, v38, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L3:
	;
	if v6 == int32(0) {
		v44 = v3
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
		v44 = v3
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v32 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v17 = v7
	v18 = v6
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v32 = v21
	v33 = v22
	goto L8
L13:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v44 = v3
	goto L1
L16:
	;
	goto L2
L17:
	;
	return int32(0)
L18:
	;
	v44 = v40
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	return v54
L2:
	;
	return int32(0)
L3:
	;
	if v8 == int32(0) {
		v54 = v3
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
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v46 != v47 {
		v54 = v3
		goto L1
	} else {
		goto L20
	}
L6:
	;
	if v14 == int32(0) {
		v54 = v3
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
		v54 = v3
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v21 == int32(0) {
		v40 = v20
		v41 = v21
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v41-v40 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v20 != v21 {
		v40 = v20
		v41 = v21
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v25 = v15
	v26 = v14
	goto L14
L14:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v40 = v29
		v41 = v30
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v40 = v29
	v41 = v30
	goto L11
L16:
	;
	v33 = int32(1)
	if v29 == v30 {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v54 = v3
	goto L1
L19:
	;
	goto L5
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v51 = F_equal(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v54 = v51
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
					*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(1362)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(1363)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(1364)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(1365)
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
							v56 = int32(4)
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
							if v58&int32(254) == int32(2) {
								v67 = v56
							} else {
								v67 = base.B2i32(v58 == int32(18)) << (uint(v56) % 32)
							}
							if v58 == int32(1) {
								v70 = v56
							} else {
								v70 = v67
							}
							v81 = v70
						} else {
							v71 = int32(1)
							if v52 != 0 {
								v81 = int32(base.Ui32(v50)>>(uint(v71)%32)) - v71
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
								v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v83 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker[0]))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
						v86 = F_makeJsonLexContextCstringLen(m, v11+int32(12), v53, v81, v84, int32(1))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v17))) = v86
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker[1]))
							v95 = F_AllocSetContextCreateInternal(m, v90, int32(_a_F_each_worker_0), int32(0), int32(_a_F_each_worker_1), int32(_a_F_each_worker_2))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v95
								v100 = F_pg_parse_json(m, v11+int32(12), v20)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									if v100 != 0 {
										F_json_errsave_error(m, v100, v11+int32(12), int32(0))
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return
										} else {
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
											F_MemoryContextDelete(m, v107)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												F_freeJsonLexContext(m, v11+int32(12))
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return
												} else {
													v114 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v114)
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
										F_MemoryContextDelete(m, v107)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											F_freeJsonLexContext(m, v11+int32(12))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return
											} else {
												v114 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v114)
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(_a_F_errcontext_msg_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[0])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[1]))
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(_a_F_errcontext_msg_1)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[2]))
	v25 = v18 * int32(100)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errcontext_msg[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[2])) = v28
	F_initStringInfo(m, v9+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L26
	}
L4:
	;
	return
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errcontext_msg[4])))
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_appendStringInfoString(m, v9+int32(16), v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errcontext_msg[5])))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[6])) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v50 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	F_appendStringInfoChar(m, v9+int32(16), int32(10))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	if v50 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v55 = v50
	goto L15
L13:
	;
	goto L14
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errcontext_msg[4])))
	if v76 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	F_enlargeStringInfo(m, v9+int32(16), v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errcontext_msg[5])))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[6])) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v68 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v68 != 0 {
		v55 = v68
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	F_pfree(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v80 = F_pstrdup(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errcontext_msg[4]))) = v80
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_pfree(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[2])) = v22
	v88 = int32(_a_F_errcontext_msg_0)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errcontext_msg[0])) = v90 - int32(1)
	m.G0 = v9 + int32(32)
	return
L26:
	;
	F_errmsg_internal(m, int32(_a_F_errcontext_msg_2), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_errcontext_msg_3), int32(1393), int32(_a_F_errcontext_msg_4))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
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
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
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
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
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
	v312 = F_fflush(m, int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L65
	} else {
		goto L98
	}
L2:
	;
	v294 = int32(_a_F_errfinish_0)
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[0])) = v296 - int32(1)
	v301 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[2])) = v301
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[3])) = v301
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[4])) = v301
	F_pg_re_throw(m)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L65
	} else {
		goto L97
	}
L3:
	;
	v17 = v13 * int32(100)
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
	v284 = m.ExcPending
	if v284 != 0 {
		goto L65
	} else {
		goto L94
	}
L6:
	;
	v23 = F_strlen(m, l0)
	mBase = m.M
	v30 = v23 + int32(1)
	goto L11
L7:
	;
	v73 = v4
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_errfinish[5]))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_errfinish[6]))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_errfinish[7]))) = v73
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_errfinish[8])))
	v78 = int32(_a_F_errfinish_1)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[9]))
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[9])) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_errfinish[11])))
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	goto L9
L11:
	;
	v32 = int32(0)
	if v30 == v32 {
		v42 = v32
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v42 = v37
	goto L10
L13:
	;
	v36 = v30 - int32(1)
	v37 = l0 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v38 != int32(47) {
		v30 = v36
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v45 = v42 + int32(1)
	goto L17
L16:
	;
	v45 = l0
	goto L17
L17:
	;
	v49 = F_strlen(m, v45)
	mBase = m.M
	v56 = v49 + int32(1)
	goto L20
L18:
	;
	if v68 != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	goto L18
L20:
	;
	v58 = int32(0)
	if v56 == v58 {
		v68 = v58
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v68 = v63
	goto L19
L22:
	;
	v62 = v56 - int32(1)
	v63 = v45 + v62
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v64 != int32(92) {
		v56 = v62
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v71 = v68 + int32(1)
	goto L26
L25:
	;
	v71 = v45
	goto L26
L26:
	;
	v73 = v71
	goto L8
L27:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[12]))
	if v212 != 0 {
		goto L68
	} else {
		goto L69
	}
L28:
	;
	if l2 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[13]))
	if v88 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[14]))
	if v92 == int32(0) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v95 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v98 == int32(0) {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v101 = v92
	goto L34
L34:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v109 == int32(0) {
		v128 = v108
		v129 = v109
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v192 = m.G0
	v194 = v192 - int32(16)
	m.G0 = v194
	F_initStringInfo(m, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L65
	} else {
		goto L66
	}
L36:
	;
	if v129-v128 != 0 {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	goto L36
L38:
	;
	if v108 != v109 {
		v128 = v108
		v129 = v109
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v113 = l2
	v114 = v101
	goto L40
L40:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if v118 == int32(0) {
		v128 = v117
		v129 = v118
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v128 = v117
	v129 = v118
	goto L37
L42:
	;
	v121 = int32(1)
	if v117 == v118 {
		v113 = v113 + v121
		v114 = v114 + v121
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v101&int32(3) == int32(0) {
		v154 = v101
		goto L49
	} else {
		goto L50
	}
L45:
	;
	goto L46
L46:
	;
	goto L35
L47:
	;
	v190 = v187 + v101 + int32(1)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v191 != 0 {
		v101 = v190
		goto L34
	} else {
		goto L64
	}
L48:
	;
	v187 = v179 - v101
	goto L47
L49:
	;
	v158 = v154
	goto L58
L50:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v138 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v187 = int32(0)
	goto L47
L52:
	;
	goto L53
L53:
	;
	v143 = v101
	goto L54
L54:
	;
	v147 = v143 + int32(1)
	if v147&int32(3) == int32(0) {
		v154 = v147
		goto L49
	} else {
		goto L56
	}
L55:
	;
	v179 = v147
	goto L48
L56:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v152 != 0 {
		v143 = v147
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v167 = int32(-2139062144)
	if (int32(16843008)-v164|v164)&v167 == v167 {
		v158 = v158 + int32(4)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v173 = v158
	goto L61
L60:
	;
	goto L59
L61:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v177 != 0 {
		v173 = v173 + int32(1)
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v179 = v173
	goto L48
L63:
	;
	goto L62
L64:
	;
	goto L27
L65:
	;
	return
L66:
	;
	F_appendStringInfoString(m, v194, int32(_a_F_errfinish_2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_errfinish[11]))) = v201
	m.G0 = v194 + int32(16)
	goto L27
L68:
	;
	v213 = v212
	goto L71
L69:
	;
	goto L70
L70:
	;
	if v77 == int32(21) {
		goto L2
	} else {
		goto L75
	}
L71:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	m.T0[v219].(func(*base.Module, int32))(m, v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L65
	} else {
		goto L73
	}
L72:
	;
	goto L70
L73:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v222 != 0 {
		v213 = v222
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	F_EmitErrorReport(m)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L65
	} else {
		goto L76
	}
L76:
	;
	F_FreeErrorDataContents(m, v17+int32(_a_F_errfinish_3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L65
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[9])) = v79
	v236 = int32(_a_F_errfinish_4)
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[1]))
	v239 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[1])) = v238 - v239
	v242 = int32(_a_F_errfinish_0)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[0])) = v244 - v239
	if v77 == int32(22) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[15]))
	if v251 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	if int32(23) <= v77 {
		goto L1
	} else {
		goto L89
	}
L81:
	;
	v260 = F_fflush(m, int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L65
	} else {
		goto L84
	}
L82:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[16]))
	if v253 != int32(2) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[16])) = int32(0)
	goto L81
L84:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[17]))
	if v263 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[17])) = int32(3)
	goto L87
L86:
	;
	goto L87
L87:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L65
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[18]))
	if v275 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L65
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	return
L93:
	;
	goto L92
L94:
	;
	F_errmsg_internal(m, int32(_a_F_errfinish_5), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L65
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_errfinish_6), int32(482), int32(_a_F_errfinish_7))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L65
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_abort(m)
	mBase = m.M
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errhint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(_a_F_errhint_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_errhint[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint[0])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_errhint[1]))
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(_a_F_errhint_1)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_errhint[2]))
	v25 = v18 * int32(100)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errhint[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint[2])) = v28
	F_initStringInfo(m, v9+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errhint[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errhint[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint[5])) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v40 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = v40
	goto L10
L8:
	;
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errhint[6])))
	if v66 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	F_enlargeStringInfo(m, v9+int32(16), v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errhint[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint[5])) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v58 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v58 != 0 {
		v45 = v58
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v70 = F_pstrdup(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_errhint[6]))) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_pfree(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errhint[2])) = v22
	v78 = int32(_a_F_errhint_0)
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_errhint[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint[0])) = v80 - int32(1)
	m.G0 = v9 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(_a_F_errhint_2), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_errhint_3), int32(1324), int32(_a_F_errhint_4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v384 int32
	_ = v384
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
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v918 int32
	_ = v918
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1115 int32
	_ = v1115
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1150 int32
	_ = v1150
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1467 int32
	_ = v1467
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1502 int32
	_ = v1502
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1650 int32
	_ = v1650
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1669 int32
	_ = v1669
	var v1685 int32
	_ = v1685
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1743 int32
	_ = v1743
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1769 int32
	_ = v1769
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1816 int32
	_ = v1816
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1835 int32
	_ = v1835
	var v1851 int32
	_ = v1851
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1903 int32
	_ = v1903
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1973 int32
	_ = v1973
	var v1991 int32
	_ = v1991
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2004 int32
	_ = v2004
	var v2010 int32
	_ = v2010
	var v2026 int32
	_ = v2026
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
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
	return v2088
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
	v2088 = v25
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
	v2088 = v31
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
	v2088 = v37
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
	v2088 = v43
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
	v2088 = v49
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
	v2088 = v55
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
	v2088 = v61
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
	v2088 = v67
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
	v2088 = v73
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
	v2088 = v79
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
	v2088 = v85
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
	v2088 = v91
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
	v2088 = v97
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
	v2088 = v103
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
	v2088 = v109
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
	v2088 = v115
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
	v2088 = v121
	goto L1
L62:
	;
	if v127 < int32(0) {
		v2088 = v127
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
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
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
	v384 = v258
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
		v384 = v377
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
	v384 = v377
	goto L64
L118:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v752
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v754)))
	if v752 < v755 {
		goto L205
	} else {
		goto L206
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
	if v461 < int32(0) {
		goto L118
	} else {
		goto L145
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
	v461 = v443
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
	v461 = int32(-1)
	goto L125
L135:
	;
	goto L136
L136:
	;
	v424 = v417 - int32(1)
	v426 = int32(*(*int8)(unsafe.Add(mBase, uint32(v410+v424))))
	if int32(0) <= v426 {
		v443 = v424
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v447 = int32(1)
	if v447 < v419 {
		v417 = v443
		v419 = v419 - v447
		goto L132
	} else {
		goto L144
	}
L138:
	;
	if v424 <= v6 {
		v443 = v424
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v431 = v424
	goto L140
L140:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v431))))
	if base.Ui32(int32(191)) < base.Ui32(v436) {
		v443 = v431
		goto L137
	} else {
		goto L142
	}
L141:
	;
	v443 = v6
	goto L137
L142:
	;
	v440 = v431 - int32(1)
	if v6 < v440 {
		v431 = v440
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	goto L133
L145:
	;
	v464 = v407 - v409
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v464 + v465
	switch v402 - int32(1) {
	case 0:
		goto L147
	case 1:
		goto L146
	default:
		goto L118
	}
L146:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L179
L147:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L150
L148:
	;
	if v598 != 0 {
		goto L118
	} else {
		goto L172
	}
L149:
	;
	v598 = v591
	goto L148
L150:
	;
	if v486 <= v487 {
		v591 = int32(-1)
		goto L149
	} else {
		goto L152
	}
L151:
	;
	v591 = int32(0)
	goto L149
L152:
	;
	v504 = int32(1)
	v505 = v486 - v504
	v507 = int32(*(*int8)(unsafe.Add(mBase, uint32(v483+v505))))
	v509 = v507 & int32(255)
	if v505 == v487 {
		v564 = v509
		v565 = v504
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if int32(252) < v564 {
		goto L162
	} else {
		goto L163
	}
L154:
	;
	if int32(0) <= v507 {
		v564 = v509
		v565 = v504
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v515 = v509 & int32(63)
	v517 = v486 - int32(2)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483+v517))))
	v521 = v519 << (uint(int32(6)) % 32)
	if base.B2i32(v517 != v487)&base.B2i32(base.Ui32(v519) < base.Ui32(int32(192))) == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v564 = v521&int32(1984) | v515
	v565 = int32(2)
	goto L153
L157:
	;
	goto L158
L158:
	;
	v534 = v521&int32(4032) | v515
	v536 = v486 - int32(3)
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483+v536))))
	if base.B2i32(v536 != v487)&base.B2i32(base.Ui32(v538) < base.Ui32(int32(224))) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v564 = v538<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v534
	v565 = int32(3)
	goto L153
L160:
	;
	goto L161
L161:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486+(v483-int32(4))))))
	v564 = v538<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v556&int32(7)<<(uint(int32(18))%32) | v534
	v565 = int32(4)
	goto L153
L162:
	;
	v598 = v565
	goto L148
L163:
	;
	goto L164
L164:
	;
	v569 = v564 - int32(97)
	if v569 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v598 = v565
	goto L148
L166:
	;
	goto L167
L167:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v569)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[1]))))
	if int32(base.Ui32(v575)>>(uint(v569&int32(7))%32))&int32(1) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v598 = v565
	goto L148
L169:
	;
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v486 - v565
	goto L171
L171:
	;
	goto L151
L172:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v599 + v464
	v604 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	if v604 != 0 {
		goto L118
	} else {
		goto L174
	}
L174:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v606 + v464
	v609 = F_slice_del(m, l0)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L3
	} else {
		goto L175
	}
L175:
	;
	if int32(0) <= v609 {
		goto L118
	} else {
		goto L176
	}
L176:
	;
	v2088 = v609
	goto L1
L177:
	;
	if v741 != 0 {
		goto L118
	} else {
		goto L201
	}
L178:
	;
	v741 = v734
	goto L177
L179:
	;
	if v629 <= v630 {
		v734 = int32(-1)
		goto L178
	} else {
		goto L181
	}
L180:
	;
	v734 = int32(0)
	goto L178
L181:
	;
	v647 = int32(1)
	v648 = v629 - v647
	v650 = int32(*(*int8)(unsafe.Add(mBase, uint32(v626+v648))))
	v652 = v650 & int32(255)
	if v648 == v630 {
		v707 = v652
		v708 = v647
		goto L182
	} else {
		goto L183
	}
L182:
	;
	if int32(382) < v707 {
		goto L191
	} else {
		goto L192
	}
L183:
	;
	if int32(0) <= v650 {
		v707 = v652
		v708 = v647
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v658 = v652 & int32(63)
	v660 = v629 - int32(2)
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626+v660))))
	v664 = v662 << (uint(int32(6)) % 32)
	if base.B2i32(v660 != v630)&base.B2i32(base.Ui32(v662) < base.Ui32(int32(192))) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v707 = v664&int32(1984) | v658
	v708 = int32(2)
	goto L182
L186:
	;
	goto L187
L187:
	;
	v677 = v664&int32(4032) | v658
	v679 = v629 - int32(3)
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626+v679))))
	if base.B2i32(v679 != v630)&base.B2i32(base.Ui32(v681) < base.Ui32(int32(224))) == int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v707 = v681<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v677
	v708 = int32(3)
	goto L182
L189:
	;
	goto L190
L190:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+(v626-int32(4))))))
	v707 = v681<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v699&int32(7)<<(uint(int32(18))%32) | v677
	v708 = int32(4)
	goto L182
L191:
	;
	v741 = v708
	goto L177
L192:
	;
	goto L193
L193:
	;
	v712 = v707 - int32(98)
	if v712 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v741 = v708
	goto L177
L195:
	;
	goto L196
L196:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v712)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[2]))))
	if int32(base.Ui32(v718)>>(uint(v712&int32(7))%32))&int32(1) == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v741 = v708
	goto L177
L198:
	;
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v629 - v708
	goto L200
L200:
	;
	goto L180
L201:
	;
	v742 = F_slice_del(m, l0)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L3
	} else {
		goto L202
	}
L202:
	;
	if int32(0) <= v742 {
		goto L118
	} else {
		goto L203
	}
L203:
	;
	v2088 = v742
	goto L1
L204:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1903
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L491
L205:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v934
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v936)))
	if v934 < v937 {
		goto L247
	} else {
		goto L248
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v752
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v755
	if v752 <= v755 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v758
	goto L205
L208:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761+v752-int32(1)))))
	if v765&int32(224) != int32(96) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	if int32(1)<<(uint(v765)%32)&int32(_a_F_estonian_UTF_8_stem_24) == int32(0) {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v778 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_25), int32(21))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L3
	} else {
		goto L211
	}
L211:
	;
	if v778 == int32(0) {
		goto L207
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v758
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v783
	switch v778 - int32(1) {
	case 0:
		goto L215
	case 1:
		goto L214
	case 2:
		goto L213
	default:
		goto L204
	}
L213:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L222
L214:
	;
	v793 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_26))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L3
	} else {
		goto L218
	}
L215:
	;
	v787 = F_slice_del(m, l0)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L3
	} else {
		goto L216
	}
L216:
	;
	if int32(0) <= v787 {
		goto L204
	} else {
		goto L217
	}
L217:
	;
	v2088 = v787
	goto L1
L218:
	;
	if int32(0) <= v793 {
		goto L204
	} else {
		goto L219
	}
L219:
	;
	v2088 = v793
	goto L1
L220:
	;
	if v925 != 0 {
		goto L205
	} else {
		goto L244
	}
L221:
	;
	v925 = v918
	goto L220
L222:
	;
	if v813 <= v814 {
		v918 = int32(-1)
		goto L221
	} else {
		goto L224
	}
L223:
	;
	v918 = int32(0)
	goto L221
L224:
	;
	v831 = int32(1)
	v832 = v813 - v831
	v834 = int32(*(*int8)(unsafe.Add(mBase, uint32(v810+v832))))
	v836 = v834 & int32(255)
	if v832 == v814 {
		v891 = v836
		v892 = v831
		goto L225
	} else {
		goto L226
	}
L225:
	;
	if int32(252) < v891 {
		goto L234
	} else {
		goto L235
	}
L226:
	;
	if int32(0) <= v834 {
		v891 = v836
		v892 = v831
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v842 = v836 & int32(63)
	v844 = v813 - int32(2)
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810+v844))))
	v848 = v846 << (uint(int32(6)) % 32)
	if base.B2i32(v844 != v814)&base.B2i32(base.Ui32(v846) < base.Ui32(int32(192))) == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v891 = v848&int32(1984) | v842
	v892 = int32(2)
	goto L225
L229:
	;
	goto L230
L230:
	;
	v861 = v848&int32(4032) | v842
	v863 = v813 - int32(3)
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810+v863))))
	if base.B2i32(v863 != v814)&base.B2i32(base.Ui32(v865) < base.Ui32(int32(224))) == int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v891 = v865<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v861
	v892 = int32(3)
	goto L225
L232:
	;
	goto L233
L233:
	;
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813+(v810-int32(4))))))
	v891 = v865<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v883&int32(7)<<(uint(int32(18))%32) | v861
	v892 = int32(4)
	goto L225
L234:
	;
	v925 = v892
	goto L220
L235:
	;
	goto L236
L236:
	;
	v896 = v891 - int32(97)
	if v896 < int32(0) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v925 = v892
	goto L220
L238:
	;
	goto L239
L239:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v896)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v902)>>(uint(v896&int32(7))%32))&int32(1) == int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v925 = v892
	goto L220
L241:
	;
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v813 - v892
	goto L243
L243:
	;
	goto L223
L244:
	;
	v926 = F_slice_del(m, l0)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L3
	} else {
		goto L245
	}
L245:
	;
	if int32(0) <= v926 {
		goto L204
	} else {
		goto L246
	}
L246:
	;
	v2088 = v926
	goto L1
L247:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v993
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)))
	if v993 < v996 {
		goto L264
	} else {
		goto L265
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v934
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v937
	if v934-int32(3) <= v937 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v940
	goto L247
L250:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945+v934-int32(1)))))
	if v949&int32(224) != int32(96) {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	if int32(1)<<(uint(v949)%32)&int32(_a_F_estonian_UTF_8_stem_27) == int32(0) {
		goto L249
	} else {
		goto L252
	}
L252:
	;
	v962 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_28), int32(12))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L3
	} else {
		goto L253
	}
L253:
	;
	if v962 == int32(0) {
		goto L249
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v940
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v967
	switch v962 - int32(1) {
	case 0:
		goto L257
	case 1:
		goto L256
	case 2:
		goto L255
	default:
		goto L247
	}
L255:
	;
	v985 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_29))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L3
	} else {
		goto L262
	}
L256:
	;
	v979 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_30))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L3
	} else {
		goto L260
	}
L257:
	;
	v973 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_31))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L3
	} else {
		goto L258
	}
L258:
	;
	if int32(0) <= v973 {
		goto L247
	} else {
		goto L259
	}
L259:
	;
	v2088 = v973
	goto L1
L260:
	;
	if int32(0) <= v979 {
		goto L247
	} else {
		goto L261
	}
L261:
	;
	v2088 = v979
	goto L1
L262:
	;
	if int32(0) <= v985 {
		goto L247
	} else {
		goto L263
	}
L263:
	;
	v2088 = v985
	goto L1
L264:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1238
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1240)))
	if v1238 < v1241 {
		goto L325
	} else {
		goto L326
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v993
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v996
	if v993 <= v996 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v999
	goto L264
L267:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002+v993-int32(1)))))
	if v1006&int32(224) != int32(96) {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	if int32(1)<<(uint(v1006)%32)&int32(_a_F_estonian_UTF_8_stem_32) == int32(0) {
		goto L266
	} else {
		goto L269
	}
L269:
	;
	v1019 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_33), int32(10))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L3
	} else {
		goto L270
	}
L270:
	;
	if v1019 == int32(0) {
		goto L266
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v999
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1024
	switch v1019 - int32(1) {
	case 0:
		goto L274
	case 1:
		goto L273
	default:
		goto L272
	}
L272:
	;
	v1229 = F_slice_del(m, l0)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L3
	} else {
		goto L323
	}
L273:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L304
L274:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L277
L275:
	;
	if v1157 == int32(0) {
		goto L272
	} else {
		goto L299
	}
L276:
	;
	v1157 = v1150
	goto L275
L277:
	;
	if v1045 <= v1046 {
		v1150 = int32(-1)
		goto L276
	} else {
		goto L279
	}
L278:
	;
	v1150 = int32(0)
	goto L276
L279:
	;
	v1063 = int32(1)
	v1064 = v1045 - v1063
	v1066 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1042+v1064))))
	v1068 = v1066 & int32(255)
	if v1064 == v1046 {
		v1123 = v1068
		v1124 = v1063
		goto L280
	} else {
		goto L281
	}
L280:
	;
	if int32(117) < v1123 {
		goto L289
	} else {
		goto L290
	}
L281:
	;
	if int32(0) <= v1066 {
		v1123 = v1068
		v1124 = v1063
		goto L280
	} else {
		goto L282
	}
L282:
	;
	v1074 = v1068 & int32(63)
	v1076 = v1045 - int32(2)
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042+v1076))))
	v1080 = v1078 << (uint(int32(6)) % 32)
	if base.B2i32(v1076 != v1046)&base.B2i32(base.Ui32(v1078) < base.Ui32(int32(192))) == int32(0) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1123 = v1080&int32(1984) | v1074
	v1124 = int32(2)
	goto L280
L284:
	;
	goto L285
L285:
	;
	v1093 = v1080&int32(4032) | v1074
	v1095 = v1045 - int32(3)
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042+v1095))))
	if base.B2i32(v1095 != v1046)&base.B2i32(base.Ui32(v1097) < base.Ui32(int32(224))) == int32(0) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1123 = v1097<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1093
	v1124 = int32(3)
	goto L280
L287:
	;
	goto L288
L288:
	;
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045+(v1042-int32(4))))))
	v1123 = v1097<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1115&int32(7)<<(uint(int32(18))%32) | v1093
	v1124 = int32(4)
	goto L280
L289:
	;
	v1157 = v1124
	goto L275
L290:
	;
	goto L291
L291:
	;
	v1128 = v1123 - int32(97)
	if v1128 < int32(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1157 = v1124
	goto L275
L293:
	;
	goto L294
L294:
	;
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1128)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1134)>>(uint(v1128&int32(7))%32))&int32(1) == int32(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1157 = v1124
	goto L275
L296:
	;
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1045 - v1124
	goto L298
L298:
	;
	goto L278
L299:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1160 + (v1024 - v1028)
	v1166 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L3
	} else {
		goto L300
	}
L300:
	;
	if v1166 != 0 {
		goto L272
	} else {
		goto L301
	}
L301:
	;
	goto L264
L302:
	;
	if v1220 < int32(0) {
		goto L264
	} else {
		goto L322
	}
L304:
	;
	goto L305
L305:
	;
	goto L306
L306:
	;
	v1176 = v1024
	v1178 = int32(4)
	goto L309
L308:
	;
	v1220 = v1202
	goto L302
L309:
	;
	if v1176 <= v999 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	goto L308
L311:
	;
	v1220 = int32(-1)
	goto L302
L312:
	;
	goto L313
L313:
	;
	v1183 = v1176 - int32(1)
	v1185 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1169+v1183))))
	if int32(0) <= v1185 {
		v1202 = v1183
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1206 = int32(1)
	if v1206 < v1178 {
		v1176 = v1202
		v1178 = v1178 - v1206
		goto L309
	} else {
		goto L321
	}
L315:
	;
	if v1183 <= v999 {
		v1202 = v1183
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1190 = v1183
	goto L317
L317:
	;
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1169+v1190))))
	if base.Ui32(int32(191)) < base.Ui32(v1195) {
		v1202 = v1190
		goto L314
	} else {
		goto L319
	}
L318:
	;
	v1202 = v999
	goto L314
L319:
	;
	v1199 = v1190 - int32(1)
	if v999 < v1199 {
		v1190 = v1199
		goto L317
	} else {
		goto L320
	}
L320:
	;
	goto L318
L321:
	;
	goto L310
L322:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1223 + (v1024 - v1168)
	goto L272
L323:
	;
	if int32(0) <= v1229 {
		goto L264
	} else {
		goto L324
	}
L324:
	;
	v2088 = v1229
	goto L1
L325:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1529
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1531)))
	if v1529 < v1532 {
		goto L407
	} else {
		goto L408
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1238
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1241
	if v1238 <= v1241 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1244
	goto L325
L328:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1247+v1238-int32(1)))))
	if v1251&int32(254) != int32(100) {
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v1258 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_34), int32(7))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L3
	} else {
		goto L330
	}
L330:
	;
	if v1258 == int32(0) {
		goto L327
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1244
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1263
	switch v1258 - int32(1) {
	case 0:
		goto L335
	case 1:
		goto L334
	case 2:
		goto L333
	case 3:
		goto L332
	default:
		goto L325
	}
L332:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L378
L333:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L344
L334:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1276 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L3
	} else {
		goto L338
	}
L335:
	;
	v1269 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_35))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L3
	} else {
		goto L336
	}
L336:
	;
	if int32(0) <= v1269 {
		goto L325
	} else {
		goto L337
	}
L337:
	;
	v2088 = v1269
	goto L1
L338:
	;
	if v1276 != 0 {
		goto L325
	} else {
		goto L339
	}
L339:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1278 + (v1263 - v1273)
	v1282 = F_slice_del(m, l0)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L3
	} else {
		goto L340
	}
L340:
	;
	if int32(0) <= v1282 {
		goto L325
	} else {
		goto L341
	}
L341:
	;
	v2088 = v1282
	goto L1
L342:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1341 = v1339 + (v1263 - v1286)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1341
	if int32(0) <= v1338 {
		goto L362
	} else {
		goto L363
	}
L344:
	;
	goto L345
L345:
	;
	goto L346
L346:
	;
	v1294 = v1263
	v1296 = int32(4)
	goto L349
L348:
	;
	v1338 = v1320
	goto L342
L349:
	;
	if v1294 <= v1244 {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L348
L351:
	;
	v1338 = int32(-1)
	goto L342
L352:
	;
	goto L353
L353:
	;
	v1301 = v1294 - int32(1)
	v1303 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1287+v1301))))
	if int32(0) <= v1303 {
		v1320 = v1301
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1324 = int32(1)
	if v1324 < v1296 {
		v1294 = v1320
		v1296 = v1296 - v1324
		goto L349
	} else {
		goto L361
	}
L355:
	;
	if v1301 <= v1244 {
		v1320 = v1301
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v1308 = v1301
	goto L357
L357:
	;
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1287+v1308))))
	if base.Ui32(int32(191)) < base.Ui32(v1313) {
		v1320 = v1308
		goto L354
	} else {
		goto L359
	}
L358:
	;
	v1320 = v1244
	goto L354
L359:
	;
	v1317 = v1308 - int32(1)
	if v1244 < v1317 {
		v1308 = v1317
		goto L357
	} else {
		goto L360
	}
L360:
	;
	goto L358
L361:
	;
	goto L350
L362:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1341 <= v1345 {
		goto L365
	} else {
		goto L366
	}
L363:
	;
	goto L364
L364:
	;
	v1376 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_36))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L3
	} else {
		goto L374
	}
L365:
	;
	v1370 = F_slice_del(m, l0)
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L3
	} else {
		goto L372
	}
L366:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1349 = int32(1)
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1347+v1341-v1349))))
	if base.Ui32(v1349) < base.Ui32((v1351-int32(115))&int32(255)) {
		goto L365
	} else {
		goto L367
	}
L367:
	;
	v1360 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_37), int32(5))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L3
	} else {
		goto L369
	}
L368:
	;
	v1366 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_38))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L3
	} else {
		goto L370
	}
L369:
	;
	switch v1360 - int32(1) {
	case 0:
		goto L368
	case 1:
		goto L365
	default:
		goto L325
	}
L370:
	;
	if int32(0) <= v1366 {
		goto L325
	} else {
		goto L371
	}
L371:
	;
	v2088 = v1366
	goto L1
L372:
	;
	if int32(0) <= v1370 {
		goto L325
	} else {
		goto L373
	}
L373:
	;
	v2088 = v1370
	goto L1
L374:
	;
	if int32(0) <= v1376 {
		goto L325
	} else {
		goto L375
	}
L375:
	;
	v2088 = v1376
	goto L1
L376:
	;
	if v1509 != 0 {
		goto L400
	} else {
		goto L401
	}
L377:
	;
	v1509 = v1502
	goto L376
L378:
	;
	if v1397 <= v1398 {
		v1502 = int32(-1)
		goto L377
	} else {
		goto L380
	}
L379:
	;
	v1502 = int32(0)
	goto L377
L380:
	;
	v1415 = int32(1)
	v1416 = v1397 - v1415
	v1418 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1394+v1416))))
	v1420 = v1418 & int32(255)
	if v1416 == v1398 {
		v1475 = v1420
		v1476 = v1415
		goto L381
	} else {
		goto L382
	}
L381:
	;
	if int32(117) < v1475 {
		goto L390
	} else {
		goto L391
	}
L382:
	;
	if int32(0) <= v1418 {
		v1475 = v1420
		v1476 = v1415
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1426 = v1420 & int32(63)
	v1428 = v1397 - int32(2)
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394+v1428))))
	v1432 = v1430 << (uint(int32(6)) % 32)
	if base.B2i32(v1428 != v1398)&base.B2i32(base.Ui32(v1430) < base.Ui32(int32(192))) == int32(0) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1475 = v1432&int32(1984) | v1426
	v1476 = int32(2)
	goto L381
L385:
	;
	goto L386
L386:
	;
	v1445 = v1432&int32(4032) | v1426
	v1447 = v1397 - int32(3)
	v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394+v1447))))
	if base.B2i32(v1447 != v1398)&base.B2i32(base.Ui32(v1449) < base.Ui32(int32(224))) == int32(0) {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1475 = v1449<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1445
	v1476 = int32(3)
	goto L381
L388:
	;
	goto L389
L389:
	;
	v1467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1397+(v1394-int32(4))))))
	v1475 = v1449<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1467&int32(7)<<(uint(int32(18))%32) | v1445
	v1476 = int32(4)
	goto L381
L390:
	;
	v1509 = v1476
	goto L376
L391:
	;
	goto L392
L392:
	;
	v1480 = v1475 - int32(97)
	if v1480 < int32(0) {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1509 = v1476
	goto L376
L394:
	;
	goto L395
L395:
	;
	v1486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1480)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1486)>>(uint(v1480&int32(7))%32))&int32(1) == int32(0) {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v1509 = v1476
	goto L376
L397:
	;
	goto L398
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397 - v1476
	goto L399
L399:
	;
	goto L379
L400:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1510 + (v1263 - v1380)
	v1516 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L3
	} else {
		goto L403
	}
L401:
	;
	goto L402
L402:
	;
	v1520 = F_slice_del(m, l0)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L3
	} else {
		goto L405
	}
L403:
	;
	if v1516 == int32(0) {
		goto L325
	} else {
		goto L404
	}
L404:
	;
	goto L402
L405:
	;
	if int32(0) <= v1520 {
		goto L325
	} else {
		goto L406
	}
L406:
	;
	v2088 = v1520
	goto L1
L407:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1705
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1707)))
	if v1705 < v1708 {
		goto L446
	} else {
		goto L447
	}
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1529
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1532
	if v1529 <= v1532 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1535
	goto L407
L410:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1538+v1529-int32(1)))))
	if v1542&int32(224) != int32(96) {
		goto L409
	} else {
		goto L411
	}
L411:
	;
	if int32(1)<<(uint(v1542)%32)&int32(_a_F_estonian_UTF_8_stem_39) == int32(0) {
		goto L409
	} else {
		goto L412
	}
L412:
	;
	v1555 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_40), int32(3))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L3
	} else {
		goto L413
	}
L413:
	;
	if v1555 == int32(0) {
		goto L409
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1535
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1560
	switch v1555 - int32(1) {
	case 0:
		goto L416
	case 1:
		goto L415
	default:
		goto L407
	}
L415:
	;
	v1697 = F_slice_del(m, l0)
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L3
	} else {
		goto L444
	}
L416:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L419
L417:
	;
	if v1692 != 0 {
		goto L407
	} else {
		goto L441
	}
L418:
	;
	v1692 = v1685
	goto L417
L419:
	;
	if v1580 <= v1581 {
		v1685 = int32(-1)
		goto L418
	} else {
		goto L421
	}
L420:
	;
	v1685 = int32(0)
	goto L418
L421:
	;
	v1598 = int32(1)
	v1599 = v1580 - v1598
	v1601 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1577+v1599))))
	v1603 = v1601 & int32(255)
	if v1599 == v1581 {
		v1658 = v1603
		v1659 = v1598
		goto L422
	} else {
		goto L423
	}
L422:
	;
	if int32(117) < v1658 {
		goto L431
	} else {
		goto L432
	}
L423:
	;
	if int32(0) <= v1601 {
		v1658 = v1603
		v1659 = v1598
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v1609 = v1603 & int32(63)
	v1611 = v1580 - int32(2)
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577+v1611))))
	v1615 = v1613 << (uint(int32(6)) % 32)
	if base.B2i32(v1611 != v1581)&base.B2i32(base.Ui32(v1613) < base.Ui32(int32(192))) == int32(0) {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1658 = v1615&int32(1984) | v1609
	v1659 = int32(2)
	goto L422
L426:
	;
	goto L427
L427:
	;
	v1628 = v1615&int32(4032) | v1609
	v1630 = v1580 - int32(3)
	v1632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577+v1630))))
	if base.B2i32(v1630 != v1581)&base.B2i32(base.Ui32(v1632) < base.Ui32(int32(224))) == int32(0) {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1658 = v1632<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1628
	v1659 = int32(3)
	goto L422
L429:
	;
	goto L430
L430:
	;
	v1650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580+(v1577-int32(4))))))
	v1658 = v1632<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1650&int32(7)<<(uint(int32(18))%32) | v1628
	v1659 = int32(4)
	goto L422
L431:
	;
	v1692 = v1659
	goto L417
L432:
	;
	goto L433
L433:
	;
	v1663 = v1658 - int32(97)
	if v1663 < int32(0) {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1692 = v1659
	goto L417
L435:
	;
	goto L436
L436:
	;
	v1669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1663)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1669)>>(uint(v1663&int32(7))%32))&int32(1) == int32(0) {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v1692 = v1659
	goto L417
L438:
	;
	goto L439
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1580 - v1659
	goto L440
L440:
	;
	goto L420
L441:
	;
	v1693 = F_slice_del(m, l0)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L3
	} else {
		goto L442
	}
L442:
	;
	if int32(0) <= v1693 {
		goto L407
	} else {
		goto L443
	}
L443:
	;
	v2088 = v1693
	goto L1
L444:
	;
	if int32(0) <= v1697 {
		goto L407
	} else {
		goto L445
	}
L445:
	;
	v2088 = v1697
	goto L1
L446:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1866
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1868)))
	if v1866 < v1869 {
		goto L204
	} else {
		goto L480
	}
L447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1705
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1708
	if v1705 <= v1708 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1711
	goto L446
L449:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1714+v1705-int32(1)))))
	if v1718 != int32(105) {
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v1723 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_41), int32(1))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L3
	} else {
		goto L451
	}
L451:
	;
	if v1723 == int32(0) {
		goto L448
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1711
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1728
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L455
L453:
	;
	if v1858 != 0 {
		goto L446
	} else {
		goto L477
	}
L454:
	;
	v1858 = v1851
	goto L453
L455:
	;
	if v1728 <= v1711 {
		v1851 = int32(-1)
		goto L454
	} else {
		goto L457
	}
L456:
	;
	v1851 = int32(0)
	goto L454
L457:
	;
	v1764 = int32(1)
	v1765 = v1728 - v1764
	v1767 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1743+v1765))))
	v1769 = v1767 & int32(255)
	if v1765 == v1711 {
		v1824 = v1769
		v1825 = v1764
		goto L458
	} else {
		goto L459
	}
L458:
	;
	if int32(117) < v1824 {
		goto L467
	} else {
		goto L468
	}
L459:
	;
	if int32(0) <= v1767 {
		v1824 = v1769
		v1825 = v1764
		goto L458
	} else {
		goto L460
	}
L460:
	;
	v1775 = v1769 & int32(63)
	v1777 = v1728 - int32(2)
	v1779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1743+v1777))))
	v1781 = v1779 << (uint(int32(6)) % 32)
	if base.B2i32(v1777 != v1711)&base.B2i32(base.Ui32(v1779) < base.Ui32(int32(192))) == int32(0) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v1824 = v1781&int32(1984) | v1775
	v1825 = int32(2)
	goto L458
L462:
	;
	goto L463
L463:
	;
	v1794 = v1781&int32(4032) | v1775
	v1796 = v1728 - int32(3)
	v1798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1743+v1796))))
	if base.B2i32(v1796 != v1711)&base.B2i32(base.Ui32(v1798) < base.Ui32(int32(224))) == int32(0) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v1824 = v1798<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1794
	v1825 = int32(3)
	goto L458
L465:
	;
	goto L466
L466:
	;
	v1816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728+(v1743-int32(4))))))
	v1824 = v1798<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1816&int32(7)<<(uint(int32(18))%32) | v1794
	v1825 = int32(4)
	goto L458
L467:
	;
	v1858 = v1825
	goto L453
L468:
	;
	goto L469
L469:
	;
	v1829 = v1824 - int32(97)
	if v1829 < int32(0) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v1858 = v1825
	goto L453
L471:
	;
	goto L472
L472:
	;
	v1835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1829)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1835)>>(uint(v1829&int32(7))%32))&int32(1) == int32(0) {
		goto L473
	} else {
		goto L474
	}
L473:
	;
	v1858 = v1825
	goto L453
L474:
	;
	goto L475
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1728 - v1825
	goto L476
L476:
	;
	goto L456
L477:
	;
	v1859 = F_slice_del(m, l0)
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L3
	} else {
		goto L478
	}
L478:
	;
	if int32(0) <= v1859 {
		goto L446
	} else {
		goto L479
	}
L479:
	;
	v2088 = v1859
	goto L1
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1866
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1869
	v1875 = v1866 - int32(1)
	if v1875 <= v1869 {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1872
	goto L204
L482:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1877+v1875))))
	if base.B2i32(v1879 != int32(117))&base.B2i32(v1879 != int32(97)) != 0 {
		goto L481
	} else {
		goto L483
	}
L483:
	;
	v1887 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_42), int32(4))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L3
	} else {
		goto L484
	}
L484:
	;
	if v1887 == int32(0) {
		goto L481
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1872
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1892
	v1894 = F_slice_del(m, l0)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L3
	} else {
		goto L486
	}
L486:
	;
	if int32(0) <= v1894 {
		goto L204
	} else {
		goto L487
	}
L487:
	;
	v2088 = v1894
	goto L1
L488:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2085
	v2088 = int32(1)
	goto L1
L489:
	;
	if v2033 != 0 {
		goto L488
	} else {
		goto L513
	}
L490:
	;
	v2033 = v2026
	goto L489
L491:
	;
	if v1903 <= v1922 {
		v2026 = int32(-1)
		goto L490
	} else {
		goto L493
	}
L492:
	;
	v2026 = int32(0)
	goto L490
L493:
	;
	v1939 = int32(1)
	v1940 = v1903 - v1939
	v1942 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1918+v1940))))
	v1944 = v1942 & int32(255)
	if v1940 == v1922 {
		v1999 = v1944
		v2000 = v1939
		goto L494
	} else {
		goto L495
	}
L494:
	;
	if int32(252) < v1999 {
		goto L503
	} else {
		goto L504
	}
L495:
	;
	if int32(0) <= v1942 {
		v1999 = v1944
		v2000 = v1939
		goto L494
	} else {
		goto L496
	}
L496:
	;
	v1950 = v1944 & int32(63)
	v1952 = v1903 - int32(2)
	v1954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1918+v1952))))
	v1956 = v1954 << (uint(int32(6)) % 32)
	if base.B2i32(v1952 != v1922)&base.B2i32(base.Ui32(v1954) < base.Ui32(int32(192))) == int32(0) {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v1999 = v1956&int32(1984) | v1950
	v2000 = int32(2)
	goto L494
L498:
	;
	goto L499
L499:
	;
	v1969 = v1956&int32(4032) | v1950
	v1971 = v1903 - int32(3)
	v1973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1918+v1971))))
	if base.B2i32(v1971 != v1922)&base.B2i32(base.Ui32(v1973) < base.Ui32(int32(224))) == int32(0) {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v1999 = v1973<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1969
	v2000 = int32(3)
	goto L494
L501:
	;
	goto L502
L502:
	;
	v1991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1903+(v1918-int32(4))))))
	v1999 = v1973<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1991&int32(7)<<(uint(int32(18))%32) | v1969
	v2000 = int32(4)
	goto L494
L503:
	;
	v2033 = v2000
	goto L489
L504:
	;
	goto L505
L505:
	;
	v2004 = v1999 - int32(97)
	if v2004 < int32(0) {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v2033 = v2000
	goto L489
L507:
	;
	goto L508
L508:
	;
	v2010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2004)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2010)>>(uint(v2004&int32(7))%32))&int32(1) == int32(0) {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v2033 = v2000
	goto L489
L510:
	;
	goto L511
L511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1903 - v2000
	goto L512
L512:
	;
	goto L492
L513:
	;
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v2035)))
	if v2034 < v2036 {
		goto L488
	} else {
		goto L514
	}
L514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2034
	v2040 = v2034 - int32(1)
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2040 <= v2041 {
		goto L488
	} else {
		goto L515
	}
L515:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2043+v2040))))
	if v2045&int32(224) != int32(96) {
		goto L488
	} else {
		goto L516
	}
L516:
	;
	if int32(1)<<(uint(v2045)%32)&int32(_a_F_estonian_UTF_8_stem_43) == int32(0) {
		goto L488
	} else {
		goto L517
	}
L517:
	;
	v2058 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_44), int32(3))
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L3
	} else {
		goto L518
	}
L518:
	;
	if v2058 == int32(0) {
		goto L488
	} else {
		goto L519
	}
L519:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2062
	switch v2058 - int32(1) {
	case 0:
		goto L522
	case 1:
		goto L521
	case 2:
		goto L520
	default:
		goto L488
	}
L520:
	;
	v2080 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_45))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L3
	} else {
		goto L527
	}
L521:
	;
	v2074 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_46))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L3
	} else {
		goto L525
	}
L522:
	;
	v2068 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_47))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L3
	} else {
		goto L523
	}
L523:
	;
	if int32(0) <= v2068 {
		goto L488
	} else {
		goto L524
	}
L524:
	;
	v2088 = v2068
	goto L1
L525:
	;
	if int32(0) <= v2074 {
		goto L488
	} else {
		goto L526
	}
L526:
	;
	v2088 = v2074
	goto L1
L527:
	;
	if v2080 < int32(0) {
		v2088 = v2080
		goto L1
	} else {
		goto L528
	}
L528:
	;
	goto L488
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
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v975 int32
	_ = v975
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1162 int32
	_ = v1162
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1211 int32
	_ = v1211
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	if l3&int32(3) == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v49 = F_exprType(m, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v23 = l3 + int32(32)
	if base.Ui32(v23) <= base.Ui32(l3) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v39
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v39
	goto L1
L5:
	;
	v29 = l3 + int32(4)
	if base.Ui32(v29) < base.Ui32(v23) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = v23
	goto L8
L7:
	;
	v31 = v29
	goto L8
L8:
	;
	v38 = F__emscripten_memset_bulkmem(m, l3, base.I32_extend8_s(int32(0)), (l3^int32(-1)+v31)&int32(-4)+int32(4))
	mBase = m.M
	goto L9
L9:
	;
	goto L1
L10:
	;
	return
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+68))
	if v53 == int32(0) {
		v94 = l1
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v107 = v94
	goto L33
L13:
	;
	if l1 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v94 = int32(0)
	goto L12
L15:
	;
	goto L16
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v59 != int32(319) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v64 = F_expression_tree_walker_impl(m, l1, int32(1503), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v72 = l1
	goto L23
L20:
	;
	if v64 == int32(0) {
		v94 = l1
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v88 = F_expression_tree_mutator_impl(m, v72, int32(1504), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L27
	}
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v81 != int32(319) {
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v94 = int32(0)
	goto L12
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v84 != 0 {
		v72 = v84
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v94 = v88
	goto L12
L28:
	;
	m.G0 = v16 + int32(80)
	return
L29:
	;
	F_bms_free(m, v597)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L10
	} else {
		goto L226
	}
L30:
	;
	v729 = l1
	v731 = v5
	goto L29
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L10
	} else {
		goto L223
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L10
	} else {
		goto L220
	}
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v116 != int32(27) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L10
	} else {
		goto L217
	}
L35:
	;
	if v116 != int32(6) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v107 = v677
	goto L33
L37:
	;
	goto L34
L38:
	;
	goto L37
L39:
	;
	v594 = F_pull_varnos(m, l0, v107)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L10
	} else {
		goto L182
	}
L40:
	;
	if l2 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if l2 != v121 {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v107
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v125 = F_find_base_rel(m, l0, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L10
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v130
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v107)+8)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v125)+108))
	if v134 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v195)
	v207 = l0
	v211 = v107
	goto L63
L47:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v135 <= int32(0) {
		v195 = int32(0)
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v195 = int32(0)
	goto L46
L50:
	;
	v138 = int32(0)
	if v138 < v135 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v141 = v135
	goto L53
L52:
	;
	v141 = v138
	goto L53
L53:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v149 = int32(0)
	goto L54
L54:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v142+v149<<(uint(int32(2))%32))))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+101)))
	if v161 != int32(1) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L49
L56:
	;
	v177 = v149 + int32(1)
	if v177 != v141 {
		v149 = v177
		goto L54
	} else {
		goto L62
	}
L57:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)+40))
	if v164 != int32(1) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v160)+44))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v168 != v132 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v170 = int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v160)+88))
	if v171 == int32(0) {
		v195 = v170
		goto L46
	} else {
		goto L60
	}
L60:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+100)))
	if v174 != 0 {
		v195 = v170
		goto L46
	} else {
		goto L61
	}
L61:
	;
	goto L56
L62:
	;
	goto L55
L63:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v207)+36))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220+v221<<(uint(int32(2))%32))))
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_examine_variable[0]))
	if v227 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L28
L65:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	switch v252 {
	case 0:
		goto L77
	case 1:
		goto L76
	default:
		goto L28
	case 6:
		goto L75
	}
L66:
	;
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v211)+8)))
	v231 = m.T0[v227].(func(*base.Module, int32, int32, int32, int32) int32)(m, v207, v225, v230, l3)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	if v231 == int32(0) {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v235 == int32(0) {
		goto L28
	} else {
		goto L69
	}
L69:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v238 != 0 {
		goto L28
	} else {
		goto L70
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_0), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_2), int32(_a_F_examine_variable_3))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	if v444 == int32(0) {
		goto L28
	} else {
		goto L121
	}
L75:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+92)))
	if v284 != 0 {
		goto L28
	} else {
		goto L87
	}
L76:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+20)))
	if v273 != 0 {
		goto L28
	} else {
		goto L84
	}
L77:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	v255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v211)+8)))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+20)))
	v257 = F_SearchSysCache3(m, int32(65), v254, v255, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1505)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v257
	if v257 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v263 = int32(*(*int16)(unsafe.Add(mBase, uint32(v211)+8)))
	v266 = F_bms_make_singleton(m, v263+int32(7))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L10
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v271 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v271)
	goto L28
L82:
	;
	v268 = F_all_rows_selectable(m, v207, v262, v266)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v268)
	goto L28
L84:
	;
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+8)))
	if v274 == int32(0) {
		goto L28
	} else {
		goto L85
	}
L85:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v280 = F_find_base_rel(m, v207, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v439 = v211 + int32(8)
	v443 = v280 + int32(140)
	goto L74
L87:
	;
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+8)))
	if v285 == int32(0) {
		goto L28
	} else {
		goto L88
	}
L88:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v225)+88))
	v295 = v290
	v296 = v207
	goto L90
L89:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+48))
	if v326 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L90:
	;
	if v295 == int32(0) {
		goto L89
	} else {
		goto L92
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L10
	} else {
		goto L94
	}
L92:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v296)+16))
	if v308 != 0 {
		v295 = v295 - int32(1)
		v296 = v308
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v225)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v313
	F_errmsg_internal(m, int32(_a_F_examine_variable_4), v16-int32(-64))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_5), int32(_a_F_examine_variable_3))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v296)+76))
	if v412 != 0 {
		goto L116
	} else {
		goto L117
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L10
	} else {
		goto L113
	}
L99:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v329 <= int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v225)+84))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v339 = int32(0)
	goto L101
L101:
	;
	v349 = v339 << (uint(int32(2)) % 32)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v333+v349)))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	if v356 == int32(0) {
		v375 = v355
		v376 = v356
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L98
L103:
	;
	if v376-v375 == int32(0) {
		goto L97
	} else {
		goto L111
	}
L104:
	;
	goto L103
L105:
	;
	if v355 != v356 {
		v375 = v355
		v376 = v356
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v360 = v352
	v361 = v332
	goto L107
L107:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+1)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	if v365 == int32(0) {
		v375 = v364
		v376 = v365
		goto L104
	} else {
		goto L109
	}
L108:
	;
	v375 = v364
	v376 = v365
	goto L104
L109:
	;
	v368 = int32(1)
	if v364 == v365 {
		v360 = v360 + v368
		v361 = v361 + v368
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v381 = v339 + int32(1)
	if v329 != v381 {
		v339 = v381
		goto L101
	} else {
		goto L112
	}
L112:
	;
	goto L102
L113:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v225)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v400
	F_errmsg_internal(m, int32(_a_F_examine_variable_6), v16+int32(16))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L10
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_7), int32(_a_F_examine_variable_3))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	v415 = v413
	goto L118
L117:
	;
	v415 = int32(0)
	goto L118
L118:
	;
	if v415 <= v339 {
		goto L38
	} else {
		goto L119
	}
L119:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v412)+12))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v417+v349)))
	if v419 <= int32(0) {
		goto L32
	} else {
		goto L120
	}
L120:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+16))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+12))
	v439 = v211 + int32(8)
	v443 = v424 + v419<<(uint(int32(2))%32) - int32(4)
	goto L74
L121:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+144))
	if v448 != 0 {
		goto L28
	} else {
		goto L122
	}
L122:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v447)+108))
	if v449 != 0 {
		goto L28
	} else {
		goto L123
	}
L123:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v447)+96))
	if v450 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v452 = v450
	goto L126
L125:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v447)+76))
	v452 = v451
	goto L126
L126:
	;
	v453 = int32(*(*int16)(unsafe.Add(mBase, uint32(v439))))
	if v452 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	if v491 == int32(0) {
		goto L31
	} else {
		goto L140
	}
L128:
	;
	goto L127
L129:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v457 <= int32(0) {
		v491 = int32(0)
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v491 = int32(0)
	goto L128
L132:
	;
	v460 = int32(0)
	if v460 < v457 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v463 = v457
	goto L135
L134:
	;
	v463 = v460
	goto L135
L135:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v452)+12))
	v468 = int32(0)
	goto L136
L136:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v464+v468<<(uint(int32(2))%32))))
	v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v476)+8)))
	if v477 == v453&int32(_a_F_examine_variable_8) {
		v491 = v476
		goto L128
	} else {
		goto L138
	}
L137:
	;
	goto L131
L138:
	;
	v480 = v468 + int32(1)
	if v480 != v463 {
		v468 = v480
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+26)))
	if v495 == int32(1) {
		goto L31
	} else {
		goto L141
	}
L141:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v447)+120))
	if v498 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v499 != int32(1) {
		goto L28
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v447)+100))
	if v541 != 0 {
		goto L160
	} else {
		goto L161
	}
L145:
	;
	v502 = int32(0)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v491)+16))
	if v504 == v502 {
		v535 = v502
		goto L147
	} else {
		goto L148
	}
L146:
	;
	if v535 == int32(0) {
		goto L28
	} else {
		goto L159
	}
L147:
	;
	goto L146
L148:
	;
	if v498 == int32(0) {
		v535 = v502
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if int32(0) < v509 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v513 = int32(0)
	goto L153
L151:
	;
	goto L152
L152:
	;
	v535 = v502
	goto L147
L153:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517+v513<<(uint(int32(2))%32))))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+4))
	if v504 == v522 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L152
L155:
	;
	v535 = int32(1)
	goto L147
L156:
	;
	goto L157
L157:
	;
	v526 = v513 + int32(1)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v526 < v527 {
		v513 = v526
		goto L153
	} else {
		goto L158
	}
L158:
	;
	goto L154
L159:
	;
	v539 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v539)
	goto L28
L160:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	if v542 != int32(1) {
		goto L28
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+40)))
	if v584 != 0 {
		goto L28
	} else {
		goto L178
	}
L163:
	;
	v545 = int32(0)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v491)+16))
	if v547 == v545 {
		v578 = v545
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v578 == int32(0) {
		goto L28
	} else {
		goto L177
	}
L165:
	;
	goto L164
L166:
	;
	if v541 == int32(0) {
		v578 = v545
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	if int32(0) < v552 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v556 = int32(0)
	goto L171
L169:
	;
	goto L170
L170:
	;
	v578 = v545
	goto L165
L171:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v560+v556<<(uint(int32(2))%32))))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	if v547 == v565 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	goto L170
L173:
	;
	v578 = int32(1)
	goto L165
L174:
	;
	goto L175
L175:
	;
	v569 = v556 + int32(1)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	if v569 < v570 {
		v556 = v569
		goto L171
	} else {
		goto L176
	}
L176:
	;
	goto L172
L177:
	;
	v582 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v582)
	goto L28
L178:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	if v585 == int32(0) {
		goto L28
	} else {
		goto L179
	}
L179:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v585)))
	if v588 != int32(6) {
		goto L28
	} else {
		goto L180
	}
L180:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v585)+28))
	if v591 == int32(0) {
		v207 = v444
		v211 = v585
		goto L63
	} else {
		goto L181
	}
L181:
	;
	goto L64
L182:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v597 = F_bms_difference(m, v594, v596)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L10
	} else {
		goto L183
	}
L183:
	;
	if v597 == int32(0) {
		goto L30
	} else {
		goto L184
	}
L184:
	;
	v603 = int32(0)
	if v597 == v603 {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v673
	v729 = v107
	v731 = v675
	goto L29
L186:
	;
	if v656 != 0 {
		goto L202
	} else {
		goto L203
	}
L187:
	;
	v656 = int32(0)
	goto L186
L188:
	;
	goto L189
L189:
	;
	v611 = int32(1)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v612 <= v611 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v615 = v611
	goto L192
L191:
	;
	v615 = v612
	goto L192
L192:
	;
	v620 = int32(0)
	v623 = int32(-1)
	goto L194
L193:
	;
	v656 = v648
	goto L186
L194:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v597+int32(8)+v620<<(uint(int32(2))%32))))
	if v630 != 0 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(76)))) = v640
	v648 = int32(1)
	goto L193
L196:
	;
	if int32(0) <= v623 {
		v648 = v603
		goto L193
	} else {
		goto L199
	}
L197:
	;
	v640 = v623
	goto L198
L198:
	;
	v642 = v620 + int32(1)
	if v642 != v615 {
		v620 = v642
		v623 = v640
		goto L194
	} else {
		goto L201
	}
L199:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v630)) {
		v648 = v603
		goto L193
	} else {
		goto L200
	}
L200:
	;
	v640 = base.I32_ctz(v630) | v620<<(uint(int32(5))%32)
	goto L198
L201:
	;
	goto L195
L202:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	if v658 != l2 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	goto L204
L204:
	;
	if l2 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L205:
	;
	v660 = l2
	goto L207
L206:
	;
	v660 = int32(0)
	goto L207
L207:
	;
	if v660 != 0 {
		goto L30
	} else {
		goto L208
	}
L208:
	;
	v661 = F_find_base_rel(m, l0, v658)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L10
	} else {
		goto L209
	}
L209:
	;
	v673 = v661
	v675 = v661
	goto L185
L210:
	;
	v665 = F_find_join_rel(m, l0, v594)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L10
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v667 = F_bms_is_member(m, l2, v594)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L10
	} else {
		goto L214
	}
L213:
	;
	v673 = v665
	v675 = v5
	goto L185
L214:
	;
	if v667 == int32(0) {
		goto L30
	} else {
		goto L215
	}
L215:
	;
	v671 = F_find_base_rel(m, l0, l2)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L10
	} else {
		goto L216
	}
L216:
	;
	v673 = v671
	v675 = v5
	goto L185
L217:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v225)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v682
	F_errmsg_internal(m, int32(_a_F_examine_variable_9), v16+int32(48))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L10
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_10), int32(_a_F_examine_variable_3))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L10
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v225)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v698
	F_errmsg_internal(m, int32(_a_F_examine_variable_11), v16+int32(32))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L10
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_12), int32(_a_F_examine_variable_3))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L10
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L223:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	v716 = int32(*(*int16)(unsafe.Add(mBase, uint32(v439))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v716
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v715
	F_errmsg_internal(m, int32(_a_F_examine_variable_13), v16)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L10
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_14), int32(_a_F_examine_variable_3))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L10
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v729
	v735 = F_exprType(m, v729)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L10
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v735
	v738 = F_exprTypmod(m, v729)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v738
	if v731 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	F_bms_free(m, v594)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L10
	} else {
		goto L355
	}
L230:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v744 = int32(0)
	if v594 == v744 {
		v785 = v744
		goto L232
	} else {
		goto L233
	}
L231:
	;
	if v785 != 0 {
		goto L245
	} else {
		goto L246
	}
L232:
	;
	goto L231
L233:
	;
	if v743 == int32(0) {
		v785 = v744
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	if v753 < v754 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v756 = v753
	goto L237
L236:
	;
	v756 = v754
	goto L237
L237:
	;
	if v756 <= int32(1) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v759 = int32(1)
	goto L240
L239:
	;
	v759 = v756
	goto L240
L240:
	;
	v760 = int32(8)
	v765 = int32(0)
	goto L241
L241:
	;
	v772 = v765 << (uint(int32(2)) % 32)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v743+v760+v772)))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v772+(v594+v760))))
	v777 = v774 & v776
	v779 = base.B2i32(v777 != int32(0))
	if v777 != 0 {
		v785 = v779
		goto L232
	} else {
		goto L243
	}
L242:
	;
	v785 = v779
	goto L232
L243:
	;
	v781 = v765 + int32(1)
	if v781 != v759 {
		v765 = v781
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v791 = F_remove_nulling_relids(m, v729, v789, int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L10
	} else {
		goto L248
	}
L246:
	;
	v793 = v729
	goto L247
L247:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v731)+108))
	if v794 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v793 = v791
	goto L247
L249:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v731)+112))
	if v1006 == int32(0) {
		goto L229
	} else {
		goto L310
	}
L250:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	if v797 <= int32(0) {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v810 = v5
	goto L252
L252:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v794)+12))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v813+v810<<(uint(int32(2))%32))))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)+84))
	if v818 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L249
L254:
	;
	v990 = v810 + int32(1)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	if v990 < v991 {
		v810 = v990
		goto L252
	} else {
		goto L309
	}
L255:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	if v821 == int32(0) {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v817)+36))
	if int32(0) < v824 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v830 = v821
	v833 = int32(0)
	v834 = v824
	goto L260
L258:
	;
	goto L259
L259:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v975 != 0 {
		goto L249
	} else {
		goto L308
	}
L260:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v817)+44))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v841+v833<<(uint(int32(2))%32))))
	if v845 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	goto L259
L262:
	;
	if v830 != 0 {
		goto L267
	} else {
		goto L268
	}
L263:
	;
	v957 = v830
	v958 = v834
	goto L264
L264:
	;
	v960 = v833 + int32(1)
	if v960 < v958 {
		v830 = v957
		v833 = v960
		v834 = v958
		goto L260
	} else {
		goto L307
	}
L265:
	;
	v946 = v830 + int32(4)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v817)+84))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v948)+12))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if base.Ui32(v946) < base.Ui32(v949+v950<<(uint(int32(2))%32)) {
		goto L304
	} else {
		goto L305
	}
L266:
	;
	v942 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v942)
	goto L265
L267:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v830)))
	if v848 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L268:
	;
	goto L269
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L10
	} else {
		goto L301
	}
L270:
	;
	v857 = F_equal(m, v793, v856)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L10
	} else {
		goto L275
	}
L271:
	;
	v856 = int32(0)
	goto L270
L272:
	;
	goto L273
L273:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v848)))
	if v852 != int32(27) {
		v856 = v848
		goto L270
	} else {
		goto L274
	}
L274:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	v856 = v855
	goto L270
L275:
	;
	if v857 == int32(0) {
		goto L265
	} else {
		goto L276
	}
L276:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817)+101)))
	if v861 != int32(1) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_examine_variable[1]))
	if v874 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L278:
	;
	if v833 != 0 {
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v817)+40))
	if v864 != int32(1) {
		goto L277
	} else {
		goto L280
	}
L280:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v817)+88))
	if v867 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817)+100)))
	if v868 != int32(1) {
		goto L277
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v871 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v871)
	goto L277
L284:
	;
	goto L283
L285:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v817)+88))
	if v902 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L286:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v817)+4))
	v881 = m.T0[v874].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v877, base.I32_extend16_s(v833+int32(1)), l3)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L10
	} else {
		goto L287
	}
L287:
	;
	if v881 == int32(0) {
		goto L285
	} else {
		goto L288
	}
L288:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v885 == int32(0) {
		goto L265
	} else {
		goto L289
	}
L289:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v888 != 0 {
		goto L249
	} else {
		goto L290
	}
L290:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L10
	} else {
		goto L291
	}
L291:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_0), int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L10
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_15), int32(_a_F_examine_variable_16))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L10
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v817)+4))
	v907 = int32(16)
	v914 = F_SearchSysCache3(m, int32(65), v906, (v833<<(uint(v907)%32)+int32(_a_F_examine_variable_17))>>(uint(v907)%32), int32(0))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L10
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v928 != 0 {
		goto L249
	} else {
		goto L300
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1505)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v914
	if v914 == int32(0) {
		goto L266
	} else {
		goto L298
	}
L298:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v817)+12))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v921)+68))
	v924 = F_all_rows_selectable(m, l0, v922, int32(0))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L10
	} else {
		goto L299
	}
L299:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v924)
	goto L296
L300:
	;
	goto L265
L301:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_18), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L10
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_19), int32(_a_F_examine_variable_16))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L10
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	v955 = v946
	goto L306
L305:
	;
	v955 = int32(0)
	goto L306
L306:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v817)+36))
	v957 = v955
	v958 = v956
	goto L264
L307:
	;
	goto L261
L308:
	;
	goto L254
L309:
	;
	goto L253
L310:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+4))
	if v1009 <= int32(0) {
		goto L229
	} else {
		goto L311
	}
L311:
	;
	v1023 = int32(0)
	goto L312
L312:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1026 != 0 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	goto L229
L314:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v1041 != 0 {
		goto L229
	} else {
		goto L318
	}
L315:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v731)+68))
	v1040 = v1026 + v1027<<(uint(int32(2))%32)
	goto L314
L316:
	;
	goto L317
L317:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1031)+52))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+12))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v731)+68))
	v1040 = v1033 + v1034<<(uint(int32(2))%32) - int32(4)
	goto L314
L318:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+12))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1042+v1023<<(uint(int32(2))%32))))
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1046)+16)))
	if v1047 != int32(101) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1194 = v1023 + int32(1)
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+4))
	if v1194 < v1195 {
		v1023 = v1194
		goto L312
	} else {
		goto L354
	}
L320:
	;
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1046)+8)))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1040)))
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051)+20)))
	if v1050 != v1052 {
		goto L319
	} else {
		goto L321
	}
L321:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+24))
	if v1054 == int32(0) {
		goto L319
	} else {
		goto L322
	}
L322:
	;
	v1057 = int32(0)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+4))
	if v1058 <= v1057 {
		goto L319
	} else {
		goto L323
	}
L323:
	;
	v1066 = v1057
	goto L324
L324:
	;
	v1074 = int32(0)
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+12))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1075+v1066<<(uint(int32(2))%32))))
	if v1079 == v1074 {
		v1086 = v1074
		goto L326
	} else {
		goto L327
	}
L325:
	;
	goto L319
L326:
	;
	v1087 = F_equal(m, v793, v1086)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L10
	} else {
		goto L329
	}
L327:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1079)))
	if v1082 != int32(27) {
		v1086 = v1079
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+4))
	v1086 = v1085
	goto L326
L329:
	;
	if v1087 != 0 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+4))
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051)+20)))
	v1091 = m.G0
	v1093 = v1091 - int32(48)
	m.G0 = v1093
	v1096 = F_SearchSysCache2(m, int32(62), v1089, v1090)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L10
	} else {
		goto L335
	}
L331:
	;
	goto L332
L332:
	;
	v1177 = v1066 + int32(1)
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+4))
	if v1177 < v1178 {
		v1066 = v1177
		goto L324
	} else {
		goto L353
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1506)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1131
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v731)+68))
	v1173 = F_all_rows_selectable(m, l0, v1171, int32(0))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L10
	} else {
		goto L352
	}
L334:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L10
	} else {
		goto L349
	}
L335:
	;
	if v1096 != 0 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v1102 = F_SysCacheGetAttr(m, int32(62), v1096, int32(6), v1093+int32(47))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L10
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L10
	} else {
		goto L346
	}
L339:
	;
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1093)+47)))
	if v1104 == int32(1) {
		goto L334
	} else {
		goto L340
	}
L340:
	;
	v1107 = F_DatumGetExpandedArray(m, v1102)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L10
	} else {
		goto L341
	}
L341:
	;
	F_deconstruct_expanded_array(m, v1107)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L10
	} else {
		goto L342
	}
L342:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+48))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1111+v1066<<(uint(int32(2))%32))))
	v1116 = F_pg_detoast_datum(m, v1115)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L10
	} else {
		goto L343
	}
L343:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1116)))
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+40)) = v1116
	v1120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+36)) = v1120
	*(*uint16)(unsafe.Add(mBase, uint32(v1093)+32)) = uint16(v1120)
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+28)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+24)) = int32(base.Ui32(v1118) >> (uint(int32(2)) % 32))
	v1131 = F_heap_copytuple(m, v1093+int32(24))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L10
	} else {
		goto L344
	}
L344:
	;
	F_ReleaseCatCache(m, v1096)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L10
	} else {
		goto L345
	}
L345:
	;
	m.G0 = v1093 + int32(48)
	goto L333
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1093))) = v1089
	F_errmsg_internal(m, int32(_a_F_examine_variable_20), v1093)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L10
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_examine_variable_21), int32(2414), int32(_a_F_examine_variable_22))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L10
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+20)) = v1089
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+16)) = int32(101)
	F_errmsg_internal(m, int32(_a_F_examine_variable_23), v1093+int32(16))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L10
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_examine_variable_21), int32(2421), int32(_a_F_examine_variable_22))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L10
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v1173)
	goto L319
L353:
	;
	goto L325
L354:
	;
	goto L313
L355:
	;
	goto L28
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
	var v34 int32
	_ = v34
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
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
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
			if v22 == int32(0) {
				v27 = v4
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = l1
				v27 = v25
			}
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
			if l2 != 0 {
				v33 = F_SPI_cursor_open_with_paramlist(m, int32(0), v29, v27, v28&int32(1))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v33
					if v33 == int32(0) {
						F_errstart_cold(m, int32(21), int32(_a_F_exec_run_select_0))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v111 = *(*int32)(unsafe.Add(mBase, _c_F_exec_run_select[0]))
							v112 = F_SPI_result_code_string(m, v111)
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v112
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v109
								F_errmsg_internal(m, int32(_a_F_exec_run_select_1), v10+int32(32))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_3), int32(_a_F_exec_run_select_4))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
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
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						if v38 != 0 {
							F_SPI_freetuptable(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v41
								v43 = int32(10)
								v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								if v44 == v41 {
									v97 = v43
									m.G0 = v10 + int32(48)
									return v97
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
									F_MemoryContextReset(m, v47)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v97 = v43
										m.G0 = v10 + int32(48)
										return v97
									}
								}
							}
						} else {
							v41 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v41
							v43 = int32(10)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							if v44 == v41 {
								v97 = v43
								m.G0 = v10 + int32(48)
								return v97
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
								F_MemoryContextReset(m, v47)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v97 = v43
									m.G0 = v10 + int32(48)
									return v97
								}
							}
						}
					}
				}
			} else {
				v54 = F_SPI_execute_plan_with_paramlist(m, v29, v27, v28&int32(1), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					if v54 != int32(5) {
						if v54 == int32(6) {
							v131 = int32(_a_F_exec_run_select_0)
							F_errstart_cold(m, int32(21), v131)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16801924))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_exec_run_select_5), int32(0))
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										F_set_errcontext_domain(m, v131)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v145
											F_errcontext_msg(m, int32(_a_F_exec_run_select_6), v10+int32(16))
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_7), int32(_a_F_exec_run_select_4))
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
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
							v62 = int32(_a_F_exec_run_select_0)
							F_errstart_cold(m, int32(21), v62)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16801924))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_exec_run_select_8), int32(0))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_set_errcontext_domain(m, v62)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
											F_errcontext_msg(m, int32(_a_F_exec_run_select_6), v10)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_9), int32(_a_F_exec_run_select_4))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
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
						v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_run_select[1]))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v90
						v93 = *(*int64)(unsafe.Add(mBase, _c_F_exec_run_select[2]))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v93
						v97 = int32(5)
						m.G0 = v10 + int32(48)
						return v97
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		if v22 == int32(0) {
			v27 = v4
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = l1
			v27 = v25
		}
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if l2 != 0 {
			v33 = F_SPI_cursor_open_with_paramlist(m, int32(0), v29, v27, v28&int32(1))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v33
				if v33 == int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_run_select_0))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v111 = *(*int32)(unsafe.Add(mBase, _c_F_exec_run_select[0]))
						v112 = F_SPI_result_code_string(m, v111)
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v109
							F_errmsg_internal(m, int32(_a_F_exec_run_select_1), v10+int32(32))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_3), int32(_a_F_exec_run_select_4))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
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
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					if v38 != 0 {
						F_SPI_freetuptable(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v41
							v43 = int32(10)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
							if v44 == v41 {
								v97 = v43
								m.G0 = v10 + int32(48)
								return v97
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
								F_MemoryContextReset(m, v47)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									v97 = v43
									m.G0 = v10 + int32(48)
									return v97
								}
							}
						}
					} else {
						v41 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v41
						v43 = int32(10)
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						if v44 == v41 {
							v97 = v43
							m.G0 = v10 + int32(48)
							return v97
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
							F_MemoryContextReset(m, v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v97 = v43
								m.G0 = v10 + int32(48)
								return v97
							}
						}
					}
				}
			}
		} else {
			v54 = F_SPI_execute_plan_with_paramlist(m, v29, v27, v28&int32(1), int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				if v54 != int32(5) {
					if v54 == int32(6) {
						v131 = int32(_a_F_exec_run_select_0)
						F_errstart_cold(m, int32(21), v131)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16801924))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_exec_run_select_5), int32(0))
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									F_set_errcontext_domain(m, v131)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v145
										F_errcontext_msg(m, int32(_a_F_exec_run_select_6), v10+int32(16))
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_7), int32(_a_F_exec_run_select_4))
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
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
						v62 = int32(_a_F_exec_run_select_0)
						F_errstart_cold(m, int32(21), v62)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16801924))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_exec_run_select_8), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_set_errcontext_domain(m, v62)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v76
										F_errcontext_msg(m, int32(_a_F_exec_run_select_6), v10)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_exec_run_select_2), int32(_a_F_exec_run_select_9), int32(_a_F_exec_run_select_4))
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
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
					v90 = *(*int32)(unsafe.Add(mBase, _c_F_exec_run_select[1]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v90
					v93 = *(*int64)(unsafe.Add(mBase, _c_F_exec_run_select[2]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v93
					v97 = int32(5)
					m.G0 = v10 + int32(48)
					return v97
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
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = l0
	goto L4
L1:
	;
	m.G0 = v7 + int32(32)
	return v73
L2:
	;
	v71 = F_pstrdup(m, l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L14
	} else {
		goto L31
	}
L3:
	;
	if v20 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v12 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	goto L5
L7:
	;
	v20 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v12 == int32(47) {
		v20 = v10
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v10 = v10 + int32(1)
	goto L4
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_expand_dynamic_library_name[0]))
	v25 = F_find_in_path(m, l0, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v45 = F_substitute_path_macro(m, l0, int32(_a_F_expand_dynamic_library_name_0), int32(_a_F_expand_dynamic_library_name_1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L14
	} else {
		goto L21
	}
L14:
	;
	return int32(0)
L15:
	;
	if v25 != 0 {
		v73 = v25
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_expand_dynamic_library_name_2)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v33 = F_psprintf(m, int32(_a_F_expand_dynamic_library_name_3), v7)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_expand_dynamic_library_name[0]))
	v37 = F_find_in_path(m, v33, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	F_pfree(m, v33)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if v37 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v73 = v37
	goto L1
L21:
	;
	v47 = F_pg_file_exists(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	if v47 != 0 {
		v73 = v45
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_pfree(m, v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a_F_expand_dynamic_library_name_2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	v57 = F_psprintf(m, int32(_a_F_expand_dynamic_library_name_3), v7+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v61 = F_substitute_path_macro(m, v57, int32(_a_F_expand_dynamic_library_name_0), int32(_a_F_expand_dynamic_library_name_1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	F_pfree(m, v57)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v65 = F_pg_file_exists(m, v61)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	if v65 != 0 {
		v73 = v61
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_pfree(m, v61)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	goto L2
L31:
	;
	v73 = v71
	goto L1
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
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
	v31 = v18
	v33 = v24
	v34 = v4
	goto L7
L5:
	;
	v130 = v18
	v133 = v4
	v134 = int32(1)
	goto L6
L6:
	;
	if v130 == int32(0) {
		v181 = v133
		goto L36
	} else {
		goto L37
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v31 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v130 = v117
	v133 = v120
	v134 = v21 + v24
	goto L6
L9:
	;
	v120 = F_lappend(m, v34, v116)
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
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
	v48 = v31 + int32(4)
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
	v117 = v56
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
	v117 = v31
	goto L9
L33:
	;
	if base.B2i32(v33 == v21) == int32(0) {
		v31 = v117
		v33 = v33 + int32(1)
		v34 = v120
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
	v193 = m.ExcPending
	if v193 != 0 {
		goto L21
	} else {
		goto L48
	}
L36:
	;
	m.G0 = v15 + int32(16)
	return v181
L37:
	;
	v144 = v130
	v147 = v133
	v148 = v134
	goto L38
L38:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+26)))
	if v153 == int32(0) {
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v181 = v162
	goto L36
L40:
	;
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v152)+8)))
	if v156 != v148 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v158 = F_flatCopyTargetEntry(m, v152)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L21
	} else {
		goto L44
	}
L42:
	;
	v161 = v152
	goto L43
L43:
	;
	v162 = F_lappend(m, v147, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L21
	} else {
		goto L45
	}
L44:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v158)+8)) = uint16(v148)
	v161 = v158
	goto L43
L45:
	;
	v165 = v144 + int32(4)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v166+v167<<(uint(int32(2))%32)) <= base.Ui32(v165) {
		v181 = v162
		goto L36
	} else {
		goto L46
	}
L46:
	;
	if v165 != 0 {
		v144 = v165
		v147 = v162
		v148 = v148 + int32(1)
		goto L38
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	F_errmsg_internal(m, int32(_a_F_expand_insert_targetlist_0), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_expand_insert_targetlist_1), int32(504), int32(_a_F_expand_insert_targetlist_2))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
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
	var v32 int32
	_ = v32
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
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
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	v8 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v15 == v8 {
		v59 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v59 = v8
	goto L1
L4:
	;
	goto L5
L5:
	;
	v29 = v8
	v32 = v8
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v29<<(uint(int32(2))%32))))
	v40 = F_bms_add_member(m, v32, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v59 = v40
	goto L1
L8:
	;
	return int32(0)
L9:
	;
	v45 = v29 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v45 < v46 {
		v29 = v45
		v32 = v40
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v63 <= v62 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v176 = v8
	goto L13
L13:
	;
	return v176
L14:
	;
	return int32(0)
L15:
	;
	goto L16
L16:
	;
	v76 = v62
	v78 = v8
	goto L17
L17:
	;
	v83 = v76 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v87 = v84 + v76<<(uint(int32(2))%32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v90 == int32(0) {
		v161 = v78
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v176 = v161
	goto L13
L19:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v83 < v164 {
		v76 = v83
		v78 = v161
		goto L17
	} else {
		goto L28
	}
L20:
	;
	v93 = F_bms_is_member(m, v83, v59)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	if v93 != 0 {
		v161 = v78
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v96 = F_lappend_int(m, v95, v83)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v101 = F_lappend(m, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
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
		goto L8
	} else {
		goto L25
	}
L25:
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
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v140 = F_lappend(m, v104, v124)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v140
	v143 = int32(5)
	v145 = l6 + v78<<(uint(v143)%32)
	v148 = l1 + v76<<(uint(v143)%32)
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v148)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v145)+24)) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v148)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v145)+16)) = v151
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v145)+8)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v155
	v161 = v78 + int32(1)
	goto L19
L28:
	;
	goto L18
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
	var v77 int32
	_ = v77
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
		v77 = v33
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v86 = v77
	goto L12
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v43 < v42 {
		v77 = v33
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+(l1+v49))))
	v68 = v63 & (v65 ^ int32(-1))
	v70 = base.B2i32(v68 == int32(0))
	if v68 != 0 {
		v77 = v70
		goto L16
	} else {
		goto L24
	}
L23:
	;
	v77 = v70
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
