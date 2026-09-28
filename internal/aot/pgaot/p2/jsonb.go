package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonbIteratorInit(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_iteratorFromContainer(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_JsonbToJsonbValue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	v5 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l0 + v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(base.Ui32(v8)>>(uint(int32(2))%32)) - v5
	return
}
func F_add_jsonb(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l3 != 0 {
		if l1 != 0 {
			v12 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v12
			v26 = int32(0)
			v27 = v12
			F_datum_to_jsonb_internal(m, l0, l1, l2, v26, v27, l4)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		} else {
			F_json_categorize_type(m, l3, int32(1), v10+int32(12), v10+int32(8))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				v26 = v24
				v27 = v25
				F_datum_to_jsonb_internal(m, l0, l1, l2, v26, v27, l4)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_add_jsonb_0), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_add_jsonb_1), int32(1061), int32(_a_F_add_jsonb_2))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_fillJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v125 int32
	_ = v125
	v13 = l0 + l1<<(uint(int32(2))%32) + int32(4)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	switch int32(base.Ui32(v14)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L5
	case 1:
		goto L4
	case 2:
		goto L2
	case 3:
		goto L3
	case 4:
		goto L6
	default:
		goto L1
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(18)
	v81 = (l3 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = l2 + v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v84 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)) = uint8(v72)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(3)
	return
L3:
	;
	v68 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+8)) = uint8(v68)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(3)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = l2 + (l3+int32(3))&int32(-4)
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = l2 + l3
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v25 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	return
L7:
	;
	v30 = l1
	v31 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v25 & int32(268435455)
	return
L10:
	;
	v38 = v30 - int32(1)
	if int32(0) <= v38 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v25&int32(268435455) - v51
	return
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0+v30<<(uint(int32(2))%32))))
	v47 = v44&int32(268435455) + v31
	if int32(0) <= v44 {
		v30 = v38
		v31 = v47
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v51 = v31
	goto L14
L14:
	;
	goto L11
L15:
	;
	v51 = v47
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v125 + (l3 - v81)
	return
L17:
	;
	v89 = l1
	v90 = int32(0)
	goto L20
L18:
	;
	goto L19
L19:
	;
	v125 = v84 & int32(268435455)
	goto L16
L20:
	;
	v97 = v89 - int32(1)
	if int32(0) <= v97 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v125 = v84&int32(268435455) - v110
	goto L16
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0+v89<<(uint(int32(2))%32))))
	v106 = v103&int32(268435455) + v90
	if int32(0) <= v103 {
		v89 = v97
		v90 = v106
		goto L20
	} else {
		goto L25
	}
L23:
	;
	v110 = v90
	goto L24
L24:
	;
	goto L21
L25:
	;
	v110 = v106
	goto L24
}
func F_jsonb_agg_transfn_worker(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = v9 + int32(12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == v3 {
		v31 = int32(0)
		if v12 == v31 {
			v39 = v31
		} else {
			v34 = v31
			v35 = v3
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
			v39 = v35
		}
		v42 = v39
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		switch v17 - int32(435) {
		case 0:
			if v12 == int32(0) {
				v42 = int32(1)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+168))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
				v34 = v24
				v35 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
				v42 = v39
			}
		case 1:
			if v12 == int32(0) {
				v42 = int32(2)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+376))
				v34 = v29
				v35 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
				v42 = v39
			}
		default:
			v31 = int32(0)
			if v12 == v31 {
				v39 = v31
			} else {
				v34 = v31
				v35 = v3
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
			}
			v42 = v39
		}
	}
	if v42 != 0 {
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v43 == int32(1) {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v48 = F_get_fn_expr_argtype(m, v46, int32(1))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int64(0)
			} else {
				if v48 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_jsonb_agg_transfn_worker_0), int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_jsonb_agg_transfn_worker_1), int32(1497), int32(_a_F_jsonb_agg_transfn_worker_2))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v56 = F_MemoryContextAllocZero(m, v54, int32(36))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v58
						F_pushJsonbValue(m, v56, int32(4), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int64(0)
						} else {
							F_json_categorize_type(m, v48, int32(1), v56+int32(28), v56+int32(32))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								v72 = v56
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
								if l1 != 0 {
									if v74&int32(1) == int32(0) {
										v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
										v84 = v82
										v85 = int32(0)
										v86 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
										F_datum_to_jsonb_internal(m, v84, v85, v72, v86, v87, int32(0))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int64(0)
										} else {
											m.G0 = v9 + int32(16)
											return base.I64_extend_i32_u(v72)
										}
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_u(v72)
									}
								} else {
									v79 = int32(1)
									if v74&v79 != 0 {
										v84 = int64(0)
										v85 = v79
									} else {
										v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
										v84 = v82
										v85 = int32(0)
									}
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
									F_datum_to_jsonb_internal(m, v84, v85, v72, v86, v87, int32(0))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int64(0)
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_u(v72)
									}
								}
							}
						}
					}
				}
			}
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v72 = v71
			v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if l1 != 0 {
				if v74&int32(1) == int32(0) {
					v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v84 = v82
					v85 = int32(0)
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
					F_datum_to_jsonb_internal(m, v84, v85, v72, v86, v87, int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v72)
					}
				} else {
					m.G0 = v9 + int32(16)
					return base.I64_extend_i32_u(v72)
				}
			} else {
				v79 = int32(1)
				if v74&v79 != 0 {
					v84 = int64(0)
					v85 = v79
				} else {
					v82 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v84 = v82
					v85 = int32(0)
				}
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
				F_datum_to_jsonb_internal(m, v84, v85, v72, v86, v87, int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int64(0)
				} else {
					m.G0 = v9 + int32(16)
					return base.I64_extend_i32_u(v72)
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v101 = m.ExcPending
		if v101 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_jsonb_agg_transfn_worker_3), int32(0))
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_jsonb_agg_transfn_worker_1), int32(1485), int32(_a_F_jsonb_agg_transfn_worker_2))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_jsonb_contains(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if (v17^v18)&int32(536870912) == int32(0) {
				v26 = F_JsonbIteratorInit(m, v10+int32(4))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v26
					v31 = F_JsonbIteratorInit(m, v15+int32(4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v31
						v38 = F_JsonbDeepContains(m, v7+int32(12), v7+int32(8))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v41 = base.I64_extend_i32_u(v38)
							m.G0 = v7 + int32(16)
							return v41
						}
					}
				}
			} else {
				v41 = int64(0)
				m.G0 = v7 + int32(16)
				return v41
			}
		}
	}
}
func F_jsonb_delete(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v20 = int32(1)
	v21 = v19 & v20
	if v19 == v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v49
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v51
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v55&int32(268435456) == v49 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v27 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v38 = int32(1)
	if v21 != 0 {
		v48 = int32(base.Ui32(v19)>>(uint(v38)%32)) - v38
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v30 = int32(16)
	goto L10
L9:
	;
	v30 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = int32(4)
	goto L13
L12:
	;
	v37 = v30
	goto L13
L13:
	;
	v48 = v37
	goto L4
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v55&int32(268435455) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L64
	}
L18:
	;
	v64 = F_JsonbIteratorInit(m, v12+int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v186 = v12
	goto L20
L20:
	;
	m.G0 = v9 - int32(-64)
	return base.I64_extend_i32_u(v186)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v64
	v70 = F_JsonbIteratorNext(m, v7+int32(-28), v9, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v70 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v21 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v181 = int32(0)
	goto L25
L25:
	;
	v182 = F_JsonbValueToJsonb(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L63
	}
L26:
	;
	v74 = int32(1)
	goto L28
L27:
	;
	v74 = int32(4)
	goto L28
L28:
	;
	v75 = v17 + v74
	v78 = v70
	goto L29
L29:
	;
	v83 = base.B2i32(v78 != int32(1))
	if v83&base.B2i32(v78 != int32(3)) != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	v181 = v173
	goto L25
L31:
	;
	v171 = F_JsonbIteratorNext(m, v7+int32(-28), v9, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L61
	}
L32:
	;
	if base.Ui32(v78) < base.Ui32(int32(4)) {
		goto L57
	} else {
		goto L58
	}
L33:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v87 != int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v48 != v90 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v48) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	if v154 != 0 {
		goto L32
	} else {
		goto L54
	}
L37:
	;
	v154 = int32(0)
	goto L36
L38:
	;
	v128 = v123
	v129 = v124
	v130 = v125
	goto L48
L39:
	;
	if (v75|v92)&int32(3) != 0 {
		v123 = v75
		v124 = v92
		v125 = v48
		goto L38
	} else {
		goto L42
	}
L40:
	;
	v116 = v75
	v117 = v92
	v118 = v48
	goto L41
L41:
	;
	if v118 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L42:
	;
	v100 = v75
	v101 = v92
	v102 = v48
	goto L43
L43:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v105 != v106 {
		v123 = v100
		v124 = v101
		v125 = v102
		goto L38
	} else {
		goto L45
	}
L44:
	;
	v116 = v111
	v117 = v109
	v118 = v113
	goto L41
L45:
	;
	v108 = int32(4)
	v109 = v101 + v108
	v111 = v100 + v108
	v113 = v102 - v108
	if base.Ui32(int32(3)) < base.Ui32(v113) {
		v100 = v111
		v101 = v109
		v102 = v113
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v123 = v116
	v124 = v117
	v125 = v118
	goto L38
L48:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	if v133 == v134 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v154 = v133 - v134
	goto L36
L50:
	;
	v136 = int32(1)
	v141 = v130 - v136
	if v141 != 0 {
		v128 = v128 + v136
		v129 = v129 + v136
		v130 = v141
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	goto L37
L54:
	;
	if v78 != int32(1) {
		goto L31
	} else {
		goto L55
	}
L55:
	;
	v158 = F_JsonbIteratorNext(m, v7+int32(-28), v9, int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L31
L57:
	;
	v165 = v9
	goto L59
L58:
	;
	v165 = int32(0)
	goto L59
L59:
	;
	F_pushJsonbValue(m, v7+int32(-24), v78, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L31
L61:
	;
	if v171 != 0 {
		v78 = v171
		goto L29
	} else {
		goto L62
	}
L62:
	;
	goto L30
L63:
	;
	v186 = v182
	goto L20
L64:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(_a_F_jsonb_delete_0), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_jsonb_delete_1), int32(_a_F_jsonb_delete_2), int32(_a_F_jsonb_delete_3))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_each_text(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_each_worker_jsonb(m, l0, int32(_a_F_jsonb_each_text_0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_jsonb_exists_any(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
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
	var v99 int32
	_ = v99
	var v107 int64
	_ = v107
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_deconstruct_array_builtin(m, v17, int32(25), v9+int32(44), v9+int32(40), v9+int32(36))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v28 <= int32(0) {
		v107 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v9 + int32(48)
	return v107
L6:
	;
	v34 = int32(0)
	v36 = v28
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v34))))
	if v42 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v107 = int64(1)
	goto L5
L9:
	;
	goto L8
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	v46 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
	v50 = v45 + v34<<(uint(int32(3))%32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v54&v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v96 = v36
	goto L12
L12:
	;
	v99 = v34 + int32(1)
	if v99 < v96 {
		v34 = v99
		v36 = v96
		goto L7
	} else {
		goto L29
	}
L13:
	;
	v57 = v46
	goto L15
L14:
	;
	v57 = int32(4)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v51 + v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v61 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v90
	v93 = F_findJsonbValueFromContainer(m, v12+int32(4), int32(1610612736), v9)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L27
	}
L17:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v67 == int32(18) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v78 = int32(1)
	if v61&v78 != 0 {
		v90 = int32(base.Ui32(v61)>>(uint(v78)%32)) - v78
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v70 = int32(16)
	goto L22
L21:
	;
	v70 = int32(0)
	goto L22
L22:
	;
	if base.Ui32((v67-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = int32(4)
	goto L25
L24:
	;
	v77 = v70
	goto L25
L25:
	;
	v90 = v77
	goto L16
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	if v93 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	v96 = v95
	goto L12
L29:
	;
	v107 = v6
	goto L5
}
func F_jsonb_extract_path(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_get_jsonb_path_all(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_gt(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v17 = F_compareJsonbContainers(m, v6+int32(4), v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v19 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v23 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(int32(0) < v17))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v17))
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v23 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v17))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v17))
					}
				}
			}
		}
	}
}
func F_jsonb_in_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(1)
	v11 = F_strlen(m, l1)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v11
	if base.Ui32(int32(268435456)) <= base.Ui32(v11) {
		v15 = int32(23)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = F_errsave_start(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				v48 = v15
				m.G0 = v7 + int32(48)
				return v48
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_jsonb_in_object_field_start_0), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(268435455)
						v33 = F_errdetail(m, int32(_a_F_jsonb_in_object_field_start_1), v7)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v16, int32(_a_F_jsonb_in_object_field_start_2), int32(276), int32(_a_F_jsonb_in_object_field_start_3))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v48 = v15
								m.G0 = v7 + int32(48)
								return v48
							}
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l1
		F_pushJsonbValue(m, l0, int32(1), v7+int32(16))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v48 = int32(0)
			m.G0 = v7 + int32(48)
			return v48
		}
	}
}
func F_jsonb_numeric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v18 = F_JsonbExtractScalar(m, v12+int32(4), v9)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			if v18 == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_cannotCastJsonbValue(m, v20, int32(_a_F_jsonb_numeric_0), v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v48 = int64(0)
					m.G0 = v9 + int32(32)
					return v48
				}
			} else {
				switch v20 {
				case 0:
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v28 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
							v48 = int64(0)
							m.G0 = v9 + int32(32)
							return v48
						}
					} else {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
						v48 = int64(0)
						m.G0 = v9 + int32(32)
						return v48
					}
				default:
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_cannotCastJsonbValue(m, v20, int32(_a_F_jsonb_numeric_0), v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v48 = int64(0)
						m.G0 = v9 + int32(32)
						return v48
					}
				case 2:
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					v41 = F_pg_detoast_datum_copy(m, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v43 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								v48 = base.I64_extend_i32_u(v41)
								m.G0 = v9 + int32(32)
								return v48
							}
						} else {
							v48 = base.I64_extend_i32_u(v41)
							m.G0 = v9 + int32(32)
							return v48
						}
					}
				}
			}
		}
	}
}
func F_jsonb_object(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v17 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v17
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v19
	F_pushJsonbValue(m, v7+int32(-32), int32(6), v17)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	switch v16 {
	case 0:
		goto L5
	case 1:
		goto L9
	case 2:
		goto L8
	default:
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L43
	}
L5:
	;
	F_pushJsonbValue(m, v7+int32(-32), int32(7), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L41
	}
L6:
	;
	F_deconstruct_array_builtin(m, v12, int32(25), v7+int32(-4), v7+int32(-8), v7+int32(-12))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v50 == int32(2) {
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	if v29&int32(1) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errmsg(m, int32(_a_F_jsonb_object_0), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_1), int32(1309), int32(_a_F_jsonb_object_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(_a_F_jsonb_object_3), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_1), int32(1316), int32(_a_F_jsonb_object_2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_jsonb_object_4), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_1), int32(1322), int32(_a_F_jsonb_object_2))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	v95 = int32(2)
	v96 = base.I32_div_s(v94, v95)
	if v95 <= v94 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v100 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	F_pfree(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L39
	}
L28:
	;
	v106 = int32(1)
	v107 = v100 << (uint(v106) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v108))))
	if v110 == v106 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v107<<(uint(int32(3))%32))))
	v118 = F_text_to_cstring(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v120 = F_strlen(m, v118)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v120
	v123 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v123
	F_pushJsonbValue(m, v7+int32(-32), v123, v9)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v131 = v107 | int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v132))))
	if v134 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v149 = int32(0)
	goto L35
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v131<<(uint(int32(3))%32))))
	v141 = F_text_to_cstring(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v149
	F_pushJsonbValue(m, v7+int32(-32), int32(2), v9)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v143 = F_strlen(m, v141)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v143
	v149 = int32(1)
	goto L35
L37:
	;
	v157 = v100 + int32(1)
	if v157 != v96 {
		v100 = v157
		goto L28
	} else {
		goto L38
	}
L38:
	;
	goto L29
L39:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	F_pfree(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L5
L41:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v184 = F_JsonbValueToJsonb(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	m.G0 = v9 - int32(-64)
	return base.I64_extend_i32_u(v184)
L43:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_jsonb_object_5), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_1), int32(1338), int32(_a_F_jsonb_object_2))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_ops__add_path_item(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
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
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v321 int64
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v14 - int32(21) {
	case 0, 1, 2, 3:
		v321 = int64(0)
		v323 = F_palloc(m, int32(24))
		mBase = m.M
		v324 = m.ExcPending
		if v324 != 0 {
			return int32(0)
		} else {
			v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v323)+8)) = v321
			*(*int32)(unsafe.Add(mBase, uint32(v323)+16)) = v325
			v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v323))) = v328
			v331 = v323
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v331
			v340 = int32(1)
			m.G0 = v12 + int32(32)
			return v340
		}
	case 4:
		v18 = v12 + int32(16)
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
		} else {
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		if int32(126) <= v23 {
			v31 = v23 - int32(1636608432)
			if v21&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v23) {
					v140 = v21
					v141 = v23
					v142 = v31
					v143 = v31
					v144 = v31
					for {
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
						v147 = v146 + v143
						v148 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v140)+8))
						v151 = v150 + v144
						v153 = int32(4)
						v155 = v148 + v142 - v151 ^ base.I32_rotl(v151, v153)
						v159 = v147 - v155 ^ base.I32_rotl(v155, int32(6))
						v160 = v151 + v147
						v161 = v155 + v160
						v162 = v159 + v161
						v166 = v160 - v159 ^ base.I32_rotl(v159, int32(8))
						v170 = v161 - v166 ^ base.I32_rotl(v166, int32(16))
						v174 = v162 - v170 ^ base.I32_rotl(v170, int32(19))
						v175 = v166 + v162
						v176 = v170 + v175
						v177 = v174 + v176
						v181 = v175 - v174 ^ base.I32_rotl(v174, v153)
						v182 = int32(12)
						v183 = v140 + v182
						v185 = v141 - v182
						if base.Ui32(int32(11)) < base.Ui32(v185) {
							v140 = v183
							v141 = v185
							v142 = v176
							v143 = v177
							v144 = v181
							continue
						} else {
							break
						}
						break
					}
					v188 = v183
					v189 = v185
					v190 = v176
					v191 = v177
					v192 = v181
				} else {
					v188 = v21
					v189 = v23
					v190 = v31
					v191 = v31
					v192 = v31
				}
				switch v189 - int32(1) {
				case 0:
					v251 = v190
					v252 = v191
					v253 = v192
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 1:
					v244 = v190
					v245 = v191
					v246 = v192
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 2:
					v237 = v190
					v238 = v191
					v239 = v192
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 3:
					v231 = v191
					v232 = v192
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)))
					v237 = v233<<(uint(int32(24))%32) + v190
					v238 = v231
					v239 = v232
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 4:
					v227 = v191
					v228 = v192
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+4)))
					v231 = v227 + v229
					v232 = v228
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)))
					v237 = v233<<(uint(int32(24))%32) + v190
					v238 = v231
					v239 = v232
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 5:
					v221 = v191
					v222 = v192
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+5)))
					v227 = v223<<(uint(int32(8))%32) + v221
					v228 = v222
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+4)))
					v231 = v227 + v229
					v232 = v228
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)))
					v237 = v233<<(uint(int32(24))%32) + v190
					v238 = v231
					v239 = v232
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 6:
					v215 = v191
					v216 = v192
					v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+6)))
					v221 = v217<<(uint(int32(16))%32) + v215
					v222 = v216
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+5)))
					v227 = v223<<(uint(int32(8))%32) + v221
					v228 = v222
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+4)))
					v231 = v227 + v229
					v232 = v228
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)))
					v237 = v233<<(uint(int32(24))%32) + v190
					v238 = v231
					v239 = v232
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 7:
					v210 = v192
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+7)))
					v215 = v211<<(uint(int32(24))%32) + v191
					v216 = v210
					v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+6)))
					v221 = v217<<(uint(int32(16))%32) + v215
					v222 = v216
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+5)))
					v227 = v223<<(uint(int32(8))%32) + v221
					v228 = v222
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+4)))
					v231 = v227 + v229
					v232 = v228
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)))
					v237 = v233<<(uint(int32(24))%32) + v190
					v238 = v231
					v239 = v232
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 8:
					v205 = v192
					v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+8)))
					v210 = v206<<(uint(int32(8))%32) + v205
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+7)))
					v215 = v211<<(uint(int32(24))%32) + v191
					v216 = v210
					v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+6)))
					v221 = v217<<(uint(int32(16))%32) + v215
					v222 = v216
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+5)))
					v227 = v223<<(uint(int32(8))%32) + v221
					v228 = v222
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+4)))
					v231 = v227 + v229
					v232 = v228
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)))
					v237 = v233<<(uint(int32(24))%32) + v190
					v238 = v231
					v239 = v232
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 9:
					v200 = v192
					v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+9)))
					v205 = v201<<(uint(int32(16))%32) + v200
					v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+8)))
					v210 = v206<<(uint(int32(8))%32) + v205
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+7)))
					v215 = v211<<(uint(int32(24))%32) + v191
					v216 = v210
					v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+6)))
					v221 = v217<<(uint(int32(16))%32) + v215
					v222 = v216
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+5)))
					v227 = v223<<(uint(int32(8))%32) + v221
					v228 = v222
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+4)))
					v231 = v227 + v229
					v232 = v228
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)))
					v237 = v233<<(uint(int32(24))%32) + v190
					v238 = v231
					v239 = v232
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				case 10:
					v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+10)))
					v200 = v196<<(uint(int32(24))%32) + v192
					v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+9)))
					v205 = v201<<(uint(int32(16))%32) + v200
					v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+8)))
					v210 = v206<<(uint(int32(8))%32) + v205
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+7)))
					v215 = v211<<(uint(int32(24))%32) + v191
					v216 = v210
					v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+6)))
					v221 = v217<<(uint(int32(16))%32) + v215
					v222 = v216
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+5)))
					v227 = v223<<(uint(int32(8))%32) + v221
					v228 = v222
					v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+4)))
					v231 = v227 + v229
					v232 = v228
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+3)))
					v237 = v233<<(uint(int32(24))%32) + v190
					v238 = v231
					v239 = v232
					v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
					v244 = v240<<(uint(int32(16))%32) + v237
					v245 = v238
					v246 = v239
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
					v251 = v247<<(uint(int32(8))%32) + v244
					v252 = v245
					v253 = v246
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
					v258 = v251 + v254
					v259 = v252
					v260 = v253
				default:
					v258 = v190
					v259 = v191
					v260 = v192
				}
			} else {
				if base.Ui32(v23) < base.Ui32(int32(12)) {
					v86 = v21
					v87 = v23
					v88 = v31
					v89 = v31
					v90 = v31
				} else {
					v38 = v21
					v39 = v23
					v40 = v31
					v41 = v31
					v42 = v31
					for {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
						v45 = v44 + v41
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
						v49 = v48 + v42
						v51 = int32(4)
						v53 = v46 + v40 - v49 ^ base.I32_rotl(v49, v51)
						v57 = v45 - v53 ^ base.I32_rotl(v53, int32(6))
						v58 = v49 + v45
						v59 = v53 + v58
						v60 = v57 + v59
						v64 = v58 - v57 ^ base.I32_rotl(v57, int32(8))
						v68 = v59 - v64 ^ base.I32_rotl(v64, int32(16))
						v72 = v60 - v68 ^ base.I32_rotl(v68, int32(19))
						v73 = v64 + v60
						v74 = v68 + v73
						v75 = v72 + v74
						v79 = v73 - v72 ^ base.I32_rotl(v72, v51)
						v80 = int32(12)
						v81 = v38 + v80
						v83 = v39 - v80
						if base.Ui32(int32(11)) < base.Ui32(v83) {
							v38 = v81
							v39 = v83
							v40 = v74
							v41 = v75
							v42 = v79
							continue
						} else {
							break
						}
						break
					}
					v86 = v81
					v87 = v83
					v88 = v74
					v89 = v75
					v90 = v79
				}
				switch v87 - int32(1) {
				case 0:
					v137 = v88
					v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
					v258 = v137 + v138
					v259 = v89
					v260 = v90
				case 1:
					v132 = v88
					v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
					v137 = v133<<(uint(int32(8))%32) + v132
					v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
					v258 = v137 + v138
					v259 = v89
					v260 = v90
				case 2:
					v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+2)))
					v132 = v128<<(uint(int32(16))%32) + v88
					v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
					v137 = v133<<(uint(int32(8))%32) + v132
					v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
					v258 = v137 + v138
					v259 = v89
					v260 = v90
				case 3:
					v125 = v89
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v258 = v126 + v88
					v259 = v125
					v260 = v90
				case 4:
					v122 = v89
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)))
					v125 = v122 + v123
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v258 = v126 + v88
					v259 = v125
					v260 = v90
				case 5:
					v117 = v89
					v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
					v122 = v118<<(uint(int32(8))%32) + v117
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)))
					v125 = v122 + v123
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v258 = v126 + v88
					v259 = v125
					v260 = v90
				case 6:
					v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+6)))
					v117 = v113<<(uint(int32(16))%32) + v89
					v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+5)))
					v122 = v118<<(uint(int32(8))%32) + v117
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+4)))
					v125 = v122 + v123
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v258 = v126 + v88
					v259 = v125
					v260 = v90
				case 7:
					v108 = v90
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
					v258 = v109 + v88
					v259 = v111 + v89
					v260 = v108
				case 8:
					v103 = v90
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+8)))
					v108 = v104<<(uint(int32(8))%32) + v103
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
					v258 = v109 + v88
					v259 = v111 + v89
					v260 = v108
				case 9:
					v98 = v90
					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+9)))
					v103 = v99<<(uint(int32(16))%32) + v98
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+8)))
					v108 = v104<<(uint(int32(8))%32) + v103
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
					v258 = v109 + v88
					v259 = v111 + v89
					v260 = v108
				case 10:
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)))
					v98 = v94<<(uint(int32(24))%32) + v90
					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+9)))
					v103 = v99<<(uint(int32(16))%32) + v98
					v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+8)))
					v108 = v104<<(uint(int32(8))%32) + v103
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
					v258 = v109 + v88
					v259 = v111 + v89
					v260 = v108
				default:
					v258 = v88
					v259 = v89
					v260 = v90
				}
			}
			v263 = int32(14)
			v265 = v259 ^ v260 - base.I32_rotl(v259, v263)
			v269 = v265 ^ v258 - base.I32_rotl(v265, int32(11))
			v273 = v269 ^ v259 - base.I32_rotl(v269, int32(25))
			v277 = v273 ^ v265 - base.I32_rotl(v273, int32(16))
			v281 = v277 ^ v269 - base.I32_rotl(v277, int32(4))
			v285 = v281 ^ v273 - base.I32_rotl(v281, v263)
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v285 ^ v277 - base.I32_rotl(v285, int32(24))
			v292 = v12 + int32(22)
			v295 = F_pg_snprintf(m, v292, int32(10), int32(_a_F_jsonb_ops__add_path_item_0), v12)
			mBase = m.M
			v298 = m.ExcPending
			if v298 != 0 {
				return int32(0)
			} else {
				v301 = int32(8)
				v302 = v292
				v303 = int32(17)
				v305 = v301 + int32(5)
				v306 = F_palloc(m, v305)
				mBase = m.M
				v307 = m.ExcPending
				if v307 != 0 {
					return int32(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v306)+4)) = uint8(v303)
					*(*int32)(unsafe.Add(mBase, uint32(v306))) = v305 << (uint(int32(2)) % 32)
					if v301 != 0 {
						base.MemoryCopy(m, v306+int32(5), v302, v301)
					} else {
					}
					v321 = base.I64_extend_i32_u(v306)
					v323 = F_palloc(m, int32(24))
					mBase = m.M
					v324 = m.ExcPending
					if v324 != 0 {
						return int32(0)
					} else {
						v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int64)(unsafe.Add(mBase, uint32(v323)+8)) = v321
						*(*int32)(unsafe.Add(mBase, uint32(v323)+16)) = v325
						v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v323))) = v328
						v331 = v323
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v331
						v340 = int32(1)
						m.G0 = v12 + int32(32)
						return v340
					}
				}
			}
		} else {
			v301 = v23
			v302 = v21
			v303 = int32(1)
			v305 = v301 + int32(5)
			v306 = F_palloc(m, v305)
			mBase = m.M
			v307 = m.ExcPending
			if v307 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v306)+4)) = uint8(v303)
				*(*int32)(unsafe.Add(mBase, uint32(v306))) = v305 << (uint(int32(2)) % 32)
				if v301 != 0 {
					base.MemoryCopy(m, v306+int32(5), v302, v301)
				} else {
				}
				v321 = base.I64_extend_i32_u(v306)
				v323 = F_palloc(m, int32(24))
				mBase = m.M
				v324 = m.ExcPending
				if v324 != 0 {
					return int32(0)
				} else {
					v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int64)(unsafe.Add(mBase, uint32(v323)+8)) = v321
					*(*int32)(unsafe.Add(mBase, uint32(v323)+16)) = v325
					v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v323))) = v328
					v331 = v323
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v331
					v340 = int32(1)
					m.G0 = v12 + int32(32)
					return v340
				}
			}
		}
	default:
		v340 = v3
		m.G0 = v12 + int32(32)
		return v340
	case 6:
		v331 = v3
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v331
		v340 = int32(1)
		m.G0 = v12 + int32(32)
		return v340
	}
}
func F_jsonb_ops__extract_nodes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v114 = l3
	goto L3
L3:
	;
	return v114
L4:
	;
	v12 = l3
	v13 = v8
	goto L7
L5:
	;
	v37 = l3
	goto L6
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v41 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v16 == int32(25) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v37 = v30
	goto L6
L9:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v21 = F_palloc(m, int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v30 = v12
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v33 != 0 {
		v12 = v30
		v13 = v33
		goto L7
	} else {
		goto L15
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(2)
	v28 = F_lappend(m, v12, v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v30 = v28
	goto L11
L15:
	;
	goto L8
L16:
	;
	v105 = F_lappend(m, v37, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L12
	} else {
		goto L34
	}
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v44 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	goto L19
L19:
	;
	v93 = F_make_scalar_key(m, l2, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L32
	}
L20:
	;
	v84 = F_make_scalar_key(m, l2, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L12
	} else {
		goto L30
	}
L21:
	;
	v83 = int32(0)
	goto L20
L22:
	;
	v56 = F_make_scalar_key(m, l2, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	v45 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v46 == v45 {
		v83 = v45
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	switch v50 - int32(21) {
	case 0, 2:
		v83 = int32(1)
		goto L20
	default:
		goto L21
	case 3:
		goto L22
	}
L25:
	;
	v59 = F_palloc(m, int32(16))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(2)
	v65 = F_make_scalar_key(m, l2, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v68 = F_palloc(m, int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(2)
	v74 = F_palloc(m, int32(24))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
	v103 = v74
	goto L16
L30:
	;
	v87 = F_palloc(m, int32(16))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = int32(2)
	v103 = v87
	goto L16
L32:
	;
	v96 = F_palloc(m, int32(16))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v96)+8)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = int32(2)
	v103 = v96
	goto L16
L34:
	;
	v114 = v105
	goto L3
}
func F_jsonb_path_ops__add_path_item(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v10 - int32(21) {
	case 0, 2:
		v28 = v9
		m.G0 = v7 + int32(32)
		return v28
	default:
		v28 = int32(0)
		m.G0 = v7 + int32(32)
		return v28
	case 4:
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
		v18 = v7 + int32(8)
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
		} else {
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v21
		F_JsonbHashScalarValue(m, v7, l0)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v28 = v9
			m.G0 = v7 + int32(32)
			return v28
		}
	case 6:
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		v28 = v9
		m.G0 = v7 + int32(32)
		return v28
	}
}
func F_jsonb_path_query_first_internal(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v23 int64
	_ = v23
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
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v21 = F_pg_detoast_datum(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(8589934592)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v10
				v33 = F_executeJsonPath(m, v18, v21, int32(1534), int32(1535), v13, base.B2i32(v23 == int64(0)), v10, l1)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					if v35 != 0 {
						v38 = F_JsonbValueToJsonb(m, v10+int32(16))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v44 = base.I64_extend_i32_u(v38)
							m.G0 = v10 + int32(80)
							return v44
						}
					} else {
						v41 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
						v44 = int64(0)
						m.G0 = v10 + int32(80)
						return v44
					}
				}
			}
		}
	}
}
func F_jsonb_path_query_tz(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_populate_record_valid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_populate_record_valid[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v9
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_jsonb_populate_record_valid[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
	v17 = F_populate_record_worker(m, l0, int32(_a_F_jsonb_populate_record_valid_0), int32(0), int32(1), v6)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
		m.G0 = v6 + int32(16)
		return v21 ^ int64(1)
	}
}
func F_jsonb_put_escaped_value(m *base.Module, l0 int32, l1 int32) {
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
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v3 {
	case 0:
		F_appendBinaryStringInfo(m, l0, int32(_a_F_jsonb_put_escaped_value_0), int32(4))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			return
		}
	case 1:
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		F_escape_json_with_len(m, l0, v4, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	case 2:
		v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+8)))
		v11 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_appendStringInfoString(m, l0, base.I32_wrap_i64(v11))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		}
	case 3:
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
		if v16 == int32(1) {
			F_appendBinaryStringInfo(m, l0, int32(_a_F_jsonb_put_escaped_value_1), int32(4))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		} else {
			F_appendBinaryStringInfo(m, l0, int32(_a_F_jsonb_put_escaped_value_2), int32(5))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				return
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_jsonb_put_escaped_value_3), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_jsonb_put_escaped_value_4), int32(363), int32(_a_F_jsonb_put_escaped_value_5))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
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
func F_jsonb_string_to_tsvector_byid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
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
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+28)) = uint32(v10)
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v17
		v22 = v8 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v22
		F_iterate_jsonb_values(m, v12, int32(2), v8+int32(24))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v29 = F_make_tsvector(m, v22)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v31 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						m.G0 = v8 + int32(32)
						return base.I64_extend_i32_u(v29)
					}
				} else {
					m.G0 = v8 + int32(32)
					return base.I64_extend_i32_u(v29)
				}
			}
		}
	}
}
func F_jsonb_strip_nulls(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+88)) = int32(0)
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+80)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v7)+72)) = v16
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v20 == int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = base.B2i32(v23 != int64(0))
	goto L5
L4:
	;
	v26 = int32(0)
	goto L5
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+7)))
	if v27&int32(16) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = F_JsonbIteratorInit(m, v10+int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v112 = v10
	goto L8
L8:
	;
	m.G0 = v7 + int32(96)
	return base.I64_extend_i32_u(v112)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+92)) = v34
	goto L10
L10:
	;
	v42 = int32(0)
	goto L12
L11:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v7)+72))
	v108 = F_JsonbValueToJsonb(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L35
	}
L12:
	;
	v51 = F_JsonbIteratorNext(m, v7+int32(92), v7+int32(40), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	goto L11
L14:
	;
	goto L13
L15:
	;
	if v42&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v7)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v53
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v55
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v7)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v7)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v59
	v42 = int32(1)
	goto L12
L17:
	;
	switch v51 {
	case 0:
		goto L14
	case 1:
		goto L16
	default:
		goto L15
	}
L18:
	;
	if v51 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v79 = int32(0)
	if base.B2i32(v26 == v79)|base.B2i32(v51 != int32(3)) == v79 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	if v67 == v66 {
		v42 = v66
		goto L12
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_pushJsonbValue(m, v7+int32(72), int32(1), v7+int32(8))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	goto L20
L26:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	if v87 == v86 {
		v42 = v86
		goto L12
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v51&int32(-2) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	F_pushJsonbValue(m, v7+int32(72), v51, v7+int32(40))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v101 = int32(0)
	F_pushJsonbValue(m, v7+int32(72), v51, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L10
L34:
	;
	v42 = v101
	goto L12
L35:
	;
	v112 = v108
	goto L8
}
func F_jsonb_typeof(m *base.Module, l0 int32) int64 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v14 = v9 + int32(4)
		v16 = v6 + int32(16)
		v17 = F_JsonbExtractScalar(m, v14, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			if v17 != 0 {
				v19 = F_JsonbTypeName(m, v16)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v31 = v19
					v32 = F_cstring_to_text(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						m.G0 = v6 + int32(48)
						return base.I64_extend_i32_u(v32)
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v22&int32(1073741824) != 0 {
					v31 = int32(_a_F_jsonb_typeof_0)
					v32 = F_cstring_to_text(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						m.G0 = v6 + int32(48)
						return base.I64_extend_i32_u(v32)
					}
				} else {
					if v22&int32(536870912) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v43
							F_errmsg_internal(m, int32(_a_F_jsonb_typeof_1), v6)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_jsonb_typeof_2), int32(163), int32(_a_F_jsonb_typeof_3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = int32(_a_F_jsonb_typeof_4)
						v32 = F_cstring_to_text(m, v31)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							m.G0 = v6 + int32(48)
							return base.I64_extend_i32_u(v32)
						}
					}
				}
			}
		}
	}
}
