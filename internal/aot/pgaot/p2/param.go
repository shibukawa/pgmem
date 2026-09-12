package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bind_param_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 < v2 {
		m.G0 = v7 + int32(80)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v12 != 0 {
			F_initStringInfo(m, v7-int32(-64))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v21 = *(*int32)(unsafe.Add(mBase, _consts[885]))
				F_appendStringInfoStringQuoted(m, v7-int32(-64), v19, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+64))
					v25 = v24
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v26 == int32(0) {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v59 = v57 + int32(1)
							if v25 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v59
								F_errcontext_msg(m, int32(484883), v7)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									m.G0 = v7 + int32(80)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v25
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v59
								F_errcontext_msg(m, int32(207627), v7+int32(16))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_pfree(m, v25)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										m.G0 = v7 + int32(80)
										return
									}
								}
							}
						}
					} else {
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
						if v29 == int32(0) {
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v59 = v57 + int32(1)
								if v25 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v59
									F_errcontext_msg(m, int32(484883), v7)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										m.G0 = v7 + int32(80)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v25
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v59
									F_errcontext_msg(m, int32(207627), v7+int32(16))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_pfree(m, v25)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											m.G0 = v7 + int32(80)
											return
										}
									}
								}
							}
						} else {
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v37 = v35 + int32(1)
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if v25 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v25
									*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v37
									*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v38
									F_errcontext_msg(m, int32(207661), v7+int32(48))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										F_pfree(m, v25)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											m.G0 = v7 + int32(80)
											return
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v37
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v38
									F_errcontext_msg(m, int32(484959), v7+int32(32))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										m.G0 = v7 + int32(80)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v25 = v2
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v26 == int32(0) {
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v59 = v57 + int32(1)
					if v25 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v59
						F_errcontext_msg(m, int32(484883), v7)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							m.G0 = v7 + int32(80)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v25
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v59
						F_errcontext_msg(m, int32(207627), v7+int32(16))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_pfree(m, v25)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								m.G0 = v7 + int32(80)
								return
							}
						}
					}
				}
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
				if v29 == int32(0) {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v59 = v57 + int32(1)
						if v25 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v59
							F_errcontext_msg(m, int32(484883), v7)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								m.G0 = v7 + int32(80)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v25
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v59
							F_errcontext_msg(m, int32(207627), v7+int32(16))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								F_pfree(m, v25)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									m.G0 = v7 + int32(80)
									return
								}
							}
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v37 = v35 + int32(1)
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v25 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v25
							*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v37
							*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v38
							F_errcontext_msg(m, int32(207661), v7+int32(48))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								F_pfree(m, v25)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									m.G0 = v7 + int32(80)
									return
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v37
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v38
							F_errcontext_msg(m, int32(484959), v7+int32(32))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								m.G0 = v7 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_generate_new_exec_param(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v8 = F_palloc0(m, int32(28))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(4294967304)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
		if v15 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v18 = v16
		} else {
			v18 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
		v22 = F_lappend_oid(m, v21, l1)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
			return v8
		}
	}
}
