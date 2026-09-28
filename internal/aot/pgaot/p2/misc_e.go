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
							F_errfinish(m, int32(_a_F_ExecARInsertTriggers_1), int32(2571), int32(_a_F_ExecARInsertTriggers_2))
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
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(34359738816)
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
	F_errfinish(m, int32(_a_F_ExecBSInsertTriggers_1), int32(2463), int32(_a_F_ExecBSInsertTriggers_2))
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
	*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(42949673408)
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
	F_errfinish(m, int32(_a_F_ExecBSUpdateTriggers_1), int32(2964), int32(_a_F_ExecBSUpdateTriggers_2))
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
	*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v42
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
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2625), int32(_a_F_ExecBuildAuxRowMark_4))
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
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2634), int32(_a_F_ExecBuildAuxRowMark_4))
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
	F_errfinish(m, int32(_a_F_ExecBuildAuxRowMark_3), int32(2644), int32(_a_F_ExecBuildAuxRowMark_4))
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
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			F_pfree(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v23 = v20 & int32(-5)
				v24 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v24)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v24
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v24)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = l0
				v34 = v23 & int32(_a_F_ExecForceStoreMinimalTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v34)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v37 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l0 - v37
				*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v36 + v37
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
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v24)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v24
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v24)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = l0
			v34 = v23 & int32(_a_F_ExecForceStoreMinimalTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v34)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v37 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l0 - v37
			*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v36 + v37
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
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
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
	return v316
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
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+4)))
	if v280&int32(2) != 0 {
		v316 = int32(0)
		goto L30
	} else {
		goto L95
	}
L33:
	;
	goto L32
L34:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v246 != int32(1) {
		goto L31
	} else {
		goto L82
	}
L35:
	;
	v231 = int32(0)
	v233 = F_ExecStoreMinimalTuple(m, v163, v115, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L80
	}
L36:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	m.T0[v228].(func(*base.Module, int32))(m, v115)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
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
	v221 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v221
	v226 = base.AtomicRmwOr32(m, v221, int32(_a_F_ExecGather_0), v221)
	goto L78
L78:
	;
	v143 = v211
	goto L46
L79:
	;
	v271 = v115
	goto L33
L80:
	;
	if v115 == int32(0) {
		v316 = v231
		goto L30
	} else {
		goto L81
	}
L81:
	;
	v271 = v115
	goto L33
L82:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v250 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+24))
	v253 = v251
	goto L85
L84:
	;
	v253 = int32(0)
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+172)) = v253
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v116)+52))
	if v255 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_ExecReScan(m, v116)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v259 = m.T0[v258].(func(*base.Module, int32) int32)(m, v116)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v249)+172)) = int32(0)
	if v259 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+4)))
	if v263&int32(2) == int32(0) {
		v271 = v259
		goto L33
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v268)
	goto L31
L94:
	;
	goto L93
L95:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v283 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v316 = v271
	goto L30
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v271
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+80))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v287)+24))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	m.T0[v291].(func(*base.Module, int32))(m, v289)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v294 = int32(_a_F_ExecGather_1)
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3]))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v288)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3])) = v297
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v287)+32))
	v303 = m.T0[v302].(func(*base.Module, int32, int32, int32) int64)(m, v287+int32(8), v288, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGather[3])) = v295
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289)+4)))
	v309 = v307 & int32(_a_F_ExecGather_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+4)) = uint16(v309)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+6)) = uint16(v312)
	v316 = v289
	goto L30
}
func F_ExecGatherMerge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v274 int32
	_ = v274
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[0]))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v15 == int32(0) {
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	if v19 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	F_MemoryContextReset(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L30
	}
L9:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGatherMerge[1])))
	if v82 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+160)))
	if v23 != int32(1) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v28 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	F_LaunchParallelWorkers(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L18
	}
L13:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v32 = F_ExecInitParallelPlan(m, v27, v22, v26, v19, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_ExecParallelReinitialize(m, v27, v28, v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v32
	v38 = v32
	goto L12
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v38 = v37
	goto L12
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v22)+164))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+164)) = v44 + v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v22)+168))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+168)) = v48 + v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if int32(0) < v52 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_ExecParallelCreateReaders(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v73
	goto L9
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v58
	v62 = F_palloc(m, v58<<(uint(int32(2))%32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v67 = v65 << (uint(int32(2)) % 32)
	if v67 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+40))
	base.MemoryCopy(m, v62, v71, v67)
	goto L9
L25:
	;
	v88 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v88)
	goto L8
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v85 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)) = uint8(v86)
	goto L25
L29:
	;
	goto L28
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v99 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v284 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104
	if v104 < v102 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v263)+24))
	v267 = F_gather_merge_readnext(m, l0, base.I32_wrap_i64(v264), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L80
	}
L35:
	;
	v111 = int32(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v152)+8)) = uint8(v153)
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = int32(0)
	goto L42
L38:
	;
	v119 = v111 << (uint(int32(4)) % 32)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v122 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v119+v120)+4)) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v124+v119)+8)) = v122
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*uint8)(unsafe.Add(mBase, uint32(v128+v119)+12)) = uint8(v122)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v134 = v111 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132+v134<<(uint(int32(2))%32))))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	m.T0[v140].(func(*base.Module, int32))(m, v138)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	if v134 != v102 {
		v111 = v134
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v102 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_binaryheap_build(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L79
	}
L44:
	;
	v161 = int32(1)
	goto L45
L45:
	;
	v169 = int32(0)
	goto L47
L46:
	;
	goto L43
L47:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[0]))
	if v177 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v102 <= int32(0) {
		goto L43
	} else {
		goto L70
	}
L49:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v169 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L51
L53:
	;
	v212 = v169 + int32(1)
	if v212 <= v102 {
		v169 = v212
		goto L47
	} else {
		goto L69
	}
L54:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190+v169<<(uint(int32(2))%32))))
	if v194 != 0 {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+106)))
	if v182 != 0 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v184 = int32(4)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v169<<(uint(v184)%32)-v184))))
	if v189 != 0 {
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
	F_load_tuple_array(m, l0, v169)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L68
	}
L61:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+4)))
	if v195&int32(2) == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v200 = F_gather_merge_readnext(m, l0, v169, v161)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v200 == int32(0) {
		goto L53
	} else {
		goto L66
	}
L66:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_binaryheap_add_unordered(m, v204, base.I64_extend_i32_s(v169))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
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
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v219 = int32(1)
	goto L71
L71:
	;
	v226 = int32(4)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216+v219<<(uint(v226)%32)-v226))))
	if v231 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L46
L73:
	;
	v234 = int32(0)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+v219<<(uint(int32(2))%32))))
	if v239 == v234 {
		v161 = v234
		goto L45
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v248 = v219 + int32(1)
	if v248 <= v102 {
		v219 = v248
		goto L71
	} else {
		goto L78
	}
L76:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+4)))
	if v242&int32(2) != 0 {
		v161 = v234
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
	v261 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v261)
	goto L31
L80:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v267 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_binaryheap_replace_first(m, v269, base.I64_extend32_s(v264))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v273 = F_binaryheap_remove_first(m, v269)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
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
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v287 <= int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v352 = int32(0)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v283)+24))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v353+base.I32_wrap_i64(v354)<<(uint(int32(2))%32))))
	if v359 == v352 {
		v398 = v352
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
	v296 = int32(0)
	goto L92
L92:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v304 = v301 + v296<<(uint(int32(4))%32)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+8))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v305 < v306 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	return int32(0)
L94:
	;
	v310 = v305
	goto L97
L95:
	;
	goto L96
L96:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v339 = v296 + int32(1)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v337+v339<<(uint(int32(2))%32))))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+8))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	m.T0[v345].(func(*base.Module, int32))(m, v343)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L101
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+8)) = v310 + int32(1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319+v310<<(uint(int32(2))%32))))
	F_pfree(m, v323)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L99
	}
L98:
	;
	goto L96
L99:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v304)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v326 < v327 {
		v310 = v326
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v339 < v348 {
		v296 = v339
		goto L92
	} else {
		goto L102
	}
L102:
	;
	goto L93
L103:
	;
	return v398
L104:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359)+4)))
	if v362&int32(2) != 0 {
		v398 = v352
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v365 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	return v359
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = v359
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+80))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v370)+24))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	m.T0[v374].(func(*base.Module, int32))(m, v372)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v377 = int32(_a_F_ExecGatherMerge_0)
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2]))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v371)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2])) = v380
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v370)+32))
	v386 = m.T0[v385].(func(*base.Module, int32, int32, int32) int64)(m, v370+int32(8), v371, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecGatherMerge[2])) = v378
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372)+4)))
	v392 = v390 & int32(_a_F_ExecGatherMerge_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v372)+4)) = uint16(v392)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	*(*uint16)(unsafe.Add(mBase, uint32(v372)+6)) = uint16(v395)
	v398 = v372
	goto L103
}
func F_ExecInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
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
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
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
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 float64
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int64
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
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
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int64
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 float64
	_ = v518
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int64
	_ = v533
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v698 float64
	_ = v698
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v723 int64
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v817 int32
	_ = v817
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+128))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+200))
	if v31 == int32(0) {
		v54 = l1
		v55 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	m.T0[v58].(func(*base.Module, int32))(m, v55)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L17
	}
L2:
	;
	v34 = F_ExecFindPartition(m, v28, l1, v31, l2, v27)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+204))
	if v38 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v46 = F_ExecGetRootToChildMap(m, v34, v27)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L14
	}
L8:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	if v41 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v43 = l2
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v43
	goto L7
L11:
	;
	v42 = int32(0)
	goto L13
L12:
	;
	v42 = l2
	goto L13
L13:
	;
	v43 = v42
	goto L10
L14:
	;
	if v46 == int32(0) {
		v54 = v34
		v55 = l2
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+204))
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
	v54 = v34
	v55 = v52
	goto L1
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+116)))
	if v63 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	if v71 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v66 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_ExecOpenIndices(m, v54, base.B2i32(v30 != int32(0)))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L3
	} else {
		goto L259
	}
L23:
	;
	m.G0 = v24 + int32(32)
	return v817
L24:
	;
	if l3 != 0 {
		goto L215
	} else {
		goto L216
	}
L25:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v54)+84))
	if v98 != 0 {
		goto L42
	} else {
		goto L43
	}
L26:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+8)))
	if v74 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
	if v77 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v88 = v71
	goto L29
L29:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+10)))
	if v89 != int32(1) {
		goto L25
	} else {
		goto L39
	}
L30:
	;
	F_ExecPendingInserts(m, v27)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v80 = F_ExecBRInsertTriggers(m, v27, v54, v55)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	if v80 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v817 = int32(0)
	goto L23
L36:
	;
	goto L37
L37:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	if v85 == int32(0) {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	v88 = v85
	goto L29
L39:
	;
	v92 = int32(0)
	v93 = F_ExecIRInsertTriggers(m, v27, v54, v55)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if v93 == int32(0) {
		v817 = v92
		goto L23
	} else {
		goto L41
	}
L41:
	;
	v704 = v55
	v708 = v92
	goto L24
L42:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+40)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	if v103 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v61)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+40)) = v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
	if v225 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L45:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v54)+104))
	if int32(2) <= v112 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+17)))
	if v106 != int32(1) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	F_ExecComputeStoredGenerated(m, v54, v27, v55, int32(3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	if v112 == v115 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v54)+84))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+52))
	v213 = m.T0[v212].(func(*base.Module, int32, int32, int32, int32) int32)(m, v27, v54, v55, v26)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L75
	}
L52:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v54)+108))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v54)+112))
	F_ExecBatchInsert(m, v28, v54, v117, v118, v112, v27, l3)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v121 = int32(_a_F_ExecInsert_0)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v27)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v54)+108))
	if v126 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v54)+104))
	v131 = F_palloc_mul(m, int32(4), v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v54)+100))
	if v140 <= v139 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+108)) = v131
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v54)+104))
	v136 = F_palloc_mul(m, int32(4), v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+112)) = v136
	goto L58
L61:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v143 = F_CreateTupleDescCopy(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L64
	}
L62:
	;
	v171 = v139
	goto L63
L63:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v54)+108))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v171<<(uint(int32(2))%32))))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+32))
	m.T0[v179].(func(*base.Module, int32, int32))(m, v177, v55)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L68
	}
L64:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v146 = F_CreateTupleDescCopy(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v149 = F_MakeSingleTupleTableSlot(m, v143, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v54)+108))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v151+v152<<(uint(int32(2))%32)))) = v149
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v158 = F_MakeSingleTupleTableSlot(m, v146, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v54)+112))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v160+v161<<(uint(int32(2))%32)))) = v158
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v54)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+100)) = v166 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v171 = v170
	goto L63
L68:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v54)+112))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182+v183<<(uint(int32(2))%32))))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+32))
	m.T0[v189].(func(*base.Module, int32, int32))(m, v187, v26)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	if v192|base.B2i32(v115 == v112) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v204 = v192
	goto L72
L71:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
	v196 = F_lappend(m, v195, v54)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L3
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+96)) = v204 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v122
	v817 = int32(0)
	goto L23
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+188)) = v196
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v27)+192))
	v200 = F_lappend(m, v199, v28)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+192)) = v200
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v204 = v203
	goto L72
L75:
	;
	if v213 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v817 = int32(0)
	goto L23
L77:
	;
	goto L78
L78:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v213)+40)) = v219
	v704 = v213
	v708 = int32(0)
	goto L24
L79:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v28)+104))
	switch v234 - int32(2) {
	case 0:
		v246 = v234
		goto L83
	default:
		goto L84
	case 3:
		goto L85
	}
L80:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+17)))
	if v228 != int32(1) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	F_ExecComputeStoredGenerated(m, v54, v27, v55, int32(3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v54)+116))
	if v247 != 0 {
		goto L89
	} else {
		goto L90
	}
L84:
	;
	v246 = int32(1)
	goto L83
L85:
	;
	v237 = int32(2)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v28)+216))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	if v241 == v237 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v244 = v237
	goto L88
L87:
	;
	v244 = int32(1)
	goto L88
L88:
	;
	v246 = v244
	goto L83
L89:
	;
	F_ExecWithCheckOptions(m, v246, v54, v55, v27)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L3
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v61)+52))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+24))
	if v251 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L91
L93:
	;
	F_ExecConstraints(m, v54, v55, v27)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+131)))
	if v255 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	if v30 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L98:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v54)+200))
	if v258 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	if v259 == int32(0) {
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v267 = F_ExecPartitionCheck(m, v54, v55, v27, int32(1))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L104
	}
L102:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+8)))
	if v262 != int32(1) {
		goto L97
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	goto L97
L105:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	if v695 == int32(0) {
		v817 = v687
		goto L23
	} else {
		goto L214
	}
L106:
	;
	v639 = F_ExecGetReturningSlot(m, v27, v54)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L3
	} else {
		goto L197
	}
L107:
	;
	v621 = int32(0)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v61)+188))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)+80))
	m.T0[v626].(func(*base.Module, int32, int32, int32, int32, int32))(m, v61, v55, v622, v621, v621)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L194
	}
L108:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v272 <= int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v275 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+22)) = uint16(v275)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+18)) = int32(-1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v54)+156))
	goto L110
L110:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[1]))
	if v304 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L3
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)) = uint8(v307)
	v313 = F_ExecCheckIndexConstraints(m, v54, v55, v27, v24+int32(24), v24+int32(18), v279)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	if v313 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	switch v30 - int32(2) {
	case 0:
		goto L121
	case 1:
		goto L120
	default:
		goto L106
	}
L118:
	;
	goto L119
L119:
	;
	v546 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L3
	} else {
		goto L182
	}
L120:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v54)+160))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v427)+20))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+64))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v427)+16))
	switch v433 {
	case 0:
		goto L150
	case 1:
		v478 = int32(0)
		goto L145
	case 2:
		goto L146
	case 3:
		goto L149
	case 4:
		goto L148
	default:
		goto L147
	}
L121:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v54)+160))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+20))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+64))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v324 = F_ExecUpdateLockMode(m, v323, v54)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v330 = F_ExecOnConflictLockRow(m, l0, v322, v24+int32(24), v328, v324, int32(1))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	if v330 == int32(0) {
		goto L110
	} else {
		goto L124
	}
L124:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecCheckTupleVisible(m, v334, v317, v322)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	v337 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+12)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v321)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v321)+4)) = v322
	if v319 == v337 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v54)+116))
	if v370 != 0 {
		goto L132
	} else {
		goto L133
	}
L127:
	;
	v343 = int32(_a_F_ExecInsert_0)
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v321)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v346
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v319)+24))
	v351 = m.T0[v350].(func(*base.Module, int32, int32, int32) int64)(m, v319, v321, v24+int32(31))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v344
	if v351 != int64(0) {
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	m.T0[v358].(func(*base.Module, int32))(m, v322)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	v361 = int32(0)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v320)+20))
	if v362 == v361 {
		v687 = v361
		goto L105
	} else {
		goto L131
	}
L131:
	;
	v365 = *(*float64)(unsafe.Add(mBase, uint32(v362)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v362)+424)) = base.F64_add(v365, float64(1))
	v687 = v361
	goto L105
L132:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	F_ExecWithCheckOptions(m, int32(3), v54, v322, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L3
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v54)+160))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+80))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376)+24))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)+8))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	m.T0[v380].(func(*base.Module, int32))(m, v378)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L3
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	v383 = int32(_a_F_ExecInsert_0)
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v377)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v386
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v376)+32))
	v392 = m.T0[v391].(func(*base.Module, int32, int32, int32) int64)(m, v376+int32(8), v377, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v384
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v378)+4)))
	v398 = v396 & int32(_a_F_ExecInsert_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v378)+4)) = uint16(v398)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	*(*uint16)(unsafe.Add(mBase, uint32(v378)+6)) = uint16(v401)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v54)+160))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+8))
	v408 = F_ExecUpdate(m, l0, v54, v24+int32(24), int32(0), v322, v407, l3)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L3
	} else {
		goto L139
	}
L138:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)+12))
	m.T0[v423].(func(*base.Module, int32))(m, v322)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L3
	} else {
		goto L143
	}
L139:
	;
	if v408 == int32(0) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v54)+152))
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+12)))
	if v413&int32(2) == int32(0) {
		goto L138
	} else {
		goto L141
	}
L141:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v408)+8))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+28))
	m.T0[v419].(func(*base.Module, int32))(m, v408)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L3
	} else {
		goto L142
	}
L142:
	;
	goto L138
L143:
	;
	v687 = v408
	goto L105
L144:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecCheckTupleVisible(m, v487, v426, v428)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L3
	} else {
		goto L165
	}
L145:
	;
	v482 = F_ExecOnConflictLockRow(m, l0, v428, v24+int32(24), v426, v478, int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L3
	} else {
		goto L163
	}
L146:
	;
	v478 = int32(1)
	goto L145
L147:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L3
	} else {
		goto L160
	}
L148:
	;
	v478 = int32(3)
	goto L145
L149:
	;
	v478 = int32(2)
	goto L145
L150:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[2]))
	if v435 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecInsert[3])))
	if v437&int32(1) == int32(0) {
		goto L22
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v426)+188))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+60))
	v447 = m.T0[v446].(func(*base.Module, int32, int32, int32, int32) int32)(m, v426, v24+int32(24), int32(_a_F_ExecInsert_2), v428)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L3
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	if v447 != 0 {
		goto L144
	} else {
		goto L156
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInsert_3), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L3
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_4), int32(3053), int32(_a_F_ExecInsert_5))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L3
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v433
	F_errmsg_internal(m, int32(_a_F_ExecInsert_6), v24)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L3
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_4), int32(3074), int32(_a_F_ExecInsert_5))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	if v482 == int32(0) {
		goto L110
	} else {
		goto L164
	}
L164:
	;
	goto L144
L165:
	;
	v490 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v431)+12)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v431)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v431)+4)) = v428
	if v429 == v490 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v54)+116))
	if v523 != 0 {
		goto L172
	} else {
		goto L173
	}
L167:
	;
	v496 = int32(_a_F_ExecInsert_0)
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0]))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v431)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v499
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v429)+24))
	v504 = m.T0[v503].(func(*base.Module, int32, int32, int32) int64)(m, v429, v431, v24+int32(31))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[0])) = v497
	if v504 != int64(0) {
		goto L166
	} else {
		goto L169
	}
L169:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v428)+8))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+12))
	m.T0[v511].(func(*base.Module, int32))(m, v428)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L3
	} else {
		goto L170
	}
L170:
	;
	v514 = int32(0)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v430)+20))
	if v515 == v514 {
		v687 = v514
		goto L105
	} else {
		goto L171
	}
L171:
	;
	v518 = *(*float64)(unsafe.Add(mBase, uint32(v515)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v515)+424)) = base.F64_add(v518, float64(1))
	v687 = v514
	goto L105
L172:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v430)+8))
	F_ExecWithCheckOptions(m, int32(3), v54, v428, v525)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L3
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v530 = F_ExecProcessReturning(m, l0, v54, int32(0), v428, v428, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L3
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	if l3 != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v532)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v532)+112)) = v533 + int64(1)
	goto L179
L178:
	;
	goto L179
L179:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v530)+8))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+28))
	m.T0[v539].(func(*base.Module, int32))(m, v530)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L3
	} else {
		goto L180
	}
L180:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v428)+8))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+12))
	m.T0[v543].(func(*base.Module, int32))(m, v428)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L3
	} else {
		goto L181
	}
L181:
	;
	v687 = v530
	goto L105
L182:
	;
	v548 = m.G0
	v550 = v548 - int32(16)
	m.G0 = v550
	v552 = int32(_a_F_ExecInsert_7)
	v553 = int32(1)
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[4]))
	v557 = v555 + v553
	if base.Ui32(v557) <= base.Ui32(v553) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v560 = v553
	goto L185
L184:
	;
	v560 = v557
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[4])) = v560
	*(*int64)(unsafe.Add(mBase, uint32(v550)+8)) = int64(74027918874902528)
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v550)+4)) = v560
	v567 = int32(0)
	v569 = F_LockAcquire(m, v550, int32(7), v567, v567)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L3
	} else {
		goto L186
	}
L186:
	;
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[4]))
	m.G0 = v550 + int32(16)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	v577 = int32(0)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v61)+188))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+84))
	m.T0[v580].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v61, v55, v576, v577, v577, v572)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L3
	} else {
		goto L187
	}
L187:
	;
	v586 = F_ExecInsertIndexTuples(m, v54, v27, int32(2), v55, v279, v24+int32(17))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L3
	} else {
		goto L188
	}
L188:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v61)+188))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+88))
	m.T0[v594].(func(*base.Module, int32, int32, int32, int32))(m, v61, v55, v572, (v588^int32(-1))&int32(1))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L3
	} else {
		goto L189
	}
L189:
	;
	v597 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L3
	} else {
		goto L190
	}
L190:
	;
	v599 = m.G0
	v601 = v599 - int32(16)
	m.G0 = v601
	*(*int32)(unsafe.Add(mBase, uint32(v601))) = v597
	*(*int64)(unsafe.Add(mBase, uint32(v601)+8)) = int64(74027918874902528)
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+4)) = v607
	v611 = F_LockRelease(m, v601, int32(7), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L3
	} else {
		goto L191
	}
L191:
	;
	m.G0 = v601 + int32(16)
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)))
	if v616 != int32(1) {
		v704 = v55
		v708 = v586
		goto L24
	} else {
		goto L192
	}
L192:
	;
	F_list_free(m, v586)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L3
	} else {
		goto L193
	}
L193:
	;
	goto L110
L194:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v629 <= int32(0) {
		v704 = v55
		v708 = v621
		goto L24
	} else {
		goto L195
	}
L195:
	;
	v632 = int32(0)
	v635 = F_ExecInsertIndexTuples(m, v54, v27, v632, v55, v632, v632)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L3
	} else {
		goto L196
	}
L196:
	;
	v704 = v55
	v708 = v635
	goto L24
L197:
	;
	v642 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[5]))
	if int32(2) <= v642 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v687 = int32(0)
	goto L105
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L3
	} else {
		goto L211
	}
L200:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsert[2]))
	if v647 != 0 {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L202
L202:
	;
	goto L198
L203:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecInsert[3])))
	if v649&int32(1) == int32(0) {
		goto L22
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v645)+188))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+60))
	v657 = m.T0[v656].(func(*base.Module, int32, int32, int32, int32) int32)(m, v645, v24+int32(24), int32(_a_F_ExecInsert_2), v639)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L3
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	if v657 == int32(0) {
		goto L199
	} else {
		goto L208
	}
L208:
	;
	F_ExecCheckTupleVisible(m, v27, v645, v639)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L3
	} else {
		goto L209
	}
L209:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v639)+8))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v663)+12))
	m.T0[v664].(func(*base.Module, int32))(m, v639)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L3
	} else {
		goto L210
	}
L210:
	;
	goto L202
L211:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInsert_3), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L3
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_4), int32(423), int32(_a_F_ExecInsert_8))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L3
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	v698 = *(*float64)(unsafe.Add(mBase, uint32(v695)+408))
	*(*float64)(unsafe.Add(mBase, uint32(v695)+408)) = base.F64_add(v698, float64(1))
	v817 = v687
	goto L23
L215:
	;
	v723 = *(*int64)(unsafe.Add(mBase, uint32(v27)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+112)) = v723 + int64(1)
	goto L217
L216:
	;
	goto L217
L217:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v28)+204))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v28)+104))
	if v728 != int32(2) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	F_ExecARInsertTriggers(m, v27, v54, v704, v708, v746)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L3
	} else {
		goto L229
	}
L219:
	;
	v746 = v727
	goto L218
L220:
	;
	goto L221
L221:
	;
	if v727 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v746 = int32(0)
	goto L218
L223:
	;
	goto L224
L224:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+2)))
	if v734 != int32(1) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v746 = v727
	goto L218
L226:
	;
	goto L227
L227:
	;
	v737 = int32(0)
	F_ExecARUpdateTriggers(m, v27, v54, v737, v737, v737, v737, v704, v737, v727, v737)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L3
	} else {
		goto L228
	}
L228:
	;
	v746 = v737
	goto L218
L229:
	;
	F_list_free(m, v708)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L3
	} else {
		goto L230
	}
L230:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v54)+116))
	if v751 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	F_ExecWithCheckOptions(m, int32(0), v54, v704, v27)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L3
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v54)+152))
	if v755 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	goto L233
L235:
	;
	if l4 != 0 {
		goto L255
	} else {
		goto L256
	}
L236:
	;
	v805 = int32(0)
	goto L235
L237:
	;
	goto L238
L238:
	;
	v759 = int32(0)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v760 == v759 {
		v780 = v759
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v783 = F_ExecProcessReturning(m, l0, v54, int32(0), v780, v704, v26)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L3
	} else {
		goto L247
	}
L240:
	;
	v763 = F_ExecGetRootToChildMap(m, v54, v27)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L3
	} else {
		goto L241
	}
L241:
	;
	if v763 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v780 = v760
	goto L239
L243:
	;
	goto L244
L244:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v763)+8))
	v768 = F_ExecGetReturningSlot(m, v27, v54)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L3
	} else {
		goto L245
	}
L245:
	;
	v770 = F_execute_attr_map_slot(m, v767, v760, v768)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L3
	} else {
		goto L246
	}
L246:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v772)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+40)) = v773
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v775)+36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v770)+36)) = uint16(v776)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v775)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+32)) = v778
	v780 = v770
	goto L239
L247:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v785 == int32(0) {
		v805 = v783
		goto L235
	} else {
		goto L248
	}
L248:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v783)+8))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v788)+28))
	m.T0[v789].(func(*base.Module, int32))(m, v783)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L3
	} else {
		goto L249
	}
L249:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v780)+8))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v792)+12))
	m.T0[v793].(func(*base.Module, int32))(m, v780)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L3
	} else {
		goto L250
	}
L250:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v780 != v796 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v796)+8))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v798)+12))
	m.T0[v799].(func(*base.Module, int32))(m, v796)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L3
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	v805 = v783
	goto L235
L254:
	;
	goto L253
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v704
	goto L257
L256:
	;
	goto L257
L257:
	;
	if l5 == int32(0) {
		v817 = v805
		goto L23
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v54
	v817 = v805
	goto L23
L259:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInsert_9), int32(0))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L3
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_ExecInsert_10), int32(1355), int32(_a_F_ExecInsert_11))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L3
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecNamedTuplestoreScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(785), int32(786))
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
	v4 = F_ExecScan(m, l0, int32(792), int32(793))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
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
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int64
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 float64
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+156))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	F_MemoryContextReset(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v20|v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v16 + int32(16)
	return v287
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[0]))
	if v32 != 0 {
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
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v19 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v114 = m.T0[l1].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L43
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
	if v38 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	v43 = F_bms_is_member(m, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v54 = int32(1)
	v55 = v38 - v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v57 = v55 + v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v58 == v54 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	if v43 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v48 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v48 != 0 {
		v287 = v47
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	m.T0[v51].(func(*base.Module, int32))(m, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v287 = v47
	goto L3
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	m.T0[v63].(func(*base.Module, int32))(m, v61)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v67 = v55 << (uint(int32(2)) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67+v68)))
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v287 = v61
	goto L3
L25:
	;
	v71 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v71)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)))
	if v73&int32(2) != 0 {
		v287 = v4
		goto L3
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84+v67)))
	if v86 == int32(0) {
		goto L11
	} else {
		goto L34
	}
L28:
	;
	v76 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v70)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v76 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	m.T0[v81].(func(*base.Module, int32))(m, v70)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v287 = v70
	goto L3
L33:
	;
	goto L32
L34:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v90)
	v94 = F_EvalPlanQualFetchRowMark(m, v19, v38, v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if base.B2i32(v89 == int32(0))|base.B2i32(v94 == int32(0)) != 0 {
		v287 = v4
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v99&int32(2) != 0 {
		v287 = v4
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v102 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v89)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v102 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	m.T0[v107].(func(*base.Module, int32))(m, v89)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v287 = v89
	goto L3
L42:
	;
	goto L41
L43:
	;
	v287 = v114
	goto L3
L44:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[0]))
	if v130 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v19 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v213
	if v21 != 0 {
		goto L87
	} else {
		goto L88
	}
L51:
	;
	if v20 == int32(0) {
		v287 = v221
		goto L3
	} else {
		goto L84
	}
L52:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
	if v216&int32(2) == int32(0) {
		goto L50
	} else {
		goto L83
	}
L53:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	m.T0[v209].(func(*base.Module, int32))(m, v206)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L82
	}
L54:
	;
	v203 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v160)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L80
	}
L55:
	;
	if v199 != 0 {
		v213 = v199
		goto L52
	} else {
		goto L79
	}
L56:
	;
	v196 = m.T0[l1].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L78
	}
L57:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+72))
	if v136 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+64))
	v141 = F_bms_is_member(m, v139, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v148 = int32(1)
	v149 = v136 - v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v151 = v149 + v150
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v152 == v148 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	if v141 == int32(0) {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v146 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v146 != 0 {
		v199 = v145
		goto L55
	} else {
		goto L64
	}
L64:
	;
	v206 = v145
	goto L53
L65:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v206 = v155
	goto L53
L66:
	;
	goto L67
L67:
	;
	v157 = v149 << (uint(int32(2)) % 32)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157+v158)))
	if v160 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v161)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+4)))
	if v163&int32(2) == int32(0) {
		goto L54
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v157)))
	if v171 == int32(0) {
		goto L56
	} else {
		goto L72
	}
L71:
	;
	v221 = int32(0)
	goto L51
L72:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v175)
	v177 = int32(0)
	v180 = F_EvalPlanQualFetchRowMark(m, v19, v136, v174)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if base.B2i32(v174 == v177)|base.B2i32(v180 == int32(0)) != 0 {
		v221 = v177
		goto L51
	} else {
		goto L74
	}
L74:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
	if v185&int32(2) != 0 {
		v221 = v177
		goto L51
	} else {
		goto L75
	}
L75:
	;
	v188 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, l0, v174)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v188 == int32(0) {
		v206 = v174
		goto L53
	} else {
		goto L77
	}
L77:
	;
	v213 = v174
	goto L52
L78:
	;
	v199 = v196
	goto L55
L79:
	;
	v221 = int32(0)
	goto L51
L80:
	;
	if v203 != 0 {
		v213 = v160
		goto L52
	} else {
		goto L81
	}
L81:
	;
	v206 = v160
	goto L53
L82:
	;
	v213 = v206
	goto L52
L83:
	;
	v221 = v213
	goto L51
L84:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	m.T0[v229].(func(*base.Module, int32))(m, v227)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v287 = v227
	goto L3
L86:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v276 != 0 {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v233 = int32(_a_F_ExecScan_0)
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1]))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v236
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v241 = m.T0[v240].(func(*base.Module, int32, int32, int32) int64)(m, v21, v22, v16+int32(15))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v20 == int32(0) {
		v287 = v213
		goto L3
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v234
	if v241 == int64(0) {
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	m.T0[v253].(func(*base.Module, int32))(m, v251)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v256 = int32(_a_F_ExecScan_0)
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1]))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v250)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v259
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v265 = m.T0[v264].(func(*base.Module, int32, int32, int32) int64)(m, v20+int32(8), v250, int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecScan[1])) = v257
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+4)))
	v271 = v269 & int32(_a_F_ExecScan_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v251)+4)) = uint16(v271)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	*(*uint16)(unsafe.Add(mBase, uint32(v251)+6)) = uint16(v274)
	v287 = v251
	goto L3
L95:
	;
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v276)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v276)+424)) = base.F64_add(v277, float64(1))
	goto L97
L96:
	;
	goto L97
L97:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	F_MemoryContextReset(m, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
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
	v21 = F_planstate_tree_walker_impl(m, l0, int32(680), l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+361)))
	if v15 != int32(1) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	F_InstrStart(m, v12)
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
	switch v23 - int32(424) {
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
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
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
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+164))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v63
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
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+361)))
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			F_pfree(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v17 = v14 & int32(-5)
				v18 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
				v24 = v17 & int32(_a_F_ExecStorePinnedBufferHeapTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v24)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v26
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v28)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
				if l2 != v30 {
					if v30 != 0 {
						F_ReleaseBuffer(m, v30)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = l2
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = l2
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
						return
					}
				} else {
					if l2 == int32(0) {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
						return
					} else {
						F_ReleaseBuffer(m, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						}
					}
				}
			}
		} else {
			v17 = v8
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
			v24 = v17 & int32(_a_F_ExecStorePinnedBufferHeapTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v26
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v28)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
			if l2 != v30 {
				if v30 != 0 {
					F_ReleaseBuffer(m, v30)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = l2
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = l2
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
					return
				}
			} else {
				if l2 == int32(0) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
					return
				} else {
					F_ReleaseBuffer(m, l2)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
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
				F_errfinish(m, int32(_a_F_ExecStorePinnedBufferHeapTuple_3), int32(1713), int32(_a_F_ExecStorePinnedBufferHeapTuple_4))
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v7 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v12 <= v7 {
		v24 = v7
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v26 = int32(0)
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
		if v30 == int32(3) {
			v33 = int32(208)
		} else {
			v33 = int32(204)
		}
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v11+v33)))
		F_ExecARUpdateTriggers(m, v25, l2, v26, v26, l3, l4, l5, v24, v35, int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			F_list_free(m, v24)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
				if v41 != 0 {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_ExecWithCheckOptions(m, int32(0), l2, l5, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
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
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		switch v15 {
		case 0:
			v24 = v7
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v26 = int32(0)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
			if v30 == int32(3) {
				v33 = int32(208)
			} else {
				v33 = int32(204)
			}
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v11+v33)))
			F_ExecARUpdateTriggers(m, v25, l2, v26, v26, l3, l4, l5, v24, v35, int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				F_list_free(m, v24)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
					if v41 != 0 {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_ExecWithCheckOptions(m, int32(0), l2, l5, v43)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				}
			}
		default:
			v17 = int32(1)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = int32(0)
			v21 = F_ExecInsertIndexTuples(m, l2, v18, v17, l5, v19, v19)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = v21
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v26 = int32(0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
				if v30 == int32(3) {
					v33 = int32(208)
				} else {
					v33 = int32(204)
				}
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v11+v33)))
				F_ExecARUpdateTriggers(m, v25, l2, v26, v26, l3, l4, l5, v24, v35, int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_list_free(m, v24)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
						if v41 != 0 {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_ExecWithCheckOptions(m, int32(0), l2, l5, v43)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
		case 2:
			v17 = int32(5)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v19 = int32(0)
			v21 = F_ExecInsertIndexTuples(m, l2, v18, v17, l5, v19, v19)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = v21
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v26 = int32(0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
				if v30 == int32(3) {
					v33 = int32(208)
				} else {
					v33 = int32(204)
				}
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v11+v33)))
				F_ExecARUpdateTriggers(m, v25, l2, v26, v26, l3, l4, l5, v24, v35, int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_list_free(m, v24)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
						if v41 != 0 {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_ExecWithCheckOptions(m, int32(0), l2, l5, v43)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
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
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	v2 = l1
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
					*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(1487)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(1488)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(1489)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(1490)
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v17)+21)) = uint8(v39)
					*(*uint8)(unsafe.Add(mBase, uint32(v17)+20)) = uint8(v2)
					v42 = F_pg_detoast_datum_packed(m, v14)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
						if v44 == int32(1) {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
							if v50 == int32(18) {
								v53 = int32(16)
							} else {
								v53 = int32(0)
							}
							if base.Ui32((v50-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v60 = int32(4)
							} else {
								v60 = v53
							}
							v73 = v60
						} else {
							v61 = int32(1)
							if v44&v61 != 0 {
								v73 = int32(base.Ui32(v44)>>(uint(v61)%32)) - v61
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
								v73 = int32(base.Ui32(v67)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v75 = v11 + int32(12)
						v76 = int32(1)
						if v44&v76 != 0 {
							v80 = v76
						} else {
							v80 = int32(4)
						}
						v83 = *(*int32)(unsafe.Add(mBase, _c_F_each_worker[0]))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
						v86 = F_makeJsonLexContextCstringLen(m, v75, v42+v80, v73, v84, int32(1))
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
								v98 = F_pg_parse_json(m, v75, v20)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									if v98 != 0 {
										F_json_errsave_error(m, v98, v75, int32(0))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
											F_MemoryContextDelete(m, v103)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												F_freeJsonLexContext(m, v11+int32(12))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v110 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
										F_MemoryContextDelete(m, v103)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											F_freeJsonLexContext(m, v11+int32(12))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												v110 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
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
func F_ean13_in(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14257(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
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
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_element_match[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v10)
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
func F_eqjoinsel_find_matches(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 float64
	_ = v104
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v125 int64
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v147 int64
	_ = v147
	var v164 int32
	_ = v164
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v200 int64
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v229 int32
	_ = v229
	var v244 int64
	_ = v244
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v251 int64
	_ = v251
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v265 int64
	_ = v265
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int64
	_ = v280
	var v290 int64
	_ = v290
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v400 int32
	_ = v400
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v427 int64
	_ = v427
	var v429 int64
	_ = v429
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v538 int32
	_ = v538
	var v540 int64
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int64
	_ = v555
	var v557 int64
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int64
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v620 int64
	_ = v620
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int64
	_ = v696
	var v698 int64
	_ = v698
	var v700 int64
	_ = v700
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int64
	_ = v737
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v767 int32
	_ = v767
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v825 int32
	_ = v825
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v897 float64
	_ = v897
	var v898 int32
	_ = v898
	var v902 int64
	_ = v902
	var v905 int32
	_ = v905
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v939 int64
	_ = v939
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int64
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 float32
	_ = v962
	var v963 int32
	_ = v963
	var v967 float32
	_ = v967
	var v975 int32
	_ = v975
	var v994 int32
	_ = v994
	var v1004 float64
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1088 float64
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1093 int64
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int64
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1141 int32
	_ = v1141
	var v1143 int64
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int64
	_ = v1158
	var v1160 int64
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int64
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1198 float32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1203 float32
	_ = v1203
	var v1226 int32
	_ = v1226
	var v1236 float64
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1257 int32
	_ = v1257
	var v1267 float64
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1290 int32
	_ = v1290
	var v1300 float64
	_ = v1300
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	v14 = int32(0)
	v28 = float64(0)
	v29 = m.G0
	v31 = v29 - int32(144)
	m.G0 = v31
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+136)) = uint8(v14)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+120)) = uint8(v14)
	v37 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+106)) = uint16(v37)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+104)) = uint8(v14)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+100)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v31)+92)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+88)) = l0
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L19
	} else {
		goto L207
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v1290
	*(*float64)(unsafe.Add(mBase, uint32(l12))) = v1300
	m.G0 = v31 + int32(144)
	return
L3:
	;
	v1036 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+13)) = uint8(v1036)
	if l2 != l3 {
		goto L167
	} else {
		goto L168
	}
L4:
	;
	v859 = v31 + int32(88)
	if l4 != 0 {
		goto L149
	} else {
		goto L150
	}
L5:
	;
	v46 = l2
	goto L7
L6:
	;
	v46 = v14
	goto L7
L7:
	;
	if v46 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if int32(0) < l7 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	if v52 <= v51 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v1290 = v14
	v1300 = v28
	goto L2
L12:
	;
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v56 = l5
	v57 = l4
	v58 = l6
	v59 = l8
	v60 = l10
	v61 = l7
	v62 = l9
	goto L12
L14:
	;
	goto L15
L15:
	;
	v56 = l6
	v57 = l4 ^ int32(1)
	v58 = l5
	v59 = l7
	v60 = l9
	v61 = l8
	v62 = l10
	goto L12
L16:
	;
	v63 = l2
	goto L18
L17:
	;
	v63 = l3
	goto L18
L18:
	;
	v65 = v31 + int32(20)
	F_fmgr_info(m, v63, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+80)) = uint8(v68)
	v70 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+66)) = uint16(v70)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+64)) = uint8(v68)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v31)+52)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+13)) = uint8(v70)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+12)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v31 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v31 + int32(88)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	F_get_typlenbyval(m, v87, v31+int32(16), v31+int32(14))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_eqjoinsel_find_matches[0]))
	v97 = F_MemoryContextAllocZero(m, v95, int32(32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+24)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v97)+28)) = v31 + int32(4)
	v104 = float64(4.294967296e+09)
	v107 = base.F64_div(base.F64_convert_i32_u(v59), float64(0.9))
	if base.F64_ge(v107, v104) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v110 = v104
	goto L25
L24:
	;
	v110 = v107
	goto L25
L25:
	;
	v111 = base.I64_trunc_sat_f64_u(v110)
	if base.Ui64(v111) <= base.Ui64(int64(2)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v114 = int64(2)
	goto L28
L27:
	;
	v114 = v111
	goto L28
L28:
	;
	v115 = int64(1)
	if v114&(v114-v115) == int64(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v125 = v114
	goto L31
L30:
	;
	v125 = v115 << (uint(int64(64)-base.I64_clz(v114)) % 64)
	goto L31
L31:
	;
	if base.Ui64(v125*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v134 = F_MemoryContextAllocExtended(m, v95, base.I32_wrap_i64(v125)*int32(24), int32(5))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L19
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+20)) = v134
	v137 = int64(1)
	if v125&(v125-v137) == int64(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v147 = v125
	goto L38
L37:
	;
	v147 = v137 << (uint(int64(64)-base.I64_clz(v125)) % 64)
	goto L38
L38:
	;
	if base.Ui64(v147*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = base.I32_wrap_i64(v147) - int32(1)
	if v147 == int64(4294967296) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	goto L1
L42:
	;
	v164 = int32(-85899346)
	goto L44
L43:
	;
	v164 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v147), float64(0.9)))
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v164
	if v59 <= int32(0) {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v188 = v14
	goto L46
L46:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v196+v188<<(uint(int32(3))%32))))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+16)) = uint8(v203)
	*(*int64)(unsafe.Add(mBase, uint32(v202)+24)) = v200
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v208 = m.T0[v207].(func(*base.Module, int32) int64)(m, v202)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	v210 = base.I32_wrap_i64(v208)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v229 = base.B2i32(base.Ui32(v211) < base.Ui32(v212))
	goto L49
L49:
	;
	if v229 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v855 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v855
	v229 = v855
	goto L49
L52:
	;
	v825 = v188 + int32(1)
	if v825 != v59 {
		v188 = v825
		goto L46
	} else {
		goto L148
	}
L53:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v788 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v787 + v788
	*(*uint8)(unsafe.Add(mBase, uint32(v767)+16)) = uint8(v788)
	*(*int32)(unsafe.Add(mBase, uint32(v767)+12)) = v210
	*(*int64)(unsafe.Add(mBase, uint32(v767))) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v767)+8)) = v188
	goto L52
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L19
	} else {
		goto L145
	}
L55:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	if v244 == int64(4294967296) {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v500 = int32(0)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v503 = v502 & v210
	v506 = v501 + v503*int32(24)
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+16)))
	if v507 == v500 {
		v767 = v506
		goto L53
	} else {
		goto L99
	}
L58:
	;
	v247 = int32(0)
	v249 = int64(2)
	v251 = v244 << (uint(int64(1)) % 64)
	if base.Ui64(v251) <= base.Ui64(v249) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v229 = int32(1)
	goto L49
L60:
	;
	v254 = v249
	goto L62
L61:
	;
	v254 = v251
	goto L62
L62:
	;
	v255 = int64(1)
	if v254&(v254-v255) == int64(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v265 = v254
	goto L65
L64:
	;
	v265 = v255 << (uint(int64(64)-base.I64_clz(v254)) % 64)
	goto L65
L65:
	;
	if base.Ui64(v265*int64(24)) < base.Ui64(int64(2147483647)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	v277 = F_MemoryContextAllocExtended(m, v272, base.I32_wrap_i64(v265)*int32(24), int32(5))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L19
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L1
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+20)) = v277
	v280 = int64(1)
	if v265&(v265-v280) == int64(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v290 = v265
	goto L72
L71:
	;
	v290 = v280 << (uint(int64(64)-base.I64_clz(v265)) % 64)
	goto L72
L72:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v290*int64(24)) {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v290
	v298 = base.I32_wrap_i64(v290) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v298
	if v290 == int64(4294967296) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v307 = int32(-85899346)
	goto L76
L75:
	;
	v307 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v290), float64(0.9)))
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v307
	if v271 != int64(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v317 = v247
	goto L81
L78:
	;
	goto L79
L79:
	;
	F_pfree(m, v270)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L19
	} else {
		goto L98
	}
L80:
	;
	v360 = v353
	v371 = v247
	goto L86
L81:
	;
	v341 = v270 + v317*int32(24)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+16)))
	if v342 != int32(1) {
		v353 = v317
		goto L80
	} else {
		goto L83
	}
L82:
	;
	v353 = int32(0)
	goto L80
L83:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	if v345&v298 == v317 {
		v353 = v317
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v349 = v317 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v349)) < base.Ui64(v271) {
		v317 = v349
		goto L81
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	v384 = v270 + v360*int32(24)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+16)))
	if v385 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L79
L88:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v384)+12))
	v400 = v389
	goto L91
L89:
	;
	goto L90
L90:
	;
	v460 = v360 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v460)) < base.Ui64(v271) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v418 = v400 & v388
	v423 = v277 + v418*int32(24)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+16)))
	if v424 != 0 {
		v400 = v418 + int32(1)
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v384)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v423)+16)) = v425
	v427 = *(*int64)(unsafe.Add(mBase, uint32(v384)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v423)+8)) = v427
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v384)))
	*(*int64)(unsafe.Add(mBase, uint32(v423))) = v429
	goto L90
L93:
	;
	goto L92
L94:
	;
	v464 = v460
	goto L96
L95:
	;
	v464 = int32(0)
	goto L96
L96:
	;
	v466 = v371 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v466)) < base.Ui64(v271) {
		v360 = v464
		v371 = v466
		goto L86
	} else {
		goto L97
	}
L97:
	;
	goto L87
L98:
	;
	goto L59
L99:
	;
	v516 = v500
	v518 = v506
	v520 = v503
	goto L100
L100:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	if v538 != v210 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v767 = v744
	goto L53
L102:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v518)+12))
	v571 = v569 & v570
	if base.Ui32(v520) < base.Ui32(v571) {
		goto L118
	} else {
		goto L119
	}
L103:
	;
	v540 = *(*int64)(unsafe.Add(mBase, uint32(v518)))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+9)))
	if v542 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+10)))
	v546 = int32(*(*int16)(unsafe.Add(mBase, uint32(v541)+12)))
	v547 = F_datum_image_eq(m, v540, v200, v545, v546)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L19
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+8)))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	v553 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v552)+16)) = uint8(v553)
	if v551 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	if v547 == int32(0) {
		goto L102
	} else {
		goto L108
	}
L108:
	;
	goto L52
L109:
	;
	v555 = v200
	goto L111
L110:
	;
	v555 = v540
	goto L111
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v552)+40)) = v555
	if v551 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v557 = v540
	goto L114
L113:
	;
	v557 = v200
	goto L114
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v552)+24)) = v557
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	v561 = m.T0[v560].(func(*base.Module, int32) int64)(m, v552)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L19
	} else {
		goto L115
	}
L115:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+16)))
	if v563 != 0 {
		goto L102
	} else {
		goto L116
	}
L116:
	;
	if v561 != int64(0) {
		goto L52
	} else {
		goto L117
	}
L117:
	;
	goto L102
L118:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v575 = v520 + v573
	goto L120
L119:
	;
	v575 = v520
	goto L120
L120:
	;
	v578 = v569 & (v520 + int32(1))
	if base.Ui32(v575-v571) < base.Ui32(v516) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v584 = v501 + v578*int32(24)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584)+16)))
	if v585 != 0 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v732 = v516 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v732) {
		goto L140
	} else {
		goto L141
	}
L124:
	;
	v601 = v578
	v602 = int32(0)
	goto L127
L125:
	;
	v638 = v584
	v647 = v578
	goto L126
L126:
	;
	if v520 != v647 {
		goto L134
	} else {
		goto L135
	}
L127:
	;
	v615 = v602 + int32(1)
	if int32(151) <= v615 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v638 = v630
	v647 = v627
	goto L126
L129:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v620 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v618), base.F64_convert_i64_u(v620)), float64(0.1)) != 0 {
		goto L51
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v627 = (v601 + int32(1)) & v569
	v630 = v501 + v627*int32(24)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+16)))
	if v631 != 0 {
		v601 = v627
		v602 = v615
		goto L127
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	goto L128
L134:
	;
	v667 = v638
	v676 = v647
	goto L137
L135:
	;
	goto L136
L136:
	;
	v767 = v518
	goto L53
L137:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v692 = v689 & (v676 - int32(1))
	v695 = v501 + v692*int32(24)
	v696 = *(*int64)(unsafe.Add(mBase, uint32(v695)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+16)) = v696
	v698 = *(*int64)(unsafe.Add(mBase, uint32(v695)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+8)) = v698
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v695)))
	*(*int64)(unsafe.Add(mBase, uint32(v667))) = v700
	if v520 != v692 {
		v667 = v695
		v676 = v692
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L136
L139:
	;
	goto L138
L140:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v737 = *(*int64)(unsafe.Add(mBase, uint32(v97)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v735), base.F64_convert_i64_u(v737)), float64(0.1)) != 0 {
		goto L51
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v744 = v501 + v578*int32(24)
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+16)))
	if v745 != 0 {
		v516 = v732
		v518 = v744
		v520 = v578
		goto L100
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	goto L101
L145:
	;
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_find_matches_0), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L19
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_find_matches_1), int32(635), int32(_a_F_eqjoinsel_find_matches_2))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L19
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	goto L3
L149:
	;
	v862 = int32(40)
	goto L151
L150:
	;
	v862 = int32(24)
	goto L151
L151:
	;
	if l4 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v868 = int32(24)
	goto L154
L153:
	;
	v868 = int32(40)
	goto L154
L154:
	;
	v886 = v14
	v887 = v14
	v897 = v28
	goto L155
L155:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v902 = *(*int64)(unsafe.Add(mBase, uint32(v898+v886<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v859+v862))) = v902
	if l8 <= int32(0) {
		v994 = v887
		v1004 = v897
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v1290 = v994
	v1300 = v1004
	goto L2
L157:
	;
	v1006 = v886 + int32(1)
	if v1006 != l7 {
		v886 = v1006
		v887 = v994
		v897 = v1004
		goto L155
	} else {
		goto L166
	}
L158:
	;
	v905 = int32(0)
	goto L159
L159:
	;
	v933 = v905 + l10
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933))))
	if v934 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v994 = v887
	v1004 = v897
	goto L157
L161:
	;
	v975 = v905 + int32(1)
	if v975 != l8 {
		v905 = v975
		goto L159
	} else {
		goto L165
	}
L162:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v935+v905<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v868+v859))) = v939
	v941 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+104)) = uint8(v941)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v31)+88))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
	v947 = m.T0[v946].(func(*base.Module, int32) int64)(m, v31+int32(88))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L19
	} else {
		goto L163
	}
L163:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+104)))
	if v949|base.B2i32(v947 == int64(0)) != 0 {
		goto L161
	} else {
		goto L164
	}
L164:
	;
	v953 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v933))) = uint8(v953)
	*(*uint8)(unsafe.Add(mBase, uint32(l9+v886))) = uint8(v953)
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v959 = int32(2)
	v962 = *(*float32)(unsafe.Add(mBase, uint32(v958+v886<<(uint(v959)%32))))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	v967 = *(*float32)(unsafe.Add(mBase, uint32(v963+v905<<(uint(v959)%32))))
	v994 = v887 + v953
	v1004 = base.F64_add(v897, base.F64_promote_f32(base.F32_mul(v962, v967)))
	goto L157
L165:
	;
	goto L160
L166:
	;
	goto L156
L167:
	;
	if v57 != 0 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L169
L169:
	;
	if v61 <= int32(0) {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	v1040 = l3
	goto L172
L171:
	;
	v1040 = l2
	goto L172
L172:
	;
	v1042 = v31 + int32(20)
	F_fmgr_info(m, v1040, v1042)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L19
	} else {
		goto L173
	}
L173:
	;
	v1045 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+80)) = uint8(v1045)
	v1047 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+66)) = uint16(v1047)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+64)) = uint8(v1045)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v31)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v1042
	goto L169
L174:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	F_pfree(m, v1268)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L19
	} else {
		goto L205
	}
L175:
	;
	v1257 = v1036
	v1267 = float64(0)
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1077 = int32(0)
	v1078 = v1036
	v1088 = float64(0)
	goto L178
L178:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v1093 = *(*int64)(unsafe.Add(mBase, uint32(v1089+v1077<<(uint(int32(3))%32))))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+4))
	v1096 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1095)+16)) = uint8(v1096)
	*(*int64)(unsafe.Add(mBase, uint32(v1095)+24)) = v1093
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1095)))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1099)))
	v1101 = m.T0[v1100].(func(*base.Module, int32) int64)(m, v1095)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L19
	} else {
		goto L180
	}
L179:
	;
	v1257 = v1226
	v1267 = v1236
	goto L174
L180:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v1104 = base.I32_wrap_i64(v1101)
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v1106 = v1104 & v1105
	v1109 = v1103 + v1106*int32(24)
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1109)+16)))
	if v1110 == int32(0) {
		v1226 = v1078
		v1236 = v1088
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1238 = v1077 + int32(1)
	if v1238 != v61 {
		v1077 = v1238
		v1078 = v1226
		v1088 = v1236
		goto L178
	} else {
		goto L204
	}
L182:
	;
	v1123 = v1106
	v1128 = v1109
	goto L183
L183:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+12))
	if v1141 != v1104 {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+8))
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v1183))))
	if v1185 != 0 {
		v1226 = v1078
		v1236 = v1088
		goto L181
	} else {
		goto L203
	}
L185:
	;
	goto L184
L186:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v1176 = v1173 & (v1123 + int32(1))
	v1179 = v1172 + v1176*int32(24)
	v1180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179)+16)))
	if v1180 != 0 {
		v1123 = v1176
		v1128 = v1179
		goto L183
	} else {
		goto L202
	}
L187:
	;
	v1143 = *(*int64)(unsafe.Add(mBase, uint32(v1128)))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+9)))
	if v1145 == int32(1) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+10)))
	v1149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1144)+12)))
	v1150 = F_datum_image_eq(m, v1143, v1093, v1148, v1149)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L19
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+8)))
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1144)))
	v1156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1155)+16)) = uint8(v1156)
	if v1154 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	if v1150 == int32(0) {
		goto L186
	} else {
		goto L192
	}
L192:
	;
	goto L185
L193:
	;
	v1158 = v1093
	goto L195
L194:
	;
	v1158 = v1143
	goto L195
L195:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1155)+40)) = v1158
	if v1154 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1160 = v1143
	goto L198
L197:
	;
	v1160 = v1093
	goto L198
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1155)+24)) = v1160
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1155)))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1162)))
	v1164 = m.T0[v1163].(func(*base.Module, int32) int64)(m, v1155)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L19
	} else {
		goto L199
	}
L199:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1155)+16)))
	if v1166 != 0 {
		goto L186
	} else {
		goto L200
	}
L200:
	;
	if v1164 != int64(0) {
		goto L185
	} else {
		goto L201
	}
L201:
	;
	goto L186
L202:
	;
	v1226 = v1078
	v1236 = v1088
	goto L181
L203:
	;
	v1187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1077+v62))) = uint8(v1187)
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v1189))) = uint8(v1187)
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+8))
	v1195 = int32(2)
	v1198 = *(*float32)(unsafe.Add(mBase, uint32(v1193+v1194<<(uint(v1195)%32))))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	v1203 = *(*float32)(unsafe.Add(mBase, uint32(v1199+v1077<<(uint(v1195)%32))))
	v1226 = v1078 + v1187
	v1236 = base.F64_add(v1088, base.F64_promote_f32(base.F32_mul(v1198, v1203)))
	goto L181
L204:
	;
	goto L179
L205:
	;
	F_pfree(m, v97)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L19
	} else {
		goto L206
	}
L206:
	;
	v1290 = v1257
	v1300 = v1267
	goto L2
L207:
	;
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_find_matches_3), int32(0))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L19
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_find_matches_1), int32(332), int32(_a_F_eqjoinsel_find_matches_4))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L19
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_eqsel(m *base.Module, l0 int32) int64 {
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_eqsel_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_reinterpret_f64(v3)
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
				F_errfinish(m, int32(_a_F_errcode_1), int32(880), int32(_a_F_errcode_2))
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
	F_errfinish(m, int32(_a_F_errcontext_msg_3), int32(1584), int32(_a_F_errcontext_msg_4))
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
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
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
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
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
	v270 = F_fflush(m, int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L46
	} else {
		goto L84
	}
L2:
	;
	v252 = int32(_a_F_errfinish_0)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[0])) = v254 - int32(1)
	v259 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[2])) = v259
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[3])) = v259
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[4])) = v259
	F_pg_re_throw(m)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L46
	} else {
		goto L83
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
	v242 = m.ExcPending
	if v242 != 0 {
		goto L46
	} else {
		goto L80
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
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[1]))
	if v179 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[9])) = v79
	v192 = int32(_a_F_errfinish_3)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[1]))
	v195 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[1])) = v194 - v195
	v198 = int32(_a_F_errfinish_0)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[0])) = v200 - v195
	if v77&int32(-2) == int32(22) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	F_FreeErrorDataContents(m, v71+int32(_a_F_errfinish_4))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L46
	} else {
		goto L63
	}
L60:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[0]))
	if v181 != int32(1) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[10]))
	F_MemoryContextReset(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L46
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	goto L58
L64:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[15]))
	if v209 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if int32(24) <= v77 {
		goto L1
	} else {
		goto L75
	}
L67:
	;
	v218 = F_fflush(m, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L46
	} else {
		goto L70
	}
L68:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[16]))
	if v211 != int32(2) {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[16])) = int32(0)
	goto L67
L70:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[17]))
	if v221 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errfinish[17])) = int32(3)
	goto L73
L72:
	;
	goto L73
L73:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L46
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_errfinish[18]))
	if v233 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L46
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	return
L79:
	;
	goto L78
L80:
	;
	F_errmsg_internal(m, int32(_a_F_errfinish_5), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L46
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_errfinish_6), int32(494), int32(_a_F_errfinish_7))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L46
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
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
	Fn14260(m, l0, l1, int32(_a_F_errhint_0), int32(1515))
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
func F_esperanto_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v902 int32
	_ = v902
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v984 int32
	_ = v984
	var v993 int32
	_ = v993
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1037 int32
	_ = v1037
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1119 int32
	_ = v1119
	var v1128 int32
	_ = v1128
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1222 int32
	_ = v1222
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1307 int32
	_ = v1307
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1398 int32
	_ = v1398
	var v1413 int32
	_ = v1413
	var v1419 int32
	_ = v1419
	var v1424 int32
	_ = v1424
	var v1428 int32
	_ = v1428
	var v1438 int32
	_ = v1438
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1594 int32
	_ = v1594
	var v1610 int32
	_ = v1610
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v9
	v13 = int32(0)
	goto L2
L1:
	;
	return v1654
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v11
	v22 = F_find_among(m, l0, int32(_a_F_esperanto_UTF_8_stem_0), int32(17), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
	if v13 != 0 {
		v1654 = int32(0)
		goto L1
	} else {
		goto L65
	}
L4:
	;
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
	switch v22 - int32(1) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	case 5:
		goto L16
	case 6:
		goto L15
	case 7:
		goto L14
	case 8:
		goto L13
	case 9:
		goto L12
	case 10:
		goto L11
	case 11:
		goto L10
	case 12:
		goto L9
	case 13:
		goto L8
	default:
		v163 = v13
		goto L7
	}
L7:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v165
	v13 = v163
	goto L2
L8:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L46
L9:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v104
	v13 = int32(0)
	goto L2
L10:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v102
	v13 = int32(1)
	goto L2
L11:
	;
	v94 = int32(1)
	v97 = F_slice_from_s(m, l0, v94, int32(_a_F_esperanto_UTF_8_stem_1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L42
	}
L12:
	;
	v87 = int32(1)
	v90 = F_slice_from_s(m, l0, v87, int32(_a_F_esperanto_UTF_8_stem_2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L40
	}
L13:
	;
	v80 = int32(1)
	v83 = F_slice_from_s(m, l0, v80, int32(_a_F_esperanto_UTF_8_stem_3))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L38
	}
L14:
	;
	v73 = int32(1)
	v76 = F_slice_from_s(m, l0, v73, int32(_a_F_esperanto_UTF_8_stem_4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L36
	}
L15:
	;
	v66 = int32(1)
	v69 = F_slice_from_s(m, l0, v66, int32(_a_F_esperanto_UTF_8_stem_5))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L34
	}
L16:
	;
	v62 = F_slice_from_s(m, l0, int32(2), int32(_a_F_esperanto_UTF_8_stem_6))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L32
	}
L17:
	;
	v56 = F_slice_from_s(m, l0, int32(2), int32(_a_F_esperanto_UTF_8_stem_7))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L30
	}
L18:
	;
	v50 = F_slice_from_s(m, l0, int32(2), int32(_a_F_esperanto_UTF_8_stem_8))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L28
	}
L19:
	;
	v44 = F_slice_from_s(m, l0, int32(2), int32(_a_F_esperanto_UTF_8_stem_9))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L26
	}
L20:
	;
	v38 = F_slice_from_s(m, l0, int32(2), int32(_a_F_esperanto_UTF_8_stem_10))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L24
	}
L21:
	;
	v32 = F_slice_from_s(m, l0, int32(2), int32(_a_F_esperanto_UTF_8_stem_11))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	if int32(0) <= v32 {
		v163 = v13
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v1654 = v32
	goto L1
L24:
	;
	if int32(0) <= v38 {
		v163 = v13
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v1654 = v38
	goto L1
L26:
	;
	if int32(0) <= v44 {
		v163 = v13
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v1654 = v44
	goto L1
L28:
	;
	if int32(0) <= v50 {
		v163 = v13
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v1654 = v50
	goto L1
L30:
	;
	if int32(0) <= v56 {
		v163 = v13
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v1654 = v56
	goto L1
L32:
	;
	if int32(0) <= v62 {
		v163 = v13
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v1654 = v62
	goto L1
L34:
	;
	if int32(0) <= v69 {
		v163 = v66
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v1654 = v69
	goto L1
L36:
	;
	if int32(0) <= v76 {
		v163 = v73
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v1654 = v76
	goto L1
L38:
	;
	if int32(0) <= v83 {
		v163 = v80
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v1654 = v83
	goto L1
L40:
	;
	if int32(0) <= v90 {
		v163 = v87
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v1654 = v90
	goto L1
L42:
	;
	if int32(0) <= v97 {
		v163 = v94
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v1654 = v97
	goto L1
L44:
	;
	if v158 < int32(0) {
		goto L4
	} else {
		goto L64
	}
L46:
	;
	goto L47
L47:
	;
	goto L48
L48:
	;
	v113 = v26
	v115 = int32(1)
	goto L51
L50:
	;
	v158 = v143
	goto L44
L51:
	;
	if v106 <= v113 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L50
L53:
	;
	v158 = int32(-1)
	goto L44
L54:
	;
	goto L55
L55:
	;
	v120 = v113 + int32(1)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v113))))
	if base.Ui32(v122) < base.Ui32(int32(192)) {
		v143 = v120
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v144 = int32(1)
	if v144 < v115 {
		v113 = v143
		v115 = v115 - v144
		goto L51
	} else {
		goto L63
	}
L57:
	;
	if v106 <= v120 {
		v143 = v120
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v129 = v120
	goto L59
L59:
	;
	v132 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105+v129))))
	if int32(-65) < v132 {
		v143 = v129
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v143 = v106
	goto L56
L61:
	;
	v136 = v129 + int32(1)
	if v136 != v106 {
		v129 = v136
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	goto L52
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v158
	v163 = v13
	goto L7
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v170 == v9 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v248
	v252 = int32(1)
	v254 = v248 - v252
	if v254 <= v9 {
		v302 = v252
		goto L87
	} else {
		goto L88
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	v248 = v9
	goto L66
L68:
	;
	goto L69
L69:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v9))))
	if v176 != int32(39) {
		v234 = v170
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	if v234 <= v9 {
		v248 = v234
		goto L66
	} else {
		goto L84
	}
L71:
	;
	v180 = v9 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v180
	v183 = int32(2)
	v185 = int32(0)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v187-v180 < v183 {
		v197 = v185
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v197 == int32(0) {
		v234 = v198
		goto L70
	} else {
		goto L76
	}
L73:
	;
	goto L72
L74:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v193 = F_memcmp(m, v191+v180, int32(_a_F_esperanto_UTF_8_stem_12), v183)
	mBase = m.M
	if v193 != 0 {
		v197 = v185
		goto L73
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9 + int32(3)
	v197 = int32(1)
	goto L73
L76:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v198 <= v201 {
		v234 = v198
		goto L70
	} else {
		goto L77
	}
L77:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v201))))
	if base.B2i32(v205&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v205)%32)&int32(_a_F_esperanto_UTF_8_stem_13) == int32(0)) != 0 {
		v234 = v198
		goto L70
	} else {
		goto L78
	}
L78:
	;
	v220 = F_find_among(m, l0, int32(_a_F_esperanto_UTF_8_stem_14), int32(6), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v220 == int32(0) {
		v234 = v222
		goto L70
	} else {
		goto L80
	}
L80:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v225 < v222 {
		v234 = v222
		goto L70
	} else {
		goto L81
	}
L81:
	;
	v229 = F_slice_from_s(m, l0, int32(1), int32(_a_F_esperanto_UTF_8_stem_15))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	if v229 < int32(0) {
		v1654 = v229
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v234 = v233
	goto L70
L84:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+v234-int32(1)))))
	v248 = v234 - base.B2i32(v244 == int32(110))
	goto L66
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v311
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v311 <= v315 {
		goto L111
	} else {
		goto L112
	}
L86:
	;
	if v288 < v289 {
		goto L108
	} else {
		goto L109
	}
L87:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v303
	v311 = v303
	v312 = v302
	v313 = v303
	goto L85
L88:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256+v254))))
	if v258 != int32(105) {
		v302 = v252
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v264 = F_find_among_b(m, l0, int32(_a_F_esperanto_UTF_8_stem_16), int32(17), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	if v264 == int32(0) {
		v302 = v252
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v269 < v268 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271+v268-int32(1)))))
	if v275 != int32(45) {
		v302 = v252
		goto L87
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v282 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v282 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v268 - int32(1)
	goto L94
L96:
	;
	v288 = int32(1)
	goto L98
L97:
	;
	v288 = v282 >> (uint(int32(31)) % 32) & v282
	goto L98
L98:
	;
	v289 = int32(0)
	v290 = base.B2i32(v13 == v289)
	v292 = base.B2i32(v288 < v289)
	if v288 < v289 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v293 = v288
	goto L101
L100:
	;
	v293 = v290
	goto L101
L101:
	;
	if v288 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v294 = v293
	goto L104
L103:
	;
	v294 = v290
	goto L104
L104:
	;
	if v288 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v298 = int32(base.Ui32(v288) >> (uint(int32(31)) % 32))
	goto L107
L106:
	;
	v298 = int32(2)
	goto L107
L107:
	;
	switch v298 {
	case 0:
		v1654 = v298
		goto L1
	default:
		goto L86
	case 2:
		v302 = v294
		goto L87
	}
L108:
	;
	return v294
L109:
	;
	goto L110
L110:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v311 = v306
	v312 = v294
	v313 = v307
	goto L85
L111:
	;
	v435 = v311 - v313
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v437 = v435 + v436
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v437
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v444 = base.B2i32(v437 <= v443)
	if v444 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L112:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v318 = v317 + v311
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318-int32(1)))))
	if v321 != int32(39) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v325 = v311 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v325
	if v325 <= v315 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v325
	v346 = int32(2)
	v348 = int32(0)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v325-v351 < v346 {
		v361 = v348
		goto L121
	} else {
		goto L122
	}
L115:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318-int32(2)))))
	if v331 != int32(108) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v335 = v311 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v335
	if v315 < v335 {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	v340 = F_slice_from_s(m, l0, int32(1), int32(_a_F_esperanto_UTF_8_stem_17))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	if int32(0) <= v340 {
		goto L111
	} else {
		goto L119
	}
L119:
	;
	v1654 = v340
	goto L1
L120:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v361 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L120
L122:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v357 = F_memcmp(m, v354+v325-v346, int32(_a_F_esperanto_UTF_8_stem_18), v346)
	mBase = m.M
	if v357 != 0 {
		v361 = v348
		goto L121
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v325 - v346
	v361 = int32(1)
	goto L121
L124:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v374 = v313 - v325
	v375 = v373 - v374
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v375
	if v375-int32(2) <= v362 {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v362 < v365 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v369 = F_slice_from_s(m, l0, int32(1), int32(_a_F_esperanto_UTF_8_stem_19))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L127
	}
L127:
	;
	if int32(0) <= v369 {
		goto L111
	} else {
		goto L128
	}
L128:
	;
	v1654 = v369
	goto L1
L129:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v423 - v374
	v428 = F_slice_from_s(m, l0, int32(1), int32(_a_F_esperanto_UTF_8_stem_20))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L140
	}
L130:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v382 = int32(1)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380+v375-v382))))
	if base.B2i32(v384&int32(224) != int32(96))|base.B2i32(v382<<(uint(v384)%32)&int32(68438676) == int32(0)) != 0 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v399 = F_find_among_b(m, l0, int32(_a_F_esperanto_UTF_8_stem_21), int32(20), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	if v399 == int32(0) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v404 < v403 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v403-int32(1)))))
	if v410 != int32(45) {
		goto L129
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v418 = F_slice_from_s(m, l0, int32(3), int32(_a_F_esperanto_UTF_8_stem_22))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v403 - int32(1)
	goto L136
L138:
	;
	if int32(0) <= v418 {
		goto L111
	} else {
		goto L139
	}
L139:
	;
	v1654 = v418
	goto L1
L140:
	;
	if v428 < int32(0) {
		v1654 = v428
		goto L1
	} else {
		goto L141
	}
L141:
	;
	goto L111
L142:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v437-int32(1)))))
	v455 = v437 - base.B2i32(v451 == int32(110))
	goto L144
L143:
	;
	v455 = v437
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v455
	if v455 <= v443 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v678 = base.B2i32(v674 < int32(0))
	if v674 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L146:
	;
	v629 = int32(0)
	if v627 <= v628 {
		v674 = v629
		goto L145
	} else {
		goto L180
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v437
	if v444 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459+v455-int32(1)))))
	if v463 != int32(101) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v467 = v455 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v467
	v627 = v467
	v628 = v443
	goto L146
L150:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v437-int32(1)))))
	v480 = v437 - base.B2i32(v476 == int32(110))
	goto L152
L151:
	;
	v480 = v437
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v480
	if v443 < v480 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483+v480-int32(1)))))
	v491 = v480 - base.B2i32(v487 == int32(106))
	goto L155
L154:
	;
	v491 = v480
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v491
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L158
L156:
	;
	if v623 != 0 {
		v674 = int32(0)
		goto L145
	} else {
		goto L179
	}
L157:
	;
	v623 = v616
	goto L156
L158:
	;
	if v491 <= v508 {
		v616 = int32(-1)
		goto L157
	} else {
		goto L160
	}
L159:
	;
	v616 = int32(0)
	goto L157
L160:
	;
	v525 = int32(1)
	v526 = v491 - v525
	v528 = int32(*(*int8)(unsafe.Add(mBase, uint32(v509+v526))))
	v530 = v528 & int32(255)
	if base.B2i32(v526 == v508)|base.B2i32(int32(0) <= v528) != 0 {
		v588 = v530
		v592 = v525
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if int32(117) < v588 {
		goto L169
	} else {
		goto L170
	}
L162:
	;
	v537 = v530 & int32(63)
	v539 = v491 - int32(2)
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509+v539))))
	v543 = v541 << (uint(int32(6)) % 32)
	if base.B2i32(v539 != v508)&base.B2i32(base.Ui32(v541) < base.Ui32(int32(192))) == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v588 = v543&int32(1984) | v537
	v592 = int32(2)
	goto L161
L164:
	;
	goto L165
L165:
	;
	v556 = v543&int32(4032) | v537
	v558 = v491 - int32(3)
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509+v558))))
	if base.B2i32(v558 != v508)&base.B2i32(base.Ui32(v560) < base.Ui32(int32(224))) == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v588 = v560<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_23) | v556
	v592 = int32(3)
	goto L161
L167:
	;
	goto L168
L168:
	;
	v578 = int32(4)
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491+v509-v578))))
	v588 = v560<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_24) | v580&int32(7)<<(uint(int32(18))%32) | v556
	v592 = v578
	goto L161
L169:
	;
	v623 = v592
	goto L156
L170:
	;
	goto L171
L171:
	;
	v594 = v588 - int32(97)
	if v594 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v623 = v592
	goto L156
L173:
	;
	goto L174
L174:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v594)>>(uint(int32(3))%32)))+uint32(_c_F_esperanto_UTF_8_stem[0]))))
	if int32(base.Ui32(v600)>>(uint(v594&int32(7))%32))&int32(1) == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v623 = v592
	goto L156
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v491 - v592
	goto L178
L178:
	;
	goto L159
L179:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v627 = v625
	v628 = v624
	goto L146
L180:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631+v627-int32(1)))))
	if v635 != int32(105) {
		v674 = v629
		goto L145
	} else {
		goto L181
	}
L181:
	;
	v639 = v627 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v639
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v645 = F_find_among_b(m, l0, int32(_a_F_esperanto_UTF_8_stem_25), int32(7), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L5
	} else {
		goto L183
	}
L182:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v653 < v652 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	if v645 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v652 = v647
	goto L182
L185:
	;
	goto L186
L186:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v650 = v648 + (v639 - v641)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v650
	v652 = v650
	goto L182
L187:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655+v652-int32(1)))))
	if v659 != int32(45) {
		v674 = v629
		goto L145
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v662 + (v437 - v436)
	v667 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v667 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	goto L189
L191:
	;
	v673 = int32(1)
	goto L193
L192:
	;
	v673 = v667 >> (uint(int32(31)) % 32) & v667
	goto L193
L193:
	;
	v674 = v673
	goto L145
L194:
	;
	v679 = v674
	goto L196
L195:
	;
	v679 = v312
	goto L196
L196:
	;
	if v674 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v680 = v679
	goto L199
L198:
	;
	v680 = v312
	goto L199
L199:
	;
	if v674 != 0 {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	v693 = int32(0)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v696-int32(2) <= v695 {
		v738 = v693
		goto L209
	} else {
		goto L210
	}
L201:
	;
	if v674 < int32(0) {
		goto L206
	} else {
		goto L207
	}
L202:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v686 = v685 + v435
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v686
	v691 = v685
	v692 = v686
	goto L200
L203:
	;
	v684 = int32(base.Ui32(v674) >> (uint(int32(31)) % 32))
	goto L205
L204:
	;
	v684 = int32(3)
	goto L205
L205:
	;
	switch v684 {
	case 0:
		v1654 = v684
		goto L1
	default:
		goto L201
	case 3:
		goto L202
	}
L206:
	;
	return v680
L207:
	;
	goto L208
L208:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v691 = v690
	v692 = v689
	goto L200
L209:
	;
	if v738 != 0 {
		v1654 = v693
		goto L1
	} else {
		goto L218
	}
L210:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v702 = int32(1)
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v700+v696-v702))))
	if base.B2i32(v704&int32(224) != int32(96))|base.B2i32(v702<<(uint(v704)%32)&int32(_a_F_esperanto_UTF_8_stem_26) == int32(0)) != 0 {
		v738 = v693
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v719 = F_find_among_b(m, l0, int32(_a_F_esperanto_UTF_8_stem_27), int32(24), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	if v719 == int32(0) {
		v738 = v693
		goto L209
	} else {
		goto L213
	}
L213:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v724 < v723 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726+v723-int32(1)))))
	if v730 != int32(45) {
		v738 = v693
		goto L209
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v738 = int32(1)
	goto L209
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v723 - int32(1)
	goto L216
L218:
	;
	v739 = v692 - v691
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v741 = v739 + v740
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v741
	v743 = int32(0)
	v746 = v741 - int32(1)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v746 <= v747 {
		v788 = v743
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if v788 != 0 {
		v1654 = v693
		goto L1
	} else {
		goto L227
	}
L220:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749+v746))))
	v753 = v751 - int32(105)
	v754 = int32(0)
	if base.B2i32(v753 == v754)|base.B2i32(v753 == int32(12)) == v754 {
		v788 = v743
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v764 = F_find_among_b(m, l0, int32(_a_F_esperanto_UTF_8_stem_28), int32(3), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L5
	} else {
		goto L222
	}
L222:
	;
	if v764 == int32(0) {
		v788 = v743
		goto L219
	} else {
		goto L223
	}
L223:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v769-int32(2) <= v768 {
		v788 = v743
		goto L219
	} else {
		goto L224
	}
L224:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773+v769-int32(1)))))
	switch v777 - int32(107) {
	case 0, 9:
		goto L225
	default:
		v788 = v743
		goto L219
	}
L225:
	;
	v783 = F_find_among_b(m, l0, int32(_a_F_esperanto_UTF_8_stem_29), int32(2), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L5
	} else {
		goto L226
	}
L226:
	;
	v788 = base.B2i32(v783 != int32(0))
	goto L219
L227:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v790 = v789 + v739
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v790
	v792 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v790
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v795 < v790 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v797+v790-int32(1)))))
	v805 = v790 - base.B2i32(v801 == int32(110))
	goto L230
L229:
	;
	v805 = v790
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v805
	if v795 < v805 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v808+v805-int32(1)))))
	v816 = v805 - base.B2i32(v812 == int32(106))
	goto L233
L232:
	;
	v816 = v805
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v816
	if v816-int32(2) <= v795 {
		v858 = v792
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if v858 != 0 {
		goto L249
	} else {
		goto L250
	}
L235:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v822+v816-int32(1)))))
	if v826 != int32(117) {
		v858 = v792
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v832 = F_find_among_b(m, l0, int32(_a_F_esperanto_UTF_8_stem_30), int32(2), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L5
	} else {
		goto L237
	}
L237:
	;
	if v832 == int32(0) {
		v858 = v792
		goto L234
	} else {
		goto L238
	}
L238:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v837 < v836 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839+v836-int32(1)))))
	if v843 != int32(45) {
		v858 = v792
		goto L234
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v850 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v850 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v836 - int32(1)
	goto L241
L243:
	;
	v856 = int32(1)
	goto L245
L244:
	;
	v856 = v850 >> (uint(int32(31)) % 32) & v850
	goto L245
L245:
	;
	v858 = v856
	goto L234
L246:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v902 = v877
	goto L266
L247:
	;
	v867 = base.B2i32(v858 < int32(0))
	if v858 < int32(0) {
		goto L252
	} else {
		goto L253
	}
L248:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v864 = v863 + v739
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v864
	v874 = v864
	v875 = v863
	goto L246
L249:
	;
	v862 = int32(base.Ui32(v858) >> (uint(int32(31)) % 32))
	goto L251
L250:
	;
	v862 = int32(6)
	goto L251
L251:
	;
	switch v862 {
	case 0:
		v1654 = v862
		goto L1
	default:
		goto L247
	case 6:
		goto L248
	}
L252:
	;
	if v858 < int32(0) {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	goto L254
L254:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v874 = v871
	v875 = v872
	goto L246
L255:
	;
	v868 = v858
	goto L257
L256:
	;
	v868 = v680
	goto L257
L257:
	;
	if v858 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v869 = v868
	goto L260
L259:
	;
	v869 = v680
	goto L260
L260:
	;
	return v869
L261:
	;
	if v1438 == int32(0) {
		v1654 = int32(0)
		goto L1
	} else {
		goto L369
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1428
	v1438 = int32(1)
	goto L261
L263:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1428 = v1424 - v1419
	goto L262
L264:
	;
	if int32(0) <= v1008 {
		goto L282
	} else {
		goto L283
	}
L265:
	;
	v1008 = int32(-1)
	goto L264
L266:
	;
	if v902 <= v892 {
		goto L265
	} else {
		goto L268
	}
L268:
	;
	v909 = int32(1)
	v910 = v902 - v909
	v912 = int32(*(*int8)(unsafe.Add(mBase, uint32(v893+v910))))
	v914 = v912 & int32(255)
	if base.B2i32(v910 == v892)|base.B2i32(int32(0) <= v912) != 0 {
		v972 = v914
		v976 = v909
		goto L269
	} else {
		goto L270
	}
L269:
	;
	if int32(117) < v972 {
		goto L277
	} else {
		goto L278
	}
L270:
	;
	v921 = v914 & int32(63)
	v923 = v902 - int32(2)
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893+v923))))
	v927 = v925 << (uint(int32(6)) % 32)
	if base.B2i32(v923 != v892)&base.B2i32(base.Ui32(v925) < base.Ui32(int32(192))) == int32(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v972 = v927&int32(1984) | v921
	v976 = int32(2)
	goto L269
L272:
	;
	goto L273
L273:
	;
	v940 = v927&int32(4032) | v921
	v942 = v902 - int32(3)
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893+v942))))
	if base.B2i32(v942 != v892)&base.B2i32(base.Ui32(v944) < base.Ui32(int32(224))) == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v972 = v944<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_23) | v940
	v976 = int32(3)
	goto L269
L275:
	;
	goto L276
L276:
	;
	v962 = int32(4)
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902+v893-v962))))
	v972 = v944<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_24) | v964&int32(7)<<(uint(int32(18))%32) | v940
	v976 = v962
	goto L269
L277:
	;
	v993 = v902 - v976
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v993
	v902 = v993
	goto L266
L278:
	;
	v978 = v972 - int32(97)
	if v978 < int32(0) {
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v978)>>(uint(int32(3))%32)))+uint32(_c_F_esperanto_UTF_8_stem[1]))))
	if int32(base.Ui32(v984)>>(uint(v978&int32(7))%32))&int32(1) == int32(0) {
		goto L277
	} else {
		goto L280
	}
L280:
	;
	v1008 = v976
	goto L264
L282:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1012 = v1011 - v1008
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1012
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1037 = v1012
	goto L287
L283:
	;
	goto L284
L284:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1148 = v878 - v877
	v1152 = v1147 - v1148
	goto L304
L285:
	;
	if int32(0) <= v1143 {
		v1419 = v1143
		goto L263
	} else {
		goto L303
	}
L286:
	;
	v1143 = int32(-1)
	goto L285
L287:
	;
	if v1037 <= v1027 {
		goto L286
	} else {
		goto L289
	}
L289:
	;
	v1044 = int32(1)
	v1045 = v1037 - v1044
	v1047 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1028+v1045))))
	v1049 = v1047 & int32(255)
	if base.B2i32(v1045 == v1027)|base.B2i32(int32(0) <= v1047) != 0 {
		v1107 = v1049
		v1111 = v1044
		goto L290
	} else {
		goto L291
	}
L290:
	;
	if int32(117) < v1107 {
		goto L298
	} else {
		goto L299
	}
L291:
	;
	v1056 = v1049 & int32(63)
	v1058 = v1037 - int32(2)
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028+v1058))))
	v1062 = v1060 << (uint(int32(6)) % 32)
	if base.B2i32(v1058 != v1027)&base.B2i32(base.Ui32(v1060) < base.Ui32(int32(192))) == int32(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1107 = v1062&int32(1984) | v1056
	v1111 = int32(2)
	goto L290
L293:
	;
	goto L294
L294:
	;
	v1075 = v1062&int32(4032) | v1056
	v1077 = v1037 - int32(3)
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028+v1077))))
	if base.B2i32(v1077 != v1027)&base.B2i32(base.Ui32(v1079) < base.Ui32(int32(224))) == int32(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1107 = v1079<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_23) | v1075
	v1111 = int32(3)
	goto L290
L296:
	;
	goto L297
L297:
	;
	v1097 = int32(4)
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037+v1028-v1097))))
	v1107 = v1079<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_24) | v1099&int32(7)<<(uint(int32(18))%32) | v1075
	v1111 = v1097
	goto L290
L298:
	;
	v1128 = v1037 - v1111
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1128
	v1037 = v1128
	goto L287
L299:
	;
	v1113 = v1107 - int32(97)
	if v1113 < int32(0) {
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1113)>>(uint(int32(3))%32)))+uint32(_c_F_esperanto_UTF_8_stem[1]))))
	if int32(base.Ui32(v1119)>>(uint(v1113&int32(7))%32))&int32(1) == int32(0) {
		goto L298
	} else {
		goto L301
	}
L301:
	;
	v1143 = v1111
	goto L285
L303:
	;
	goto L284
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1152
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1152 <= v1160 {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1281 = v1280 - v1148
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1281
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1307 = v1281
	goto L352
L306:
	;
	goto L305
L307:
	;
	goto L332
L308:
	;
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152+v1159-int32(1)))))
	if v1165 != int32(45) {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1168 = int32(1)
	v1169 = v1152 - v1168
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1169
	goto L312
L310:
	;
	if v1222 < int32(0) {
		goto L306
	} else {
		goto L329
	}
L312:
	;
	goto L313
L313:
	;
	goto L314
L314:
	;
	v1177 = v1169
	v1179 = v1168
	goto L317
L316:
	;
	v1222 = v1204
	goto L310
L317:
	;
	if v1177 <= v1160 {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	goto L316
L319:
	;
	v1222 = int32(-1)
	goto L310
L320:
	;
	goto L321
L321:
	;
	v1184 = v1177 - int32(1)
	v1186 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1159+v1184))))
	if base.B2i32(int32(0) <= v1186)|base.B2i32(v1184 <= v1160) != 0 {
		v1204 = v1184
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1208 = int32(1)
	if v1208 < v1179 {
		v1177 = v1204
		v1179 = v1179 - v1208
		goto L317
	} else {
		goto L328
	}
L323:
	;
	v1192 = v1184
	goto L324
L324:
	;
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1159+v1192))))
	if base.Ui32(int32(191)) < base.Ui32(v1197) {
		v1204 = v1192
		goto L322
	} else {
		goto L326
	}
L325:
	;
	v1204 = v1160
	goto L322
L326:
	;
	v1201 = v1192 - int32(1)
	if v1160 < v1201 {
		v1192 = v1201
		goto L324
	} else {
		goto L327
	}
L327:
	;
	goto L325
L328:
	;
	goto L318
L329:
	;
	v1428 = v1222
	goto L262
L330:
	;
	if int32(0) <= v1276 {
		v1152 = v1276
		goto L304
	} else {
		goto L349
	}
L332:
	;
	goto L333
L333:
	;
	goto L334
L334:
	;
	v1231 = v1152
	v1233 = int32(1)
	goto L337
L336:
	;
	v1276 = v1258
	goto L330
L337:
	;
	if v1231 <= v1160 {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	goto L336
L339:
	;
	v1276 = int32(-1)
	goto L330
L340:
	;
	goto L341
L341:
	;
	v1238 = v1231 - int32(1)
	v1240 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1159+v1238))))
	if base.B2i32(int32(0) <= v1240)|base.B2i32(v1238 <= v1160) != 0 {
		v1258 = v1238
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1262 = int32(1)
	if v1262 < v1233 {
		v1231 = v1258
		v1233 = v1233 - v1262
		goto L337
	} else {
		goto L348
	}
L343:
	;
	v1246 = v1238
	goto L344
L344:
	;
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1159+v1246))))
	if base.Ui32(int32(191)) < base.Ui32(v1251) {
		v1258 = v1246
		goto L342
	} else {
		goto L346
	}
L345:
	;
	v1258 = v1160
	goto L342
L346:
	;
	v1255 = v1246 - int32(1)
	if v1160 < v1255 {
		v1246 = v1255
		goto L344
	} else {
		goto L347
	}
L347:
	;
	goto L345
L348:
	;
	goto L338
L349:
	;
	goto L306
L350:
	;
	if v1413 < int32(0) {
		v1438 = int32(0)
		goto L261
	} else {
		goto L368
	}
L351:
	;
	v1413 = int32(-1)
	goto L350
L352:
	;
	if v1307 <= v1297 {
		goto L351
	} else {
		goto L354
	}
L354:
	;
	v1314 = int32(1)
	v1315 = v1307 - v1314
	v1317 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1298+v1315))))
	v1319 = v1317 & int32(255)
	if base.B2i32(v1315 == v1297)|base.B2i32(int32(0) <= v1317) != 0 {
		v1377 = v1319
		v1381 = v1314
		goto L355
	} else {
		goto L356
	}
L355:
	;
	if int32(57) < v1377 {
		goto L363
	} else {
		goto L364
	}
L356:
	;
	v1326 = v1319 & int32(63)
	v1328 = v1307 - int32(2)
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1298+v1328))))
	v1332 = v1330 << (uint(int32(6)) % 32)
	if base.B2i32(v1328 != v1297)&base.B2i32(base.Ui32(v1330) < base.Ui32(int32(192))) == int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1377 = v1332&int32(1984) | v1326
	v1381 = int32(2)
	goto L355
L358:
	;
	goto L359
L359:
	;
	v1345 = v1332&int32(4032) | v1326
	v1347 = v1307 - int32(3)
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1298+v1347))))
	if base.B2i32(v1347 != v1297)&base.B2i32(base.Ui32(v1349) < base.Ui32(int32(224))) == int32(0) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1377 = v1349<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_23) | v1345
	v1381 = int32(3)
	goto L355
L361:
	;
	goto L362
L362:
	;
	v1367 = int32(4)
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1307+v1298-v1367))))
	v1377 = v1349<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_24) | v1369&int32(7)<<(uint(int32(18))%32) | v1345
	v1381 = v1367
	goto L355
L363:
	;
	v1398 = v1307 - v1381
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1398
	v1307 = v1398
	goto L352
L364:
	;
	v1383 = v1377 - int32(48)
	if v1383 < int32(0) {
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v1389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1383)>>(uint(int32(3))%32)))+uint32(_c_F_esperanto_UTF_8_stem[2]))))
	if int32(base.Ui32(v1389)>>(uint(v1383&int32(7))%32))&int32(1) == int32(0) {
		goto L363
	} else {
		goto L366
	}
L366:
	;
	v1413 = v1381
	goto L350
L368:
	;
	v1419 = v1413
	goto L263
L369:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1448 = v1446 + (v874 - v875)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1448
	v1450 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1448
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1448 <= v1453 {
		v1647 = v1450
		goto L370
	} else {
		goto L371
	}
L370:
	;
	if v1647 <= int32(0) {
		v1654 = v1647
		goto L1
	} else {
		goto L412
	}
L371:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1457 = int32(1)
	v1459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1455+v1448-v1457))))
	if base.B2i32(v1459&int32(224) != int32(96))|base.B2i32(v1457<<(uint(v1459)%32)&int32(_a_F_esperanto_UTF_8_stem_26) == int32(0)) != 0 {
		v1647 = v1450
		goto L370
	} else {
		goto L372
	}
L372:
	;
	v1474 = F_find_among_b(m, l0, int32(_a_F_esperanto_UTF_8_stem_31), int32(19), int32(0))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L5
	} else {
		goto L376
	}
L373:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1627 < v1625 {
		goto L406
	} else {
		goto L407
	}
L374:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1479 < v1478 {
		goto L378
	} else {
		goto L379
	}
L375:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1625 = v1476
	goto L373
L376:
	;
	switch v1474 {
	case 0:
		v1647 = v1474
		goto L370
	case 1:
		goto L374
	default:
		goto L375
	}
L377:
	;
	v1622 = v1620 + (v1478 - v1477)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1622
	v1625 = v1622
	goto L373
L378:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1481+v1478-int32(1)))))
	if v1485 == int32(45) {
		v1620 = v1477
		goto L377
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L384
L381:
	;
	goto L380
L382:
	;
	if v1617 != 0 {
		v1647 = int32(0)
		goto L370
	} else {
		goto L405
	}
L383:
	;
	v1617 = v1610
	goto L382
L384:
	;
	if v1501 <= v1502 {
		v1610 = int32(-1)
		goto L383
	} else {
		goto L386
	}
L385:
	;
	v1610 = int32(0)
	goto L383
L386:
	;
	v1519 = int32(1)
	v1520 = v1501 - v1519
	v1522 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1503+v1520))))
	v1524 = v1522 & int32(255)
	if base.B2i32(v1520 == v1502)|base.B2i32(int32(0) <= v1522) != 0 {
		v1582 = v1524
		v1586 = v1519
		goto L387
	} else {
		goto L388
	}
L387:
	;
	if int32(57) < v1582 {
		goto L395
	} else {
		goto L396
	}
L388:
	;
	v1531 = v1524 & int32(63)
	v1533 = v1501 - int32(2)
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1503+v1533))))
	v1537 = v1535 << (uint(int32(6)) % 32)
	if base.B2i32(v1533 != v1502)&base.B2i32(base.Ui32(v1535) < base.Ui32(int32(192))) == int32(0) {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1582 = v1537&int32(1984) | v1531
	v1586 = int32(2)
	goto L387
L390:
	;
	goto L391
L391:
	;
	v1550 = v1537&int32(4032) | v1531
	v1552 = v1501 - int32(3)
	v1554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1503+v1552))))
	if base.B2i32(v1552 != v1502)&base.B2i32(base.Ui32(v1554) < base.Ui32(int32(224))) == int32(0) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1582 = v1554<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_23) | v1550
	v1586 = int32(3)
	goto L387
L393:
	;
	goto L394
L394:
	;
	v1572 = int32(4)
	v1574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1501+v1503-v1572))))
	v1582 = v1554<<(uint(int32(12))%32)&int32(_a_F_esperanto_UTF_8_stem_24) | v1574&int32(7)<<(uint(int32(18))%32) | v1550
	v1586 = v1572
	goto L387
L395:
	;
	v1617 = v1586
	goto L382
L396:
	;
	goto L397
L397:
	;
	v1588 = v1582 - int32(48)
	if v1588 < int32(0) {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v1617 = v1586
	goto L382
L399:
	;
	goto L400
L400:
	;
	v1594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1588)>>(uint(int32(3))%32)))+uint32(_c_F_esperanto_UTF_8_stem[2]))))
	if int32(base.Ui32(v1594)>>(uint(v1588&int32(7))%32))&int32(1) == int32(0) {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1617 = v1586
	goto L382
L402:
	;
	goto L403
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1501 - v1586
	goto L404
L404:
	;
	goto L385
L405:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1620 = v1618
	goto L377
L406:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1629+v1625-int32(1)))))
	v1637 = v1625 - base.B2i32(v1633 == int32(45))
	goto L408
L407:
	;
	v1637 = v1625
	goto L408
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1637
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1637
	v1641 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1641 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1644 = int32(1)
	goto L411
L410:
	;
	v1644 = v1641
	goto L411
L411:
	;
	v1647 = v1644
	goto L370
L412:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1650
	v1654 = int32(1)
	goto L1
}
func F_estonian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
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
	var v259 int32
	_ = v259
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
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
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v919 int32
	_ = v919
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1152 int32
	_ = v1152
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1194 int32
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1233 int32
	_ = v1233
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1313 int32
	_ = v1313
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1491 int32
	_ = v1491
	var v1507 int32
	_ = v1507
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1675 int32
	_ = v1675
	var v1691 int32
	_ = v1691
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1721 int32
	_ = v1721
	var v1727 int32
	_ = v1727
	var v1744 int32
	_ = v1744
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1823 int32
	_ = v1823
	var v1827 int32
	_ = v1827
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
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1900 int32
	_ = v1900
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1950 int32
	_ = v1950
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1995 int32
	_ = v1995
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2007 int32
	_ = v2007
	var v2023 int32
	_ = v2023
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2041 int32
	_ = v2041
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	v11 = F_find_among(m, l0, int32(_a_F_estonian_UTF_8_stem_0), int32(290), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v2086
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v160 = v6
	goto L67
L3:
	;
	return int32(0)
L4:
	;
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = v17
	goto L2
L6:
	;
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 < v20 {
		v135 = v20
		goto L2
	} else {
		goto L8
	}
L8:
	;
	switch v11 - int32(1) {
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
	v128 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L62
	}
L11:
	;
	v122 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L60
	}
L12:
	;
	v116 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_3))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L58
	}
L13:
	;
	v110 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_4))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L56
	}
L14:
	;
	v104 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_5))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L54
	}
L15:
	;
	v98 = F_slice_from_s(m, l0, int32(7), int32(_a_F_estonian_UTF_8_stem_6))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L52
	}
L16:
	;
	v92 = F_slice_from_s(m, l0, int32(7), int32(_a_F_estonian_UTF_8_stem_7))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L50
	}
L17:
	;
	v86 = F_slice_from_s(m, l0, int32(6), int32(_a_F_estonian_UTF_8_stem_8))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L48
	}
L18:
	;
	v80 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_9))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L46
	}
L19:
	;
	v74 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_10))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L44
	}
L20:
	;
	v68 = F_slice_from_s(m, l0, int32(6), int32(_a_F_estonian_UTF_8_stem_11))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L42
	}
L21:
	;
	v62 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_12))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L40
	}
L22:
	;
	v56 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_13))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L38
	}
L23:
	;
	v50 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_14))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L36
	}
L24:
	;
	v44 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_15))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L34
	}
L25:
	;
	v38 = F_slice_from_s(m, l0, int32(5), int32(_a_F_estonian_UTF_8_stem_16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L32
	}
L26:
	;
	v32 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_17))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L30
	}
L27:
	;
	v26 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_18))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	if int32(0) <= v26 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v2086 = v26
	goto L1
L30:
	;
	if int32(0) <= v32 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	v2086 = v32
	goto L1
L32:
	;
	if int32(0) <= v38 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	v2086 = v38
	goto L1
L34:
	;
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v2086 = v44
	goto L1
L36:
	;
	if int32(0) <= v50 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v2086 = v50
	goto L1
L38:
	;
	if int32(0) <= v56 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v2086 = v56
	goto L1
L40:
	;
	if int32(0) <= v62 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v2086 = v62
	goto L1
L42:
	;
	if int32(0) <= v68 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v2086 = v68
	goto L1
L44:
	;
	if int32(0) <= v74 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	v2086 = v74
	goto L1
L46:
	;
	if int32(0) <= v80 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	v2086 = v80
	goto L1
L48:
	;
	if int32(0) <= v86 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	v2086 = v86
	goto L1
L50:
	;
	if int32(0) <= v92 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	v2086 = v92
	goto L1
L52:
	;
	if int32(0) <= v98 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	v2086 = v98
	goto L1
L54:
	;
	if int32(0) <= v104 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v2086 = v104
	goto L1
L56:
	;
	if int32(0) <= v110 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	v2086 = v110
	goto L1
L58:
	;
	if int32(0) <= v116 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v2086 = v116
	goto L1
L60:
	;
	if int32(0) <= v122 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v2086 = v122
	goto L1
L62:
	;
	if v128 < int32(0) {
		v2086 = v128
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
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v386 < v388 {
		goto L116
	} else {
		goto L117
	}
L65:
	;
	if v255 < int32(0) {
		goto L64
	} else {
		goto L90
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
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v259 = v258 + v255
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v259
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v282 = v259
	goto L93
L91:
	;
	if v378 < int32(0) {
		goto L64
	} else {
		goto L115
	}
L92:
	;
	v378 = v349
	goto L91
L93:
	;
	if v273 <= v282 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v378 = int32(-1)
	goto L91
L96:
	;
	goto L97
L97:
	;
	v289 = int32(1)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282+v274))))
	if base.Ui32(v291) < base.Ui32(int32(192)) {
		v348 = v291
		v349 = v289
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if int32(252) < v348 {
		goto L92
	} else {
		goto L111
	}
L99:
	;
	v295 = v282 + int32(1)
	if v295 == v273 {
		v348 = v291
		v349 = v289
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v274))))
	v300 = v298 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v291) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304+v274))))
	v316 = v314 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v291) {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	v304 = v282 + int32(2)
	if v304 != v273 {
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v348 = v291<<(uint(int32(6))%32)&int32(1984) | v300
	v349 = int32(2)
	goto L98
L105:
	;
	goto L104
L106:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274+v320))))
	v348 = v333&int32(63) | (v291<<(uint(int32(18))%32)&int32(_a_F_estonian_UTF_8_stem_19) | v300<<(uint(int32(12))%32) | v316<<(uint(int32(6))%32))
	v349 = int32(4)
	goto L98
L107:
	;
	v320 = v282 + int32(3)
	if v320 != v273 {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v348 = v291<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v300<<(uint(int32(6))%32) | v316
	v349 = int32(3)
	goto L98
L110:
	;
	goto L109
L111:
	;
	v353 = v348 - int32(97)
	if v353 < int32(0) {
		goto L92
	} else {
		goto L112
	}
L112:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v353)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v359)>>(uint(v353&int32(7))%32))&int32(1) == int32(0) {
		goto L92
	} else {
		goto L113
	}
L113:
	;
	v367 = v349 + v282
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v367
	v282 = v367
	goto L93
L115:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v381 + v378
	goto L64
L116:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v753
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v753 < v755 {
		goto L198
	} else {
		goto L199
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v388
	v393 = v386 - int32(1)
	if v393 <= v388 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	goto L116
L119:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v393))))
	if v397 != int32(105) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v403 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_21), int32(2), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	if v403 == int32(0) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v408
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L125
L123:
	;
	if v463 < int32(0) {
		goto L116
	} else {
		goto L142
	}
L125:
	;
	goto L126
L126:
	;
	goto L127
L127:
	;
	v418 = v408
	v420 = int32(4)
	goto L130
L129:
	;
	v463 = v445
	goto L123
L130:
	;
	if v418 <= v6 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	goto L129
L132:
	;
	v463 = int32(-1)
	goto L123
L133:
	;
	goto L134
L134:
	;
	v425 = v418 - int32(1)
	v427 = int32(*(*int8)(unsafe.Add(mBase, uint32(v411+v425))))
	if base.B2i32(int32(0) <= v427)|base.B2i32(v425 <= v6) != 0 {
		v445 = v425
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v449 = int32(1)
	if v449 < v420 {
		v418 = v445
		v420 = v420 - v449
		goto L130
	} else {
		goto L141
	}
L136:
	;
	v433 = v425
	goto L137
L137:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411+v433))))
	if base.Ui32(int32(191)) < base.Ui32(v438) {
		v445 = v433
		goto L135
	} else {
		goto L139
	}
L138:
	;
	v445 = v6
	goto L135
L139:
	;
	v442 = v433 - int32(1)
	if v6 < v442 {
		v433 = v442
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	goto L131
L142:
	;
	v466 = v408 - v410
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v466 + v467
	switch v403 - int32(1) {
	case 0:
		goto L144
	case 1:
		goto L143
	default:
		goto L116
	}
L143:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L174
L144:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L147
L145:
	;
	if v600 != 0 {
		goto L116
	} else {
		goto L168
	}
L146:
	;
	v600 = v593
	goto L145
L147:
	;
	if v484 <= v485 {
		v593 = int32(-1)
		goto L146
	} else {
		goto L149
	}
L148:
	;
	v593 = int32(0)
	goto L146
L149:
	;
	v502 = int32(1)
	v503 = v484 - v502
	v505 = int32(*(*int8)(unsafe.Add(mBase, uint32(v486+v503))))
	v507 = v505 & int32(255)
	if base.B2i32(v503 == v485)|base.B2i32(int32(0) <= v505) != 0 {
		v565 = v507
		v569 = v502
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if int32(252) < v565 {
		goto L158
	} else {
		goto L159
	}
L151:
	;
	v514 = v507 & int32(63)
	v516 = v484 - int32(2)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486+v516))))
	v520 = v518 << (uint(int32(6)) % 32)
	if base.B2i32(v516 != v485)&base.B2i32(base.Ui32(v518) < base.Ui32(int32(192))) == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v565 = v520&int32(1984) | v514
	v569 = int32(2)
	goto L150
L153:
	;
	goto L154
L154:
	;
	v533 = v520&int32(4032) | v514
	v535 = v484 - int32(3)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486+v535))))
	if base.B2i32(v535 != v485)&base.B2i32(base.Ui32(v537) < base.Ui32(int32(224))) == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v565 = v537<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v533
	v569 = int32(3)
	goto L150
L156:
	;
	goto L157
L157:
	;
	v555 = int32(4)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484+v486-v555))))
	v565 = v537<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v557&int32(7)<<(uint(int32(18))%32) | v533
	v569 = v555
	goto L150
L158:
	;
	v600 = v569
	goto L145
L159:
	;
	goto L160
L160:
	;
	v571 = v565 - int32(97)
	if v571 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v600 = v569
	goto L145
L162:
	;
	goto L163
L163:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v571)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[1]))))
	if int32(base.Ui32(v577)>>(uint(v571&int32(7))%32))&int32(1) == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v600 = v569
	goto L145
L165:
	;
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v484 - v569
	goto L167
L167:
	;
	goto L148
L168:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v601 + v466
	v607 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	if v607 != 0 {
		goto L116
	} else {
		goto L170
	}
L170:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v609 + v466
	v612 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v612 {
		goto L116
	} else {
		goto L171
	}
L171:
	;
	v2086 = v612
	goto L1
L172:
	;
	if v743 != 0 {
		goto L116
	} else {
		goto L195
	}
L173:
	;
	v743 = v736
	goto L172
L174:
	;
	if v627 <= v628 {
		v736 = int32(-1)
		goto L173
	} else {
		goto L176
	}
L175:
	;
	v736 = int32(0)
	goto L173
L176:
	;
	v645 = int32(1)
	v646 = v627 - v645
	v648 = int32(*(*int8)(unsafe.Add(mBase, uint32(v629+v646))))
	v650 = v648 & int32(255)
	if base.B2i32(v646 == v628)|base.B2i32(int32(0) <= v648) != 0 {
		v708 = v650
		v712 = v645
		goto L177
	} else {
		goto L178
	}
L177:
	;
	if int32(382) < v708 {
		goto L185
	} else {
		goto L186
	}
L178:
	;
	v657 = v650 & int32(63)
	v659 = v627 - int32(2)
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+v659))))
	v663 = v661 << (uint(int32(6)) % 32)
	if base.B2i32(v659 != v628)&base.B2i32(base.Ui32(v661) < base.Ui32(int32(192))) == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v708 = v663&int32(1984) | v657
	v712 = int32(2)
	goto L177
L180:
	;
	goto L181
L181:
	;
	v676 = v663&int32(4032) | v657
	v678 = v627 - int32(3)
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629+v678))))
	if base.B2i32(v678 != v628)&base.B2i32(base.Ui32(v680) < base.Ui32(int32(224))) == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v708 = v680<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v676
	v712 = int32(3)
	goto L177
L183:
	;
	goto L184
L184:
	;
	v698 = int32(4)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627+v629-v698))))
	v708 = v680<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v700&int32(7)<<(uint(int32(18))%32) | v676
	v712 = v698
	goto L177
L185:
	;
	v743 = v712
	goto L172
L186:
	;
	goto L187
L187:
	;
	v714 = v708 - int32(98)
	if v714 < int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v743 = v712
	goto L172
L189:
	;
	goto L190
L190:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v714)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[2]))))
	if int32(base.Ui32(v720)>>(uint(v714&int32(7))%32))&int32(1) == int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v743 = v712
	goto L172
L192:
	;
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v627 - v712
	goto L194
L194:
	;
	goto L175
L195:
	;
	v744 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v744 {
		goto L116
	} else {
		goto L196
	}
L196:
	;
	v2086 = v744
	goto L1
L197:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1900
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L463
L198:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v934
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v934 < v936 {
		goto L236
	} else {
		goto L237
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v753
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v755
	if v753 <= v755 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v758
	goto L198
L201:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v763 = int32(1)
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761+v753-v763))))
	if base.B2i32(v765&int32(224) != int32(96))|base.B2i32(v763<<(uint(v765)%32)&int32(_a_F_estonian_UTF_8_stem_24) == int32(0)) != 0 {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v780 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_25), int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L3
	} else {
		goto L203
	}
L203:
	;
	if v780 == int32(0) {
		goto L200
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v758
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v785
	switch v780 - int32(1) {
	case 0:
		goto L207
	case 1:
		goto L206
	case 2:
		goto L205
	default:
		goto L197
	}
L205:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L213
L206:
	;
	v794 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_26))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L3
	} else {
		goto L209
	}
L207:
	;
	v789 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v789 {
		goto L197
	} else {
		goto L208
	}
L208:
	;
	v2086 = v789
	goto L1
L209:
	;
	if int32(0) <= v794 {
		goto L197
	} else {
		goto L210
	}
L210:
	;
	v2086 = v794
	goto L1
L211:
	;
	if v926 != 0 {
		goto L198
	} else {
		goto L234
	}
L212:
	;
	v926 = v919
	goto L211
L213:
	;
	if v810 <= v811 {
		v919 = int32(-1)
		goto L212
	} else {
		goto L215
	}
L214:
	;
	v919 = int32(0)
	goto L212
L215:
	;
	v828 = int32(1)
	v829 = v810 - v828
	v831 = int32(*(*int8)(unsafe.Add(mBase, uint32(v812+v829))))
	v833 = v831 & int32(255)
	if base.B2i32(v829 == v811)|base.B2i32(int32(0) <= v831) != 0 {
		v891 = v833
		v895 = v828
		goto L216
	} else {
		goto L217
	}
L216:
	;
	if int32(252) < v891 {
		goto L224
	} else {
		goto L225
	}
L217:
	;
	v840 = v833 & int32(63)
	v842 = v810 - int32(2)
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812+v842))))
	v846 = v844 << (uint(int32(6)) % 32)
	if base.B2i32(v842 != v811)&base.B2i32(base.Ui32(v844) < base.Ui32(int32(192))) == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v891 = v846&int32(1984) | v840
	v895 = int32(2)
	goto L216
L219:
	;
	goto L220
L220:
	;
	v859 = v846&int32(4032) | v840
	v861 = v810 - int32(3)
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812+v861))))
	if base.B2i32(v861 != v811)&base.B2i32(base.Ui32(v863) < base.Ui32(int32(224))) == int32(0) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v891 = v863<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v859
	v895 = int32(3)
	goto L216
L222:
	;
	goto L223
L223:
	;
	v881 = int32(4)
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810+v812-v881))))
	v891 = v863<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v883&int32(7)<<(uint(int32(18))%32) | v859
	v895 = v881
	goto L216
L224:
	;
	v926 = v895
	goto L211
L225:
	;
	goto L226
L226:
	;
	v897 = v891 - int32(97)
	if v897 < int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v926 = v895
	goto L211
L228:
	;
	goto L229
L229:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v897)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v903)>>(uint(v897&int32(7))%32))&int32(1) == int32(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v926 = v895
	goto L211
L231:
	;
	goto L232
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v810 - v895
	goto L233
L233:
	;
	goto L214
L234:
	;
	v927 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v927 {
		goto L197
	} else {
		goto L235
	}
L235:
	;
	v2086 = v927
	goto L1
L236:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v994
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v994 < v996 {
		goto L252
	} else {
		goto L253
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v934
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v936
	if v934-int32(3) <= v936 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v939
	goto L236
L239:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v946 = int32(1)
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v944+v934-v946))))
	if base.B2i32(v948&int32(224) != int32(96))|base.B2i32(v946<<(uint(v948)%32)&int32(_a_F_estonian_UTF_8_stem_27) == int32(0)) != 0 {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v963 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_28), int32(12), int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L3
	} else {
		goto L241
	}
L241:
	;
	if v963 == int32(0) {
		goto L238
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v939
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v968
	switch v963 - int32(1) {
	case 0:
		goto L245
	case 1:
		goto L244
	case 2:
		goto L243
	default:
		goto L236
	}
L243:
	;
	v986 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_29))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L3
	} else {
		goto L250
	}
L244:
	;
	v980 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_30))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L3
	} else {
		goto L248
	}
L245:
	;
	v974 = F_slice_from_s(m, l0, int32(4), int32(_a_F_estonian_UTF_8_stem_31))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L3
	} else {
		goto L246
	}
L246:
	;
	if int32(0) <= v974 {
		goto L236
	} else {
		goto L247
	}
L247:
	;
	v2086 = v974
	goto L1
L248:
	;
	if int32(0) <= v980 {
		goto L236
	} else {
		goto L249
	}
L249:
	;
	v2086 = v980
	goto L1
L250:
	;
	if int32(0) <= v986 {
		goto L236
	} else {
		goto L251
	}
L251:
	;
	v2086 = v986
	goto L1
L252:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1242
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1242 < v1244 {
		goto L309
	} else {
		goto L310
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v994
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v996
	if v994 <= v996 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v999
	goto L252
L255:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1004 = int32(1)
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002+v994-v1004))))
	if base.B2i32(v1006&int32(224) != int32(96))|base.B2i32(v1004<<(uint(v1006)%32)&int32(_a_F_estonian_UTF_8_stem_32) == int32(0)) != 0 {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1021 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_33), int32(10), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L3
	} else {
		goto L257
	}
L257:
	;
	if v1021 == int32(0) {
		goto L254
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v999
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1026
	switch v1021 - int32(1) {
	case 0:
		goto L261
	case 1:
		goto L260
	default:
		goto L259
	}
L259:
	;
	v1233 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1233 {
		goto L252
	} else {
		goto L308
	}
L260:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L290
L261:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L264
L262:
	;
	if v1159 == int32(0) {
		goto L259
	} else {
		goto L285
	}
L263:
	;
	v1159 = v1152
	goto L262
L264:
	;
	if v1043 <= v1044 {
		v1152 = int32(-1)
		goto L263
	} else {
		goto L266
	}
L265:
	;
	v1152 = int32(0)
	goto L263
L266:
	;
	v1061 = int32(1)
	v1062 = v1043 - v1061
	v1064 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1045+v1062))))
	v1066 = v1064 & int32(255)
	if base.B2i32(v1062 == v1044)|base.B2i32(int32(0) <= v1064) != 0 {
		v1124 = v1066
		v1128 = v1061
		goto L267
	} else {
		goto L268
	}
L267:
	;
	if int32(117) < v1124 {
		goto L275
	} else {
		goto L276
	}
L268:
	;
	v1073 = v1066 & int32(63)
	v1075 = v1043 - int32(2)
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045+v1075))))
	v1079 = v1077 << (uint(int32(6)) % 32)
	if base.B2i32(v1075 != v1044)&base.B2i32(base.Ui32(v1077) < base.Ui32(int32(192))) == int32(0) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1124 = v1079&int32(1984) | v1073
	v1128 = int32(2)
	goto L267
L270:
	;
	goto L271
L271:
	;
	v1092 = v1079&int32(4032) | v1073
	v1094 = v1043 - int32(3)
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045+v1094))))
	if base.B2i32(v1094 != v1044)&base.B2i32(base.Ui32(v1096) < base.Ui32(int32(224))) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1124 = v1096<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1092
	v1128 = int32(3)
	goto L267
L273:
	;
	goto L274
L274:
	;
	v1114 = int32(4)
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043+v1045-v1114))))
	v1124 = v1096<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1116&int32(7)<<(uint(int32(18))%32) | v1092
	v1128 = v1114
	goto L267
L275:
	;
	v1159 = v1128
	goto L262
L276:
	;
	goto L277
L277:
	;
	v1130 = v1124 - int32(97)
	if v1130 < int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1159 = v1128
	goto L262
L279:
	;
	goto L280
L280:
	;
	v1136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1130)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1136)>>(uint(v1130&int32(7))%32))&int32(1) == int32(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1159 = v1128
	goto L262
L282:
	;
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1043 - v1128
	goto L284
L284:
	;
	goto L265
L285:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1162 + (v1026 - v1030)
	v1169 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9), int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L3
	} else {
		goto L286
	}
L286:
	;
	if v1169 != 0 {
		goto L259
	} else {
		goto L287
	}
L287:
	;
	goto L252
L288:
	;
	if v1224 < int32(0) {
		goto L252
	} else {
		goto L307
	}
L290:
	;
	goto L291
L291:
	;
	goto L292
L292:
	;
	v1179 = v1026
	v1181 = int32(4)
	goto L295
L294:
	;
	v1224 = v1206
	goto L288
L295:
	;
	if v1179 <= v999 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	goto L294
L297:
	;
	v1224 = int32(-1)
	goto L288
L298:
	;
	goto L299
L299:
	;
	v1186 = v1179 - int32(1)
	v1188 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1172+v1186))))
	if base.B2i32(int32(0) <= v1188)|base.B2i32(v1186 <= v999) != 0 {
		v1206 = v1186
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1210 = int32(1)
	if v1210 < v1181 {
		v1179 = v1206
		v1181 = v1181 - v1210
		goto L295
	} else {
		goto L306
	}
L301:
	;
	v1194 = v1186
	goto L302
L302:
	;
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1172+v1194))))
	if base.Ui32(int32(191)) < base.Ui32(v1199) {
		v1206 = v1194
		goto L300
	} else {
		goto L304
	}
L303:
	;
	v1206 = v999
	goto L300
L304:
	;
	v1203 = v1194 - int32(1)
	if v999 < v1203 {
		v1194 = v1203
		goto L302
	} else {
		goto L305
	}
L305:
	;
	goto L303
L306:
	;
	goto L296
L307:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1227 + (v1026 - v1171)
	goto L259
L308:
	;
	v2086 = v1233
	goto L1
L309:
	;
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1534
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1534 < v1536 {
		goto L386
	} else {
		goto L387
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1242
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1244
	if v1242 <= v1244 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1247
	goto L309
L312:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1250+v1242-int32(1)))))
	if v1254&int32(254) != int32(100) {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1262 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_34), int32(7), int32(0))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L3
	} else {
		goto L314
	}
L314:
	;
	if v1262 == int32(0) {
		goto L311
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1247
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1267
	switch v1262 - int32(1) {
	case 0:
		goto L319
	case 1:
		goto L318
	case 2:
		goto L317
	case 3:
		goto L316
	default:
		goto L309
	}
L316:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L359
L317:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L327
L318:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1281 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9), int32(0))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L3
	} else {
		goto L322
	}
L319:
	;
	v1273 = F_slice_from_s(m, l0, int32(3), int32(_a_F_estonian_UTF_8_stem_35))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L3
	} else {
		goto L320
	}
L320:
	;
	if int32(0) <= v1273 {
		goto L309
	} else {
		goto L321
	}
L321:
	;
	v2086 = v1273
	goto L1
L322:
	;
	if v1281 != 0 {
		goto L309
	} else {
		goto L323
	}
L323:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1283 + (v1267 - v1277)
	v1287 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1287 {
		goto L309
	} else {
		goto L324
	}
L324:
	;
	v2086 = v1287
	goto L1
L325:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1346 = v1344 + (v1267 - v1290)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1346
	if int32(0) <= v1343 {
		goto L344
	} else {
		goto L345
	}
L327:
	;
	goto L328
L328:
	;
	goto L329
L329:
	;
	v1298 = v1267
	v1300 = int32(4)
	goto L332
L331:
	;
	v1343 = v1325
	goto L325
L332:
	;
	if v1298 <= v1247 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	goto L331
L334:
	;
	v1343 = int32(-1)
	goto L325
L335:
	;
	goto L336
L336:
	;
	v1305 = v1298 - int32(1)
	v1307 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1291+v1305))))
	if base.B2i32(int32(0) <= v1307)|base.B2i32(v1305 <= v1247) != 0 {
		v1325 = v1305
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1329 = int32(1)
	if v1329 < v1300 {
		v1298 = v1325
		v1300 = v1300 - v1329
		goto L332
	} else {
		goto L343
	}
L338:
	;
	v1313 = v1305
	goto L339
L339:
	;
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291+v1313))))
	if base.Ui32(int32(191)) < base.Ui32(v1318) {
		v1325 = v1313
		goto L337
	} else {
		goto L341
	}
L340:
	;
	v1325 = v1247
	goto L337
L341:
	;
	v1322 = v1313 - int32(1)
	if v1247 < v1322 {
		v1313 = v1322
		goto L339
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	goto L333
L344:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1346 <= v1350 {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	goto L346
L346:
	;
	v1381 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_36))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L3
	} else {
		goto L355
	}
L347:
	;
	v1376 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1376 {
		goto L309
	} else {
		goto L354
	}
L348:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1354 = int32(1)
	v1356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1352+v1346-v1354))))
	if base.Ui32(v1354) < base.Ui32((v1356-int32(115))&int32(255)) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v1366 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_37), int32(5), int32(0))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L3
	} else {
		goto L351
	}
L350:
	;
	v1372 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_38))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L3
	} else {
		goto L352
	}
L351:
	;
	switch v1366 - int32(1) {
	case 0:
		goto L350
	case 1:
		goto L347
	default:
		goto L309
	}
L352:
	;
	if int32(0) <= v1372 {
		goto L309
	} else {
		goto L353
	}
L353:
	;
	v2086 = v1372
	goto L1
L354:
	;
	v2086 = v1376
	goto L1
L355:
	;
	if int32(0) <= v1381 {
		goto L309
	} else {
		goto L356
	}
L356:
	;
	v2086 = v1381
	goto L1
L357:
	;
	if v1514 != 0 {
		goto L380
	} else {
		goto L381
	}
L358:
	;
	v1514 = v1507
	goto L357
L359:
	;
	if v1398 <= v1399 {
		v1507 = int32(-1)
		goto L358
	} else {
		goto L361
	}
L360:
	;
	v1507 = int32(0)
	goto L358
L361:
	;
	v1416 = int32(1)
	v1417 = v1398 - v1416
	v1419 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1400+v1417))))
	v1421 = v1419 & int32(255)
	if base.B2i32(v1417 == v1399)|base.B2i32(int32(0) <= v1419) != 0 {
		v1479 = v1421
		v1483 = v1416
		goto L362
	} else {
		goto L363
	}
L362:
	;
	if int32(117) < v1479 {
		goto L370
	} else {
		goto L371
	}
L363:
	;
	v1428 = v1421 & int32(63)
	v1430 = v1398 - int32(2)
	v1432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400+v1430))))
	v1434 = v1432 << (uint(int32(6)) % 32)
	if base.B2i32(v1430 != v1399)&base.B2i32(base.Ui32(v1432) < base.Ui32(int32(192))) == int32(0) {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1479 = v1434&int32(1984) | v1428
	v1483 = int32(2)
	goto L362
L365:
	;
	goto L366
L366:
	;
	v1447 = v1434&int32(4032) | v1428
	v1449 = v1398 - int32(3)
	v1451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400+v1449))))
	if base.B2i32(v1449 != v1399)&base.B2i32(base.Ui32(v1451) < base.Ui32(int32(224))) == int32(0) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1479 = v1451<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1447
	v1483 = int32(3)
	goto L362
L368:
	;
	goto L369
L369:
	;
	v1469 = int32(4)
	v1471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1398+v1400-v1469))))
	v1479 = v1451<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1471&int32(7)<<(uint(int32(18))%32) | v1447
	v1483 = v1469
	goto L362
L370:
	;
	v1514 = v1483
	goto L357
L371:
	;
	goto L372
L372:
	;
	v1485 = v1479 - int32(97)
	if v1485 < int32(0) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1514 = v1483
	goto L357
L374:
	;
	goto L375
L375:
	;
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1485)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1491)>>(uint(v1485&int32(7))%32))&int32(1) == int32(0) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1514 = v1483
	goto L357
L377:
	;
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1398 - v1483
	goto L379
L379:
	;
	goto L360
L380:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1515 + (v1267 - v1385)
	v1522 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_23), int32(9), int32(0))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L3
	} else {
		goto L383
	}
L381:
	;
	goto L382
L382:
	;
	v1526 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1526 {
		goto L309
	} else {
		goto L385
	}
L383:
	;
	if v1522 == int32(0) {
		goto L309
	} else {
		goto L384
	}
L384:
	;
	goto L382
L385:
	;
	v2086 = v1526
	goto L1
L386:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1709
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1709 < v1711 {
		goto L421
	} else {
		goto L422
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1534
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1536
	if v1534 <= v1536 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1539
	goto L386
L389:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1544 = int32(1)
	v1546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1542+v1534-v1544))))
	if base.B2i32(v1546&int32(224) != int32(96))|base.B2i32(v1544<<(uint(v1546)%32)&int32(_a_F_estonian_UTF_8_stem_39) == int32(0)) != 0 {
		goto L388
	} else {
		goto L390
	}
L390:
	;
	v1561 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_40), int32(3), int32(0))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L3
	} else {
		goto L391
	}
L391:
	;
	if v1561 == int32(0) {
		goto L388
	} else {
		goto L392
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1539
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1566
	switch v1561 - int32(1) {
	case 0:
		goto L394
	case 1:
		goto L393
	default:
		goto L386
	}
L393:
	;
	v1702 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1702 {
		goto L386
	} else {
		goto L420
	}
L394:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L397
L395:
	;
	if v1698 != 0 {
		goto L386
	} else {
		goto L418
	}
L396:
	;
	v1698 = v1691
	goto L395
L397:
	;
	if v1582 <= v1583 {
		v1691 = int32(-1)
		goto L396
	} else {
		goto L399
	}
L398:
	;
	v1691 = int32(0)
	goto L396
L399:
	;
	v1600 = int32(1)
	v1601 = v1582 - v1600
	v1603 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1584+v1601))))
	v1605 = v1603 & int32(255)
	if base.B2i32(v1601 == v1583)|base.B2i32(int32(0) <= v1603) != 0 {
		v1663 = v1605
		v1667 = v1600
		goto L400
	} else {
		goto L401
	}
L400:
	;
	if int32(117) < v1663 {
		goto L408
	} else {
		goto L409
	}
L401:
	;
	v1612 = v1605 & int32(63)
	v1614 = v1582 - int32(2)
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1584+v1614))))
	v1618 = v1616 << (uint(int32(6)) % 32)
	if base.B2i32(v1614 != v1583)&base.B2i32(base.Ui32(v1616) < base.Ui32(int32(192))) == int32(0) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1663 = v1618&int32(1984) | v1612
	v1667 = int32(2)
	goto L400
L403:
	;
	goto L404
L404:
	;
	v1631 = v1618&int32(4032) | v1612
	v1633 = v1582 - int32(3)
	v1635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1584+v1633))))
	if base.B2i32(v1633 != v1583)&base.B2i32(base.Ui32(v1635) < base.Ui32(int32(224))) == int32(0) {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1663 = v1635<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1631
	v1667 = int32(3)
	goto L400
L406:
	;
	goto L407
L407:
	;
	v1653 = int32(4)
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1582+v1584-v1653))))
	v1663 = v1635<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1655&int32(7)<<(uint(int32(18))%32) | v1631
	v1667 = v1653
	goto L400
L408:
	;
	v1698 = v1667
	goto L395
L409:
	;
	goto L410
L410:
	;
	v1669 = v1663 - int32(97)
	if v1669 < int32(0) {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1698 = v1667
	goto L395
L412:
	;
	goto L413
L413:
	;
	v1675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1669)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1675)>>(uint(v1669&int32(7))%32))&int32(1) == int32(0) {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1698 = v1667
	goto L395
L415:
	;
	goto L416
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1582 - v1667
	goto L417
L417:
	;
	goto L398
L418:
	;
	v1699 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1699 {
		goto L386
	} else {
		goto L419
	}
L419:
	;
	v2086 = v1699
	goto L1
L420:
	;
	v2086 = v1702
	goto L1
L421:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1864
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1864 < v1866 {
		goto L197
	} else {
		goto L453
	}
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1709
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1711
	if v1711 < v1709 {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1714
	v1727 = v1709 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1727
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1727
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L430
L424:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1717+v1709-int32(1)))))
	if v1721 == int32(105) {
		goto L423
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1714
	goto L421
L427:
	;
	goto L426
L428:
	;
	if v1858 != 0 {
		goto L421
	} else {
		goto L451
	}
L429:
	;
	v1858 = v1851
	goto L428
L430:
	;
	if v1727 <= v1714 {
		v1851 = int32(-1)
		goto L429
	} else {
		goto L432
	}
L431:
	;
	v1851 = int32(0)
	goto L429
L432:
	;
	v1760 = int32(1)
	v1761 = v1727 - v1760
	v1763 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1744+v1761))))
	v1765 = v1763 & int32(255)
	if base.B2i32(v1761 == v1714)|base.B2i32(int32(0) <= v1763) != 0 {
		v1823 = v1765
		v1827 = v1760
		goto L433
	} else {
		goto L434
	}
L433:
	;
	if int32(117) < v1823 {
		goto L441
	} else {
		goto L442
	}
L434:
	;
	v1772 = v1765 & int32(63)
	v1774 = v1727 - int32(2)
	v1776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1744+v1774))))
	v1778 = v1776 << (uint(int32(6)) % 32)
	if base.B2i32(v1774 != v1714)&base.B2i32(base.Ui32(v1776) < base.Ui32(int32(192))) == int32(0) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1823 = v1778&int32(1984) | v1772
	v1827 = int32(2)
	goto L433
L436:
	;
	goto L437
L437:
	;
	v1791 = v1778&int32(4032) | v1772
	v1793 = v1727 - int32(3)
	v1795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1744+v1793))))
	if base.B2i32(v1793 != v1714)&base.B2i32(base.Ui32(v1795) < base.Ui32(int32(224))) == int32(0) {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v1823 = v1795<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1791
	v1827 = int32(3)
	goto L433
L439:
	;
	goto L440
L440:
	;
	v1813 = int32(4)
	v1815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1727+v1744-v1813))))
	v1823 = v1795<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1815&int32(7)<<(uint(int32(18))%32) | v1791
	v1827 = v1813
	goto L433
L441:
	;
	v1858 = v1827
	goto L428
L442:
	;
	goto L443
L443:
	;
	v1829 = v1823 - int32(97)
	if v1829 < int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v1858 = v1827
	goto L428
L445:
	;
	goto L446
L446:
	;
	v1835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1829)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[3]))))
	if int32(base.Ui32(v1835)>>(uint(v1829&int32(7))%32))&int32(1) == int32(0) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v1858 = v1827
	goto L428
L448:
	;
	goto L449
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1727 - v1827
	goto L450
L450:
	;
	goto L431
L451:
	;
	v1859 = F_slice_del(m, l0)
	mBase = m.M
	if v1859 < int32(0) {
		v2086 = v1859
		goto L1
	} else {
		goto L452
	}
L452:
	;
	goto L421
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1864
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1866
	v1872 = v1864 - int32(1)
	if v1872 <= v1866 {
		goto L454
	} else {
		goto L455
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1869
	goto L197
L455:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1874+v1872))))
	if base.B2i32(v1876 != int32(117))&base.B2i32(v1876 != int32(97)) != 0 {
		goto L454
	} else {
		goto L456
	}
L456:
	;
	v1885 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_41), int32(4), int32(0))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L3
	} else {
		goto L457
	}
L457:
	;
	if v1885 == int32(0) {
		goto L454
	} else {
		goto L458
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1869
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1890
	v1892 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1892 {
		goto L197
	} else {
		goto L459
	}
L459:
	;
	v2086 = v1892
	goto L1
L460:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2083
	v2086 = int32(1)
	goto L1
L461:
	;
	if v2030 != 0 {
		goto L460
	} else {
		goto L484
	}
L462:
	;
	v2030 = v2023
	goto L461
L463:
	;
	if v1900 <= v1915 {
		v2023 = int32(-1)
		goto L462
	} else {
		goto L465
	}
L464:
	;
	v2023 = int32(0)
	goto L462
L465:
	;
	v1932 = int32(1)
	v1933 = v1900 - v1932
	v1935 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1916+v1933))))
	v1937 = v1935 & int32(255)
	if base.B2i32(v1933 == v1915)|base.B2i32(int32(0) <= v1935) != 0 {
		v1995 = v1937
		v1999 = v1932
		goto L466
	} else {
		goto L467
	}
L466:
	;
	if int32(252) < v1995 {
		goto L474
	} else {
		goto L475
	}
L467:
	;
	v1944 = v1937 & int32(63)
	v1946 = v1900 - int32(2)
	v1948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1916+v1946))))
	v1950 = v1948 << (uint(int32(6)) % 32)
	if base.B2i32(v1946 != v1915)&base.B2i32(base.Ui32(v1948) < base.Ui32(int32(192))) == int32(0) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v1995 = v1950&int32(1984) | v1944
	v1999 = int32(2)
	goto L466
L469:
	;
	goto L470
L470:
	;
	v1963 = v1950&int32(4032) | v1944
	v1965 = v1900 - int32(3)
	v1967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1916+v1965))))
	if base.B2i32(v1965 != v1915)&base.B2i32(base.Ui32(v1967) < base.Ui32(int32(224))) == int32(0) {
		goto L471
	} else {
		goto L472
	}
L471:
	;
	v1995 = v1967<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_20) | v1963
	v1999 = int32(3)
	goto L466
L472:
	;
	goto L473
L473:
	;
	v1985 = int32(4)
	v1987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1900+v1916-v1985))))
	v1995 = v1967<<(uint(int32(12))%32)&int32(_a_F_estonian_UTF_8_stem_22) | v1987&int32(7)<<(uint(int32(18))%32) | v1963
	v1999 = v1985
	goto L466
L474:
	;
	v2030 = v1999
	goto L461
L475:
	;
	goto L476
L476:
	;
	v2001 = v1995 - int32(97)
	if v2001 < int32(0) {
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v2030 = v1999
	goto L461
L478:
	;
	goto L479
L479:
	;
	v2007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2001)>>(uint(int32(3))%32)))+uint32(_c_F_estonian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2007)>>(uint(v2001&int32(7))%32))&int32(1) == int32(0) {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v2030 = v1999
	goto L461
L481:
	;
	goto L482
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1900 - v1999
	goto L483
L483:
	;
	goto L464
L484:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2031 < v2032 {
		goto L460
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2031
	v2036 = v2031 - int32(1)
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2036 <= v2037 {
		goto L460
	} else {
		goto L486
	}
L486:
	;
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2039+v2036))))
	if base.B2i32(v2041&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2041)%32)&int32(_a_F_estonian_UTF_8_stem_42) == int32(0)) != 0 {
		goto L460
	} else {
		goto L487
	}
L487:
	;
	v2056 = F_find_among_b(m, l0, int32(_a_F_estonian_UTF_8_stem_43), int32(3), int32(0))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L3
	} else {
		goto L488
	}
L488:
	;
	if v2056 == int32(0) {
		goto L460
	} else {
		goto L489
	}
L489:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2060
	switch v2056 - int32(1) {
	case 0:
		goto L492
	case 1:
		goto L491
	case 2:
		goto L490
	default:
		goto L460
	}
L490:
	;
	v2078 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_44))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L3
	} else {
		goto L497
	}
L491:
	;
	v2072 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_45))
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L3
	} else {
		goto L495
	}
L492:
	;
	v2066 = F_slice_from_s(m, l0, int32(1), int32(_a_F_estonian_UTF_8_stem_46))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L3
	} else {
		goto L493
	}
L493:
	;
	if int32(0) <= v2066 {
		goto L460
	} else {
		goto L494
	}
L494:
	;
	v2086 = v2066
	goto L1
L495:
	;
	if int32(0) <= v2072 {
		goto L460
	} else {
		goto L496
	}
L496:
	;
	v2086 = v2072
	goto L1
L497:
	;
	if v2078 < int32(0) {
		v2086 = v2078
		goto L1
	} else {
		goto L498
	}
L498:
	;
	goto L460
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
	var v27 int64
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
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
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
					v27 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)))
					v29 = F_SearchSysCacheCopy(m, int32(82), v27, int64(0))
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
								v79 = F_OidFunctionCall1Coll(m, v76, int32(0), base.I64_extend_i32_u(v12))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									if v79 != int64(0) {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
										if v87 == int32(0) {
											F_pfree(m, v29)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v12)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													v98 = int32(0)
													m.G0 = v9 + int32(16)
													return v98
												}
											}
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if int32(0) < v90 {
												v98 = v12
												m.G0 = v9 + int32(16)
												return v98
											} else {
												F_pfree(m, v29)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v12)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v98 = int32(0)
														m.G0 = v9 + int32(16)
														return v98
													}
												}
											}
										}
									} else {
										F_pfree(m, v29)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v12)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v98 = int32(0)
												m.G0 = v9 + int32(16)
												return v98
											}
										}
									}
								}
							} else {
								v83 = F_std_typanalyze(m, v12)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									if v83 == int32(0) {
										F_pfree(m, v29)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v12)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v98 = int32(0)
												m.G0 = v9 + int32(16)
												return v98
											}
										}
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
										if v87 == int32(0) {
											F_pfree(m, v29)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v12)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													v98 = int32(0)
													m.G0 = v9 + int32(16)
													return v98
												}
											}
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if int32(0) < v90 {
												v98 = v12
												m.G0 = v9 + int32(16)
												return v98
											} else {
												F_pfree(m, v29)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v12)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return int32(0)
													} else {
														v98 = int32(0)
														m.G0 = v9 + int32(16)
														return v98
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
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v107
								F_errmsg_internal(m, int32(_a_F_examine_expression_0), v9)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_examine_expression_1), int32(681), int32(_a_F_examine_expression_2))
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
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v178 int32
	_ = v178
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
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
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v836 int64
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
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
	var v906 int64
	_ = v906
	var v907 int64
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
	var v960 int64
	_ = v960
	var v977 int32
	_ = v977
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1110 int64
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1228 int32
	_ = v1228
	v5 = int32(0)
	v14 = int64(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v14
	v27 = F_exprType(m, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v27
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
	if v31 == int32(0) {
		v73 = l1
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v87 = v73
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
	v73 = int32(0)
	goto L3
L6:
	;
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v37 != int32(321) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = F_expression_tree_walker_impl(m, l1, int32(1702), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v50 = l1
	goto L14
L11:
	;
	if v42 == int32(0) {
		v73 = l1
		goto L3
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v67 = F_expression_tree_mutator_impl(m, v50, int32(1703), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v60 != int32(321) {
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v73 = int32(0)
	goto L3
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v63 != 0 {
		v50 = v63
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v73 = v67
	goto L3
L19:
	;
	m.G0 = v17 + int32(80)
	return
L20:
	;
	F_bms_free(m, v589)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L211
	}
L21:
	;
	v723 = l1
	v724 = v5
	goto L20
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L208
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L205
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v97 != int32(27) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L202
	}
L26:
	;
	if v97 != int32(6) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v87 = v670
	goto L24
L28:
	;
	goto L25
L29:
	;
	goto L28
L30:
	;
	v586 = F_pull_varnos(m, l0, v87)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
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
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if l2 != v102 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v87
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v106 = F_find_base_rel(m, l0, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v111
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+8)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v106)+116))
	if v115 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v178)
	v191 = l0
	v195 = v87
	goto L54
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v116 <= int32(0) {
		v178 = int32(0)
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v178 = int32(0)
	goto L37
L41:
	;
	v119 = int32(0)
	if v119 < v116 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v122 = v116
	goto L44
L43:
	;
	v122 = v119
	goto L44
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v126 = int32(0)
	goto L45
L45:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v123+v126<<(uint(int32(2))%32))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+101)))
	if v143 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L40
L47:
	;
	v159 = v126 + int32(1)
	if v159 != v122 {
		v126 = v159
		goto L45
	} else {
		goto L53
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+40))
	if v146 != int32(1) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)+44))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v150 != v113 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v152 = int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v142)+88))
	if v153 == int32(0) {
		v178 = v152
		goto L37
	} else {
		goto L51
	}
L51:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+100)))
	if v156 != 0 {
		v178 = v152
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
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v191)+44))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(int32(2))%32))))
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_examine_variable[0]))
	if v212 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L19
L56:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	switch v237 {
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
	v215 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195)+8)))
	v216 = m.T0[v212].(func(*base.Module, int32, int32, int32, int32) int32)(m, v191, v210, v215, l3)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v216 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v220 == int32(0) {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v223 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_0), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_2), int32(_a_F_examine_variable_3))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
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
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	if v434 == int32(0) {
		goto L19
	} else {
		goto L109
	}
L66:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+92)))
	if v269 != 0 {
		goto L19
	} else {
		goto L78
	}
L67:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+20)))
	if v258 != 0 {
		goto L19
	} else {
		goto L75
	}
L68:
	;
	v239 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v210)+16)))
	v240 = int64(*(*int16)(unsafe.Add(mBase, uint32(v195)+8)))
	v241 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v210)+20)))
	v242 = F_SearchSysCache3(m, int32(65), v239, v240, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1704)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v242
	if v242 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195)+8)))
	v251 = F_bms_make_singleton(m, v248+int32(7))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v256 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v256)
	goto L19
L73:
	;
	v253 = F_all_rows_selectable(m, v191, v247, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v253)
	goto L19
L75:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+8)))
	if v259 == int32(0) {
		goto L19
	} else {
		goto L76
	}
L76:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v265 = F_find_base_rel(m, v191, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v427 = v195 + int32(8)
	v433 = v265 + int32(148)
	goto L65
L78:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+8)))
	if v270 == int32(0) {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v210)+88))
	v277 = v191
	v280 = v275
	goto L81
L80:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+48))
	if v312 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L81:
	;
	if v280 == int32(0) {
		goto L80
	} else {
		goto L83
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L85
	}
L83:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v277)+16))
	if v294 != 0 {
		v277 = v294
		v280 = v280 - int32(1)
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v210)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v299
	F_errmsg_internal(m, int32(_a_F_examine_variable_4), v17-int32(-64))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_5), int32(_a_F_examine_variable_3))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
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
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v277)+84))
	if v401 == int32(0) {
		goto L29
	} else {
		goto L106
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L103
	}
L90:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v315 <= int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v210)+84))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v325 = int32(0)
	goto L92
L92:
	;
	v336 = v325 << (uint(int32(2)) % 32)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v319+v336)))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if base.B2i32(v342 == int32(0))|base.B2i32(v342 != v345) != 0 {
		v363 = v342
		v364 = v345
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L89
L94:
	;
	if v363-v364 == int32(0) {
		goto L88
	} else {
		goto L101
	}
L95:
	;
	goto L94
L96:
	;
	v348 = v339
	v349 = v318
	goto L97
L97:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	if v353 == int32(0) {
		v363 = v353
		v364 = v352
		goto L95
	} else {
		goto L99
	}
L98:
	;
	v363 = v353
	v364 = v352
	goto L95
L99:
	;
	v356 = int32(1)
	if v353 == v352 {
		v348 = v348 + v356
		v349 = v349 + v356
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v369 = v325 + int32(1)
	if v315 != v369 {
		v325 = v369
		goto L92
	} else {
		goto L102
	}
L102:
	;
	goto L93
L103:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v210)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v389
	F_errmsg_internal(m, int32(_a_F_examine_variable_6), v17+int32(16))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_7), int32(_a_F_examine_variable_3))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
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
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	if v404 <= v325 {
		goto L29
	} else {
		goto L107
	}
L107:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v406+v336)))
	if v408 <= int32(0) {
		goto L23
	} else {
		goto L108
	}
L108:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+16))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+12))
	v427 = v195 + int32(8)
	v433 = v413 + v408<<(uint(int32(2))%32) - int32(4)
	goto L65
L109:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+144))
	if v438 != 0 {
		goto L19
	} else {
		goto L110
	}
L110:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v437)+108))
	if v439 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)+96))
	if v440 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v442 = v440
	goto L114
L113:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437)+76))
	v442 = v441
	goto L114
L114:
	;
	v443 = int32(*(*int16)(unsafe.Add(mBase, uint32(v427))))
	if v442 != 0 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	if v481 == int32(0) {
		goto L22
	} else {
		goto L128
	}
L116:
	;
	goto L115
L117:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v447 <= int32(0) {
		v481 = int32(0)
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v481 = int32(0)
	goto L116
L120:
	;
	v450 = int32(0)
	if v450 < v447 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v453 = v447
	goto L123
L122:
	;
	v453 = v450
	goto L123
L123:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	v458 = int32(0)
	goto L124
L124:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v454+v458<<(uint(int32(2))%32))))
	v467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466)+8)))
	if v467 == v443&int32(_a_F_examine_variable_8) {
		v481 = v466
		goto L116
	} else {
		goto L126
	}
L125:
	;
	goto L119
L126:
	;
	v470 = v458 + int32(1)
	if v470 != v453 {
		v458 = v470
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+26)))
	if v485 == int32(1) {
		goto L22
	} else {
		goto L129
	}
L129:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v437)+120))
	if v488 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v489 != int32(1) {
		goto L19
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v437)+100))
	if v532 != 0 {
		goto L147
	} else {
		goto L148
	}
L133:
	;
	v492 = int32(0)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v481)+16))
	if base.B2i32(v494 == v492)|base.B2i32(v488 == v492) != 0 {
		v527 = v492
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v527 == int32(0) {
		goto L19
	} else {
		goto L146
	}
L135:
	;
	goto L134
L136:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if int32(0) < v500 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v504 = int32(0)
	goto L140
L138:
	;
	goto L139
L139:
	;
	v527 = v492
	goto L135
L140:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508+v504<<(uint(int32(2))%32))))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	if v494 == v513 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L139
L142:
	;
	v527 = int32(1)
	goto L135
L143:
	;
	goto L144
L144:
	;
	v517 = v504 + int32(1)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	if v517 < v518 {
		v504 = v517
		goto L140
	} else {
		goto L145
	}
L145:
	;
	goto L141
L146:
	;
	v530 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v530)
	goto L19
L147:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	if v533 != int32(1) {
		goto L19
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+40)))
	if v576 != 0 {
		goto L19
	} else {
		goto L164
	}
L150:
	;
	v536 = int32(0)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v481)+16))
	if base.B2i32(v538 == v536)|base.B2i32(v532 == v536) != 0 {
		v571 = v536
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v571 == int32(0) {
		goto L19
	} else {
		goto L163
	}
L152:
	;
	goto L151
L153:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	if int32(0) < v544 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v548 = int32(0)
	goto L157
L155:
	;
	goto L156
L156:
	;
	v571 = v536
	goto L152
L157:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v532)+12))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v552+v548<<(uint(int32(2))%32))))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+4))
	if v538 == v557 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	goto L156
L159:
	;
	v571 = int32(1)
	goto L152
L160:
	;
	goto L161
L161:
	;
	v561 = v548 + int32(1)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	if v561 < v562 {
		v548 = v561
		goto L157
	} else {
		goto L162
	}
L162:
	;
	goto L158
L163:
	;
	v574 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v574)
	goto L19
L164:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v577 == int32(0) {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v577)))
	if v580 != int32(6) {
		goto L19
	} else {
		goto L166
	}
L166:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v577)+28))
	if v583 == int32(0) {
		v191 = v434
		v195 = v577
		goto L54
	} else {
		goto L167
	}
L167:
	;
	goto L55
L168:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v589 = F_bms_difference(m, v586, v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	if v589 == int32(0) {
		goto L21
	} else {
		goto L170
	}
L170:
	;
	v595 = int32(0)
	if v589 == v595 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v666
	v723 = v87
	v724 = v668
	goto L20
L172:
	;
	if v649 != 0 {
		goto L187
	} else {
		goto L188
	}
L173:
	;
	v649 = int32(0)
	goto L172
L174:
	;
	goto L175
L175:
	;
	v603 = int32(1)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4))
	if v604 <= v603 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v607 = v603
	goto L178
L177:
	;
	v607 = v604
	goto L178
L178:
	;
	v612 = int32(0)
	v614 = int32(-1)
	goto L180
L179:
	;
	v649 = v641
	goto L172
L180:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v589+int32(8)+v612<<(uint(int32(2))%32))))
	if v622 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(76)))) = v633
	v641 = int32(1)
	goto L179
L182:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v622)))|base.B2i32(int32(0) <= v614) != 0 {
		v641 = v595
		goto L179
	} else {
		goto L185
	}
L183:
	;
	v633 = v614
	goto L184
L184:
	;
	v635 = v612 + int32(1)
	if v635 != v607 {
		v612 = v635
		v614 = v633
		goto L180
	} else {
		goto L186
	}
L185:
	;
	v633 = base.I32_ctz(v622) | v612<<(uint(int32(5))%32)
	goto L184
L186:
	;
	goto L181
L187:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if l2 != v651 {
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
	v653 = l2
	goto L192
L191:
	;
	v653 = int32(0)
	goto L192
L192:
	;
	if v653 != 0 {
		goto L21
	} else {
		goto L193
	}
L193:
	;
	v654 = F_find_base_rel(m, l0, v651)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v666 = v654
	v668 = v654
	goto L171
L195:
	;
	v658 = F_find_join_rel(m, l0, v586)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v660 = F_bms_is_member(m, l2, v586)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L199
	}
L198:
	;
	v666 = v658
	v668 = v5
	goto L171
L199:
	;
	if v660 == int32(0) {
		goto L21
	} else {
		goto L200
	}
L200:
	;
	v664 = F_find_base_rel(m, l0, l2)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v666 = v664
	v668 = v5
	goto L171
L202:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v210)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v675
	F_errmsg_internal(m, int32(_a_F_examine_variable_9), v17+int32(32))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_10), int32(_a_F_examine_variable_3))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
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
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v210)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v691
	F_errmsg_internal(m, int32(_a_F_examine_variable_11), v17+int32(48))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_12), int32(_a_F_examine_variable_3))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
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
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	v709 = int32(*(*int16)(unsafe.Add(mBase, uint32(v427))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v708
	F_errmsg_internal(m, int32(_a_F_examine_variable_13), v17)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_14), int32(_a_F_examine_variable_3))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v723
	v728 = F_exprType(m, v723)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v728
	v731 = F_exprTypmod(m, v723)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v731
	if v724 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	F_bms_free(m, v586)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L1
	} else {
		goto L346
	}
L215:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v737 = int32(0)
	if base.B2i32(v586 == v737)|base.B2i32(v736 == v737) != 0 {
		v782 = v737
		goto L217
	} else {
		goto L218
	}
L216:
	;
	if v782 != 0 {
		goto L229
	} else {
		goto L230
	}
L217:
	;
	goto L216
L218:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v736)+4))
	if v747 < v748 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v750 = v747
	goto L221
L220:
	;
	v750 = v748
	goto L221
L221:
	;
	if v750 <= int32(1) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v753 = int32(1)
	goto L224
L223:
	;
	v753 = v750
	goto L224
L224:
	;
	v754 = int32(8)
	v759 = int32(0)
	goto L225
L225:
	;
	v766 = v759 << (uint(int32(2)) % 32)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v736+v754+v766)))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v586+v754+v766)))
	v771 = v768 & v770
	v773 = base.B2i32(v771 != int32(0))
	if v771 != 0 {
		v782 = v773
		goto L217
	} else {
		goto L227
	}
L226:
	;
	v782 = v773
	goto L217
L227:
	;
	v775 = v759 + int32(1)
	if v775 != v753 {
		v759 = v775
		goto L225
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v785 = F_remove_nulling_relids(m, v723, v783, int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L232
	}
L230:
	;
	v787 = v723
	goto L231
L231:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v724)+116))
	if v788 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v787 = v785
	goto L231
L233:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v724)+120))
	if v1010 == int32(0) {
		goto L214
	} else {
		goto L293
	}
L234:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v788)+4))
	if v791 <= int32(0) {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v803 = v5
	goto L236
L236:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v788)+12))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v808+v803<<(uint(int32(2))%32))))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v812)+84))
	if v813 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	goto L233
L238:
	;
	v993 = v803 + int32(1)
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v788)+4))
	if v993 < v994 {
		v803 = v993
		goto L236
	} else {
		goto L292
	}
L239:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v813)+12))
	if v816 == int32(0) {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v812)+36))
	if int32(0) < v819 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v824 = v816
	v825 = v819
	v836 = int64(0)
	goto L244
L242:
	;
	goto L243
L243:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v977 != 0 {
		goto L233
	} else {
		goto L291
	}
L244:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v812)+44))
	v838 = base.I32_wrap_i64(v836)
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v837+v838<<(uint(int32(2))%32))))
	if v842 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	goto L243
L246:
	;
	if v824 != 0 {
		goto L251
	} else {
		goto L252
	}
L247:
	;
	v957 = v824
	v958 = v825
	goto L248
L248:
	;
	v960 = v836 + int64(1)
	if v960 < base.I64_extend_i32_s(v958) {
		v824 = v957
		v825 = v958
		v836 = v960
		goto L244
	} else {
		goto L290
	}
L249:
	;
	v946 = v824 + int32(4)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v812)+84))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v948)+12))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	if base.Ui32(v946) < base.Ui32(v949+v950<<(uint(int32(2))%32)) {
		goto L287
	} else {
		goto L288
	}
L250:
	;
	v942 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v942)
	goto L249
L251:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	if v845 == int32(0) {
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
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L284
	}
L254:
	;
	v854 = F_equal(m, v787, v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L259
	}
L255:
	;
	v853 = int32(0)
	goto L254
L256:
	;
	goto L257
L257:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	if v849 != int32(27) {
		v853 = v845
		goto L254
	} else {
		goto L258
	}
L258:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	v853 = v852
	goto L254
L259:
	;
	if v854 == int32(0) {
		goto L249
	} else {
		goto L260
	}
L260:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812)+101)))
	if base.B2i32(v858 != int32(1))|base.B2i32(v836 != int64(0)) != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_examine_variable[1]))
	if v874 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L262:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v812)+40))
	if v864 != int32(1) {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v812)+88))
	if v867 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812)+100)))
	if v868 != int32(1) {
		goto L261
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v871 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+28)) = uint8(v871)
	goto L261
L267:
	;
	goto L266
L268:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v812)+88))
	if v902 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L269:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	v881 = m.T0[v874].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v877, base.I32_extend16_s(v838+int32(1)), l3)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	if v881 == int32(0) {
		goto L268
	} else {
		goto L271
	}
L271:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v885 == int32(0) {
		goto L249
	} else {
		goto L272
	}
L272:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v888 != 0 {
		goto L233
	} else {
		goto L273
	}
L273:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_0), int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_15), int32(_a_F_examine_variable_16))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
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
	v906 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v812)+4)))
	v907 = int64(48)
	v914 = F_SearchSysCache3(m, int32(65), v906, (v836<<(uint(v907)%64)-int64(-281474976710656))>>(uint(v907)%64), int64(0))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v928 != 0 {
		goto L233
	} else {
		goto L283
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1704)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v914
	if v914 == int32(0) {
		goto L250
	} else {
		goto L281
	}
L281:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v812)+12))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v921)+76))
	v924 = F_all_rows_selectable(m, l0, v922, int32(0))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v924)
	goto L279
L283:
	;
	goto L249
L284:
	;
	F_errmsg_internal(m, int32(_a_F_examine_variable_17), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_examine_variable_1), int32(_a_F_examine_variable_18), int32(_a_F_examine_variable_16))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
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
	v955 = v946
	goto L289
L288:
	;
	v955 = int32(0)
	goto L289
L289:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v812)+36))
	v957 = v955
	v958 = v956
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
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+4))
	if v1013 <= int32(0) {
		goto L214
	} else {
		goto L294
	}
L294:
	;
	v1026 = int32(0)
	goto L295
L295:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v1031 != 0 {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	goto L214
L297:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v1046 != 0 {
		goto L214
	} else {
		goto L301
	}
L298:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v724)+76))
	v1045 = v1031 + v1032<<(uint(int32(2))%32)
	goto L297
L299:
	;
	goto L300
L300:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+52))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+12))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v724)+76))
	v1045 = v1038 + v1039<<(uint(int32(2))%32) - int32(4)
	goto L297
L301:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+12))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1047+v1026<<(uint(int32(2))%32))))
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051)+16)))
	if v1052 != int32(101) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1210 = v1026 + int32(1)
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+4))
	if v1210 < v1211 {
		v1026 = v1210
		goto L295
	} else {
		goto L345
	}
L303:
	;
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051)+8)))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1045)))
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056)+20)))
	if v1055 != v1057 {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+24))
	if v1059 == int32(0) {
		goto L302
	} else {
		goto L305
	}
L305:
	;
	v1062 = int32(0)
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	if v1063 <= v1062 {
		goto L302
	} else {
		goto L306
	}
L306:
	;
	v1067 = v1062
	goto L307
L307:
	;
	v1080 = int32(0)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+12))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1081+v1067<<(uint(int32(2))%32))))
	if v1085 == v1080 {
		v1092 = v1080
		goto L309
	} else {
		goto L310
	}
L308:
	;
	goto L302
L309:
	;
	v1093 = F_equal(m, v787, v1092)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1085)))
	if v1088 != int32(27) {
		v1092 = v1085
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+4))
	v1092 = v1091
	goto L309
L312:
	;
	if v1093 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056)+20)))
	v1097 = m.G0
	v1099 = v1097 - int32(48)
	m.G0 = v1099
	v1104 = F_SearchSysCache2(m, int32(62), base.I64_extend_i32_u(v1095), base.I64_extend_i32_u(v1096))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L318
	}
L314:
	;
	goto L315
L315:
	;
	v1192 = v1067 + int32(1)
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+4))
	if v1192 < v1193 {
		v1067 = v1192
		goto L307
	} else {
		goto L344
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1147
	if v1147 != 0 {
		goto L340
	} else {
		goto L341
	}
L317:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L337
	}
L318:
	;
	if v1104 != 0 {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	v1110 = F_SysCacheGetAttr(m, int32(62), v1104, int32(6), v1099+int32(47))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
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
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L1
	} else {
		goto L334
	}
L322:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099)+47)))
	if v1112 == int32(1) {
		goto L317
	} else {
		goto L323
	}
L323:
	;
	v1115 = F_DatumGetExpandedArray(m, v1110)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	F_deconstruct_expanded_array(m, v1115)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+52))
	if v1119 != 0 {
		goto L327
	} else {
		goto L328
	}
L326:
	;
	F_ReleaseCatCache(m, v1104)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L333
	}
L327:
	;
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067+v1119))))
	if v1122 != 0 {
		v1147 = int32(0)
		goto L326
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+48))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1123+v1067<<(uint(int32(3))%32))))
	v1128 = F_pg_detoast_datum(m, v1127)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L331
	}
L330:
	;
	goto L329
L331:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1128)))
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+40)) = v1128
	v1132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+36)) = v1132
	*(*uint16)(unsafe.Add(mBase, uint32(v1099)+32)) = uint16(v1132)
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+28)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+24)) = int32(base.Ui32(v1130) >> (uint(int32(2)) % 32))
	v1143 = F_heap_copytuple(m, v1099+int32(24))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1147 = v1143
	goto L326
L333:
	;
	m.G0 = v1099 + int32(48)
	goto L316
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1099))) = v1095
	F_errmsg_internal(m, int32(_a_F_examine_variable_19), v1099)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(_a_F_examine_variable_20), int32(2466), int32(_a_F_examine_variable_21))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+20)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v1099)+16)) = int32(101)
	F_errmsg_internal(m, int32(_a_F_examine_variable_22), v1099+int32(16))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_examine_variable_20), int32(2473), int32(_a_F_examine_variable_21))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1705)
	goto L342
L341:
	;
	goto L342
L342:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v724)+76))
	v1188 = F_all_rows_selectable(m, l0, v1186, int32(0))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+29)) = uint8(v1188)
	goto L302
L344:
	;
	goto L308
L345:
	;
	goto L296
L346:
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
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l1
				v25 = v23
			} else {
				v25 = v4
			}
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
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
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v36 != 0 {
							F_SPI_freetuptable(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v39
								v41 = int32(10)
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
							*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v39
							v41 = int32(10)
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v83
						v86 = *(*int64)(unsafe.Add(mBase, _c_F_exec_run_select[2]))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v86
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
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l1
			v25 = v23
		} else {
			v25 = v4
		}
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
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
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v36 != 0 {
						F_SPI_freetuptable(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v39
							v41 = int32(10)
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
						*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v39
						v41 = int32(10)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
					*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v83
					v86 = *(*int64)(unsafe.Add(mBase, _c_F_exec_run_select[2]))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v86
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
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
	v10 = Fn14265(m, l0, int32(47))
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
	v65 = v39 + v40<<(uint(int32(3))%32) + v33*int32(100) - int32(72)
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
	v74 = int32(1)
	v76 = F_makeConst(m, int32(23), int32(-1), int32(0), int32(4), int64(0), v74, v74)
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
	v94 = F_makeConst(m, v86, v88, v89, v90, int64(0), int32(1), v93)
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
	F_errfinish(m, int32(_a_F_expand_insert_targetlist_1), int32(506), int32(_a_F_expand_insert_targetlist_2))
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_mul_size(m, int32(4), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = l1 + v6
		v11 = F_mul_size(m, int32(4), v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = F_repalloc0(m, v4, v7, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v13
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v19 = F_mul_size(m, int32(4), v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v22 = F_mul_size(m, int32(4), v10)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = F_repalloc0(m, v16, v19, v22)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v24
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							if v27 != 0 {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v30 = F_mul_size(m, int32(4), v29)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									v33 = F_mul_size(m, int32(4), v10)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										v35 = F_repalloc0(m, v27, v30, v33)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return
										} else {
											v40 = v35
											*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v10
											*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v40
											return
										}
									}
								}
							} else {
								v38 = F_palloc0_mul(m, int32(4), v10)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									v40 = v38
									*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v10
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v40
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
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	v113 = v23 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v113 < v114 {
		v23 = v113
		goto L4
	} else {
		goto L44
	}
L7:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v107 = F_lappend(m, v106, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	v104 = l3
	v105 = v88
	goto L7
L29:
	;
	goto L30
L30:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+32)))
	if v92 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v104 = l3
	v105 = v88
	goto L7
L32:
	;
	goto L33
L33:
	;
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v88)+24))
	if v93 == int64(0) {
		v104 = l3
		v105 = v88
		goto L7
	} else {
		goto L34
	}
L34:
	;
	goto L6
L35:
	;
	v104 = l2
	v105 = v96
	goto L7
L36:
	;
	goto L37
L37:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+32)))
	if v100 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v104 = l2
	v105 = v96
	goto L7
L39:
	;
	goto L40
L40:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v96)+24))
	if v101 != int64(0) {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v104 = l2
	v105 = v96
	goto L7
L42:
	;
	return
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v107
	goto L6
L44:
	;
	goto L5
}
