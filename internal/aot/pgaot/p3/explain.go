package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExplainOpenWorker(m *base.Module, l0 int32, l1 int32) {
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v12
	v15 = l0 << (uint(int32(4)) % 32)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v17 = v15 + v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+l0))))
	if v20 == int32(0) {
		F_initStringInfo(m, v17)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v25 + v15
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			switch v28 - int32(1) {
			case 0:
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v41 + int32(2)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v45 != 0 {
					F_ExplainPropertyInteger(m, int32(228229), int32(0), base.I64_extend_i32_s(l0), l1)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v73 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if v77 == int32(0) {
								F_ExplainIndentText(m, l1)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_appendStringInfo(m, v82, int32(730177), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v53 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v73 == int32(0) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
						if v77 == int32(0) {
							F_ExplainIndentText(m, l1)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v82, int32(730177), v9)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
									m.G0 = v9 + int32(16)
									return
								}
							}
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
							m.G0 = v9 + int32(16)
							return
						}
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			case 1:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				v33 = F_lcons_int(m, int32(0), v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v33
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v41 + int32(2)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v45 != 0 {
						F_ExplainPropertyInteger(m, int32(228229), int32(0), base.I64_extend_i32_s(l0), l1)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v53 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							if v73 == int32(0) {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
								if v77 == int32(0) {
									F_ExplainIndentText(m, l1)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
										F_appendStringInfo(m, v82, int32(730177), v9)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
											m.G0 = v9 + int32(16)
											return
										}
									}
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
									m.G0 = v9 + int32(16)
									return
								}
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v73 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if v77 == int32(0) {
								F_ExplainIndentText(m, l1)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_appendStringInfo(m, v82, int32(730177), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				}
			case 2:
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				v38 = F_lcons_int(m, int32(0), v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v38
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v41 + int32(2)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v45 != 0 {
						F_ExplainPropertyInteger(m, int32(228229), int32(0), base.I64_extend_i32_s(l0), l1)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
							v53 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							if v73 == int32(0) {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
								if v77 == int32(0) {
									F_ExplainIndentText(m, l1)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
										F_appendStringInfo(m, v82, int32(730177), v9)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
											m.G0 = v9 + int32(16)
											return
										}
									}
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
									m.G0 = v9 + int32(16)
									return
								}
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v73 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if v77 == int32(0) {
								F_ExplainIndentText(m, l1)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_appendStringInfo(m, v82, int32(730177), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				}
			default:
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v45 != 0 {
					F_ExplainPropertyInteger(m, int32(228229), int32(0), base.I64_extend_i32_s(l0), l1)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						v53 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						if v73 == int32(0) {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
							if v77 == int32(0) {
								F_ExplainIndentText(m, l1)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
									F_appendStringInfo(m, v82, int32(730177), v9)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					v53 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v51+l0))) = uint8(v53)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					if v73 == int32(0) {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
						if v77 == int32(0) {
							F_ExplainIndentText(m, l1)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								F_appendStringInfo(m, v82, int32(730177), v9)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
									m.G0 = v9 + int32(16)
									return
								}
							}
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
							m.G0 = v9 + int32(16)
							return
						}
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		switch v60 - int32(1) {
		case 0:
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v68 + int32(2)
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v73 == int32(0) {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				if v77 == int32(0) {
					F_ExplainIndentText(m, l1)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_appendStringInfo(m, v82, int32(730177), v9)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		case 1, 2:
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v56+l0<<(uint(int32(2))%32))))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			v65 = F_lcons_int(m, v63, v64)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v65
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v68 + int32(2)
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				if v73 == int32(0) {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
					if v77 == int32(0) {
						F_ExplainIndentText(m, l1)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_appendStringInfo(m, v82, int32(730177), v9)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
								m.G0 = v9 + int32(16)
								return
							}
						}
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		default:
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v73 == int32(0) {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				if v77 == int32(0) {
					F_ExplainIndentText(m, l1)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_appendStringInfo(m, v82, int32(730177), v9)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v88 + int32(1)
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_ExplainPreScanNode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v43 int32
	_ = v43
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	switch v7 - int32(333) {
	case 0:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
		v26 = F_bms_add_member(m, v24, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+152))
			if v29 != 0 {
				v30 = F_bms_add_member(m, v26, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
					v33 = v30
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
					if v34 == int32(0) {
						v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							return v57
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						v40 = F_bms_add_member(m, v33, v39)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v52 = v40
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
							v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								return v57
							}
						}
					}
				}
			} else {
				v33 = v26
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
				if v34 == int32(0) {
					v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						return v57
					}
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					v40 = F_bms_add_member(m, v33, v39)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v52 = v40
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
						v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							return v57
						}
					}
				}
			}
		}
	case 1:
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
		v44 = F_bms_add_members(m, v42, v43)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v52 = v44
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
			v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				return v57
			}
		}
	case 2:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
		v48 = F_bms_add_members(m, v46, v47)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			v52 = v48
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
			v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				return v57
			}
		}
	default:
		v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			return v57
		}
	case 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
		v12 = F_bms_add_member(m, v10, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v52 = v12
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
			v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				return v57
			}
		}
	case 21:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+116))
		v18 = F_bms_add_members(m, v16, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v52 = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
			v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				return v57
			}
		}
	case 22:
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+100))
		v22 = F_bms_add_members(m, v20, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v52 = v22
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v52
			v57 = F_planstate_tree_walker_impl(m, l0, int32(548), l1)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				return v57
			}
		}
	}
}
func F_ExplainPropertyUInteger(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = l2
	v15 = F_pg_snprintf(m, v8+int32(16), int32(32), int32(38298), v8)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		F_ExplainProperty(m, l0, l1, v8+int32(16), int32(1), l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			m.G0 = v8 + int32(48)
			return
		}
	}
}
func F_ExplainSubPlans(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v5 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 < v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = l1
	v15 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v15<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v26 = F_bms_is_member(m, v24, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v32 = F_bms_add_member(m, v30, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	v43 = v12
	goto L10
L10:
	;
	v45 = v15 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 < v46 {
		v12 = v43
		v15 = v45
		goto L4
	} else {
		goto L15
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+48)) = v32
	v35 = F_lcons(m, v23, v12)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	F_ExplainNode(m, v37, v35, l2, v38, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v41 = F_list_delete_first(m, v35)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v43 = v41
	goto L10
L15:
	;
	goto L5
}
