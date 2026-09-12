package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pa_free_worker_info(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != 0 {
		F_shm_mq_detach(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v6 != 0 {
				F_shm_mq_detach(m, v6)
				mBase = m.M
				v8 = m.ExcPending
				if v8 != 0 {
					return
				} else {
					v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
					if v9 == int32(1) {
						v13 = *(*int32)(unsafe.Add(mBase, _consts[629]))
						v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
						v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						F_stream_cleanup_files(m, v14, v16)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v19 != 0 {
								F_dsm_detach(m, v19)
								mBase = m.M
								v21 = m.ExcPending
								if v21 != 0 {
									return
								} else {
									v22 = int32(4425172)
									v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
									v25 = F_list_delete_ptr(m, v24, l0)
									mBase = m.M
									v26 = m.ExcPending
									if v26 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
										F_pfree(m, l0)
										mBase = m.M
										v29 = m.ExcPending
										if v29 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v22 = int32(4425172)
								v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
								v25 = F_list_delete_ptr(m, v24, l0)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
									F_pfree(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v19 != 0 {
							F_dsm_detach(m, v19)
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								v22 = int32(4425172)
								v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
								v25 = F_list_delete_ptr(m, v24, l0)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
									F_pfree(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v22 = int32(4425172)
							v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
							v25 = F_list_delete_ptr(m, v24, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
								F_pfree(m, l0)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			} else {
				v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
				if v9 == int32(1) {
					v13 = *(*int32)(unsafe.Add(mBase, _consts[629]))
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					F_stream_cleanup_files(m, v14, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v19 != 0 {
							F_dsm_detach(m, v19)
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								v22 = int32(4425172)
								v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
								v25 = F_list_delete_ptr(m, v24, l0)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
									F_pfree(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v22 = int32(4425172)
							v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
							v25 = F_list_delete_ptr(m, v24, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
								F_pfree(m, l0)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v19 != 0 {
						F_dsm_detach(m, v19)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							v22 = int32(4425172)
							v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
							v25 = F_list_delete_ptr(m, v24, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
								F_pfree(m, l0)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v22 = int32(4425172)
						v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
						v25 = F_list_delete_ptr(m, v24, l0)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
							F_pfree(m, l0)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 != 0 {
			F_shm_mq_detach(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
				if v9 == int32(1) {
					v13 = *(*int32)(unsafe.Add(mBase, _consts[629]))
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					F_stream_cleanup_files(m, v14, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v19 != 0 {
							F_dsm_detach(m, v19)
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								v22 = int32(4425172)
								v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
								v25 = F_list_delete_ptr(m, v24, l0)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
									F_pfree(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v22 = int32(4425172)
							v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
							v25 = F_list_delete_ptr(m, v24, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
								F_pfree(m, l0)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v19 != 0 {
						F_dsm_detach(m, v19)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							v22 = int32(4425172)
							v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
							v25 = F_list_delete_ptr(m, v24, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
								F_pfree(m, l0)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v22 = int32(4425172)
						v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
						v25 = F_list_delete_ptr(m, v24, l0)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
							F_pfree(m, l0)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
			if v9 == int32(1) {
				v13 = *(*int32)(unsafe.Add(mBase, _consts[629]))
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				F_stream_cleanup_files(m, v14, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v19 != 0 {
						F_dsm_detach(m, v19)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							v22 = int32(4425172)
							v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
							v25 = F_list_delete_ptr(m, v24, l0)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
								F_pfree(m, l0)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v22 = int32(4425172)
						v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
						v25 = F_list_delete_ptr(m, v24, l0)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
							F_pfree(m, l0)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v19 != 0 {
					F_dsm_detach(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = int32(4425172)
						v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
						v25 = F_list_delete_ptr(m, v24, l0)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
							F_pfree(m, l0)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v22 = int32(4425172)
					v24 = *(*int32)(unsafe.Add(mBase, _consts[630]))
					v25 = F_list_delete_ptr(m, v24, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[630])) = v25
						F_pfree(m, l0)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
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
