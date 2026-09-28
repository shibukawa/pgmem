package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EOH_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	m.T0[v5].(func(*base.Module, int32, int32, int32))(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_ER_get_flat_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
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
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v9 != int32(2249) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v28&int32(17) == int32(1) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(0) <= v12 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = F_expanded_record_fetch_tupdesc(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v22 = v15
	goto L6
L6:
	;
	F_assign_record_type_typmod(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v22 = v18
	goto L6
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v25
	goto L1
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	return v34
L11:
	;
	goto L12
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v36 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v28&int32(4) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v166 = v36
	goto L15
L15:
	;
	return v166
L16:
	;
	F_deconstruct_expanded_record(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	v46 = v28
	goto L18
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v46&int32(16) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v46 = v45
	goto L18
L20:
	;
	if int32(0) < v47 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v111 = v47
	goto L22
L22:
	;
	v117 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v119 = int32(24)
	if v111 <= v117 {
		v151 = v119
		v152 = v117
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v57 = int32(0)
	v58 = v47
	goto L26
L24:
	;
	v99 = v47
	v105 = v46
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v105 & int32(-17)
	v111 = v99
	goto L22
L26:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v57))))
	if v66 != 0 {
		v89 = v58
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v99 = v89
	v105 = v96
	goto L25
L28:
	;
	v94 = v57 + int32(1)
	if v94 < v89 {
		v57 = v94
		v58 = v89
		goto L26
	} else {
		goto L34
	}
L29:
	;
	v68 = v57 << (uint(int32(3)) % 32)
	v69 = v48 + int32(28) + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)))
	if v70 != 0 {
		v89 = v58
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+2)))
	if v71 != int32(_a_F_ER_get_flat_size_0) {
		v89 = v58
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v74+v68)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v76)))))
	if v78 != int32(1) {
		v89 = v58
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v81 = int32(1)
	v83 = int32(0)
	F_expanded_record_set_field_internal(m, l0, v57+v81, v76, v83, v81, v83)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v89 = v88
	goto L28
L34:
	;
	goto L27
L35:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v157 = F_heap_compute_data_size(m, v48, v156, v118)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L43
	}
L36:
	;
	v124 = int32(0)
	goto L37
L37:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v118))))
	if v132 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v142 = base.I32_div_s(v138+int32(7), int32(8))
	v151 = (v142 + int32(30)) & int32(-8)
	v152 = int32(1)
	goto L35
L39:
	;
	v136 = v124 + int32(1)
	if v111 != v136 {
		v124 = v136
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	v151 = v119
	v152 = v117
	goto L35
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)) = uint8(v152)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v157
	v162 = v157 + v151
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v162
	v166 = v162
	goto L15
}
func F_ER_mc_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v3 == int32(0) {
		return
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v6
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
		if v8 <= v6 {
			return
		} else {
			v12 = v8 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v12
			if v12 != 0 {
				return
			} else {
				F_FreeTupleDesc(m, v3)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_ExecARDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if l4 == int32(0) {
		if v8 != 0 {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
			if v17 != 0 {
				v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if l3 == int32(0) {
						v27 = int32(0)
						v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v38 = int32(0)
							v40 = int32(1)
							F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = int32(0)
							v40 = int32(1)
							F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
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
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
					if v20 != int32(1) {
						return
					} else {
						v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if l3 == int32(0) {
								v27 = int32(0)
								v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
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
		} else {
			if l4 == int32(0) {
				return
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
				if v20 != int32(1) {
					return
				} else {
					v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						if l3 == int32(0) {
							v27 = int32(0)
							v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v38 = int32(0)
								v40 = int32(1)
								F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = int32(0)
								v40 = int32(1)
								F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
		if v11 == int32(0) {
			if v8 != 0 {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
				if v17 != 0 {
					v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						if l3 == int32(0) {
							v27 = int32(0)
							v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v38 = int32(0)
								v40 = int32(1)
								F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = int32(0)
								v40 = int32(1)
								F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
						if v20 != int32(1) {
							return
						} else {
							v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								if l3 == int32(0) {
									v27 = int32(0)
									v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										v38 = int32(0)
										v40 = int32(1)
										F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = int32(0)
										v40 = int32(1)
										F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
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
			} else {
				if l4 == int32(0) {
					return
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
					if v20 != int32(1) {
						return
					} else {
						v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if l3 == int32(0) {
								v27 = int32(0)
								v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
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
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
			if v14 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_ExecARDeleteTriggers_0), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecARDeleteTriggers_1), int32(2832), int32(_a_F_ExecARDeleteTriggers_2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
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
				if v8 != 0 {
					v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
					if v17 != 0 {
						v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if l3 == int32(0) {
								v27 = int32(0)
								v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
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
							v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
							if v20 != int32(1) {
								return
							} else {
								v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
								mBase = m.M
								v24 = m.ExcPending
								if v24 != 0 {
									return
								} else {
									if l3 == int32(0) {
										v27 = int32(0)
										v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
											return
										} else {
											v38 = int32(0)
											v40 = int32(1)
											F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return
											} else {
												return
											}
										}
									} else {
										F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											v38 = int32(0)
											v40 = int32(1)
											F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
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
				} else {
					if l4 == int32(0) {
						return
					} else {
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
						if v20 != int32(1) {
							return
						} else {
							v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								if l3 == int32(0) {
									v27 = int32(0)
									v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										v38 = int32(0)
										v40 = int32(1)
										F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = int32(0)
										v40 = int32(1)
										F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
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
func F_ExecARUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if l8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L33
	}
L2:
	;
	if v13 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8)+1)))
	if v19 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8)+2)))
	if v20 == int32(1) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	return
L8:
	;
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)))
	if v23 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if l8 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8)+1)))
	if v26 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8)+2)))
	if v27 != int32(1) {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	v30 = l2
	goto L18
L17:
	;
	v30 = l1
	goto L18
L18:
	;
	v31 = F_ExecGetTriggerOldSlot(m, l0, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	if l5 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v57 = F_ExecGetAllUpdatedCols(m, l1, l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L19
	} else {
		goto L31
	}
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	m.T0[v52].(func(*base.Module, int32))(m, v31)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
	} else {
		goto L30
	}
L23:
	;
	if l4 == int32(0) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_ExecForceStoreHeapTuple(m, l5, v31, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L19
	} else {
		goto L29
	}
L26:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if v37 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v40 = int32(0)
	v46 = F_GetTupleForTrigger(m, l0, v40, v30, l4, int32(3), v31, v40, v40, v40, v40)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	goto L21
L29:
	;
	goto L21
L30:
	;
	goto L21
L31:
	;
	F_AfterTriggerSaveEvent(m, l0, l1, l2, l3, int32(2), int32(1), v31, l6, l7, v57, l8, l9)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L7
L33:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_ExecARUpdateTriggers_0), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_ExecARUpdateTriggers_1), int32(3179), int32(_a_F_ExecARUpdateTriggers_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecBSDeleteTriggers(m *base.Module, l0 int32, l1 int32) {
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
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+21)))
	if v22 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+56))
	v28 = F_before_stmt_triggers_fired(m, v26, int32(4))
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
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(38654706112)
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
	if v47&int32(75) != int32(10) {
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
	F_errmsg(m, int32(_a_F_ExecBSDeleteTriggers_0), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_ExecBSDeleteTriggers_1), int32(2692), int32(_a_F_ExecBSDeleteTriggers_2))
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
func F_ExecCloseIndices(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v20 = v13 << (uint(int32(2)) % 32)
	v21 = v11 + v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10+v20)))
	F_index_insert_cleanup(m, v22, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_relation_close(m, v27, int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	v34 = v13 + int32(1)
	if v34 != v7 {
		v13 = v34
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_ExecGetRootToChildMap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
	if v7 == int32(0) {
		v10 = int32(_a_F_ExecGetRootToChildMap_0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGetRootToChildMap[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecGetRootToChildMap[0])) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+131)))
		v26 = F_build_attrmap_by_name_if_req(m, v16, v13, (v21^int32(-1))&int32(1))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			if v26 != 0 {
				v30 = F_convert_tuples_by_name_attrmap(m, v16, v13, v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v30
					*(*int32)(unsafe.Add(mBase, _c_F_ExecGetRootToChildMap[0])) = v11
					v35 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v35)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
					return v42
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_ExecGetRootToChildMap[0])) = v11
				v35 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v35)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
				return v42
			}
		}
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		return v42
	}
}
func F_ExecIRDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v14 = F_ExecGetTriggerOldSlot(m, l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(90194313664)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v28
	F_ExecForceStoreHeapTuple(m, l2, v14, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v34 <= int32(0) {
		v91 = int32(1)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v11 + int32(48)
	return v91
L5:
	;
	v44 = int32(0)
	goto L6
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v48 = v45 + v44*int32(60)
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+12)))
	if v49&int32(75) != int32(73) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v91 = v82
	goto L4
L8:
	;
	v82 = int32(1)
	v84 = v44 + v82
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v84 < v85 {
		v44 = v84
		goto L6
	} else {
		goto L22
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v55 = int32(0)
	v57 = F_TriggerEnabled(m, l0, l1, v48, v54, v55, v14, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v57 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v14
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v68 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v71 = v68
	goto L14
L13:
	;
	v69 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v73 = F_ExecCallTriggerFunc(m, v11+int32(4), v44, v66, v67, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v71 = v69
	goto L14
L16:
	;
	if v73 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v91 = int32(0)
	goto L4
L18:
	;
	goto L19
L19:
	;
	if l2 == v73 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	F_pfree(m, v73)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L8
L22:
	;
	goto L7
}
func F_ExecInitGenerated(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v22 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L14
	} else {
		goto L58
	}
L2:
	;
	m.G0 = v18 + int32(16)
	return
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)))
	if v29 != int32(1) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l2 != int32(2) {
		v41 = int32(0)
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v42 = int32(_a_F_ExecInitGenerated_0)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitGenerated[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitGenerated[0])) = v45
	v49 = F_palloc0(m, v25<<(uint(int32(2))%32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L16
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+13)))
	if v37 != 0 {
		v41 = int32(0)
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v38 = F_ExecGetUpdatedCols(m, l0, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	return
L15:
	;
	v41 = v38
	goto L8
L16:
	;
	if int32(0) < v25 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if l2 == int32(2) {
		goto L55
	} else {
		goto L56
	}
L18:
	;
	v57 = int32(0)
	v59 = v4
	goto L21
L19:
	;
	goto L20
L20:
	;
	F_pfree(m, v49)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L14
	} else {
		goto L53
	}
L21:
	;
	v70 = v57 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v71<<(uint(int32(3))%32)+v57*int32(100))+118)))
	if v78 == int32(0) {
		v160 = v59
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v160 != 0 {
		v187 = v160
		v188 = v49
		goto L17
	} else {
		goto L52
	}
L23:
	;
	if v70 != v25 {
		v57 = v70
		v59 = v160
		goto L21
	} else {
		goto L51
	}
L24:
	;
	v81 = F_build_column_default(m, v20, v70)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	if v81 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v41 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(0)
	F_pull_varattnos(m, v81, int32(1), v18+int32(12))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v78 == int32(115) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v93 = int32(0)
	if base.B2i32(v41 == v93)|base.B2i32(v92 == v93) != 0 {
		v138 = v93
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v138 == int32(0) {
		v160 = v59
		goto L23
	} else {
		goto L44
	}
L32:
	;
	goto L31
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v103 < v104 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v106 = v103
	goto L36
L35:
	;
	v106 = v104
	goto L36
L36:
	;
	if v106 <= int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v109 = int32(1)
	goto L39
L38:
	;
	v109 = v106
	goto L39
L39:
	;
	v110 = int32(8)
	v115 = int32(0)
	goto L40
L40:
	;
	v122 = v115 << (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v92+v110+v122)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v41+v110+v122)))
	v127 = v124 & v126
	v129 = base.B2i32(v127 != int32(0))
	if v127 != 0 {
		v138 = v129
		goto L32
	} else {
		goto L42
	}
L41:
	;
	v138 = v129
	goto L32
L42:
	;
	v131 = v115 + int32(1)
	if v131 != v109 {
		v115 = v131
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L29
L45:
	;
	v146 = F_ExecPrepareExpr(m, v81, l1)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L14
	} else {
		goto L48
	}
L46:
	;
	v151 = v59
	goto L47
L47:
	;
	if l2 != int32(2) {
		v160 = v151
		goto L23
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49+v57<<(uint(int32(2))%32)))) = v146
	v151 = v59 + int32(1)
	goto L47
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v157 = F_bms_add_member(m, v154, v57+int32(8))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v157
	v160 = v151
	goto L23
L51:
	;
	goto L22
L52:
	;
	goto L20
L53:
	;
	v180 = int32(0)
	v187 = v180
	v188 = v180
	goto L17
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitGenerated[0])) = v43
	goto L2
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v188
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v201)
	goto L54
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v188
	goto L54
L58:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v229 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ExecInitGenerated_1), v18)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ExecInitGenerated_2), int32(493), int32(_a_F_ExecInitGenerated_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecInitNullTupleSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = F_ExecInitExtraTupleSlot(m, l0, l1, int32(_a_F_ExecInitNullTupleSlot_0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		m.T0[v11].(func(*base.Module, int32))(m, v6)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v15 = int32(3)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v20 = v18 << (uint(v15) % 32)
			if v14&v15|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v20)) == int32(0) {
				if v20 == int32(0) {
				} else {
					v30 = v20 + v14
					v32 = v14 + int32(4)
					if base.Ui32(v32) < base.Ui32(v30) {
						v34 = v30
					} else {
						v34 = v32
					}
					v40 = (v14^int32(-1)+v34)&int32(-4) + int32(4)
					if v40 == int32(0) {
					} else {
						base.MemoryFill(m, v14, int32(0), v40)
					}
				}
			} else {
				v40 = v20
				if v40 == int32(0) {
				} else {
					base.MemoryFill(m, v14, int32(0), v40)
				}
			}
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			if v49 != 0 {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
				base.MemoryFill(m, v50, int32(1), v49)
			} else {
			}
			v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v55 = v53 & int32(_a_F_ExecInitNullTupleSlot_1)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)) = uint16(v55)
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v58)
			return v6
		}
	}
}
func F_ExecMaterial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMaterial[0]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v15 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v56 = int32(0)
	if base.B2i32(v53 == v56)|base.B2i32(v14 == int32(1)) == v56 {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v18 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v42 = v15
	goto L9
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v47*int32(24))+4)))
	goto L20
L10:
	;
	v21 = int32(1)
	v53 = v21
	v54 = v21
	v55 = v2
	goto L6
L11:
	;
	goto L12
L12:
	;
	v23 = int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMaterial[1]))
	v28 = F_tuplestore_begin_heap(m, v23, int32(0), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	F_tuplestore_set_eflags(m, v28, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v33&int32(16) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v36 = F_tuplestore_alloc_read_pointer(m, v28, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v28
	if v28 == int32(0) {
		v53 = int32(1)
		v54 = v23
		v55 = v2
		goto L6
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v42 = v28
	goto L9
L20:
	;
	v53 = v51
	v54 = int32(0)
	v55 = v42
	goto L6
L21:
	;
	return v115
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	m.T0[v112].(func(*base.Module, int32))(m, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L53
	}
L23:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v84 != 0 {
		v110 = v83
		goto L22
	} else {
		goto L37
	}
L24:
	;
	v79 = F_tuplestore_gettupleslot(m, v55, base.B2i32(v14 == int32(1)), int32(0), v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L34
	}
L25:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v63 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v53 != 0 {
		v83 = v74
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v66 = int32(0)
	v68 = F_tuplestore_advance(m, v55, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v75 = v73
	goto L24
L31:
	;
	if v68 == int32(0) {
		v115 = v66
		goto L21
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v75 = v74
	goto L24
L34:
	;
	if v79 != 0 {
		v115 = v75
		goto L21
	} else {
		goto L35
	}
L35:
	;
	if v14 != int32(1) {
		v110 = v75
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v83 = v75
	goto L23
L37:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+52))
	if v86 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_ExecReScan(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v90 = m.T0[v89].(func(*base.Module, int32) int32)(m, v85)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	if v54 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	if v90 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
	if v92&int32(2) == int32(0) {
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v97)
	return int32(0)
L47:
	;
	goto L46
L48:
	;
	F_tuplestore_puttupleslot(m, v55, v90)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+32))
	m.T0[v106].(func(*base.Module, int32, int32))(m, v83, v90)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	return v83
L53:
	;
	v115 = v110
	goto L21
}
func F_ExecMaterializesOutput(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	v3 = l0 - int32(352)
	return base.B2i32(base.Ui32(v3) < base.Ui32(int32(15))) & int32(base.Ui32(int32(_a_F_ExecMaterializesOutput_0))>>(uint(v3)%32))
}
func F_ExecOpenIndices(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+116)))
	if v13 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_RelationGetIndexList(m, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v16 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = F_palloc_mul(m, int32(4), v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v27 = F_palloc_mul(m, int32(4), v20)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) < v32 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_list_free(m, v16)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L22
	}
L12:
	;
	v45 = v36 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46)))
	v50 = F_index_open(m, v48, int32(3))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v52 = F_BuildIndexInfo(m, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if l1 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+v45))) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45+v27))) = v52
	v68 = v36 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v68 < v69 {
		v36 = v68
		goto L12
	} else {
		goto L21
	}
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+116)))
	if v56 != int32(1) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+192))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+15)))
	if v60 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_BuildSpeculativeIndexInfo(m, v50, v52)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	goto L13
L22:
	;
	goto L1
}
func F_ExecPendingInserts(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v11 = int32(0)
	goto L1
L1:
	;
	v15 = int32(0)
	if v8 == v15 {
		v25 = v15
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v7 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v19 <= v11 {
		v25 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v25 = v21 + v11<<(uint(int32(2))%32)
	goto L3
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v33+v11<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+108))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+108)))
	F_ExecBatchInsert(m, v46, v47, v48, v49, v50, l0, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L14
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	F_list_free(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.B2i32(v25 == int32(0))|base.B2i32(v30 <= v11) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v33 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	return
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	F_list_free(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	return
L14:
	;
	v11 = v11 + int32(1)
	goto L1
}
func F_ExecProcessReturning(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	if l2 != 0 {
		if l3 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l5
			v23 = int32(8)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
			if v24&int32(2) == int32(0) {
				v34 = int32(0)
				v35 = v23
				*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v34
				if l4 != 0 {
					v47 = int32(0)
					v48 = l4
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
					v54 = v50&int32(231) | v35 | v47
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
					m.T0[v59].(func(*base.Module, int32))(m, v57)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = int32(_a_F_ExecProcessReturning_0)
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
						v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
							v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
							v77 = v75 & int32(_a_F_ExecProcessReturning_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
							return v57
						}
					}
				} else {
					v38 = int32(0)
					v39 = int32(16)
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
					if v40&int32(4) == v38 {
						v47 = v39
						v48 = v38
						*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						v54 = v50&int32(231) | v35 | v47
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
						m.T0[v59].(func(*base.Module, int32))(m, v57)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = int32(_a_F_ExecProcessReturning_0)
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
								v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
								v77 = v75 & int32(_a_F_ExecProcessReturning_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
								return v57
							}
						}
					} else {
						v45 = F_ExecGetAllNullSlot(m, v9, l1)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = v39
							v48 = v45
							*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
							v54 = v50&int32(231) | v35 | v47
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
							m.T0[v59].(func(*base.Module, int32))(m, v57)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = int32(_a_F_ExecProcessReturning_0)
								v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
								v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
									v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
									v77 = v75 & int32(_a_F_ExecProcessReturning_1)
									*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
									*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
									return v57
								}
							}
						}
					}
				}
			} else {
				v30 = F_ExecGetAllNullSlot(m, v9, l1)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v30
					v35 = v23
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v34
					if l4 != 0 {
						v47 = int32(0)
						v48 = l4
						*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						v54 = v50&int32(231) | v35 | v47
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
						m.T0[v59].(func(*base.Module, int32))(m, v57)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = int32(_a_F_ExecProcessReturning_0)
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
								v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
								v77 = v75 & int32(_a_F_ExecProcessReturning_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
								return v57
							}
						}
					} else {
						v38 = int32(0)
						v39 = int32(16)
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						if v40&int32(4) == v38 {
							v47 = v39
							v48 = v38
							*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
							v54 = v50&int32(231) | v35 | v47
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
							m.T0[v59].(func(*base.Module, int32))(m, v57)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = int32(_a_F_ExecProcessReturning_0)
								v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
								v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
									v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
									v77 = v75 & int32(_a_F_ExecProcessReturning_1)
									*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
									*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
									return v57
								}
							}
						} else {
							v45 = F_ExecGetAllNullSlot(m, v9, l1)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = v39
								v48 = v45
								*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
								v54 = v50&int32(231) | v35 | v47
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
								m.T0[v59].(func(*base.Module, int32))(m, v57)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = int32(_a_F_ExecProcessReturning_0)
									v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
									v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
										v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
										v77 = v75 & int32(_a_F_ExecProcessReturning_1)
										*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
										*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
										return v57
									}
								}
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l3
			v34 = l3
			v35 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v34
			if l4 != 0 {
				v47 = int32(0)
				v48 = l4
				*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
				v54 = v50&int32(231) | v35 | v47
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
				m.T0[v59].(func(*base.Module, int32))(m, v57)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v62 = int32(_a_F_ExecProcessReturning_0)
					v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
					v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
						v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
						v77 = v75 & int32(_a_F_ExecProcessReturning_1)
						*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
						return v57
					}
				}
			} else {
				v38 = int32(0)
				v39 = int32(16)
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
				if v40&int32(4) == v38 {
					v47 = v39
					v48 = v38
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
					v54 = v50&int32(231) | v35 | v47
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
					m.T0[v59].(func(*base.Module, int32))(m, v57)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = int32(_a_F_ExecProcessReturning_0)
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
						v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
							v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
							v77 = v75 & int32(_a_F_ExecProcessReturning_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
							return v57
						}
					}
				} else {
					v45 = F_ExecGetAllNullSlot(m, v9, l1)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = v39
						v48 = v45
						*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						v54 = v50&int32(231) | v35 | v47
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
						m.T0[v59].(func(*base.Module, int32))(m, v57)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = int32(_a_F_ExecProcessReturning_0)
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
								v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
								v77 = v75 & int32(_a_F_ExecProcessReturning_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
								return v57
							}
						}
					}
				}
			}
		}
	} else {
		if l4 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l4
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l5
		if l3 == int32(0) {
			v23 = int32(8)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
			if v24&int32(2) == int32(0) {
				v34 = int32(0)
				v35 = v23
				*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v34
				if l4 != 0 {
					v47 = int32(0)
					v48 = l4
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
					v54 = v50&int32(231) | v35 | v47
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
					m.T0[v59].(func(*base.Module, int32))(m, v57)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = int32(_a_F_ExecProcessReturning_0)
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
						v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
							v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
							v77 = v75 & int32(_a_F_ExecProcessReturning_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
							return v57
						}
					}
				} else {
					v38 = int32(0)
					v39 = int32(16)
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
					if v40&int32(4) == v38 {
						v47 = v39
						v48 = v38
						*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						v54 = v50&int32(231) | v35 | v47
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
						m.T0[v59].(func(*base.Module, int32))(m, v57)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = int32(_a_F_ExecProcessReturning_0)
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
								v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
								v77 = v75 & int32(_a_F_ExecProcessReturning_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
								return v57
							}
						}
					} else {
						v45 = F_ExecGetAllNullSlot(m, v9, l1)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = v39
							v48 = v45
							*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
							v54 = v50&int32(231) | v35 | v47
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
							m.T0[v59].(func(*base.Module, int32))(m, v57)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = int32(_a_F_ExecProcessReturning_0)
								v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
								v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
									v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
									v77 = v75 & int32(_a_F_ExecProcessReturning_1)
									*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
									*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
									return v57
								}
							}
						}
					}
				}
			} else {
				v30 = F_ExecGetAllNullSlot(m, v9, l1)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v30
					v35 = v23
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v34
					if l4 != 0 {
						v47 = int32(0)
						v48 = l4
						*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						v54 = v50&int32(231) | v35 | v47
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
						m.T0[v59].(func(*base.Module, int32))(m, v57)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = int32(_a_F_ExecProcessReturning_0)
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
								v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
								v77 = v75 & int32(_a_F_ExecProcessReturning_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
								return v57
							}
						}
					} else {
						v38 = int32(0)
						v39 = int32(16)
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						if v40&int32(4) == v38 {
							v47 = v39
							v48 = v38
							*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
							v54 = v50&int32(231) | v35 | v47
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
							m.T0[v59].(func(*base.Module, int32))(m, v57)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = int32(_a_F_ExecProcessReturning_0)
								v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
								v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
									v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
									v77 = v75 & int32(_a_F_ExecProcessReturning_1)
									*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
									*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
									return v57
								}
							}
						} else {
							v45 = F_ExecGetAllNullSlot(m, v9, l1)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = v39
								v48 = v45
								*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
								v54 = v50&int32(231) | v35 | v47
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
								m.T0[v59].(func(*base.Module, int32))(m, v57)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = int32(_a_F_ExecProcessReturning_0)
									v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
									v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
										v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
										v77 = v75 & int32(_a_F_ExecProcessReturning_1)
										*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
										*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
										return v57
									}
								}
							}
						}
					}
				}
			}
		} else {
			v34 = l3
			v35 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v34
			if l4 != 0 {
				v47 = int32(0)
				v48 = l4
				*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
				v54 = v50&int32(231) | v35 | v47
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
				m.T0[v59].(func(*base.Module, int32))(m, v57)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v62 = int32(_a_F_ExecProcessReturning_0)
					v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
					v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
						v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
						v77 = v75 & int32(_a_F_ExecProcessReturning_1)
						*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
						*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
						return v57
					}
				}
			} else {
				v38 = int32(0)
				v39 = int32(16)
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
				if v40&int32(4) == v38 {
					v47 = v39
					v48 = v38
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
					v54 = v50&int32(231) | v35 | v47
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
					m.T0[v59].(func(*base.Module, int32))(m, v57)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = int32(_a_F_ExecProcessReturning_0)
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
						v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
							v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
							v77 = v75 & int32(_a_F_ExecProcessReturning_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
							*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
							return v57
						}
					}
				} else {
					v45 = F_ExecGetAllNullSlot(m, v9, l1)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = v39
						v48 = v45
						*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v48
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
						v54 = v50&int32(231) | v35 | v47
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
						m.T0[v59].(func(*base.Module, int32))(m, v57)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = int32(_a_F_ExecProcessReturning_0)
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v65
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int64)(m, v10+int32(8), v56, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v63
								v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)))
								v77 = v75 & int32(_a_F_ExecProcessReturning_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+4)) = uint16(v77)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
								*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)) = uint16(v80)
								return v57
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecProject(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	m.T0[v8].(func(*base.Module, int32))(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(_a_F_ExecProject_0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProject[0]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecProject[0])) = v16
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v22 = m.T0[v21].(func(*base.Module, int32, int32, int32) int64)(m, l0+int32(8), v5, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_ExecProject[0])) = v14
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v28 = v26 & int32(_a_F_ExecProject_1)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)) = uint16(v28)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v31)
			return v6
		}
	}
}
func F_ExecSecLabelStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v204 int32
	_ = v204
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int64
	_ = v255
	var v257 int32
	_ = v257
	var v263 int64
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v340 int64
	_ = v340
	var v342 int64
	_ = v342
	var v344 int64
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int64
	_ = v359
	var v361 int32
	_ = v361
	var v367 int64
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSecLabelStmt[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v15 == v3 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L29
	} else {
		goto L150
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L29
	} else {
		goto L146
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L29
	} else {
		goto L142
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L29
	} else {
		goto L138
	}
L5:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v87 - int32(1) {
	case 0, 5, 8, 11, 13, 17, 18, 20, 21, 22, 28, 29, 33, 34, 36, 37, 38, 41, 42, 49, 51:
		goto L27
	default:
		goto L28
	}
L6:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v14 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v20 != int32(1) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v81 = v24
	goto L5
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v27 <= int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v30 = int32(0)
	if v30 < v27 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v33 = v27
	goto L15
L14:
	;
	v33 = v30
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v39 = v3
	goto L16
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34+v39<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if base.B2i32(v50 == int32(0))|base.B2i32(v50 != v53) != 0 {
		v71 = v50
		v72 = v53
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L2
L18:
	;
	if v71-v72 == int32(0) {
		v81 = v46
		goto L5
	} else {
		goto L25
	}
L19:
	;
	goto L18
L20:
	;
	v56 = v15
	v57 = v47
	goto L21
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v61 == int32(0) {
		v71 = v61
		v72 = v60
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v71 = v61
	v72 = v60
	goto L19
L23:
	;
	v64 = int32(1)
	if v61 == v60 {
		v56 = v56 + v64
		v57 = v57 + v64
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v77 = v39 + int32(1)
	if v33 != v77 {
		v39 = v77
		goto L16
	} else {
		goto L26
	}
L26:
	;
	goto L17
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecSecLabelStmt[1])))
	F_get_object_address(m, l0, v87, v106, v11+int32(44), int32(4), v111&base.B2i32(v87 == int32(22)))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L34
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return
L30:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_0), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(162), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSecLabelStmt[2]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v121 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v123
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_check_object_ownership(m, v118, v120, v11+int32(16), v119, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v130 == int32(6) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+48))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+119)))
	v137 = v135 - int32(99)
	if base.B2i32(base.Ui32(int32(19)) < base.Ui32(v137))|base.B2i32(int32(1)<<(uint(v137)%32)&int32(_a_F_ExecSecLabelStmt_3) == int32(0)) != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	m.T0[v149].(func(*base.Module, int32, int32))(m, l0, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L29
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v154 = int32(0)
	v155 = m.G0
	v157 = v155 - int32(288)
	m.G0 = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v162 = int32(1)
	if v159 <= int32(3591) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	F_relation_close(m, v444, int32(3))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L29
	} else {
		goto L133
	}
L42:
	;
	if v233 != 0 {
		goto L67
	} else {
		goto L68
	}
L43:
	;
	goto L42
L44:
	;
	v233 = int32(0)
	goto L43
L45:
	;
	if base.B2i32(base.Ui32(v159-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v159-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v233 = v162
		goto L43
	} else {
		goto L66
	}
L46:
	;
	if v159 <= int32(2670) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if v159 <= int32(_a_F_ExecSecLabelStmt_4) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	switch v159 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v233 = v162
		goto L43
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L44
	default:
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v174 = v159 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v174))|base.B2i32(int32(1)<<(uint(v174)%32)&int32(226492515) == int32(0)) != 0 {
		goto L45
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v159-int32(2396)) {
		goto L44
	} else {
		goto L53
	}
L53:
	;
	v233 = v162
	goto L43
L54:
	;
	v233 = v162
	goto L43
L55:
	;
	if base.Ui32(v159-int32(3592)) < base.Ui32(int32(2)) {
		v233 = v162
		goto L43
	} else {
		goto L64
	}
L56:
	;
	v187 = v159 - int32(_a_F_ExecSecLabelStmt_5)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v187))|base.B2i32(int32(1)<<(uint(v187)%32)&int32(963) == int32(0)) != 0 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	switch v159 - int32(_a_F_ExecSecLabelStmt_6) {
	case 0, 1, 2, 3, 4, 59, 60:
		v233 = v162
		goto L43
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L44
	default:
		goto L60
	}
L59:
	;
	v233 = v162
	goto L43
L60:
	;
	if base.Ui32(v159-int32(_a_F_ExecSecLabelStmt_7)) < base.Ui32(int32(3)) {
		v233 = v162
		goto L43
	} else {
		goto L61
	}
L61:
	;
	v204 = v159 - int32(_a_F_ExecSecLabelStmt_8)
	if base.Ui32(int32(15)) < base.Ui32(v204) {
		goto L44
	} else {
		goto L62
	}
L62:
	;
	if int32(1)<<(uint(v204)%32)&int32(_a_F_ExecSecLabelStmt_9) != 0 {
		v233 = v162
		goto L43
	} else {
		goto L63
	}
L63:
	;
	goto L44
L64:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v159-int32(4060)) {
		goto L44
	} else {
		goto L65
	}
L65:
	;
	v233 = v162
	goto L43
L66:
	;
	goto L44
L67:
	;
	v234 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v234
	v238 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+16)) = v238
	v240 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+24)) = v240
	v242 = F_cstring_to_text(m, v152)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L29
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v332 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+12)) = uint8(v332)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v332
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)) = uint8(v332)
	v340 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+16)) = v340
	v342 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+24)) = v342
	v344 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+32)) = v344
	v346 = F_cstring_to_text(m, v152)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L29
	} else {
		goto L101
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v157)+32)) = base.I64_extend_i32_u(v242)
	if v153 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v246 = F_cstring_to_text(m, v153)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L29
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v251 = v157 - int32(-64)
	v255 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	F_ScanKeyInit(m, v251, int32(1), int32(3), int32(184), v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L29
	} else {
		goto L75
	}
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v157)+40)) = base.I64_extend_i32_u(v246)
	goto L73
L75:
	;
	v263 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	F_ScanKeyInit(m, v157+int32(120), int32(2), int32(3), int32(184), v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L29
	} else {
		goto L76
	}
L76:
	;
	v268 = int32(3)
	v271 = F_cstring_to_text(m, v152)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L29
	} else {
		goto L77
	}
L77:
	;
	F_ScanKeyInit(m, v157+int32(176), v268, v268, int32(67), base.I64_extend_i32_u(v271))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L29
	} else {
		goto L78
	}
L78:
	;
	v278 = F_table_open(m, int32(3592), int32(3))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L29
	} else {
		goto L80
	}
L79:
	;
	v444 = v278
	goto L41
L80:
	;
	v284 = F_systable_beginscan(m, v278, int32(3593), int32(1), int32(0), int32(3), v251)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L29
	} else {
		goto L81
	}
L81:
	;
	v286 = F_systable_getnext(m, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L29
	} else {
		goto L82
	}
L82:
	;
	if v286 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v153 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v309 = v154
	goto L85
L85:
	;
	F_systable_endscan(m, v284)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L29
	} else {
		goto L93
	}
L86:
	;
	F_simple_heap_delete(m, v278, v286+int32(4))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L29
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)) = uint8(v296)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v278)+52))
	v305 = F_heap_modify_tuple(m, v286, v300, v157+int32(16), v157+int32(8), v157)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L29
	} else {
		goto L91
	}
L89:
	;
	F_systable_endscan(m, v284)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L29
	} else {
		goto L90
	}
L90:
	;
	goto L79
L91:
	;
	F_CatalogTupleUpdate(m, v278, v286+int32(4), v305)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L29
	} else {
		goto L92
	}
L92:
	;
	v309 = v305
	goto L85
L93:
	;
	v312 = int32(0)
	if base.B2i32(v153 == v312)|v309 == v312 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v278)+52))
	v322 = F_heap_form_tuple(m, v317, v157+int32(16), v157+int32(8))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L29
	} else {
		goto L97
	}
L95:
	;
	v326 = v309
	goto L96
L96:
	;
	if v326 == int32(0) {
		goto L79
	} else {
		goto L99
	}
L97:
	;
	F_CatalogTupleInsert(m, v278, v322)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L29
	} else {
		goto L98
	}
L98:
	;
	v326 = v322
	goto L96
L99:
	;
	F_pfree(m, v326)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L29
	} else {
		goto L100
	}
L100:
	;
	goto L79
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v157)+40)) = base.I64_extend_i32_u(v346)
	if v153 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v350 = F_cstring_to_text(m, v153)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L29
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v355 = v157 - int32(-64)
	v359 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	F_ScanKeyInit(m, v355, int32(1), int32(3), int32(184), v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L29
	} else {
		goto L106
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v157)+48)) = base.I64_extend_i32_u(v350)
	goto L104
L106:
	;
	v367 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	F_ScanKeyInit(m, v157+int32(120), int32(2), int32(3), int32(184), v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L29
	} else {
		goto L107
	}
L107:
	;
	v372 = int32(3)
	v375 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	F_ScanKeyInit(m, v157+int32(176), v372, v372, int32(65), v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L29
	} else {
		goto L108
	}
L108:
	;
	v383 = F_cstring_to_text(m, v152)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L29
	} else {
		goto L109
	}
L109:
	;
	F_ScanKeyInit(m, v157+int32(232), int32(4), int32(3), int32(67), base.I64_extend_i32_u(v383))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L29
	} else {
		goto L110
	}
L110:
	;
	v390 = F_table_open(m, int32(3596), int32(3))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L29
	} else {
		goto L112
	}
L111:
	;
	v444 = v390
	goto L41
L112:
	;
	v396 = F_systable_beginscan(m, v390, int32(3597), int32(1), int32(0), int32(4), v355)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L29
	} else {
		goto L113
	}
L113:
	;
	v398 = F_systable_getnext(m, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L29
	} else {
		goto L114
	}
L114:
	;
	if v398 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if v153 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v421 = v154
	goto L117
L117:
	;
	F_systable_endscan(m, v396)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L29
	} else {
		goto L125
	}
L118:
	;
	F_simple_heap_delete(m, v390, v398+int32(4))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L29
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)) = uint8(v408)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v390)+52))
	v417 = F_heap_modify_tuple(m, v398, v412, v157+int32(16), v157+int32(8), v157)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L29
	} else {
		goto L123
	}
L121:
	;
	F_systable_endscan(m, v396)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L29
	} else {
		goto L122
	}
L122:
	;
	goto L111
L123:
	;
	F_CatalogTupleUpdate(m, v390, v398+int32(4), v417)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L29
	} else {
		goto L124
	}
L124:
	;
	v421 = v417
	goto L117
L125:
	;
	v424 = int32(0)
	if base.B2i32(v153 == v424)|v421 == v424 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v390)+52))
	v434 = F_heap_form_tuple(m, v429, v157+int32(16), v157+int32(8))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L29
	} else {
		goto L129
	}
L127:
	;
	v438 = v421
	goto L128
L128:
	;
	if v438 == int32(0) {
		goto L111
	} else {
		goto L131
	}
L129:
	;
	F_CatalogTupleInsert(m, v390, v434)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L29
	} else {
		goto L130
	}
L130:
	;
	v438 = v434
	goto L128
L131:
	;
	F_pfree(m, v438)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L29
	} else {
		goto L132
	}
L132:
	;
	goto L111
L133:
	;
	m.G0 = v157 + int32(288)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v454 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	F_relation_close(m, v454, int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L29
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	m.G0 = v11 + int32(48)
	return
L137:
	;
	goto L136
L138:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L29
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_10), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L29
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(133), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L29
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L29
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_11), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L29
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(137), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L29
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L29
	} else {
		goto L147
	}
L147:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v508
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_12), v11+int32(32))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L29
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(156), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L29
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L29
	} else {
		goto L151
	}
L151:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v527)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v528 + int32(4)
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_13), v11)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L29
	} else {
		goto L152
	}
L152:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+48))
	v537 = int32(*(*int8)(unsafe.Add(mBase, uint32(v536)+119)))
	F_errdetail_relkind_not_supported(m, v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L29
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(206), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L29
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExpireTreeKnownAssignedTransactionIds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ExpireTreeKnownAssignedTransactionIds[0]))
	v11 = F_LWLockAcquire(m, v7+int32(512), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		F_KnownAssignedXidsRemoveTree(m, l0, l1, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExpireTreeKnownAssignedTransactionIds[1]))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
			if v18 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v17 + base.I64_extend_i32_s(l3-base.I32_wrap_i64(v17))
			} else {
				v21 = int32(3)
				if base.B2i32(base.Ui32(l3) < base.Ui32(v21))|base.B2i32(base.Ui32(v18) < base.Ui32(v21)) == int32(0) {
					if v18-l3 < int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v17 + base.I64_extend_i32_s(l3-base.I32_wrap_i64(v17))
					} else {
					}
				} else {
					if base.Ui32(l3) <= base.Ui32(v18) {
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v17 + base.I64_extend_i32_s(l3-base.I32_wrap_i64(v17))
					}
				}
			}
			v37 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
			*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v37 + int64(1)
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_ExpireTreeKnownAssignedTransactionIds[0]))
			F_LWLockRelease(m, v42+int32(512))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F___extenddftf2(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v54 int64
	_ = v54
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	v3 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = base.I64_reinterpret_f64(l1)
	v15 = v13 & int64(4503599627370495)
	v19 = int64(base.Ui64(v13)>>(uint(int64(52))%64)) & int64(2047)
	if v19 != v3 {
		if v19 != int64(2047) {
			v74 = v19 + int64(15360)
			v75 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
			v77 = v15 << (uint(int64(60)) % 64)
		} else {
			v74 = int64(32767)
			v75 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
			v77 = v15 << (uint(int64(60)) % 64)
		}
	} else {
		if v15 == int64(0) {
			v37 = int64(0)
			v74 = v37
			v75 = v3
			v77 = v37
		} else {
			v39 = int64(0)
			v41 = base.I32_wrap_i64(base.I64_clz(v15))
			v43 = v41 + int32(49)
			if v43&int32(64) != 0 {
				v62 = int64(0)
				v63 = v15 << (uint(base.I64_extend_i32_u(v41+int32(-15))) % 64)
			} else {
				if v43 == int32(0) {
					v62 = v15
					v63 = v39
				} else {
					v54 = base.I64_extend_i32_u(v43)
					v62 = v15 << (uint(v54) % 64)
					v63 = v39<<(uint(v54)%64) | int64(base.Ui64(v15)>>(uint(base.I64_extend_i32_u(int32(64)-v43))%64))
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v62
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v63
			v67 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
			v73 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v74 = base.I64_extend_i32_u(int32(_a_F___extenddftf2_0) - v41)
			v75 = v67 ^ int64(281474976710656)
			v77 = v73
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v77
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v13&int64(-9223372036854775807-1) | v74<<(uint(int64(48))%64) | v75
	m.G0 = v11 + int32(16)
	return
}
func F__equalA_Expr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != v5 {
		v25 = v3
		return v25
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v9 = F_equal(m, v7, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v25 = v3
				return v25
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v17 = F_equal(m, v15, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 == int32(0) {
						v25 = v3
						return v25
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v23 = F_equal(m, v21, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = v23
							return v25
						}
					}
				}
			}
		}
	}
}
func F__equalA_Indices(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = int32(0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v4 != v5 {
		v19 = v3
		return v19
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v9 = F_equal(m, v7, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v19 = v3
				return v19
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v17 = F_equal(m, v15, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = v17
					return v19
				}
			}
		}
	}
}
func F__equalCoerceViaIO(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	var v18 int32
	_ = v18
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v18 = v3
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v12 != v13 {
				v18 = v3
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v18 = base.B2i32(v15 == v16)
			}
		}
		return v18
	}
}
func F__equalNullTest(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	var v18 int32
	_ = v18
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v18 = v3
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v12 != v13 {
				v18 = v3
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
				v18 = base.B2i32(v15 == v16)
			}
		}
		return v18
	}
}
func F__equalRelabelType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v21 = v3
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v12 != v13 {
				v21 = v3
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v15 != v16 {
					v21 = v3
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v21 = base.B2i32(v18 == v19)
				}
			}
		}
		return v21
	}
}
func F_each_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+32))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_each_array_start_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_each_array_start_1), int32(2176), int32(_a_F_each_array_start_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		return int32(0)
	}
}
func F_each_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
	if v5 != 0 {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
		if v6 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l1
		} else {
		}
		return int32(0)
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_each_scalar_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_each_scalar_1), int32(2190), int32(_a_F_each_scalar_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
func F_ec_member_matches_indexcol(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v9 = v7 << (uint(int32(2)) % 32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+v11)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	if v14 != int32(403) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v13 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18+v9)))
	v21 = int32(0)
	if v17 == v21 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L4:
	;
	v59 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v27 <= int32(0) {
		v53 = v21
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v59 = v53
	goto L3
L8:
	;
	v30 = int32(0)
	if v30 < v27 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v33 = v27
	goto L11
L10:
	;
	v33 = v30
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v36 = int32(0)
	goto L12
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34+v36<<(uint(int32(2))%32))))
	v45 = base.B2i32(v44 == v20)
	if v44 == v20 {
		v53 = v45
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v53 = v45
	goto L7
L14:
	;
	v47 = v36 + int32(1)
	if v47 != v33 {
		v36 = v47
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return int32(0)
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v69 = F_match_index_to_operand(m, v68, v7, v10)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v13 == v64 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	return int32(0)
L20:
	;
	return int32(0)
L21:
	;
	return v69
}
func F_elements_array_element_end(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+7)) = uint8(v3)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v12 == int32(1) {
		v15 = int32(_a_F_elements_array_element_end_0)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v18
		if l1 == int32(0) {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
			if v29 == int32(1) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v33 = F_cstring_to_text(m, v32)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = base.I64_extend_i32_u(v33)
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v39)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v55 = F_heap_form_tuple(m, v50, v7+int32(8), v7+int32(7))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_tuplestore_puttuple(m, v57, v55)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_MemoryContextReset(m, v62)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return int32(0)
							}
						}
					}
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
				v45 = F_cstring_to_text_with_len(m, v41, v43-v41)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = base.I64_extend_i32_u(v45)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v55 = F_heap_form_tuple(m, v50, v7+int32(8), v7+int32(7))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_tuplestore_puttuple(m, v57, v55)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_MemoryContextReset(m, v62)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return int32(0)
							}
						}
					}
				}
			}
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v22 != int32(1) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
				if v29 == int32(1) {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v33 = F_cstring_to_text(m, v32)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = base.I64_extend_i32_u(v33)
						v39 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v39)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v55 = F_heap_form_tuple(m, v50, v7+int32(8), v7+int32(7))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_tuplestore_puttuple(m, v57, v55)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								F_MemoryContextReset(m, v62)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(16)
									return int32(0)
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
					v45 = F_cstring_to_text_with_len(m, v41, v43-v41)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = base.I64_extend_i32_u(v45)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v55 = F_heap_form_tuple(m, v50, v7+int32(8), v7+int32(7))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_tuplestore_puttuple(m, v57, v55)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								F_MemoryContextReset(m, v62)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(16)
									return int32(0)
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
				v27 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+7)) = uint8(v27)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v55 = F_heap_form_tuple(m, v50, v7+int32(8), v7+int32(7))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_tuplestore_puttuple(m, v57, v55)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						F_MemoryContextReset(m, v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		m.G0 = v7 + int32(16)
		return int32(0)
	}
}
func F_elements_object_start(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14258(m, l0, int32(_a_F_elements_object_start_0), int32(2427), int32(_a_F_elements_object_start_1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_elements_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v9 != 0 {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
		if v10 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
		} else {
		}
		m.G0 = v6 + int32(16)
		return int32(0)
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v28
				F_errmsg(m, int32(_a_F_elements_scalar_0), v6)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_elements_scalar_1), int32(2442), int32(_a_F_elements_scalar_2))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
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
func F_eq_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6-v7 < l1 {
		v77 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v77
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = v10 + v7
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v73 != 0 {
		v77 = v4
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v73 = int32(0)
	goto L3
L5:
	;
	v47 = v42
	v48 = v43
	v49 = v44
	goto L15
L6:
	;
	if (v11|l2)&int32(3) != 0 {
		v42 = v11
		v43 = l2
		v44 = l1
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v35 = v11
	v36 = l2
	v37 = l1
	goto L8
L8:
	;
	if v37 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v19 = v11
	v20 = l2
	v21 = l1
	goto L10
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v24 != v25 {
		v42 = v19
		v43 = v20
		v44 = v21
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v35 = v30
	v36 = v28
	v37 = v32
	goto L8
L12:
	;
	v27 = int32(4)
	v28 = v20 + v27
	v30 = v19 + v27
	v32 = v21 - v27
	if base.Ui32(int32(3)) < base.Ui32(v32) {
		v19 = v30
		v20 = v28
		v21 = v32
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v42 = v35
	v43 = v36
	v44 = v37
	goto L5
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v52 == v53 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v73 = v52 - v53
	goto L3
L17:
	;
	v55 = int32(1)
	v60 = v49 - v55
	if v60 != 0 {
		v47 = v47 + v55
		v48 = v48 + v55
		v49 = v60
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1 + v7
	v77 = int32(1)
	goto L1
}
func F_eqjoinsel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 float64
	_ = v22
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 float32
	_ = v72
	var v74 float32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v101 float64
	_ = v101
	var v102 float64
	_ = v102
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 float64
	_ = v117
	var v120 int32
	_ = v120
	var v127 float64
	_ = v127
	var v130 float64
	_ = v130
	var v131 int32
	_ = v131
	var v136 float64
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 float32
	_ = v148
	var v150 float32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v177 float64
	_ = v177
	var v178 float64
	_ = v178
	var v182 int32
	_ = v182
	var v183 float64
	_ = v183
	var v186 float64
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 float64
	_ = v193
	var v196 int32
	_ = v196
	var v203 float64
	_ = v203
	var v206 float64
	_ = v206
	var v207 int32
	_ = v207
	var v212 float64
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
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
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int64
	_ = v399
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v424 float32
	_ = v424
	var v425 float32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v443 float64
	_ = v443
	var v444 float64
	_ = v444
	var v452 float64
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v465 int32
	_ = v465
	var v478 int32
	_ = v478
	var v486 float64
	_ = v486
	var v487 float64
	_ = v487
	var v503 float32
	_ = v503
	var v504 float64
	_ = v504
	var v507 int32
	_ = v507
	var v508 float64
	_ = v508
	var v510 int32
	_ = v510
	var v514 float32
	_ = v514
	var v515 float64
	_ = v515
	var v518 int32
	_ = v518
	var v519 float64
	_ = v519
	var v521 float64
	_ = v521
	var v523 float64
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v552 float64
	_ = v552
	var v553 float64
	_ = v553
	var v569 float32
	_ = v569
	var v570 float64
	_ = v570
	var v573 int32
	_ = v573
	var v574 float64
	_ = v574
	var v576 float64
	_ = v576
	var v598 float64
	_ = v598
	var v599 float64
	_ = v599
	var v612 float64
	_ = v612
	var v620 float64
	_ = v620
	var v623 float64
	_ = v623
	var v656 float64
	_ = v656
	var v657 float64
	_ = v657
	var v663 float64
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v693 int32
	_ = v693
	var v701 float64
	_ = v701
	var v702 float64
	_ = v702
	var v718 float32
	_ = v718
	var v719 float64
	_ = v719
	var v722 int32
	_ = v722
	var v723 float64
	_ = v723
	var v725 int32
	_ = v725
	var v729 float32
	_ = v729
	var v730 float64
	_ = v730
	var v733 int32
	_ = v733
	var v734 float64
	_ = v734
	var v736 float64
	_ = v736
	var v738 float64
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v767 float64
	_ = v767
	var v768 float64
	_ = v768
	var v784 float32
	_ = v784
	var v785 float64
	_ = v785
	var v788 int32
	_ = v788
	var v789 float64
	_ = v789
	var v791 float64
	_ = v791
	var v813 float64
	_ = v813
	var v814 float64
	_ = v814
	var v827 float64
	_ = v827
	var v836 float64
	_ = v836
	var v839 float64
	_ = v839
	var v863 float64
	_ = v863
	var v865 float64
	_ = v865
	var v877 float64
	_ = v877
	var v880 float64
	_ = v880
	var v884 float64
	_ = v884
	var v892 float64
	_ = v892
	var v893 float64
	_ = v893
	var v901 float64
	_ = v901
	var v902 int32
	_ = v902
	var v903 float64
	_ = v903
	var v909 float64
	_ = v909
	var v910 float64
	_ = v910
	var v917 float64
	_ = v917
	var v918 float64
	_ = v918
	var v924 float64
	_ = v924
	var v931 float64
	_ = v931
	var v933 float64
	_ = v933
	var v935 float32
	_ = v935
	var v938 float64
	_ = v938
	var v941 float32
	_ = v941
	var v944 float64
	_ = v944
	var v948 float64
	_ = v948
	var v985 float64
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
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
	var v1065 int32
	_ = v1065
	var v1077 float64
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1096 float64
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 float64
	_ = v1098
	var v1099 float64
	_ = v1099
	var v1100 float64
	_ = v1100
	var v1102 float64
	_ = v1102
	var v1106 float64
	_ = v1106
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1132 float64
	_ = v1132
	var v1140 float64
	_ = v1140
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1170 int32
	_ = v1170
	v2 = int32(0)
	v22 = float64(0)
	v36 = m.G0
	v38 = v36 - int32(240)
	m.G0 = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+132)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v38)+44)) = v2
	v52 = v38 + int32(200)
	v54 = v38 + int32(168)
	F_get_join_variables(m, v44, v43, v42, v52, v54, v38+int32(43))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v62 = v38 + int32(167)
	v63 = int32(0)
	v64 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v63)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v68 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v138 = v38 + int32(166)
	v139 = int32(0)
	v140 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v139)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	if v144 != 0 {
		goto L38
	} else {
		goto L39
	}
L4:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+28)))
	if v106 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+22)))
	v71 = v69 + v70
	v72 = *(*float32)(unsafe.Add(mBase, uint32(v71)+8))
	v74 = *(*float32)(unsafe.Add(mBase, uint32(v71)+16))
	v101 = base.F64_promote_f32(v74)
	v102 = base.F64_promote_f32(v72)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v76 == int32(16) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v101 = float64(2)
	v102 = v64
	goto L4
L9:
	;
	goto L10
L10:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v80 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v87 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+84))
	if v83 != int32(5) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v101 = float64(-1)
	v102 = v64
	goto L4
L14:
	;
	v101 = float64(0)
	v102 = v64
	goto L4
L15:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v90 != int32(6) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+8)))
	switch v94 - int32(_a_F_eqjoinsel_0) {
	case 0:
		goto L17
	default:
		goto L14
	case 5:
		v101 = float64(-1)
		v102 = v64
		goto L4
	}
L17:
	;
	v101 = float64(1)
	v102 = v64
	goto L4
L18:
	;
	v107 = base.F64_neg(base.F64_sub(float64(1), v102))
	goto L20
L19:
	;
	v107 = v101
	goto L20
L20:
	;
	if base.F64_gt(v107, float64(0)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v110 = F_clamp_row_est(m, v107)
	mBase = m.M
	v136 = v110
	goto L3
L22:
	;
	goto L23
L23:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v111 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v114 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v114)
	v136 = float64(200)
	goto L3
L25:
	;
	goto L26
L26:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v111)+128))
	if base.F64_le(v117, float64(0)) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v120 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v120)
	v136 = float64(200)
	goto L3
L28:
	;
	goto L29
L29:
	;
	if base.F64_lt(v107, float64(0)) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v127 = F_clamp_row_est(m, base.F64_mul(v117, base.F64_neg(v107)))
	mBase = m.M
	v136 = v127
	goto L3
L31:
	;
	goto L32
L32:
	;
	if base.F64_lt(v117, float64(200)) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v130 = F_clamp_row_est(m, v117)
	mBase = m.M
	v136 = v130
	goto L3
L34:
	;
	goto L35
L35:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v131)
	v136 = float64(200)
	goto L3
L36:
	;
	v213 = F_get_opcode(m, v41)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L69
	}
L37:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+28)))
	if v182 != 0 {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+22)))
	v147 = v145 + v146
	v148 = *(*float32)(unsafe.Add(mBase, uint32(v147)+8))
	v150 = *(*float32)(unsafe.Add(mBase, uint32(v147)+16))
	v177 = base.F64_promote_f32(v150)
	v178 = base.F64_promote_f32(v148)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v152 == int32(16) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v177 = float64(2)
	v178 = v140
	goto L37
L42:
	;
	goto L43
L43:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v156 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v163 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+84))
	if v159 != int32(5) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v177 = float64(-1)
	v178 = v140
	goto L37
L47:
	;
	v177 = float64(0)
	v178 = v140
	goto L37
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v166 != int32(6) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+8)))
	switch v170 - int32(_a_F_eqjoinsel_0) {
	case 0:
		goto L50
	default:
		goto L47
	case 5:
		v177 = float64(-1)
		v178 = v140
		goto L37
	}
L50:
	;
	v177 = float64(1)
	v178 = v140
	goto L37
L51:
	;
	v183 = base.F64_neg(base.F64_sub(float64(1), v178))
	goto L53
L52:
	;
	v183 = v177
	goto L53
L53:
	;
	if base.F64_gt(v183, float64(0)) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v186 = F_clamp_row_est(m, v183)
	mBase = m.M
	v212 = v186
	goto L36
L55:
	;
	goto L56
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v187 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v190)
	v212 = float64(200)
	goto L36
L58:
	;
	goto L59
L59:
	;
	v193 = *(*float64)(unsafe.Add(mBase, uint32(v187)+128))
	if base.F64_le(v193, float64(0)) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v196)
	v212 = float64(200)
	goto L36
L61:
	;
	goto L62
L62:
	;
	if base.F64_lt(v183, float64(0)) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v203 = F_clamp_row_est(m, base.F64_mul(v193, base.F64_neg(v183)))
	mBase = m.M
	v212 = v203
	goto L36
L64:
	;
	goto L65
L65:
	;
	if base.F64_lt(v193, float64(200)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v206 = F_clamp_row_est(m, v193)
	mBase = m.M
	v212 = v206
	goto L36
L67:
	;
	goto L68
L68:
	;
	v207 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v207)
	v212 = float64(200)
	goto L36
L69:
	;
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+120)) = v215
	v217 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+112)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v38)+104)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v38)+96)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v38)+88)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v38)+56)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v38)+72)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v215
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v38)+208))
	if v236 == v215 {
		v259 = v215
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v38)+208))
	if v260 == int32(0) {
		v304 = v2
		v305 = v2
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v38)+176))
	if v239 == int32(0) {
		v259 = v215
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v245 = int32(0)
	v247 = F_get_attstatsslot(m, v38+int32(88), v236, int32(1), v245, v245)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v247 == int32(0) {
		v259 = v215
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v38)+176))
	v255 = int32(0)
	v257 = F_get_attstatsslot(m, v38+int32(48), v253, int32(1), v255, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v259 = v257
	goto L70
L76:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+166)))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+167)))
	v417 = int32(0)
	if base.B2i32(v411&int32(1) == v417)|base.B2i32(v409 == v417) == v417 {
		goto L125
	} else {
		goto L126
	}
L77:
	;
	v396 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+160)) = v396
	v399 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+152)) = v399
	*(*int64)(unsafe.Add(mBase, uint32(v38)+144)) = v399
	*(*int64)(unsafe.Add(mBase, uint32(v38)+136)) = v399
	v406 = v396
	v408 = v2
	v409 = v392
	v410 = v393
	v411 = v394
	v412 = v395
	goto L76
L78:
	;
	v392 = v2
	v393 = v387
	v394 = v388
	v395 = v265
	goto L77
L79:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+22)))
	v387 = v382 + v383
	v388 = int32(0)
	goto L78
L80:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v306)+16))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+22)))
	v311 = v309 + v310
	if v259 == int32(0) {
		v392 = v304
		v393 = v311
		v394 = v2
		v395 = v305
		goto L77
	} else {
		goto L100
	}
L81:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v38)+176))
	if v307 != 0 {
		goto L79
	} else {
		goto L99
	}
L82:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v38)+176))
	if v306 != 0 {
		goto L80
	} else {
		goto L98
	}
L83:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+22)))
	v265 = v263 + v264
	if v259 == int32(0) {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+229)))
	if v270 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v286 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L93
	}
L86:
	;
	v278 = v260
	goto L88
L87:
	;
	if v213 == int32(0) {
		v304 = v2
		v305 = v265
		goto L82
	} else {
		goto L89
	}
L88:
	;
	v282 = F_get_attstatsslot(m, v38+int32(88), v278, int32(1), int32(0), int32(3))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L92
	}
L89:
	;
	v273 = F_get_func_leakproof(m, v213)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if v273 == int32(0) {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v38)+208))
	v278 = v277
	goto L88
L92:
	;
	v304 = v282
	v305 = v265
	goto L82
L93:
	;
	if v286 == int32(0) {
		v304 = v2
		v305 = v265
		goto L82
	} else {
		goto L94
	}
L94:
	;
	v290 = F_get_func_name(m, v213)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v290
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_1), v38+int32(32))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_2), int32(_a_F_eqjoinsel_3), int32(_a_F_eqjoinsel_4))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v304 = v2
	v305 = v265
	goto L82
L98:
	;
	v392 = v304
	v393 = v2
	v394 = v2
	v395 = v305
	goto L77
L99:
	;
	v387 = v2
	v388 = int32(0)
	goto L78
L100:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+197)))
	if v314 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	F_fmgr_info(m, v213, v38+int32(136))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L116
	}
L102:
	;
	v338 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L111
	}
L103:
	;
	if v213 == int32(0) {
		v392 = v304
		v393 = v311
		v394 = v2
		v395 = v305
		goto L77
	} else {
		goto L106
	}
L104:
	;
	v324 = v306
	goto L105
L105:
	;
	v325 = int32(1)
	v331 = F_get_attstatsslot(m, v38+int32(48), v324, v325, int32(0), int32(3))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L109
	}
L106:
	;
	v319 = F_get_func_leakproof(m, v213)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if v319 == int32(0) {
		goto L102
	} else {
		goto L108
	}
L108:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v38)+176))
	v324 = v323
	goto L105
L109:
	;
	if v304&v331 == int32(1) {
		goto L101
	} else {
		goto L110
	}
L110:
	;
	v392 = v304
	v393 = v311
	v394 = v331
	v395 = v305
	goto L77
L111:
	;
	if v338 == int32(0) {
		v392 = v304
		v393 = v311
		v394 = v2
		v395 = v305
		goto L77
	} else {
		goto L112
	}
L112:
	;
	v342 = F_get_func_name(m, v213)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v342
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_1), v38+int32(16))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_2), int32(_a_F_eqjoinsel_3), int32(_a_F_eqjoinsel_4))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v392 = v304
	v393 = v311
	v394 = v2
	v395 = v305
	goto L77
L116:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v360 = F_palloc0(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	v363 = F_palloc0(m, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	if v365+v366 < int32(200) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v406 = v360
	v408 = v363
	v409 = int32(1)
	v410 = v311
	v411 = v325
	v412 = v305
	goto L76
L120:
	;
	goto L121
L121:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	v373 = F_exprType(m, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v379 = F_get_op_hash_functions_ext(m, v41, v373, v38+int32(132), v38+int32(128))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v406 = v360
	v408 = v363
	v409 = int32(1)
	v410 = v311
	v411 = v325
	v412 = v305
	goto L76
L124:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	switch v986 {
	case 0, 1, 2:
		v1106 = v985
		goto L232
	default:
		goto L230
	case 4, 5:
		goto L233
	}
L125:
	;
	v424 = *(*float32)(unsafe.Add(mBase, uint32(v412)+8))
	v425 = *(*float32)(unsafe.Add(mBase, uint32(v410)+8))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v38)+132))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v38)+128))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	F_eqjoinsel_find_matches(m, v38+int32(136), v40, v428, v429, int32(0), v38+int32(88), v38+int32(48), v435, v436, v406, v408, v38+int32(44), v38+int32(232))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if v412 != 0 {
		goto L221
	} else {
		goto L222
	}
L128:
	;
	v443 = float64(0)
	v444 = *(*float64)(unsafe.Add(mBase, uint32(v38)+232))
	if base.F64_lt(v444, v443) != 0 {
		v452 = v443
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	if v453 <= int32(0) {
		v656 = v22
		v657 = v22
		goto L132
	} else {
		goto L133
	}
L130:
	;
	if base.F64_gt(v444, float64(1)) == int32(0) {
		v452 = v444
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v452 = float64(1)
	goto L129
L132:
	;
	v663 = float64(0)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	if v665 <= int32(0) {
		v863 = v663
		v865 = v663
		goto L166
	} else {
		goto L167
	}
L133:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v38)+108))
	if v453 == int32(1) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v612 = float64(0)
	if base.F64_lt(v598, v612) != 0 {
		v620 = v612
		goto L161
	} else {
		goto L162
	}
L135:
	;
	v569 = *(*float32)(unsafe.Add(mBase, uint32(v456+v531<<(uint(int32(2))%32))))
	v570 = base.F64_promote_f32(v569)
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531+v406))))
	if v573 != 0 {
		goto L155
	} else {
		goto L156
	}
L136:
	;
	v531 = int32(0)
	v552 = v22
	v553 = v22
	goto L135
L137:
	;
	goto L138
L138:
	;
	v465 = int32(0)
	v478 = v2
	v486 = v22
	v487 = v22
	goto L139
L139:
	;
	v503 = *(*float32)(unsafe.Add(mBase, uint32(v456+v465<<(uint(int32(2))%32))))
	v504 = base.F64_promote_f32(v503)
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465+v406))))
	if v507 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	if v453&int32(1) == int32(0) {
		v598 = v519
		v599 = v523
		goto L134
	} else {
		goto L154
	}
L141:
	;
	v508 = base.F64_add(v486, v504)
	goto L143
L142:
	;
	v508 = v486
	goto L143
L143:
	;
	v510 = v465 | int32(1)
	v514 = *(*float32)(unsafe.Add(mBase, uint32(v456+v510<<(uint(int32(2))%32))))
	v515 = base.F64_promote_f32(v514)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v510))))
	if v518 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v519 = base.F64_add(v508, v515)
	goto L146
L145:
	;
	v519 = v508
	goto L146
L146:
	;
	if v507 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v521 = v487
	goto L149
L148:
	;
	v521 = base.F64_add(v487, v504)
	goto L149
L149:
	;
	if v518 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v523 = v521
	goto L152
L151:
	;
	v523 = base.F64_add(v521, v515)
	goto L152
L152:
	;
	v524 = int32(2)
	v525 = v465 + v524
	v527 = v478 + v524
	if v527 != v453&int32(2147483646) {
		v465 = v525
		v478 = v527
		v486 = v519
		v487 = v523
		goto L139
	} else {
		goto L153
	}
L153:
	;
	goto L140
L154:
	;
	v531 = v525
	v552 = v519
	v553 = v523
	goto L135
L155:
	;
	v574 = base.F64_add(v552, v570)
	goto L157
L156:
	;
	v574 = v552
	goto L157
L157:
	;
	if v573 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v576 = v553
	goto L160
L159:
	;
	v576 = base.F64_add(v553, v570)
	goto L160
L160:
	;
	v598 = v574
	v599 = v576
	goto L134
L161:
	;
	if base.F64_lt(v599, float64(0)) != 0 {
		v656 = v22
		v657 = v620
		goto L132
	} else {
		goto L164
	}
L162:
	;
	if base.F64_gt(v598, float64(1)) == int32(0) {
		v620 = v598
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v620 = float64(1)
	goto L161
L164:
	;
	v623 = float64(1)
	if base.F64_gt(v599, v623) != 0 {
		v656 = v623
		v657 = v620
		goto L132
	} else {
		goto L165
	}
L165:
	;
	v656 = v599
	v657 = v620
	goto L132
L166:
	;
	v877 = float64(1)
	v880 = base.F64_sub(base.F64_sub(base.F64_sub(v877, base.F64_promote_f32(v425)), v863), v865)
	v884 = base.F64_sub(base.F64_sub(base.F64_sub(v877, base.F64_promote_f32(v424)), v657), v656)
	if base.F64_lt(v884, float64(0)) != 0 {
		v892 = v22
		goto L200
	} else {
		goto L201
	}
L167:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v665 == int32(1) {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v827 = float64(0)
	if base.F64_lt(v813, v827) != 0 {
		v836 = v827
		goto L195
	} else {
		goto L196
	}
L169:
	;
	v784 = *(*float32)(unsafe.Add(mBase, uint32(v668+v746<<(uint(int32(2))%32))))
	v785 = base.F64_promote_f32(v784)
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746+v408))))
	if v788 != 0 {
		goto L189
	} else {
		goto L190
	}
L170:
	;
	v746 = int32(0)
	v767 = v663
	v768 = float64(0)
	goto L169
L171:
	;
	goto L172
L172:
	;
	v677 = int32(0)
	v680 = v677
	v693 = v677
	v701 = v663
	v702 = float64(0)
	goto L173
L173:
	;
	v718 = *(*float32)(unsafe.Add(mBase, uint32(v668+v680<<(uint(int32(2))%32))))
	v719 = base.F64_promote_f32(v718)
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680+v408))))
	if v722 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	if v665&int32(1) == int32(0) {
		v813 = v734
		v814 = v738
		goto L168
	} else {
		goto L188
	}
L175:
	;
	v723 = base.F64_add(v701, v719)
	goto L177
L176:
	;
	v723 = v701
	goto L177
L177:
	;
	v725 = v680 | int32(1)
	v729 = *(*float32)(unsafe.Add(mBase, uint32(v668+v725<<(uint(int32(2))%32))))
	v730 = base.F64_promote_f32(v729)
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725+v408))))
	if v733 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v734 = base.F64_add(v723, v730)
	goto L180
L179:
	;
	v734 = v723
	goto L180
L180:
	;
	if v722 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v736 = v702
	goto L183
L182:
	;
	v736 = base.F64_add(v702, v719)
	goto L183
L183:
	;
	if v733 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v738 = v736
	goto L186
L185:
	;
	v738 = base.F64_add(v736, v730)
	goto L186
L186:
	;
	v739 = int32(2)
	v740 = v680 + v739
	v742 = v693 + v739
	if v742 != v665&int32(2147483646) {
		v680 = v740
		v693 = v742
		v701 = v734
		v702 = v738
		goto L173
	} else {
		goto L187
	}
L187:
	;
	goto L174
L188:
	;
	v746 = v740
	v767 = v734
	v768 = v738
	goto L169
L189:
	;
	v789 = base.F64_add(v767, v785)
	goto L191
L190:
	;
	v789 = v767
	goto L191
L191:
	;
	if v788 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v791 = v768
	goto L194
L193:
	;
	v791 = base.F64_add(v768, v785)
	goto L194
L194:
	;
	v813 = v789
	v814 = v791
	goto L168
L195:
	;
	if base.F64_lt(v814, float64(0)) != 0 {
		v863 = v836
		v865 = v827
		goto L166
	} else {
		goto L198
	}
L196:
	;
	if base.F64_gt(v813, float64(1)) == int32(0) {
		v836 = v813
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v836 = float64(1)
	goto L195
L198:
	;
	v839 = float64(1)
	if base.F64_gt(v814, v839) != 0 {
		v863 = v836
		v865 = v839
		goto L166
	} else {
		goto L199
	}
L199:
	;
	v863 = v836
	v865 = v814
	goto L166
L200:
	;
	v893 = float64(0)
	if base.F64_lt(v880, v893) != 0 {
		v901 = v893
		goto L203
	} else {
		goto L204
	}
L201:
	;
	if base.F64_gt(v884, float64(1)) == int32(0) {
		v892 = v884
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v892 = float64(1)
	goto L200
L203:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v38)+44))
	v903 = base.F64_convert_i32_s(v665)
	if base.F64_lt(v903, v212) != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	if base.F64_gt(v880, float64(1)) == int32(0) {
		v901 = v880
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v901 = float64(1)
	goto L203
L206:
	;
	v909 = base.F64_add(v452, base.F64_div(base.F64_mul(v656, v901), base.F64_sub(v212, v903)))
	goto L208
L207:
	;
	v909 = v452
	goto L208
L208:
	;
	v910 = base.F64_convert_i32_s(v902)
	if base.F64_lt(v910, v212) != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v917 = base.F64_add(base.F64_div(base.F64_mul(v892, base.F64_add(v865, v901)), base.F64_sub(v212, v910)), v909)
	goto L211
L210:
	;
	v917 = v909
	goto L211
L211:
	;
	v918 = base.F64_convert_i32_s(v453)
	if base.F64_lt(v918, v136) != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v924 = base.F64_add(v452, base.F64_div(base.F64_mul(v865, v892), base.F64_sub(v136, v918)))
	goto L214
L213:
	;
	v924 = v452
	goto L214
L214:
	;
	if base.F64_gt(v136, v910) != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v931 = base.F64_add(base.F64_div(base.F64_mul(base.F64_add(v656, v892), v901), base.F64_sub(v136, v910)), v924)
	goto L217
L216:
	;
	v931 = v924
	goto L217
L217:
	;
	if base.F64_lt(v917, v931) != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v933 = v917
	goto L220
L219:
	;
	v933 = v931
	goto L220
L220:
	;
	v985 = v933
	goto L124
L221:
	;
	v935 = *(*float32)(unsafe.Add(mBase, uint32(v412)+8))
	v938 = base.F64_promote_f32(v935)
	goto L223
L222:
	;
	v938 = float64(0)
	goto L223
L223:
	;
	if v410 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v941 = *(*float32)(unsafe.Add(mBase, uint32(v410)+8))
	v944 = base.F64_promote_f32(v941)
	goto L226
L225:
	;
	v944 = float64(0)
	goto L226
L226:
	;
	if base.F64_gt(v136, v212) != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v948 = v136
	goto L229
L228:
	;
	v948 = v212
	goto L229
L229:
	;
	v985 = base.F64_div(base.F64_mul(base.F64_sub(float64(1), v938), base.F64_sub(float64(1), v944)), v948)
	goto L124
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L290
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L287
	}
L232:
	;
	F_free_attstatsslot(m, v38+int32(88))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L266
	}
L233:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v987 == int32(0) {
		goto L231
	} else {
		goto L234
	}
L234:
	;
	v992 = int32(0)
	if v987 == v992 {
		goto L237
	} else {
		goto L238
	}
L235:
	;
	if v1052 == int32(0) {
		goto L231
	} else {
		goto L256
	}
L236:
	;
	if v1046 != 0 {
		goto L251
	} else {
		goto L252
	}
L237:
	;
	v1046 = int32(0)
	goto L236
L238:
	;
	goto L239
L239:
	;
	v1000 = int32(1)
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v987)+4))
	if v1001 <= v1000 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1004 = v1000
	goto L242
L241:
	;
	v1004 = v1001
	goto L242
L242:
	;
	v1009 = int32(0)
	v1011 = int32(-1)
	goto L244
L243:
	;
	v1046 = v1038
	goto L236
L244:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v987+int32(8)+v1009<<(uint(int32(2))%32))))
	if v1019 != 0 {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(232)))) = v1030
	v1038 = int32(1)
	goto L243
L246:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v1019)))|base.B2i32(int32(0) <= v1011) != 0 {
		v1038 = v992
		goto L243
	} else {
		goto L249
	}
L247:
	;
	v1030 = v1011
	goto L248
L248:
	;
	v1032 = v1009 + int32(1)
	if v1032 != v1004 {
		v1009 = v1032
		v1011 = v1030
		goto L244
	} else {
		goto L250
	}
L249:
	;
	v1030 = base.I32_ctz(v1019) | v1009<<(uint(int32(5))%32)
	goto L248
L250:
	;
	goto L245
L251:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v38)+232))
	v1048 = F_find_base_rel(m, v44, v1047)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v1050 = F_find_join_rel(m, v44, v987)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L255
	}
L254:
	;
	v1052 = v1048
	goto L235
L255:
	;
	v1052 = v1050
	goto L235
L256:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v38)+128))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v38)+132))
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+43)))
	if v1057 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1099 = *(*float64)(unsafe.Add(mBase, uint32(v1052)+16))
	v1100 = base.F64_mul(v985, v1099)
	if base.F64_gt(v1100, v1098) != 0 {
		goto L263
	} else {
		goto L264
	}
L258:
	;
	v1065 = int32(1)
	v1077 = F_eqjoinsel_semi(m, v38+int32(136), v40, v1056, v1055, int32(0), v38+int32(168), v136, v212, v414&v1065, v413&v1065, v38+int32(88), v38+int32(48), v412, v409, v411&v1065, v406, v408, v38+int32(44), v1052)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v1081 = int32(1)
	v1096 = F_eqjoinsel_semi(m, v38+int32(136), v40, v1056, v1055, v1081, v38+int32(200), v212, v136, v413&v1081, v414&v1081, v38+int32(48), v38+int32(88), v410, v411&v1081, v409, v408, v406, v38+int32(44), v1052)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L262
	}
L261:
	;
	v1098 = v1077
	goto L257
L262:
	;
	v1098 = v1096
	goto L257
L263:
	;
	v1102 = v1098
	goto L265
L264:
	;
	v1102 = v1100
	goto L265
L265:
	;
	v1106 = v1102
	goto L232
L266:
	;
	F_free_attstatsslot(m, v38+int32(48))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v38)+208))
	if v1117 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v38)+212))
	m.T0[v1118].(func(*base.Module, int32))(m, v1117)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v38)+176))
	if v1121 != 0 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	goto L270
L272:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v38)+180))
	m.T0[v1122].(func(*base.Module, int32))(m, v1121)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	if v406 != 0 {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	goto L274
L276:
	;
	F_pfree(m, v406)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	if v408 != 0 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	goto L278
L280:
	;
	F_pfree(m, v408)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	m.G0 = v38 + int32(240)
	v1132 = float64(0)
	if base.F64_lt(v1106, v1132) != 0 {
		v1140 = v1132
		goto L284
	} else {
		goto L285
	}
L283:
	;
	goto L282
L284:
	;
	return base.I64_reinterpret_f64(v1140)
L285:
	;
	if base.F64_gt(v1106, float64(1)) == int32(0) {
		v1140 = v1106
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1140 = float64(1)
	goto L284
L287:
	;
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_5), int32(0))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_2), int32(_a_F_eqjoinsel_6), int32(_a_F_eqjoinsel_7))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L1
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
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v1161
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_8), v38)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_2), int32(2559), int32(_a_F_eqjoinsel_9))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errcode_for_file_access(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_errcode_for_file_access[0]))
	if int32(0) <= v5 {
		v9 = v5 * int32(100)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_errcode_for_file_access[1])))
		switch v13 - int32(2) {
		case 0, 61, 67:
			v25 = int32(16797828)
		default:
			v25 = int32(2600)
		case 18:
			v25 = int32(33686021)
		case 27:
			v25 = int32(_a_F_errcode_for_file_access_0)
		case 29, 52, 53:
			v25 = int32(151027844)
		case 31, 39:
			v25 = int32(197)
		case 35:
			v25 = int32(50463237)
		case 42:
			v25 = int32(16908805)
		case 46:
			v25 = int32(_a_F_errcode_for_file_access_1)
		case 49:
			v25 = int32(_a_F_errcode_for_file_access_2)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_errcode_for_file_access[2]))) = v25
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_errcode_for_file_access[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_errcode_for_file_access_3), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_errcode_for_file_access_4), int32(903), int32(_a_F_errcode_for_file_access_5))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
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
func F_errhint_internal(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn14260(m, l0, l1, int32(_a_F_errhint_internal_0), int32(1537))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_errmsg(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn14261(m, l0, l1, int32(_a_F_errmsg_0), int32(1100))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_errorConflictingDefElem(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_errorConflictingDefElem_0), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				F_parser_errposition(m, l1, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_errorConflictingDefElem_1), int32(375), int32(_a_F_errorConflictingDefElem_2))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
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
}
func F_error_severity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = l0 - int32(10)
	if base.Ui32(v3) <= base.Ui32(int32(14)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_error_severity[0])))
		v10 = v8
	} else {
		v10 = int32(_a_F_error_severity_0)
	}
	return v10
}
func F_esc_decode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v88 int64
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v8 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v13 = l0 + l1
	v14 = l0
	v16 = l2
	v21 = v8
	goto L5
L3:
	;
	v88 = v8
	goto L4
L4:
	;
	m.G0 = v11 + int32(16)
	return v88
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v22 != int32(92) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v88 = v77
	goto L4
L7:
	;
	v77 = v21 + int64(1)
	if base.Ui32(v75) < base.Ui32(v13) {
		v14 = v75
		v16 = v16 + int32(1)
		v21 = v77
		goto L5
	} else {
		goto L18
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v22)
	v75 = v14 + int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v29 = v14 + int32(3)
	if base.Ui32(v13) <= base.Ui32(v29) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v63 = v14 + int32(1)
	if base.Ui32(v13) <= base.Ui32(v63) {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v31&int32(252) != int32(48) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
	if v36&int32(248) != int32(48) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v41&int32(248) != int32(48) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v55 = v41 + (v36<<(uint(int32(3))%32)&int32(56) | v31<<(uint(int32(6))%32)) - int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v55)
	v75 = v14 + int32(4)
	goto L7
L16:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v65 != int32(92) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v68 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v68)
	v75 = v14 + int32(2)
	goto L7
L18:
	;
	goto L6
L19:
	;
	return int64(0)
L20:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_esc_decode_0)
	F_errmsg(m, int32(_a_F_esc_decode_1), v11)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_esc_decode_2), int32(755), int32(_a_F_esc_decode_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_esc_enc_len(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	v3 = int64(0)
	if l1 != 0 {
		v6 = l0
		v8 = v3
		for {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v13 == int32(92) {
				v16 = int64(2)
			} else {
				v16 = int64(1)
			}
			if base.I32_extend8_s(v13) <= int32(0) {
				v20 = int64(4)
			} else {
				v20 = v16
			}
			v21 = v8 + v20
			v23 = v6 + int32(1)
			if base.Ui32(v23) < base.Ui32(l0+l1) {
				v6 = v23
				v8 = v21
				continue
			} else {
				break
			}
			break
		}
		v27 = v21
	} else {
		v27 = v3
	}
	return v27
}
func F_estimate_ln_dweight(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v375 int64
	_ = v375
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 float64
	_ = v396
	var v407 int64
	_ = v407
	var v422 int32
	_ = v422
	var v424 int64
	_ = v424
	var v434 int64
	_ = v434
	var v438 int64
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 float64
	_ = v447
	var v449 float64
	_ = v449
	var v462 float64
	_ = v462
	var v465 float64
	_ = v465
	var v470 float64
	_ = v470
	var v471 float64
	_ = v471
	var v472 float64
	_ = v472
	var v473 float64
	_ = v473
	var v478 float64
	_ = v478
	var v479 float64
	_ = v479
	var v480 float64
	_ = v480
	var v505 float64
	_ = v505
	var v517 float64
	_ = v517
	var v539 float64
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 float64
	_ = v565
	var v567 float64
	_ = v567
	var v578 int64
	_ = v578
	var v593 int32
	_ = v593
	var v595 int64
	_ = v595
	var v605 int64
	_ = v605
	var v609 int64
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 float64
	_ = v618
	var v620 float64
	_ = v620
	var v633 float64
	_ = v633
	var v636 float64
	_ = v636
	var v641 float64
	_ = v641
	var v642 float64
	_ = v642
	var v643 float64
	_ = v643
	var v644 float64
	_ = v644
	var v649 float64
	_ = v649
	var v650 float64
	_ = v650
	var v651 float64
	_ = v651
	var v676 float64
	_ = v676
	var v688 float64
	_ = v688
	var v710 float64
	_ = v710
	var v713 int32
	_ = v713
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 != 0 {
		v713 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v713
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(0) {
		v713 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = int32(1)
	v19 = int32(-1)
	v20 = int32(0)
	if base.B2i32(v19 < v16)&base.B2i32(v20 < v12) == v20 {
		v51 = v16
		v55 = v20
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if v12 <= int32(0) {
		v713 = v2
		goto L1
	} else {
		goto L118
	}
L5:
	;
	if v193 < int32(0) {
		goto L4
	} else {
		goto L48
	}
L6:
	;
	v193 = v183
	goto L5
L7:
	;
	if int32(0)|base.B2i32(v19 <= v51) != 0 {
		v87 = v19
		v89 = v20
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v32 = v16
	v36 = v20
	goto L9
L9:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v36<<(uint(int32(1))%32)))))
	if v42 != 0 {
		v183 = int32(1)
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v51 = v46
	v55 = v44
	goto L7
L11:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v32 - v43
	if v46 <= v19 {
		v51 = v46
		v55 = v44
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v44 < v12 {
		v32 = v46
		v36 = v44
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if v51 != v87 {
		v129 = v55
		v130 = v89
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v68 = v19
	v70 = v20
	goto L16
L16:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70<<(uint(int32(1))%32))+uint32(_c_F_estimate_ln_dweight[0]))))
	if v75 != 0 {
		v183 = int32(-1)
		goto L6
	} else {
		goto L18
	}
L17:
	;
	v87 = v79
	v89 = v77
	goto L14
L18:
	;
	v76 = int32(1)
	v77 = v70 + v76
	v79 = v68 - v76
	if v79 <= v51 {
		v87 = v79
		v89 = v77
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if v77 < v18 {
		v68 = v79
		v70 = v77
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	if v12 < v129 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v98 = v55
	v99 = v89
	goto L23
L23:
	;
	if base.B2i32(v12 <= v98)|base.B2i32(v18 <= v99) != 0 {
		v129 = v98
		v130 = v99
		goto L21
	} else {
		goto L25
	}
L24:
	;
	if base.I32_extend16_s(v115) < base.I32_extend16_s(v113) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v104 = int32(1)
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v98<<(uint(v104)%32)))))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99<<(uint(v104)%32))+uint32(_c_F_estimate_ln_dweight[0]))))
	if v113 == v115 {
		v98 = v98 + v104
		v99 = v99 + v104
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v122 = int32(1)
	goto L29
L28:
	;
	v122 = int32(-1)
	goto L29
L29:
	;
	v193 = v122
	goto L5
L30:
	;
	v133 = v129
	goto L32
L31:
	;
	v133 = v12
	goto L32
L32:
	;
	v140 = v129
	goto L33
L33:
	;
	if v133 == v140 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v183 = v166
	goto L6
L35:
	;
	if v18 < v130 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v166 = int32(1)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v140<<(uint(v166)%32)))))
	if v172 == int32(0) {
		v140 = v140 + v166
		goto L33
	} else {
		goto L47
	}
L38:
	;
	v145 = v130
	goto L40
L39:
	;
	v145 = v18
	goto L40
L40:
	;
	v153 = v130
	goto L41
L41:
	;
	if v145 == v153 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v183 = int32(-1)
	goto L6
L43:
	;
	v193 = int32(0)
	goto L5
L44:
	;
	goto L45
L45:
	;
	v157 = int32(1)
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153<<(uint(v157)%32))+uint32(_c_F_estimate_ln_dweight[0]))))
	if v162 == int32(0) {
		v153 = v153 + v157
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	goto L34
L48:
	;
	v197 = int32(2)
	v198 = int32(0)
	if base.B2i32(v198 < v16)&base.B2i32(v198 < v12) == v198 {
		v230 = v16
		v234 = v198
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if int32(0) < v372 {
		goto L4
	} else {
		goto L92
	}
L50:
	;
	v372 = v362
	goto L49
L51:
	;
	if int32(0)|base.B2i32(v198 <= v230) != 0 {
		v266 = v198
		v268 = v198
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v211 = v16
	v215 = v198
	goto L53
L53:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v215<<(uint(int32(1))%32)))))
	if v221 != 0 {
		v362 = int32(1)
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v230 = v225
	v234 = v223
	goto L51
L55:
	;
	v222 = int32(1)
	v223 = v215 + v222
	v225 = v211 - v222
	if v225 <= v198 {
		v230 = v225
		v234 = v223
		goto L51
	} else {
		goto L56
	}
L56:
	;
	if v223 < v12 {
		v211 = v225
		v215 = v223
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	if v230 != v266 {
		v308 = v234
		v309 = v268
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v247 = v198
	v249 = v198
	goto L60
L60:
	;
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249<<(uint(int32(1))%32))+uint32(_c_F_estimate_ln_dweight[1]))))
	if v254 != 0 {
		v362 = int32(-1)
		goto L50
	} else {
		goto L62
	}
L61:
	;
	v266 = v258
	v268 = v256
	goto L58
L62:
	;
	v255 = int32(1)
	v256 = v249 + v255
	v258 = v247 - v255
	if v258 <= v230 {
		v266 = v258
		v268 = v256
		goto L58
	} else {
		goto L63
	}
L63:
	;
	if v256 < v197 {
		v247 = v258
		v249 = v256
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	if v12 < v308 {
		goto L74
	} else {
		goto L75
	}
L66:
	;
	v277 = v234
	v278 = v268
	goto L67
L67:
	;
	if base.B2i32(v12 <= v277)|base.B2i32(v197 <= v278) != 0 {
		v308 = v277
		v309 = v278
		goto L65
	} else {
		goto L69
	}
L68:
	;
	if base.I32_extend16_s(v294) < base.I32_extend16_s(v292) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v283 = int32(1)
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v277<<(uint(v283)%32)))))
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278<<(uint(v283)%32))+uint32(_c_F_estimate_ln_dweight[1]))))
	if v292 == v294 {
		v277 = v277 + v283
		v278 = v278 + v283
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v301 = int32(1)
	goto L73
L72:
	;
	v301 = int32(-1)
	goto L73
L73:
	;
	v372 = v301
	goto L49
L74:
	;
	v312 = v308
	goto L76
L75:
	;
	v312 = v12
	goto L76
L76:
	;
	v319 = v308
	goto L77
L77:
	;
	if v312 == v319 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v362 = v345
	goto L50
L79:
	;
	if v197 < v309 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	v345 = int32(1)
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v319<<(uint(v345)%32)))))
	if v351 == int32(0) {
		v319 = v319 + v345
		goto L77
	} else {
		goto L91
	}
L82:
	;
	v324 = v309
	goto L84
L83:
	;
	v324 = v197
	goto L84
L84:
	;
	v332 = v309
	goto L85
L85:
	;
	if v324 == v332 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v362 = int32(-1)
	goto L50
L87:
	;
	v372 = int32(0)
	goto L49
L88:
	;
	goto L89
L89:
	;
	v336 = int32(1)
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v332<<(uint(v336)%32))+uint32(_c_F_estimate_ln_dweight[1]))))
	if v341 == int32(0) {
		v332 = v332 + v336
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	goto L78
L92:
	;
	v375 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v375
	F_sub_var(m, l0, int32(_a_F_estimate_ln_dweight_0), v9+int32(8))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	return int32(0)
L94:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if int32(0) < v388 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v394))))
	v396 = base.F64_convert_i32_s(v395)
	v407 = base.I64_reinterpret_f64(v396)
	if v407 <= int64(4503599627370495) {
		goto L102
	} else {
		goto L103
	}
L96:
	;
	v542 = v2
	goto L97
L97:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v543 != 0 {
		goto L114
	} else {
		goto L115
	}
L98:
	;
	v542 = v391<<(uint(int32(2))%32) + base.I32_trunc_sat_f64_s(v539)
	goto L97
L99:
	;
	v539 = v517
	goto L98
L100:
	;
	v443 = v441 + int32(_a_F_estimate_ln_dweight_1)
	v447 = base.F64_convert_i32_s(int32(base.Ui32(v443)>>(uint(int32(20))%32)) + v440)
	v449 = base.F64_mul(v447, float64(0.30102999566361177))
	v462 = base.F64_add(base.F64_reinterpret_i64(v438&int64(4294967295)|base.I64_extend_i32_u(v443&int32(_a_F_estimate_ln_dweight_2)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v465 = base.F64_mul(v462, base.F64_mul(v462, float64(0.5)))
	v470 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v462, v465)) & int64(-4294967296))
	v471 = float64(0.4342944818781689)
	v472 = base.F64_mul(v470, v471)
	v473 = base.F64_add(v449, v472)
	v478 = base.F64_div(v462, base.F64_add(v462, float64(2)))
	v479 = base.F64_mul(v478, v478)
	v480 = base.F64_mul(v479, v479)
	v505 = base.F64_add(base.F64_mul(v478, base.F64_add(v465, base.F64_add(base.F64_mul(v480, base.F64_add(base.F64_mul(v480, base.F64_add(base.F64_mul(v480, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v479, base.F64_add(base.F64_mul(v480, base.F64_add(base.F64_mul(v480, base.F64_add(base.F64_mul(v480, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v462, v470), v465))
	v517 = base.F64_add(v473, base.F64_add(base.F64_add(v472, base.F64_sub(v449, v473)), base.F64_add(base.F64_mul(v505, v471), base.F64_add(base.F64_mul(v447, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v505, v470), float64(2.5082946711645275e-11))))))
	goto L99
L101:
	;
	v434 = base.I64_reinterpret_f64(base.F64_mul(v396, float64(1.8014398509481984e+16)))
	v438 = v434
	v440 = int32(-1077)
	v441 = base.I32_wrap_i64(int64(base.Ui64(v434) >> (uint(int64(32)) % 64)))
	goto L100
L102:
	;
	if base.F64_eq(v396, float64(0)) != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v407) {
		v517 = v396
		goto L99
	} else {
		goto L109
	}
L105:
	;
	v539 = base.F64_div(float64(-1), base.F64_mul(v396, v396))
	goto L98
L106:
	;
	goto L107
L107:
	;
	if int64(0) <= v407 {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	v539 = base.F64_div(base.F64_sub(v396, v396), float64(0))
	goto L98
L109:
	;
	v422 = int32(-1023)
	v424 = int64(base.Ui64(v407) >> (uint(int64(32)) % 64))
	if v424 != int64(1072693248) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v438 = v407
	v440 = v422
	v441 = base.I32_wrap_i64(v424)
	goto L100
L111:
	;
	goto L112
L112:
	;
	if base.I32_wrap_i64(v407) != 0 {
		v438 = v407
		v440 = v422
		v441 = int32(1072693248)
		goto L100
	} else {
		goto L113
	}
L113:
	;
	v539 = float64(0)
	goto L98
L114:
	;
	F_pfree(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L93
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v713 = v542
	goto L1
L117:
	;
	goto L116
L118:
	;
	v549 = v16 << (uint(int32(2)) % 32)
	v550 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15))))
	if v12 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v553 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+2)))
	v559 = v553 + v550*int32(_a_F_estimate_ln_dweight_3)
	v560 = v549 - int32(4)
	goto L121
L120:
	;
	v559 = v550
	v560 = v549
	goto L121
L121:
	;
	v565 = F_log(m, base.F64_convert_i32_s(v559))
	mBase = m.M
	v567 = base.F64_abs(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v560), float64(2.302585092994046)), v565))
	v578 = base.I64_reinterpret_f64(v567)
	if v578 <= int64(4503599627370495) {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v713 = base.I32_trunc_sat_f64_s(v710)
	goto L1
L123:
	;
	v710 = v688
	goto L122
L124:
	;
	v614 = v612 + int32(_a_F_estimate_ln_dweight_1)
	v618 = base.F64_convert_i32_s(int32(base.Ui32(v614)>>(uint(int32(20))%32)) + v611)
	v620 = base.F64_mul(v618, float64(0.30102999566361177))
	v633 = base.F64_add(base.F64_reinterpret_i64(v609&int64(4294967295)|base.I64_extend_i32_u(v614&int32(_a_F_estimate_ln_dweight_2)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v636 = base.F64_mul(v633, base.F64_mul(v633, float64(0.5)))
	v641 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v633, v636)) & int64(-4294967296))
	v642 = float64(0.4342944818781689)
	v643 = base.F64_mul(v641, v642)
	v644 = base.F64_add(v620, v643)
	v649 = base.F64_div(v633, base.F64_add(v633, float64(2)))
	v650 = base.F64_mul(v649, v649)
	v651 = base.F64_mul(v650, v650)
	v676 = base.F64_add(base.F64_mul(v649, base.F64_add(v636, base.F64_add(base.F64_mul(v651, base.F64_add(base.F64_mul(v651, base.F64_add(base.F64_mul(v651, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v650, base.F64_add(base.F64_mul(v651, base.F64_add(base.F64_mul(v651, base.F64_add(base.F64_mul(v651, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v633, v641), v636))
	v688 = base.F64_add(v644, base.F64_add(base.F64_add(v643, base.F64_sub(v620, v644)), base.F64_add(base.F64_mul(v676, v642), base.F64_add(base.F64_mul(v618, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v676, v641), float64(2.5082946711645275e-11))))))
	goto L123
L125:
	;
	v605 = base.I64_reinterpret_f64(base.F64_mul(v567, float64(1.8014398509481984e+16)))
	v609 = v605
	v611 = int32(-1077)
	v612 = base.I32_wrap_i64(int64(base.Ui64(v605) >> (uint(int64(32)) % 64)))
	goto L124
L126:
	;
	if base.F64_eq(v567, float64(0)) != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v578) {
		v688 = v567
		goto L123
	} else {
		goto L133
	}
L129:
	;
	v710 = base.F64_div(float64(-1), base.F64_mul(v567, v567))
	goto L122
L130:
	;
	goto L131
L131:
	;
	if int64(0) <= v578 {
		goto L125
	} else {
		goto L132
	}
L132:
	;
	v710 = base.F64_div(base.F64_sub(v567, v567), float64(0))
	goto L122
L133:
	;
	v593 = int32(-1023)
	v595 = int64(base.Ui64(v578) >> (uint(int64(32)) % 64))
	if v595 != int64(1072693248) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v609 = v578
	v611 = v593
	v612 = base.I32_wrap_i64(v595)
	goto L124
L135:
	;
	goto L136
L136:
	;
	if base.I32_wrap_i64(v578) != 0 {
		v609 = v578
		v611 = v593
		v612 = int32(1072693248)
		goto L124
	} else {
		goto L137
	}
L137:
	;
	v710 = float64(0)
	goto L122
}
func F_examine_opclause_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 == int32(27) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v16 = v15
	} else {
		v16 = v11
	}
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v17 == int32(27) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v22 = v20
		v23 = v21
	} else {
		v22 = v10
		v23 = v17
	}
	if v23 == int32(7) {
		v29 = v22
		v30 = v16
		if l1 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
		} else {
		}
		if l2 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
		} else {
		}
		v33 = int32(1)
		if l3 == int32(0) {
			v41 = v33
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(base.B2i32(v23 == int32(7)))
			v41 = v33
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v26 != int32(7) {
			v41 = int32(0)
		} else {
			v29 = v16
			v30 = v22
			if l1 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
			} else {
			}
			if l2 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
			} else {
			}
			v33 = int32(1)
			if l3 == int32(0) {
				v41 = v33
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(base.B2i32(v23 == int32(7)))
				v41 = v33
			}
		}
	}
	return v41
}
func F_executeComparison(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
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
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v367 int64
	_ = v367
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v397 int64
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v575 int32
	_ = v575
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int64
	_ = v628
	var v629 int64
	_ = v629
	var v631 int64
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	v17 = m.G0
	v19 = v17 - int32(224)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v22 != v23 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errhint(m, int32(_a_F_executeComparison_0), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L23
	} else {
		goto L252
	}
L2:
	;
	m.G0 = v19 + int32(224)
	return v684
L3:
	;
	v25 = int32(0)
	if base.B2i32(v22 == v25)|base.B2i32(v23 == v25) == v25 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v35 = int32(2)
	switch v22 {
	case 0:
		v637 = v22
		goto L9
	case 1:
		goto L21
	case 2:
		goto L22
	case 3:
		goto L18
	default:
		goto L19
	case 16, 17, 18:
		v684 = v35
		goto L2
	case 32:
		goto L20
	}
L6:
	;
	v684 = int32(2)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v684 = base.B2i32(v21 == int32(9))
	goto L2
L9:
	;
	switch v21 - int32(8) {
	case 0:
		goto L242
	case 1:
		goto L248
	case 2:
		goto L247
	case 3:
		goto L246
	case 4:
		goto L245
	case 5:
		goto L244
	default:
		goto L243
	}
L10:
	;
	v635 = F_date_cmp_timestamp_internal(m, base.I32_wrap_i64(v302), v301)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L23
	} else {
		goto L241
	}
L11:
	;
	v631 = F_DirectFunctionCall2Coll(m, v627, int32(0), v629, v628)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L23
	} else {
		goto L240
	}
L12:
	;
	if v300 <= int32(1183) {
		goto L222
	} else {
		goto L223
	}
L13:
	;
	if v299&int32(1) != 0 {
		goto L212
	} else {
		goto L213
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L23
	} else {
		goto L209
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L23
	} else {
		goto L206
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L23
	} else {
		goto L203
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L23
	} else {
		goto L200
	}
L18:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v484 != 0 {
		goto L194
	} else {
		goto L195
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L23
	} else {
		goto L191
	}
L20:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+35)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v301 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v302 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v303 <= int32(1183) {
		goto L129
	} else {
		goto L130
	}
L21:
	;
	if v21 == int32(8) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+8)))
	v39 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)))
	v40 = F_DirectFunctionCall2Coll(m, int32(1467), int32(0), v38, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v637 = base.I32_wrap_i64(v40)
	goto L9
L25:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v47 != v48 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_executeComparison[0]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	goto L50
L28:
	;
	v684 = int32(0)
	goto L2
L29:
	;
	goto L30
L30:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v47) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v684 = base.B2i32(v114 == int32(0))
	goto L2
L32:
	;
	v114 = int32(0)
	goto L31
L33:
	;
	v88 = v83
	v89 = v84
	v90 = v85
	goto L43
L34:
	;
	if (v51|v52)&int32(3) != 0 {
		v83 = v51
		v84 = v52
		v85 = v47
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v76 = v51
	v77 = v52
	v78 = v47
	goto L36
L36:
	;
	if v78 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v60 = v51
	v61 = v52
	v62 = v47
	goto L38
L38:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v65 != v66 {
		v83 = v60
		v84 = v61
		v85 = v62
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v76 = v71
	v77 = v69
	v78 = v73
	goto L36
L40:
	;
	v68 = int32(4)
	v69 = v61 + v68
	v71 = v60 + v68
	v73 = v62 - v68
	if base.Ui32(int32(3)) < base.Ui32(v73) {
		v60 = v71
		v61 = v69
		v62 = v73
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v83 = v76
	v84 = v77
	v85 = v78
	goto L33
L43:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v93 == v94 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v114 = v93 - v94
	goto L31
L45:
	;
	v96 = int32(1)
	v101 = v90 - v96
	if v101 != 0 {
		v88 = v88 + v96
		v89 = v89 + v96
		v90 = v101
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	v232 = base.B2i32(v119 < v117)
	if v119 < v117 {
		goto L96
	} else {
		goto L97
	}
L50:
	;
	if v123 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_executeComparison[0]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	goto L52
L52:
	;
	if v128 == int32(6) {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v132 = F_pg_server_to_any(m, v120, v119, int32(6))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	v135 = F_pg_server_to_any(m, v118, v117, int32(6))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v137 = base.B2i32(v132 == v120)
	if v137 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v140 = F_strlen(m, v132)
	mBase = m.M
	v141 = v140
	goto L58
L57:
	;
	v141 = v119
	goto L58
L58:
	;
	v142 = base.B2i32(v135 == v118)
	if v142 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v145 = F_strlen(m, v135)
	mBase = m.M
	v146 = v145
	goto L61
L60:
	;
	v146 = v117
	goto L61
L61:
	;
	v147 = base.B2i32(v141 < v146)
	if v141 < v146 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v148 = v141
	goto L64
L63:
	;
	v148 = v146
	goto L64
L64:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v148) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	if v210 != 0 {
		goto L83
	} else {
		goto L84
	}
L66:
	;
	v210 = int32(0)
	goto L65
L67:
	;
	v184 = v179
	v185 = v180
	v186 = v181
	goto L77
L68:
	;
	if (v132|v135)&int32(3) != 0 {
		v179 = v132
		v180 = v135
		v181 = v148
		goto L67
	} else {
		goto L71
	}
L69:
	;
	v172 = v132
	v173 = v135
	v174 = v148
	goto L70
L70:
	;
	if v174 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v156 = v132
	v157 = v135
	v158 = v148
	goto L72
L72:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v161 != v162 {
		v179 = v156
		v180 = v157
		v181 = v158
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v172 = v167
	v173 = v165
	v174 = v169
	goto L70
L74:
	;
	v164 = int32(4)
	v165 = v157 + v164
	v167 = v156 + v164
	v169 = v158 - v164
	if base.Ui32(int32(3)) < base.Ui32(v169) {
		v156 = v167
		v157 = v165
		v158 = v169
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v179 = v172
	v180 = v173
	v181 = v174
	goto L67
L77:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if v189 == v190 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v210 = v189 - v190
	goto L65
L79:
	;
	v192 = int32(1)
	v197 = v186 - v192
	if v197 != 0 {
		v184 = v184 + v192
		v185 = v185 + v192
		v186 = v197
		goto L77
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	goto L66
L83:
	;
	v213 = v210
	goto L85
L84:
	;
	v213 = base.B2i32(v146 < v141) - v147
	goto L85
L85:
	;
	if v142&base.B2i32(v132 == v120) != 0 {
		v637 = v213
		goto L9
	} else {
		goto L86
	}
L86:
	;
	if v137 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_pfree(m, v132)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L23
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v142 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	F_pfree(m, v135)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L23
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if v213 != 0 {
		v637 = v213
		goto L9
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	goto L49
L96:
	;
	v233 = v119
	goto L98
L97:
	;
	v233 = v117
	goto L98
L98:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v233) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	if v295 != 0 {
		goto L117
	} else {
		goto L118
	}
L100:
	;
	v295 = int32(0)
	goto L99
L101:
	;
	v269 = v264
	v270 = v265
	v271 = v266
	goto L111
L102:
	;
	if (v120|v118)&int32(3) != 0 {
		v264 = v120
		v265 = v118
		v266 = v233
		goto L101
	} else {
		goto L105
	}
L103:
	;
	v257 = v120
	v258 = v118
	v259 = v233
	goto L104
L104:
	;
	if v259 == int32(0) {
		goto L100
	} else {
		goto L110
	}
L105:
	;
	v241 = v120
	v242 = v118
	v243 = v233
	goto L106
L106:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	if v246 != v247 {
		v264 = v241
		v265 = v242
		v266 = v243
		goto L101
	} else {
		goto L108
	}
L107:
	;
	v257 = v252
	v258 = v250
	v259 = v254
	goto L104
L108:
	;
	v249 = int32(4)
	v250 = v242 + v249
	v252 = v241 + v249
	v254 = v243 - v249
	if base.Ui32(int32(3)) < base.Ui32(v254) {
		v241 = v252
		v242 = v250
		v243 = v254
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v264 = v257
	v265 = v258
	v266 = v259
	goto L101
L111:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v274 == v275 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v295 = v274 - v275
	goto L99
L113:
	;
	v277 = int32(1)
	v282 = v271 - v277
	if v282 != 0 {
		v269 = v269 + v277
		v270 = v270 + v277
		v271 = v282
		goto L111
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	goto L112
L116:
	;
	goto L100
L117:
	;
	v298 = v295
	goto L119
L118:
	;
	v298 = base.B2i32(v117 < v119) - v232
	goto L119
L119:
	;
	v637 = v298
	goto L9
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L23
	} else {
		goto L188
	}
L121:
	;
	if v303 == int32(1114) {
		goto L12
	} else {
		goto L187
	}
L122:
	;
	if v300 <= int32(1183) {
		goto L174
	} else {
		goto L175
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L23
	} else {
		goto L168
	}
L124:
	;
	if v300 == int32(1114) {
		v684 = v35
		goto L2
	} else {
		goto L167
	}
L125:
	;
	if v299&int32(1) == int32(0) {
		goto L15
	} else {
		goto L165
	}
L126:
	;
	if v300 == int32(1184) {
		v684 = v35
		goto L2
	} else {
		goto L163
	}
L127:
	;
	if v300 <= int32(1183) {
		goto L152
	} else {
		goto L153
	}
L128:
	;
	if v300 <= int32(1183) {
		goto L137
	} else {
		goto L138
	}
L129:
	;
	switch v303 - int32(1082) {
	case 0:
		goto L128
	case 1:
		goto L127
	default:
		goto L121
	}
L130:
	;
	goto L131
L131:
	;
	if v303 == int32(1184) {
		goto L122
	} else {
		goto L132
	}
L132:
	;
	if v303 != int32(1266) {
		goto L120
	} else {
		goto L133
	}
L133:
	;
	v312 = int32(1575)
	if int32(1183) < v300 {
		goto L126
	} else {
		goto L134
	}
L134:
	;
	switch v300 - int32(1082) {
	case 0:
		v684 = v35
		goto L2
	case 1:
		goto L125
	default:
		goto L124
	}
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L23
	} else {
		goto L147
	}
L136:
	;
	if v300 == int32(1114) {
		goto L10
	} else {
		goto L146
	}
L137:
	;
	switch v300 - int32(1082) {
	case 0:
		v627 = int32(1576)
		v628 = v301
		v629 = v302
		goto L11
	case 1:
		v684 = v35
		goto L2
	default:
		goto L136
	}
L138:
	;
	goto L139
L139:
	;
	if v300 != int32(1184) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v300 != int32(1266) {
		goto L135
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if v299&int32(1) == int32(0) {
		goto L17
	} else {
		goto L144
	}
L143:
	;
	v684 = v35
	goto L2
L144:
	;
	v331 = F_date_cmp_timestamptz_internal(m, base.I32_wrap_i64(v302), v301)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L23
	} else {
		goto L145
	}
L145:
	;
	v637 = v331
	goto L9
L146:
	;
	goto L135
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v300
	F_errmsg_internal(m, int32(_a_F_executeComparison_1), v19+int32(48))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L23
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(4076), int32(_a_F_executeComparison_3))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L23
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L23
	} else {
		goto L160
	}
L151:
	;
	if v300 == int32(1114) {
		v684 = v35
		goto L2
	} else {
		goto L159
	}
L152:
	;
	switch v300 - int32(1082) {
	case 0:
		v684 = v35
		goto L2
	case 1:
		v627 = int32(1577)
		v628 = v301
		v629 = v302
		goto L11
	default:
		goto L151
	}
L153:
	;
	goto L154
L154:
	;
	if v300 == int32(1184) {
		v684 = v35
		goto L2
	} else {
		goto L155
	}
L155:
	;
	if v300 != int32(1266) {
		goto L150
	} else {
		goto L156
	}
L156:
	;
	if v299&int32(1) == int32(0) {
		goto L16
	} else {
		goto L157
	}
L157:
	;
	v367 = F_DirectFunctionCall1Coll(m, int32(1564), int32(0), v302)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L23
	} else {
		goto L158
	}
L158:
	;
	v627 = int32(1575)
	v628 = v301
	v629 = v367
	goto L11
L159:
	;
	goto L150
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v300
	F_errmsg_internal(m, int32(_a_F_executeComparison_1), v19+int32(80))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L23
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(_a_F_executeComparison_4), int32(_a_F_executeComparison_3))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L23
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
	if v300 != int32(1266) {
		goto L123
	} else {
		goto L164
	}
L164:
	;
	v627 = v312
	v628 = v301
	v629 = v302
	goto L11
L165:
	;
	v397 = F_DirectFunctionCall1Coll(m, int32(1564), int32(0), v301)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L23
	} else {
		goto L166
	}
L166:
	;
	v627 = v312
	v628 = v397
	v629 = v302
	goto L11
L167:
	;
	goto L123
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v300
	F_errmsg_internal(m, int32(_a_F_executeComparison_1), v19+int32(112))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L23
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(_a_F_executeComparison_5), int32(_a_F_executeComparison_3))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L23
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L23
	} else {
		goto L184
	}
L172:
	;
	if v300 == int32(1114) {
		goto L13
	} else {
		goto L183
	}
L173:
	;
	if v299&int32(1) == int32(0) {
		goto L14
	} else {
		goto L181
	}
L174:
	;
	switch v300 - int32(1082) {
	case 0:
		goto L173
	case 1:
		v684 = v35
		goto L2
	default:
		goto L172
	}
L175:
	;
	goto L176
L176:
	;
	if v300 == int32(1184) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v627 = int32(1578)
	v628 = v301
	v629 = v302
	goto L11
L178:
	;
	goto L179
L179:
	;
	if v300 != int32(1266) {
		goto L171
	} else {
		goto L180
	}
L180:
	;
	v684 = v35
	goto L2
L181:
	;
	v431 = F_date_cmp_timestamptz_internal(m, base.I32_wrap_i64(v301), v302)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L23
	} else {
		goto L182
	}
L182:
	;
	v637 = int32(0) - v431
	goto L9
L183:
	;
	goto L171
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v300
	F_errmsg_internal(m, int32(_a_F_executeComparison_1), v19+int32(176))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L23
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(_a_F_executeComparison_6), int32(_a_F_executeComparison_3))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L23
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	goto L120
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v303
	F_errmsg_internal(m, int32(_a_F_executeComparison_1), v19+int32(32))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L23
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(_a_F_executeComparison_7), int32(_a_F_executeComparison_3))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L23
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
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v472
	F_errmsg_internal(m, int32(_a_F_executeComparison_8), v19)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L23
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(3673), int32(_a_F_executeComparison_9))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L23
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
	v485 = int32(1)
	goto L196
L195:
	;
	v485 = int32(-1)
	goto L196
L196:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v484 != v487 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v489 = v485
	goto L199
L198:
	;
	v489 = int32(0)
	goto L199
L199:
	;
	v637 = v489
	goto L9
L200:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L23
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = int32(_a_F_executeComparison_10)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = int32(_a_F_executeComparison_11)
	F_errmsg(m, int32(_a_F_executeComparison_12), v19-int32(-64))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L23
	} else {
		goto L202
	}
L202:
	;
	goto L1
L203:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L23
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+100)) = int32(_a_F_executeComparison_13)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = int32(_a_F_executeComparison_14)
	F_errmsg(m, int32(_a_F_executeComparison_12), v19+int32(96))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L23
	} else {
		goto L205
	}
L205:
	;
	goto L1
L206:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L23
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = int32(_a_F_executeComparison_13)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = int32(_a_F_executeComparison_14)
	F_errmsg(m, int32(_a_F_executeComparison_12), v19+int32(128))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L23
	} else {
		goto L208
	}
L208:
	;
	goto L1
L209:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L23
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+196)) = int32(_a_F_executeComparison_10)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = int32(_a_F_executeComparison_11)
	F_errmsg(m, int32(_a_F_executeComparison_12), v19+int32(192))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L23
	} else {
		goto L211
	}
L211:
	;
	goto L1
L212:
	;
	v557 = F_timestamp_cmp_timestamptz_internal(m, v301, v302)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L23
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L23
	} else {
		goto L216
	}
L215:
	;
	v637 = int32(0) - v557
	goto L9
L216:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L23
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+212)) = int32(_a_F_executeComparison_10)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+208)) = int32(_a_F_executeComparison_15)
	F_errmsg(m, int32(_a_F_executeComparison_12), v19+int32(208))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L23
	} else {
		goto L218
	}
L218:
	;
	goto L1
L219:
	;
	if v299&int32(1) != 0 {
		goto L233
	} else {
		goto L234
	}
L220:
	;
	v604 = F_date_cmp_timestamp_internal(m, base.I32_wrap_i64(v301), v302)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L23
	} else {
		goto L232
	}
L221:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L23
	} else {
		goto L229
	}
L222:
	;
	switch v300 - int32(1082) {
	case 0:
		goto L220
	case 1:
		v684 = v35
		goto L2
	default:
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	if v300 == int32(1184) {
		goto L219
	} else {
		goto L227
	}
L225:
	;
	if v300 != int32(1114) {
		goto L221
	} else {
		goto L226
	}
L226:
	;
	v627 = int32(1578)
	v628 = v301
	v629 = v302
	goto L11
L227:
	;
	if v300 == int32(1266) {
		v684 = v35
		goto L2
	} else {
		goto L228
	}
L228:
	;
	goto L221
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v300
	F_errmsg_internal(m, int32(_a_F_executeComparison_1), v19+int32(144))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L23
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(_a_F_executeComparison_16), int32(_a_F_executeComparison_3))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L23
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	v637 = int32(0) - v604
	goto L9
L233:
	;
	v609 = F_timestamp_cmp_timestamptz_internal(m, v302, v301)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L23
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L23
	} else {
		goto L237
	}
L236:
	;
	v637 = v609
	goto L9
L237:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L23
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+164)) = int32(_a_F_executeComparison_10)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = int32(_a_F_executeComparison_15)
	F_errmsg(m, int32(_a_F_executeComparison_12), v19+int32(160))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L23
	} else {
		goto L239
	}
L239:
	;
	goto L1
L240:
	;
	v637 = base.I32_wrap_i64(v631)
	goto L9
L241:
	;
	v637 = v635
	goto L9
L242:
	;
	v684 = base.B2i32(v637 == int32(0))
	goto L2
L243:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L23
	} else {
		goto L249
	}
L244:
	;
	v684 = base.B2i32(int32(0) <= v637)
	goto L2
L245:
	;
	v684 = base.B2i32(v637 <= int32(0))
	goto L2
L246:
	;
	v684 = base.B2i32(int32(0) < v637)
	goto L2
L247:
	;
	v684 = int32(base.Ui32(v637) >> (uint(int32(31)) % 32))
	goto L2
L248:
	;
	v684 = base.B2i32(v637 != int32(0))
	goto L2
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v21
	F_errmsg_internal(m, int32(_a_F_executeComparison_17), v19+int32(16))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L23
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(3697), int32(_a_F_executeComparison_9))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L23
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errfinish(m, int32(_a_F_executeComparison_2), int32(3992), int32(_a_F_executeComparison_18))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L23
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_expandNSItemVars(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v221 int32
	_ = v221
	v6 = int32(0)
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if int32(0) < v24 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = v6
	v39 = v6
	goto L10
L8:
	;
	v221 = v6
	goto L9
L9:
	;
	return v221
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v45 = v42 + v36<<(uint(int32(5))%32)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+30)))
	if v46 != 0 {
		v202 = v39
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v221 = v202
	goto L9
L12:
	;
	v206 = v36 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v206 < v207 {
		v36 = v206
		v39 = v202
		goto L10
	} else {
		goto L38
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v36<<(uint(int32(2))%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 == int32(0) {
		v202 = v39
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+4)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v61 = F_makeVar(m, v56, v57, v58, v59, v60, l2)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v67
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+28)))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+40)) = uint16(v69)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	if v73 == int32(0) {
		v151 = l0
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v72 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v77 = v73 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v73) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v88 = l0
	v89 = int32(0)
	goto L22
L20:
	;
	v116 = l0
	goto L21
L21:
	;
	v132 = v116
	v133 = int32(0)
	goto L26
L22:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v107 = v89 + int32(8)
	if v107 != v73&int32(-8) {
		v88 = v105
		v89 = v107
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v77 == int32(0) {
		v151 = v105
		goto L17
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v116 = v105
	goto L21
L26:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v144 = v133 + int32(1)
	if v144 != v77 {
		v132 = v142
		v133 = v144
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v151 = v142
	goto L17
L28:
	;
	goto L27
L29:
	;
	v182 = F_lappend(m, v39, v61)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L15
	} else {
		goto L35
	}
L30:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	if v163 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v166 < v72 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168+v72<<(uint(int32(2))%32)-int32(4))))
	if v174 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	v178 = F_bms_union(m, v177, v174)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v178
	goto L29
L35:
	;
	if l4 == int32(0) {
		v202 = v182
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v187 = F_lappend(m, v186, v51)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v187
	v202 = v182
	goto L12
L38:
	;
	goto L11
}
func F_expand_generated_columns_in_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	if v7 == int32(0) {
		v41 = l0
		return v41
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+18)))
		if v10 != int32(1) {
			v41 = l0
			return v41
		} else {
			v14 = F_palloc0(m, int32(136))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(101)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				v24 = F_makeAlias(m, v20+int32(4), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v24
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v29
					v32 = F_get_generated_columns(m, l1, l2, v26)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if v32 == int32(0) {
							v41 = l0
							return v41
						} else {
							v36 = int32(0)
							v39 = F_ReplaceVarsFromTargetList(m, l0, l2, v14, v32, v36, int32(1), l2, v36)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v39
								return v41
							}
						}
					}
				}
			}
		}
	}
}
